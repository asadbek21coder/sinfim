package getownerdashboard

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/modules/organization/domain/membership"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type Request struct {
	OrganizationID *string `query:"organization_id" validate:"omitempty,uuid"`
}

type MetricDTO struct {
	ActiveCourses    int `json:"active_courses"`
	ActiveClasses    int `json:"active_classes"`
	ActiveStudents   int `json:"active_students"`
	NewLeads         int `json:"new_leads"`
	PendingHomework  int `json:"pending_homework"`
	PendingAccess    int `json:"pending_access"`
	NeedsRevision    int `json:"needs_revision"`
	CompletedLessons int `json:"completed_lessons"`
}

type PendingHomeworkDTO struct {
	SubmissionID    string    `json:"submission_id" bun:"submission_id"`
	StudentFullName string    `json:"student_full_name" bun:"student_full_name"`
	ClassName       string    `json:"class_name" bun:"class_name"`
	LessonTitle     string    `json:"lesson_title" bun:"lesson_title"`
	HomeworkTitle   string    `json:"homework_title" bun:"homework_title"`
	SubmissionType  string    `json:"submission_type" bun:"submission_type"`
	SubmittedAt     time.Time `json:"submitted_at" bun:"submitted_at"`
}

type PendingAccessDTO struct {
	ClassID       string  `json:"class_id" bun:"class_id"`
	ClassName     string  `json:"class_name" bun:"class_name"`
	StudentUserID string  `json:"student_user_id" bun:"student_user_id"`
	StudentName   string  `json:"student_name" bun:"student_name"`
	PhoneNumber   string  `json:"phone_number" bun:"phone_number"`
	AccessStatus  string  `json:"access_status" bun:"access_status"`
	PaymentStatus string  `json:"payment_status" bun:"payment_status"`
	Note          *string `json:"note" bun:"note"`
}

type CourseProgressDTO struct {
	CourseID        string `json:"course_id" bun:"course_id"`
	CourseTitle     string `json:"course_title" bun:"course_title"`
	ClassCount      int    `json:"class_count" bun:"class_count"`
	StudentCount    int    `json:"student_count" bun:"student_count"`
	LessonCount     int    `json:"lesson_count" bun:"lesson_count"`
	CompletionCount int    `json:"completion_count" bun:"completion_count"`
	PendingHomework int    `json:"pending_homework" bun:"pending_homework"`
	ProgressPercent int    `json:"progress_percent" bun:"progress_percent"`
}

type RecentActivityDTO struct {
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	CreatedAt time.Time `json:"created_at"`
}

type OrganizationDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Role string `json:"role"`
}

