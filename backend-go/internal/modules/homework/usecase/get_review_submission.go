package usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type GetReviewSubmissionRequest struct {
	ID string `query:"id" validate:"required,uuid"`
}

type GetReviewSubmissionResponse struct {
	Item        ReviewSubmissionSummaryDTO `json:"item"`
	Definition  DefinitionDTO              `json:"definition"`
	QuizAnswers []QuizAnswerDTO            `json:"quiz_answers"`
}

type GetReviewSubmissionUseCase = ucdef.UserAction[*GetReviewSubmissionRequest, *GetReviewSubmissionResponse]

func NewGetReviewSubmission(db *bun.DB) GetReviewSubmissionUseCase {
	return &getReviewSubmission{db: db}
}

type getReviewSubmission struct{ db *bun.DB }

func (uc *getReviewSubmission) OperationID() string { return "get-homework-review-submission" }

func (uc *getReviewSubmission) Execute(ctx context.Context, in *GetReviewSubmissionRequest) (*GetReviewSubmissionResponse, error) {
	item, definition, err := uc.load(ctx, in.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if err := ensureReviewAccess(ctx, uc.db, item.OrganizationID, item.ClassID); err != nil {
		return nil, errx.Wrap(err)
	}
	answers, err := uc.loadQuizAnswers(ctx, item.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &GetReviewSubmissionResponse{Item: *item, Definition: *definition, QuizAnswers: answers}, nil
}

func (uc *getReviewSubmission) load(ctx context.Context, id string) (*ReviewSubmissionSummaryDTO, *DefinitionDTO, error) {
	var row struct {
		ReviewSubmissionSummaryDTO
		DefID                  string  `bun:"def_id"`
		DefOrganizationID      string  `bun:"def_organization_id"`
		DefCourseID            string  `bun:"def_course_id"`
		DefLessonID            string  `bun:"def_lesson_id"`
		DefTitle               string  `bun:"def_title"`
		DefInstructions        *string `bun:"def_instructions"`
		DefSubmissionType      string  `bun:"def_submission_type"`
		DefStatus              string  `bun:"def_status"`
		DefMaxScore            int     `bun:"def_max_score"`
		DefDueDaysAfterPublish *int    `bun:"def_due_days_after_publish"`
		DefAllowResubmission   bool    `bun:"def_allow_resubmission"`
	}
	err := uc.db.NewSelect().TableExpr("homework.submissions AS hs").
		ColumnExpr("hs.id, hs.organization_id, hs.definition_id, hs.lesson_id, hs.class_id, hs.student_user_id, hs.submission_type, hs.status, hs.attempt_number, hs.text_answer, hs.file_url, hs.audio_url, hs.score, hs.max_score, hs.auto_scored, hs.submitted_at, hs.reviewed_at, hs.reviewer_user_id, hs.feedback").
		ColumnExpr("o.name AS organization_name, c.title AS course_title, l.title AS lesson_title, cl.name AS class_name, hd.title AS homework_title, u.full_name AS student_full_name, u.phone_number AS student_phone").
		ColumnExpr("hd.id AS def_id, hd.organization_id AS def_organization_id, hd.course_id AS def_course_id, hd.lesson_id AS def_lesson_id, hd.title AS def_title, hd.instructions AS def_instructions, hd.submission_type AS def_submission_type, hd.status AS def_status, hd.max_score AS def_max_score, hd.due_days_after_publish AS def_due_days_after_publish, hd.allow_resubmission AS def_allow_resubmission").
		Join("JOIN homework.definitions AS hd ON hd.id = hs.definition_id").
		Join("JOIN catalog.lessons AS l ON l.id = hs.lesson_id").
		Join("JOIN catalog.courses AS c ON c.id = hd.course_id").
		Join("JOIN classroom.classes AS cl ON cl.id = hs.class_id").
		Join("JOIN organization.organizations AS o ON o.id = hs.organization_id").
		Join("JOIN auth.users AS u ON u.id = hs.student_user_id").
		Where("hs.id = ?", id).
		Limit(1).
		Scan(ctx, &row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil, errx.New("submission not found", errx.WithType(errx.T_NotFound), errx.WithCode(CodeDefinitionNotFound))
	}
	if err != nil {
		return nil, nil, err
	}
	definition := &DefinitionDTO{ID: row.DefID, OrganizationID: row.DefOrganizationID, CourseID: row.DefCourseID, LessonID: row.DefLessonID, Title: row.DefTitle, Instructions: row.DefInstructions, SubmissionType: row.DefSubmissionType, Status: row.DefStatus, MaxScore: row.DefMaxScore, DueDaysAfterPublish: row.DefDueDaysAfterPublish, AllowResubmission: row.DefAllowResubmission}
	return &row.ReviewSubmissionSummaryDTO, definition, nil
}

func (uc *getReviewSubmission) loadQuizAnswers(ctx context.Context, submissionID string) ([]QuizAnswerDTO, error) {
	items := make([]QuizAnswerDTO, 0)
	err := uc.db.NewSelect().TableExpr("homework.quiz_answers AS qa").
		ColumnExpr("qa.question_id, qq.prompt AS question_prompt, qa.selected_option_id, qo.label AS selected_label, qa.is_correct, qq.points").
		Join("JOIN homework.quiz_questions AS qq ON qq.id = qa.question_id").
		Join("LEFT JOIN homework.quiz_options AS qo ON qo.id = qa.selected_option_id").
		Where("qa.submission_id = ?", submissionID).
		Order("qq.order_number ASC").
		Scan(ctx, &items)
	return items, err
}
