package usecase

import (
	"context"
	"time"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type SaveDefinitionRequest struct {
	LessonID            string              `json:"lesson_id" validate:"required,uuid"`
	Title               string              `json:"title" validate:"required,min=2,max=180"`
	Instructions        *string             `json:"instructions" validate:"omitempty,max=4000"`
	SubmissionType      string              `json:"submission_type" validate:"required,oneof=text file audio quiz"`
	Status              string              `json:"status" validate:"omitempty,oneof=draft published archived"`
	MaxScore            int                 `json:"max_score" validate:"omitempty,min=0,max=10000"`
	DueDaysAfterPublish *int                `json:"due_days_after_publish" validate:"omitempty,min=0,max=1000"`
	AllowResubmission   bool                `json:"allow_resubmission"`
	QuizQuestions       []QuizQuestionInput `json:"quiz_questions" validate:"omitempty,dive"`
}

type DefinitionResponse struct {
	Item          DefinitionDTO     `json:"item"`
	QuizQuestions []QuizQuestionDTO `json:"quiz_questions"`
}

type SaveDefinitionUseCase = ucdef.UserAction[*SaveDefinitionRequest, *DefinitionResponse]

func NewSaveDefinition(db *bun.DB) SaveDefinitionUseCase { return &saveDefinition{db: db} }

type saveDefinition struct{ db *bun.DB }

type definitionRow struct {
	bun.BaseModel       `bun:"table:homework.definitions"`
	ID                  string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrganizationID      string    `bun:"organization_id,type:uuid,notnull"`
	CourseID            string    `bun:"course_id,type:uuid,notnull"`
	LessonID            string    `bun:"lesson_id,type:uuid,notnull"`
	Title               string    `bun:"title,notnull"`
	Instructions        *string   `bun:"instructions"`
	SubmissionType      string    `bun:"submission_type,notnull"`
	Status              string    `bun:"status,notnull"`
	MaxScore            int       `bun:"max_score,notnull"`
	DueDaysAfterPublish *int      `bun:"due_days_after_publish"`
	AllowResubmission   bool      `bun:"allow_resubmission,notnull"`
	CreatedAt           time.Time `bun:"created_at,notnull"`
	UpdatedAt           time.Time `bun:"updated_at,notnull"`
}

type questionRow struct {
	bun.BaseModel `bun:"table:homework.quiz_questions"`
	ID            string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	DefinitionID  string    `bun:"definition_id,type:uuid,notnull"`
	Prompt        string    `bun:"prompt,notnull"`
	OrderNumber   int       `bun:"order_number,notnull"`
	Points        int       `bun:"points,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull"`
	UpdatedAt     time.Time `bun:"updated_at,notnull"`
}

type optionRow struct {
	bun.BaseModel `bun:"table:homework.quiz_options"`
	ID            string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	QuestionID    string    `bun:"question_id,type:uuid,notnull"`
	Label         string    `bun:"label,notnull"`
	IsCorrect     bool      `bun:"is_correct,notnull"`
	OrderNumber   int       `bun:"order_number,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull"`
	UpdatedAt     time.Time `bun:"updated_at,notnull"`
}

func (uc *saveDefinition) OperationID() string { return "save-homework-definition" }

func (uc *saveDefinition) Execute(ctx context.Context, in *SaveDefinitionRequest) (*DefinitionResponse, error) {
	lesson, err := loadLessonContext(ctx, uc.db, in.LessonID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if err := ensureTeacherAccess(ctx, uc.db, lesson.OrganizationID); err != nil {
		return nil, errx.Wrap(err)
	}
	status := in.Status
	if status == "" {
		status = "draft"
	}
	maxScore := in.MaxScore
	if maxScore == 0 {
		maxScore = 100
	}
	now := time.Now()
	row := &definitionRow{OrganizationID: lesson.OrganizationID, CourseID: lesson.CourseID, LessonID: lesson.LessonID, Title: in.Title, Instructions: trimPtr(in.Instructions), SubmissionType: in.SubmissionType, Status: status, MaxScore: maxScore, DueDaysAfterPublish: in.DueDaysAfterPublish, AllowResubmission: in.AllowResubmission, CreatedAt: now, UpdatedAt: now}
	err = uc.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if err := tx.NewInsert().Model(row).
			On("CONFLICT (lesson_id) DO UPDATE").
			Set("title = EXCLUDED.title").
			Set("instructions = EXCLUDED.instructions").
			Set("submission_type = EXCLUDED.submission_type").
			Set("status = EXCLUDED.status").
			Set("max_score = EXCLUDED.max_score").
			Set("due_days_after_publish = EXCLUDED.due_days_after_publish").
			Set("allow_resubmission = EXCLUDED.allow_resubmission").
			Set("updated_at = EXCLUDED.updated_at").
			Returning("*").Scan(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().TableExpr("homework.quiz_questions").Where("definition_id = ?", row.ID).Exec(ctx); err != nil {
			return err
		}
		if row.SubmissionType != "quiz" {
			return nil
		}
		for questionIndex, question := range in.QuizQuestions {
			order := question.OrderNumber
			if order == 0 {
				order = questionIndex + 1
			}
			points := question.Points
			if points == 0 {
				points = 1
			}
			q := &questionRow{DefinitionID: row.ID, Prompt: question.Prompt, OrderNumber: order, Points: points, CreatedAt: now, UpdatedAt: now}
			if err := tx.NewInsert().Model(q).Returning("*").Scan(ctx); err != nil {
				return err
			}
			for optionIndex, option := range question.Options {
				optionOrder := option.OrderNumber
				if optionOrder == 0 {
					optionOrder = optionIndex + 1
				}
				o := &optionRow{QuestionID: q.ID, Label: option.Label, IsCorrect: option.IsCorrect, OrderNumber: optionOrder, CreatedAt: now, UpdatedAt: now}
				if _, err := tx.NewInsert().Model(o).Exec(ctx); err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	item, err := getDefinition(ctx, uc.db, row.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	quiz, err := loadQuiz(ctx, uc.db, row.ID, true)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &DefinitionResponse{Item: *item, QuizQuestions: quiz}, nil
}
