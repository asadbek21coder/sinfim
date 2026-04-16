package getstudentdashboard

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"time"

	"go-enterprise-blueprint/internal/modules/learning/usecase/shared"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type Request struct {
	ClassID *string `query:"class_id" validate:"omitempty,uuid"`
}

type StudentDTO struct {
	ID string `json:"id"`
}

type OrganizationDTO struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Slug    string  `json:"slug"`
	LogoURL *string `json:"logo_url"`
}

type ClassDTO struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CourseID      string `json:"course_id"`
	CourseTitle   string `json:"course_title"`
	AccessStatus  string `json:"access_status"`
	PaymentStatus string `json:"payment_status"`
}

type ProgressDTO struct {
	CompletedLessons int `json:"completed_lessons"`
	TotalLessons     int `json:"total_lessons"`
	Percentage       int `json:"percentage"`
}

type LessonDTO struct {
	LessonID      string     `json:"lesson_id"`
	Title         string     `json:"title"`
	Description   *string    `json:"description"`
	Status        string     `json:"status"`
	AvailableAt   *time.Time `json:"available_at"`
	OrderNumber   int        `json:"order_number"`
	PublishDay    int        `json:"publish_day"`
	HasVideo      bool       `json:"has_video"`
	HasMaterials  bool       `json:"has_materials"`
	MaterialCount int        `json:"material_count"`
	Completed     bool       `json:"completed"`
}