type Response struct {
	Organization    OrganizationDTO      `json:"organization"`
	Metrics         MetricDTO            `json:"metrics"`
	PendingHomework []PendingHomeworkDTO `json:"pending_homework"`
	PendingAccess   []PendingAccessDTO   `json:"pending_access"`
	CourseProgress  []CourseProgressDTO  `json:"course_progress"`
	RecentActivity  []RecentActivityDTO  `json:"recent_activity"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(db *bun.DB) UseCase { return &usecase{db: db} }

type usecase struct{ db *bun.DB }

func (uc *usecase) OperationID() string { return "get-owner-dashboard" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	org, err := uc.resolveOrganization(ctx, in.OrganizationID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	metrics, err := uc.metrics(ctx, org.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	pendingHomework, err := uc.pendingHomework(ctx, org.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	pendingAccess, err := uc.pendingAccess(ctx, org.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	progress, err := uc.courseProgress(ctx, org.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	activity, err := uc.recentActivity(ctx, org.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{Organization: org, Metrics: metrics, PendingHomework: pendingHomework, PendingAccess: pendingAccess, CourseProgress: progress, RecentActivity: activity}, nil
}

func (uc *usecase) resolveOrganization(ctx context.Context, requested *string) (OrganizationDTO, error) {
	userCtx := auth.MustUserContext(ctx)
	query := uc.db.NewSelect().TableExpr("organization.organizations AS o").ColumnExpr("o.id, o.name, o.slug")
	role := "PLATFORM_ADMIN"
	if requested != nil && auth.HasPermission(userCtx, auth.PermissionUserRead) {
		var org OrganizationDTO
		err := query.Where("o.id = ?", *requested).Limit(1).Scan(ctx, &org)
		if errors.Is(err, sql.ErrNoRows) {
			return OrganizationDTO{}, errx.New("organization not found", errx.WithType(errx.T_NotFound), errx.WithCode("ORGANIZATION_NOT_FOUND"))
		}
		org.Role = role
		return org, err
	}
	query = query.ColumnExpr("m.role AS role").Join("JOIN auth.user_memberships AS m ON m.organization_id = o.id").Where("m.user_id = ?", userCtx.UserID).Where("m.is_active = true")
	if requested != nil {
		query = query.Where("o.id = ?", *requested)
	}
	query = query.Where("m.role IN (?)", bun.In([]string{membership.RoleOwner, membership.RoleTeacher, membership.RoleMentor})).Order("m.created_at ASC")
	var org OrganizationDTO
	err := query.Limit(1).Scan(ctx, &org)
	if errors.Is(err, sql.ErrNoRows) {
		return OrganizationDTO{}, errx.New("organization workspace not found", errx.WithType(errx.T_Forbidden), errx.WithCode(auth.CodeForbidden))
	}
	return org, err
}

func (uc *usecase) metrics(ctx context.Context, organizationID string) (MetricDTO, error) {
	var item MetricDTO
	err := uc.db.NewSelect().ColumnExpr("(SELECT COUNT(*) FROM catalog.courses c WHERE c.organization_id = ? AND c.status <> 'archived') AS active_courses", organizationID).
		ColumnExpr("(SELECT COUNT(*) FROM classroom.classes cl WHERE cl.organization_id = ? AND cl.status = 'active') AS active_classes", organizationID).
		ColumnExpr("(SELECT COUNT(DISTINCT e.student_user_id) FROM classroom.enrollments e WHERE e.organization_id = ? AND e.status = 'active') AS active_students", organizationID).
		ColumnExpr("(SELECT COUNT(*) FROM lead.leads l WHERE l.organization_id = ? AND l.status = 'new') AS new_leads", organizationID).
		ColumnExpr("(SELECT COUNT(*) FROM homework.submissions hs WHERE hs.organization_id = ? AND hs.status = 'submitted') AS pending_homework", organizationID).
		ColumnExpr("(SELECT COUNT(*) FROM classroom.access_grants ag WHERE ag.organization_id = ? AND (ag.access_status = 'pending' OR ag.payment_status IN ('unknown', 'pending'))) AS pending_access", organizationID).
		ColumnExpr("(SELECT COUNT(*) FROM homework.submissions hs WHERE hs.organization_id = ? AND hs.status = 'needs_revision') AS needs_revision", organizationID).
		ColumnExpr("(SELECT COUNT(*) FROM learning.lesson_completions lc WHERE lc.organization_id = ?) AS completed_lessons", organizationID).
		Scan(ctx, &item)
	return item, err
}

func (uc *usecase) pendingHomework(ctx context.Context, organizationID string) ([]PendingHomeworkDTO, error) {
	items := make([]PendingHomeworkDTO, 0)
	err := uc.db.NewSelect().TableExpr("homework.submissions AS hs").
		ColumnExpr("hs.id AS submission_id, u.full_name AS student_full_name, cl.name AS class_name, l.title AS lesson_title, hd.title AS homework_title, hs.submission_type, hs.submitted_at").
		Join("JOIN auth.users AS u ON u.id = hs.student_user_id").
		Join("JOIN classroom.classes AS cl ON cl.id = hs.class_id").
		Join("JOIN catalog.lessons AS l ON l.id = hs.lesson_id").
		Join("JOIN homework.definitions AS hd ON hd.id = hs.definition_id").
		Where("hs.organization_id = ?", organizationID).
		Where("hs.status = 'submitted'").
		Order("hs.submitted_at DESC").
		Limit(6).
		Scan(ctx, &items)
	return items, err
}

func (uc *usecase) pendingAccess(ctx context.Context, organizationID string) ([]PendingAccessDTO, error) {
	items := make([]PendingAccessDTO, 0)
	err := uc.db.NewSelect().TableExpr("classroom.access_grants AS ag").
		ColumnExpr("cl.id AS class_id, cl.name AS class_name, ag.student_user_id, u.full_name AS student_name, u.phone_number, ag.access_status, ag.payment_status, ag.note").
		Join("JOIN classroom.classes AS cl ON cl.id = ag.class_id").
		Join("JOIN auth.users AS u ON u.id = ag.student_user_id").
		Where("ag.organization_id = ?", organizationID).
		Where("ag.access_status = 'pending' OR ag.payment_status IN ('unknown', 'pending')").
		Order("ag.created_at DESC").
		Limit(6).
		Scan(ctx, &items)
	return items, err
}

func (uc *usecase) courseProgress(ctx context.Context, organizationID string) ([]CourseProgressDTO, error) {
	items := make([]CourseProgressDTO, 0)
	err := uc.db.NewSelect().TableExpr("catalog.courses AS c").
		ColumnExpr("c.id AS course_id, c.title AS course_title").
		ColumnExpr("COUNT(DISTINCT cl.id) AS class_count").
		ColumnExpr("COUNT(DISTINCT e.student_user_id) AS student_count").
		ColumnExpr("COUNT(DISTINCT l.id) AS lesson_count").
		ColumnExpr("COUNT(DISTINCT lc.id) AS completion_count").
		ColumnExpr("COUNT(DISTINCT hs.id) FILTER (WHERE hs.status = 'submitted') AS pending_homework").
		ColumnExpr("CASE WHEN COUNT(DISTINCT l.id) * GREATEST(COUNT(DISTINCT e.student_user_id), 1) = 0 THEN 0 ELSE ROUND(COUNT(DISTINCT lc.id)::numeric / (COUNT(DISTINCT l.id) * GREATEST(COUNT(DISTINCT e.student_user_id), 1)) * 100)::int END AS progress_percent").
		Join("LEFT JOIN classroom.classes AS cl ON cl.course_id = c.id AND cl.status = 'active'").
		Join("LEFT JOIN classroom.enrollments AS e ON e.class_id = cl.id AND e.status = 'active'").
		Join("LEFT JOIN catalog.lessons AS l ON l.course_id = c.id AND l.status = 'published'").
		Join("LEFT JOIN learning.lesson_completions AS lc ON lc.lesson_id = l.id AND lc.class_id = cl.id").
		Join("LEFT JOIN homework.submissions AS hs ON hs.lesson_id = l.id AND hs.class_id = cl.id").
		Where("c.organization_id = ?", organizationID).
		Where("c.status <> 'archived'").
		Group("c.id").
		Order("c.created_at DESC").
		Limit(6).
		Scan(ctx, &items)
	return items, err
}

func (uc *usecase) recentActivity(ctx context.Context, organizationID string) ([]RecentActivityDTO, error) {
	items := make([]RecentActivityDTO, 0)
	err := uc.db.NewSelect().
		ColumnExpr("type, title, subtitle, created_at").
		TableExpr("(SELECT 'lead' AS type, l.full_name AS title, l.phone_number AS subtitle, l.created_at FROM lead.leads l WHERE l.organization_id = ? UNION ALL SELECT 'homework' AS type, u.full_name AS title, hd.title AS subtitle, hs.submitted_at AS created_at FROM homework.submissions hs JOIN auth.users u ON u.id = hs.student_user_id JOIN homework.definitions hd ON hd.id = hs.definition_id WHERE hs.organization_id = ? UNION ALL SELECT 'access' AS type, u.full_name AS title, ag.access_status || '/' || ag.payment_status AS subtitle, ag.updated_at AS created_at FROM classroom.access_grants ag JOIN auth.users u ON u.id = ag.student_user_id WHERE ag.organization_id = ?) AS activity", organizationID, organizationID, organizationID).
		Order("created_at DESC").
		Limit(8).
		Scan(ctx, &items)
	return items, err
}
