package usecase

import (
	"context"
	"database/sql"
	"errors"

	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type GetStudentHomeworkRequest struct {
	ClassID  string `query:"class_id" validate:"required,uuid"`
	LessonID string `query:"lesson_id" validate:"required,uuid"`
}

type GetStudentHomeworkResponse struct {
	Item          *DefinitionDTO    `json:"item"`
	QuizQuestions []QuizQuestionDTO `json:"quiz_questions"`
	Submission    *SubmissionDTO    `json:"submission"`
}

type GetStudentHomeworkUseCase = ucdef.UserAction[*GetStudentHomeworkRequest, *GetStudentHomeworkResponse]

func NewGetStudentHomework(db *bun.DB) GetStudentHomeworkUseCase { return &getStudentHomework{db: db} }

type getStudentHomework struct{ db *bun.DB }

func (uc *getStudentHomework) OperationID() string { return "get-student-homework" }

func (uc *getStudentHomework) Execute(ctx context.Context, in *GetStudentHomeworkRequest) (*GetStudentHomeworkResponse, error) {
	if _, err := ensureStudentLessonAccess(ctx, uc.db, in.ClassID, in.LessonID); err != nil {
		return nil, errx.Wrap(err)
	}
	item := new(DefinitionDTO)
	err := uc.db.NewSelect().TableExpr("homework.definitions AS hd").
		ColumnExpr("hd.id, hd.organization_id, hd.course_id, hd.lesson_id, hd.title, hd.instructions, hd.submission_type, hd.status, hd.max_score, hd.due_days_after_publish, hd.allow_resubmission, hd.created_at, hd.updated_at").
		Where("hd.lesson_id = ?", in.LessonID).
		Where("hd.status = 'published'").
		Limit(1).
		Scan(ctx, item)
	if errors.Is(err, sql.ErrNoRows) {
		return &GetStudentHomeworkResponse{Item: nil, QuizQuestions: []QuizQuestionDTO{}, Submission: nil}, nil
	}
	if err != nil {
		return nil, errx.Wrap(err)
	}
	quiz, err := loadQuiz(ctx, uc.db, item.ID, false)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	submission, err := uc.loadSubmission(ctx, item.ID, in.ClassID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &GetStudentHomeworkResponse{Item: item, QuizQuestions: quiz, Submission: submission}, nil
}

func (uc *getStudentHomework) loadSubmission(ctx context.Context, definitionID string, classID string) (*SubmissionDTO, error) {
	userID := auth.MustUserContext(ctx).UserID
	item := new(SubmissionDTO)
	err := uc.db.NewSelect().TableExpr("homework.submissions AS hs").
		ColumnExpr("hs.id, hs.organization_id, hs.definition_id, hs.lesson_id, hs.class_id, hs.student_user_id, hs.submission_type, hs.status, hs.attempt_number, hs.text_answer, hs.file_url, hs.audio_url, hs.score, hs.max_score, hs.auto_scored, hs.submitted_at, hs.reviewed_at, hs.reviewer_user_id, hs.feedback").
		Where("hs.definition_id = ?", definitionID).
		Where("hs.class_id = ?", classID).
		Where("hs.student_user_id = ?", userID).
		Limit(1).
		Scan(ctx, item)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return item, err
}
