package usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type GetLessonHomeworkRequest struct {
	LessonID string `query:"lesson_id" validate:"required,uuid"`
}

type GetLessonHomeworkResponse struct {
	Item          *DefinitionDTO    `json:"item"`
	QuizQuestions []QuizQuestionDTO `json:"quiz_questions"`
}

type GetLessonHomeworkUseCase = ucdef.UserAction[*GetLessonHomeworkRequest, *GetLessonHomeworkResponse]

func NewGetLessonHomework(db *bun.DB) GetLessonHomeworkUseCase { return &getLessonHomework{db: db} }

type getLessonHomework struct{ db *bun.DB }

func (uc *getLessonHomework) OperationID() string { return "get-lesson-homework" }

func (uc *getLessonHomework) Execute(ctx context.Context, in *GetLessonHomeworkRequest) (*GetLessonHomeworkResponse, error) {
	lesson, err := loadLessonContext(ctx, uc.db, in.LessonID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if err := ensureTeacherAccess(ctx, uc.db, lesson.OrganizationID); err != nil {
		return nil, errx.Wrap(err)
	}
	item := new(DefinitionDTO)
	err = uc.db.NewSelect().TableExpr("homework.definitions AS hd").
		ColumnExpr("hd.id, hd.organization_id, hd.course_id, hd.lesson_id, hd.title, hd.instructions, hd.submission_type, hd.status, hd.max_score, hd.due_days_after_publish, hd.allow_resubmission, hd.created_at, hd.updated_at").
		Where("hd.lesson_id = ?", in.LessonID).
		Limit(1).
		Scan(ctx, item)
	if errors.Is(err, sql.ErrNoRows) {
		return &GetLessonHomeworkResponse{Item: nil, QuizQuestions: []QuizQuestionDTO{}}, nil
	}
	if err != nil {
		return nil, errx.Wrap(err)
	}
	quiz, err := loadQuiz(ctx, uc.db, item.ID, true)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &GetLessonHomeworkResponse{Item: item, QuizQuestions: quiz}, nil
}
