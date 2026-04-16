package getlessondetail

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/modules/learning/usecase/shared"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type Request struct {
	ClassID  string `query:"class_id" validate:"required,uuid"`
	LessonID string `query:"lesson_id" validate:"required,uuid"`
}

type LessonDTO struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      string     `json:"status"`
	AvailableAt *time.Time `json:"available_at"`
	OrderNumber int        `json:"order_number"`
	PublishDay  int        `json:"publish_day"`
	Completed   bool       `json:"completed"`
}

type Response struct {
	Lesson    LessonDTO            `json:"lesson"`
	Video     *shared.VideoDTO     `json:"video"`
	Materials []shared.MaterialDTO `json:"materials"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(db *bun.DB) UseCase { return &usecase{db: db} }

type usecase struct{ db *bun.DB }

func (uc *usecase) OperationID() string { return "get-learning-lesson-detail" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	userID := auth.MustUserContext(ctx).UserID
	classCtx, lessonRow, err := uc.loadLessonContext(ctx, userID, in.ClassID, in.LessonID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if classCtx.AccessStatus != "active" {
		return nil, errx.New("student class access is not active", errx.WithType(errx.T_Forbidden), errx.WithCode(shared.CodeAccessDenied))
	}
	availability := shared.ComputeAvailability(classCtx.StartDate, classCtx.LessonCadence, lessonRow.PublishDay, lessonRow.Status, time.Now())
	lessonDTO := LessonDTO{ID: lessonRow.ID, Title: lessonRow.Title, Description: lessonRow.Description, Status: availability.Status, AvailableAt: availability.AvailableAt, OrderNumber: lessonRow.OrderNumber, PublishDay: lessonRow.PublishDay, Completed: lessonRow.CompletedAt != nil}
	if availability.Status == "locked" {
		return &Response{Lesson: lessonDTO, Video: nil, Materials: []shared.MaterialDTO{}}, nil
	}
	video, err := uc.getVideo(ctx, lessonRow.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	materials, err := uc.listMaterials(ctx, lessonRow.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{Lesson: lessonDTO, Video: video, Materials: materials}, nil
}

func (uc *usecase) loadLessonContext(ctx context.Context, userID string, classID string, lessonID string) (shared.ClassContext, shared.LessonRow, error) {
	var row struct {
		shared.ClassContext
		LessonID          string     `bun:"lesson_id"`
		LessonOrgID       string     `bun:"lesson_org_id"`
		LessonCourseID    string     `bun:"lesson_course_id"`
		LessonTitle       string     `bun:"lesson_title"`
		LessonDescription *string    `bun:"lesson_description"`
		OrderNumber       int        `bun:"order_number"`
		PublishDay        int        `bun:"publish_day"`
		LessonStatus      string     `bun:"lesson_status"`
		CompletedAt       *time.Time `bun:"completed_at"`
	}
	err := uc.db.NewSelect().TableExpr("classroom.enrollments AS e").
		ColumnExpr("o.id AS organization_id, o.name AS organization, o.slug AS organization_slug, o.logo_url AS logo_url").
		ColumnExpr("cl.id AS class_id, cl.name AS class_name, cl.start_date AS start_date, cl.lesson_cadence AS lesson_cadence").
		ColumnExpr("c.id AS course_id, c.title AS course_title").
		ColumnExpr("COALESCE(ag.access_status, 'pending') AS access_status, COALESCE(ag.payment_status, 'unknown') AS payment_status").
		ColumnExpr("l.id AS lesson_id, l.organization_id AS lesson_org_id, l.course_id AS lesson_course_id, l.title AS lesson_title, l.description AS lesson_description, l.order_number AS order_number, l.publish_day AS publish_day, l.status AS lesson_status").
		ColumnExpr("lc.completed_at AS completed_at").
		Join("JOIN classroom.classes AS cl ON cl.id = e.class_id").
		Join("JOIN catalog.courses AS c ON c.id = cl.course_id").
		Join("JOIN organization.organizations AS o ON o.id = cl.organization_id").
		Join("JOIN catalog.lessons AS l ON l.course_id = c.id").
		Join("LEFT JOIN classroom.access_grants AS ag ON ag.class_id = e.class_id AND ag.student_user_id = e.student_user_id").
		Join("LEFT JOIN learning.lesson_completions AS lc ON lc.lesson_id = l.id AND lc.class_id = e.class_id AND lc.student_user_id = e.student_user_id").
		Where("e.student_user_id = ?", userID).
		Where("e.status = 'active'").
		Where("cl.id = ?", classID).
		Where("l.id = ?", lessonID).
		Limit(1).
		Scan(ctx, &row)
	if errors.Is(err, sql.ErrNoRows) {
		return shared.ClassContext{}, shared.LessonRow{}, errx.New("lesson not found for student class", errx.WithType(errx.T_NotFound), errx.WithCode(shared.CodeLessonNotFound))
	}
	if err != nil {
		return shared.ClassContext{}, shared.LessonRow{}, errx.Wrap(err)
	}
	lesson := shared.LessonRow{ID: row.LessonID, OrganizationID: row.LessonOrgID, CourseID: row.LessonCourseID, Title: row.LessonTitle, Description: row.LessonDescription, OrderNumber: row.OrderNumber, PublishDay: row.PublishDay, Status: row.LessonStatus, CompletedAt: row.CompletedAt}
	return row.ClassContext, lesson, nil
}

func (uc *usecase) getVideo(ctx context.Context, lessonID string) (*shared.VideoDTO, error) {
	item := new(shared.VideoDTO)
	err := uc.db.NewSelect().TableExpr("catalog.lesson_videos AS lv").
		ColumnExpr("lv.provider, lv.stream_ref, lv.telegram_channel_id, lv.telegram_message_id, lv.embed_url, lv.duration_seconds").
		Where("lv.lesson_id = ?", lessonID).
		Limit(1).
		Scan(ctx, item)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return item, err
}

func (uc *usecase) listMaterials(ctx context.Context, lessonID string) ([]shared.MaterialDTO, error) {
	items := make([]shared.MaterialDTO, 0)
	err := uc.db.NewSelect().TableExpr("catalog.lesson_materials AS lm").
		ColumnExpr("lm.id, lm.title, lm.material_type, lm.source_type, lm.url, lm.file_id, lm.order_number").
		Where("lm.lesson_id = ?", lessonID).
		Order("lm.order_number ASC").
		Scan(ctx, &items)
	return items, err
}
