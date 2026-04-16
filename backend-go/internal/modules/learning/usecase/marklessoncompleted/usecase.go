package marklessoncompleted

import (
	"context"
	"time"

	"go-enterprise-blueprint/internal/modules/learning/usecase/getlessondetail"
	"go-enterprise-blueprint/internal/modules/learning/usecase/shared"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type completion struct {
	bun.BaseModel  `bun:"table:learning.lesson_completions"`
	ID             string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrganizationID string    `bun:"organization_id,type:uuid,notnull"`
	ClassID        string    `bun:"class_id,type:uuid,notnull"`
	LessonID       string    `bun:"lesson_id,type:uuid,notnull"`
	StudentUserID  string    `bun:"student_user_id,notnull"`
	CompletedAt    time.Time `bun:"completed_at,notnull"`
	CreatedAt      time.Time `bun:"created_at,notnull"`
	UpdatedAt      time.Time `bun:"updated_at,notnull"`
}

type Request struct {
	ClassID  string `json:"class_id" validate:"required,uuid"`
	LessonID string `json:"lesson_id" validate:"required,uuid"`
}

type Response struct {
	LessonID      string    `json:"lesson_id"`
	ClassID       string    `json:"class_id"`
	StudentUserID string    `json:"student_user_id"`
	Completed     bool      `json:"completed"`
	CompletedAt   time.Time `json:"completed_at"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(db *bun.DB) UseCase { return &usecase{db: db, detail: getlessondetail.New(db)} }

type usecase struct {
	db     *bun.DB
	detail getlessondetail.UseCase
}

func (uc *usecase) OperationID() string { return "mark-lesson-completed" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	userID := auth.MustUserContext(ctx).UserID
	detail, err := uc.detail.Execute(ctx, &getlessondetail.Request{ClassID: in.ClassID, LessonID: in.LessonID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if detail.Lesson.Status == "locked" {
		return nil, errx.New("lesson is locked", errx.WithType(errx.T_Forbidden), errx.WithCode(shared.CodeLessonLocked))
	}
	organizationID, orgErr := uc.organizationIDForClass(ctx, in.ClassID)
	if orgErr != nil {
		return nil, errx.Wrap(orgErr)
	}
	now := time.Now()
	item := &completion{OrganizationID: organizationID, ClassID: in.ClassID, LessonID: in.LessonID, StudentUserID: userID, CompletedAt: now, CreatedAt: now, UpdatedAt: now}
	err = uc.db.NewInsert().Model(item).
		On("CONFLICT (class_id, lesson_id, student_user_id) DO UPDATE").
		Set("updated_at = EXCLUDED.updated_at").
		Returning("*").
		Scan(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{LessonID: in.LessonID, ClassID: in.ClassID, StudentUserID: userID, Completed: true, CompletedAt: item.CompletedAt}, nil
}

func (uc *usecase) organizationIDForClass(ctx context.Context, classID string) (string, error) {
	var organizationID string
	err := uc.db.NewSelect().TableExpr("classroom.classes AS cl").ColumnExpr("cl.organization_id").Where("cl.id = ?", classID).Limit(1).Scan(ctx, &organizationID)
	return organizationID, err
}
