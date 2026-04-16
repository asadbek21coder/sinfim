package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type SubmitHomeworkRequest struct {
	DefinitionID string            `json:"definition_id" validate:"required,uuid"`
	ClassID      string            `json:"class_id" validate:"required,uuid"`
	TextAnswer   *string           `json:"text_answer" validate:"omitempty,max=10000"`
	FileURL      *string           `json:"file_url" validate:"omitempty,max=2000"`
	AudioURL     *string           `json:"audio_url" validate:"omitempty,max=2000"`
	QuizAnswers  []QuizAnswerInput `json:"quiz_answers" validate:"omitempty,dive"`
}

type SubmitHomeworkResponse struct {
	Submission SubmissionDTO `json:"submission"`
	QuizScore  *int          `json:"quiz_score"`
}

type SubmitHomeworkUseCase = ucdef.UserAction[*SubmitHomeworkRequest, *SubmitHomeworkResponse]

func NewSubmitHomework(db *bun.DB) SubmitHomeworkUseCase { return &submitHomework{db: db} }

type submitHomework struct{ db *bun.DB }

type submissionRow struct {
	bun.BaseModel  `bun:"table:homework.submissions"`
	ID             string     `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrganizationID string     `bun:"organization_id,type:uuid,notnull"`
	DefinitionID   string     `bun:"definition_id,type:uuid,notnull"`
	LessonID       string     `bun:"lesson_id,type:uuid,notnull"`
	ClassID        string     `bun:"class_id,type:uuid,notnull"`
	StudentUserID  string     `bun:"student_user_id,notnull"`
	SubmissionType string     `bun:"submission_type,notnull"`
	Status         string     `bun:"status,notnull"`
	AttemptNumber  int        `bun:"attempt_number,notnull"`
	TextAnswer     *string    `bun:"text_answer"`
	FileURL        *string    `bun:"file_url"`
	AudioURL       *string    `bun:"audio_url"`
	Score          *int       `bun:"score"`
	MaxScore       int        `bun:"max_score,notnull"`
	AutoScored     bool       `bun:"auto_scored,notnull"`
	SubmittedAt    time.Time  `bun:"submitted_at,notnull"`
	ReviewedAt     *time.Time `bun:"reviewed_at"`
	ReviewerUserID *string    `bun:"reviewer_user_id"`
	Feedback       *string    `bun:"feedback"`
	CreatedAt      time.Time  `bun:"created_at,notnull"`
	UpdatedAt      time.Time  `bun:"updated_at,notnull"`
}

type answerRow struct {
	bun.BaseModel    `bun:"table:homework.quiz_answers"`
	ID               string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	SubmissionID     string    `bun:"submission_id,type:uuid,notnull"`
	QuestionID       string    `bun:"question_id,type:uuid,notnull"`
	SelectedOptionID *string   `bun:"selected_option_id,type:uuid"`
	IsCorrect        bool      `bun:"is_correct,notnull"`
	CreatedAt        time.Time `bun:"created_at,notnull"`
	UpdatedAt        time.Time `bun:"updated_at,notnull"`
}

func (uc *submitHomework) OperationID() string { return "submit-homework" }

func (uc *submitHomework) Execute(ctx context.Context, in *SubmitHomeworkRequest) (*SubmitHomeworkResponse, error) {
	userID := auth.MustUserContext(ctx).UserID
	definition, err := getDefinition(ctx, uc.db, in.DefinitionID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if definition.Status != "published" {
		return nil, errx.New("homework definition is not published", errx.WithType(errx.T_Forbidden), errx.WithCode(CodeAccessDenied))
	}
	if _, err := ensureStudentLessonAccess(ctx, uc.db, in.ClassID, definition.LessonID); err != nil {
		return nil, errx.Wrap(err)
	}
	if !definition.AllowResubmission {
		exists, err := uc.hasSubmission(ctx, in.DefinitionID, in.ClassID, userID)
		if err != nil {
			return nil, errx.Wrap(err)
		}
		if exists {
			return nil, errx.New("homework already submitted", errx.WithType(errx.T_Conflict), errx.WithCode(CodeAlreadySubmitted))
		}
	}
	now := time.Now()
	var quizScore *int
	row := &submissionRow{OrganizationID: definition.OrganizationID, DefinitionID: definition.ID, LessonID: definition.LessonID, ClassID: in.ClassID, StudentUserID: userID, SubmissionType: definition.SubmissionType, Status: "submitted", AttemptNumber: 1, TextAnswer: trimPtr(in.TextAnswer), FileURL: trimPtr(in.FileURL), AudioURL: trimPtr(in.AudioURL), MaxScore: definition.MaxScore, SubmittedAt: now, CreatedAt: now, UpdatedAt: now}
	err = uc.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		attempt, err := nextAttempt(ctx, tx, definition.ID, in.ClassID, userID)
		if err != nil {
			return err
		}
		row.AttemptNumber = attempt
		if definition.SubmissionType == "quiz" {
			score, err := scoreQuiz(ctx, tx, definition.ID, in.QuizAnswers)
			if err != nil {
				return err
			}
			row.Score = &score
			row.AutoScored = true
			row.Status = "reviewed"
			quizScore = &score
		}
		if err := tx.NewInsert().Model(row).
			On("CONFLICT (class_id, student_user_id, definition_id) DO UPDATE").
			Set("attempt_number = EXCLUDED.attempt_number").
			Set("text_answer = EXCLUDED.text_answer").
			Set("file_url = EXCLUDED.file_url").
			Set("audio_url = EXCLUDED.audio_url").
			Set("score = EXCLUDED.score").
			Set("auto_scored = EXCLUDED.auto_scored").
			Set("status = EXCLUDED.status").
			Set("submitted_at = EXCLUDED.submitted_at").
			Set("updated_at = EXCLUDED.updated_at").
			Returning("*").Scan(ctx); err != nil {
			return err
		}
		if definition.SubmissionType != "quiz" {
			return nil
		}
		if _, err := tx.NewDelete().TableExpr("homework.quiz_answers").Where("submission_id = ?", row.ID).Exec(ctx); err != nil {
			return err
		}
		return insertAnswers(ctx, tx, row.ID, in.QuizAnswers)
	})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	submission, err := uc.loadSubmission(ctx, row.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &SubmitHomeworkResponse{Submission: *submission, QuizScore: quizScore}, nil
}

func (uc *submitHomework) hasSubmission(ctx context.Context, definitionID string, classID string, userID string) (bool, error) {
	var exists bool
	err := uc.db.NewSelect().TableExpr("homework.submissions AS hs").ColumnExpr("COUNT(*) > 0").Where("hs.definition_id = ?", definitionID).Where("hs.class_id = ?", classID).Where("hs.student_user_id = ?", userID).Scan(ctx, &exists)
	return exists, err
}

func (uc *submitHomework) loadSubmission(ctx context.Context, submissionID string) (*SubmissionDTO, error) {
	item := new(SubmissionDTO)
	err := uc.db.NewSelect().TableExpr("homework.submissions AS hs").
		ColumnExpr("hs.id, hs.organization_id, hs.definition_id, hs.lesson_id, hs.class_id, hs.student_user_id, hs.submission_type, hs.status, hs.attempt_number, hs.text_answer, hs.file_url, hs.audio_url, hs.score, hs.max_score, hs.auto_scored, hs.submitted_at, hs.reviewed_at, hs.reviewer_user_id, hs.feedback").
		Where("hs.id = ?", submissionID).
		Limit(1).
		Scan(ctx, item)
	return item, err
}

func nextAttempt(ctx context.Context, tx bun.Tx, definitionID string, classID string, userID string) (int, error) {
	var current int
	err := tx.NewSelect().TableExpr("homework.submissions AS hs").ColumnExpr("COALESCE(MAX(hs.attempt_number), 0)").Where("hs.definition_id = ?", definitionID).Where("hs.class_id = ?", classID).Where("hs.student_user_id = ?", userID).Scan(ctx, &current)
	return current + 1, err
}

func scoreQuiz(ctx context.Context, tx bun.Tx, definitionID string, answers []QuizAnswerInput) (int, error) {
	score := 0
	for _, answer := range answers {
		if answer.SelectedOptionID == nil {
			continue
		}
		var points int
		err := tx.NewSelect().TableExpr("homework.quiz_options AS qo").
			ColumnExpr("qq.points").
			Join("JOIN homework.quiz_questions AS qq ON qq.id = qo.question_id").
			Where("qq.definition_id = ?", definitionID).
			Where("qq.id = ?", answer.QuestionID).
			Where("qo.id = ?", *answer.SelectedOptionID).
			Where("qo.is_correct = true").
			Limit(1).
			Scan(ctx, &points)
		if errors.Is(err, sql.ErrNoRows) {
			continue
		}
		if err != nil {
			return 0, err
		}
		score += points
	}
	return score, nil
}

func insertAnswers(ctx context.Context, tx bun.Tx, submissionID string, answers []QuizAnswerInput) error {
	now := time.Now()
	for _, answer := range answers {
		correct := false
		if answer.SelectedOptionID != nil {
			err := tx.NewSelect().TableExpr("homework.quiz_options AS qo").ColumnExpr("qo.is_correct").Where("qo.id = ?", *answer.SelectedOptionID).Where("qo.question_id = ?", answer.QuestionID).Limit(1).Scan(ctx, &correct)
			if errors.Is(err, sql.ErrNoRows) {
				return errx.New("selected quiz option does not belong to question", errx.WithType(errx.T_Validation), errx.WithCode(CodeInvalidQuiz))
			}
			if err != nil {
				return err
			}
		}
		row := &answerRow{SubmissionID: submissionID, QuestionID: answer.QuestionID, SelectedOptionID: answer.SelectedOptionID, IsCorrect: correct, CreatedAt: now, UpdatedAt: now}
		if _, err := tx.NewInsert().Model(row).Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