type Response struct {
	Student      StudentDTO      `json:"student"`
	Organization OrganizationDTO `json:"organization"`
	Class        ClassDTO        `json:"class"`
	Progress     ProgressDTO     `json:"progress"`
	Locked       bool            `json:"locked"`
	Lessons      []LessonDTO     `json:"lessons"`
	Classes      []ClassDTO      `json:"classes"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(db *bun.DB) UseCase { return &usecase{db: db} }

type usecase struct{ db *bun.DB }

func (uc *usecase) OperationID() string { return "get-student-dashboard" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	userID := auth.MustUserContext(ctx).UserID
	classes, err := uc.listClasses(ctx, userID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if len(classes) == 0 {
		return nil, errx.New("student enrollment not found", errx.WithType(errx.T_NotFound), errx.WithCode(shared.CodeEnrollmentNotFound))
	}
	selected := classes[0]
	if in.ClassID != nil {
		found := false
		for _, item := range classes {
			if item.ClassID == *in.ClassID {
				selected = item
				found = true
				break
			}
		}
		if !found {
			return nil, errx.New("student enrollment not found", errx.WithType(errx.T_NotFound), errx.WithCode(shared.CodeEnrollmentNotFound))
		}
	}
	classDTOs := make([]ClassDTO, 0, len(classes))
	for _, item := range classes {
		classDTOs = append(classDTOs, toClassDTO(item))
	}
	lessons, err := uc.listLessons(ctx, userID, selected)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	lessonDTOs, progress := buildLessons(selected, lessons)
	locked := selected.AccessStatus != "active"
	if locked {
		lessonDTOs = nil
	}
	return &Response{Student: StudentDTO{ID: userID}, Organization: OrganizationDTO{ID: selected.OrganizationID, Name: selected.Organization, Slug: selected.OrganizationSlug, LogoURL: selected.LogoURL}, Class: toClassDTO(selected), Progress: progress, Locked: locked, Lessons: lessonDTOs, Classes: classDTOs}, nil
}

func (uc *usecase) listClasses(ctx context.Context, userID string) ([]shared.ClassContext, error) {
	items := make([]shared.ClassContext, 0)
	err := uc.db.NewSelect().TableExpr("classroom.enrollments AS e").
		ColumnExpr("o.id AS organization_id").
		ColumnExpr("o.name AS organization").
		ColumnExpr("o.slug AS organization_slug").
		ColumnExpr("o.logo_url AS logo_url").
		ColumnExpr("cl.id AS class_id").
		ColumnExpr("cl.name AS class_name").
		ColumnExpr("cl.start_date AS start_date").
		ColumnExpr("cl.lesson_cadence AS lesson_cadence").
		ColumnExpr("c.id AS course_id").
		ColumnExpr("c.title AS course_title").
		ColumnExpr("COALESCE(ag.access_status, 'pending') AS access_status").
		ColumnExpr("COALESCE(ag.payment_status, 'unknown') AS payment_status").
		Join("JOIN classroom.classes AS cl ON cl.id = e.class_id").
		Join("JOIN catalog.courses AS c ON c.id = cl.course_id").
		Join("JOIN organization.organizations AS o ON o.id = cl.organization_id").
		Join("LEFT JOIN classroom.access_grants AS ag ON ag.class_id = e.class_id AND ag.student_user_id = e.student_user_id").
		Where("e.student_user_id = ?", userID).
		Where("e.status = 'active'").
		Where("cl.status = 'active'").
		Order("e.created_at DESC").
		Scan(ctx, &items)
	return items, err
}

func (uc *usecase) listLessons(ctx context.Context, userID string, classCtx shared.ClassContext) ([]shared.LessonRow, error) {
	items := make([]shared.LessonRow, 0)
	err := uc.db.NewSelect().TableExpr("catalog.lessons AS l").
		ColumnExpr("l.id, l.organization_id, l.course_id, l.title, l.description, l.order_number, l.publish_day, l.status").
		ColumnExpr("COUNT(DISTINCT lv.id) > 0 AS has_video").
		ColumnExpr("COUNT(DISTINCT lm.id) AS material_count").
		ColumnExpr("lc.completed_at AS completed_at").
		Join("LEFT JOIN catalog.lesson_videos AS lv ON lv.lesson_id = l.id").
		Join("LEFT JOIN catalog.lesson_materials AS lm ON lm.lesson_id = l.id").
		Join("LEFT JOIN learning.lesson_completions AS lc ON lc.lesson_id = l.id AND lc.class_id = ? AND lc.student_user_id = ?", classCtx.ClassID, userID).
		Where("l.course_id = ?", classCtx.CourseID).
		Where("l.status = 'published'").
		Group("l.id", "lc.completed_at").
		Order("l.order_number ASC").
		Scan(ctx, &items)
	return items, err
}

func buildLessons(classCtx shared.ClassContext, lessons []shared.LessonRow) ([]LessonDTO, ProgressDTO) {
	items := make([]LessonDTO, 0, len(lessons))
	completed := 0
	for _, row := range lessons {
		availability := shared.ComputeAvailability(classCtx.StartDate, classCtx.LessonCadence, row.PublishDay, row.Status, time.Now())
		isCompleted := row.CompletedAt != nil
		if isCompleted {
			completed++
		}
		items = append(items, LessonDTO{LessonID: row.ID, Title: row.Title, Description: row.Description, Status: availability.Status, AvailableAt: availability.AvailableAt, OrderNumber: row.OrderNumber, PublishDay: row.PublishDay, HasVideo: row.HasVideo, HasMaterials: row.MaterialCount > 0, MaterialCount: row.MaterialCount, Completed: isCompleted})
	}
	percentage := 0
	if len(lessons) > 0 {
		percentage = int(math.Round(float64(completed) / float64(len(lessons)) * 100))
	}
	return items, ProgressDTO{CompletedLessons: completed, TotalLessons: len(lessons), Percentage: percentage}
}

func toClassDTO(item shared.ClassContext) ClassDTO {
	return ClassDTO{ID: item.ClassID, Name: item.ClassName, CourseID: item.CourseID, CourseTitle: item.CourseTitle, AccessStatus: item.AccessStatus, PaymentStatus: item.PaymentStatus}
}

func IsNoRows(err error) bool { return errors.Is(err, sql.ErrNoRows) }
