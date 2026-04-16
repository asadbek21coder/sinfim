package usecase

import "time"

const (
	CodeDefinitionNotFound = "HOMEWORK_DEFINITION_NOT_FOUND"
	CodeLessonNotFound     = "LESSON_NOT_FOUND"
	CodeAccessDenied       = "ACCESS_DENIED"
	CodeAlreadySubmitted   = "HOMEWORK_ALREADY_SUBMITTED"
	CodeInvalidQuiz        = "INVALID_QUIZ_SUBMISSION"
)

type DefinitionDTO struct {
	ID                  string     `json:"id" bun:"id"`
	OrganizationID      string     `json:"organization_id" bun:"organization_id"`
	CourseID            string     `json:"course_id" bun:"course_id"`
	LessonID            string     `json:"lesson_id" bun:"lesson_id"`
	Title               string     `json:"title" bun:"title"`
	Instructions        *string    `json:"instructions" bun:"instructions"`
	SubmissionType      string     `json:"submission_type" bun:"submission_type"`
	Status              string     `json:"status" bun:"status"`
	MaxScore            int        `json:"max_score" bun:"max_score"`
	DueDaysAfterPublish *int       `json:"due_days_after_publish" bun:"due_days_after_publish"`
	AllowResubmission   bool       `json:"allow_resubmission" bun:"allow_resubmission"`
	CreatedAt           *time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at" bun:"updated_at"`
}

type QuizQuestionDTO struct {
	ID          string          `json:"id" bun:"id"`
	Prompt      string          `json:"prompt" bun:"prompt"`
	OrderNumber int             `json:"order_number" bun:"order_number"`
	Points      int             `json:"points" bun:"points"`
	Options     []QuizOptionDTO `json:"options" bun:"-"`
}

type QuizOptionDTO struct {
	ID          string `json:"id" bun:"id"`
	Label       string `json:"label" bun:"label"`
	IsCorrect   bool   `json:"is_correct" bun:"is_correct"`
	OrderNumber int    `json:"order_number" bun:"order_number"`
}

type SubmissionDTO struct {
	ID             string     `json:"id" bun:"id"`
	OrganizationID string     `json:"organization_id" bun:"organization_id"`
	DefinitionID   string     `json:"definition_id" bun:"definition_id"`
	LessonID       string     `json:"lesson_id" bun:"lesson_id"`
	ClassID        string     `json:"class_id" bun:"class_id"`
	StudentUserID  string     `json:"student_user_id" bun:"student_user_id"`
	SubmissionType string     `json:"submission_type" bun:"submission_type"`
	Status         string     `json:"status" bun:"status"`
	AttemptNumber  int        `json:"attempt_number" bun:"attempt_number"`
	TextAnswer     *string    `json:"text_answer" bun:"text_answer"`
	FileURL        *string    `json:"file_url" bun:"file_url"`
	AudioURL       *string    `json:"audio_url" bun:"audio_url"`
	Score          *int       `json:"score" bun:"score"`
	MaxScore       int        `json:"max_score" bun:"max_score"`
	AutoScored     bool       `json:"auto_scored" bun:"auto_scored"`
	SubmittedAt    *time.Time `json:"submitted_at" bun:"submitted_at"`
	ReviewedAt     *time.Time `json:"reviewed_at" bun:"reviewed_at"`
	ReviewerUserID *string    `json:"reviewer_user_id" bun:"reviewer_user_id"`
	Feedback       *string    `json:"feedback" bun:"feedback"`
}

type ReviewSubmissionSummaryDTO struct {
	SubmissionDTO
	OrganizationName string `json:"organization_name" bun:"organization_name"`
	CourseTitle      string `json:"course_title" bun:"course_title"`
	LessonTitle      string `json:"lesson_title" bun:"lesson_title"`
	ClassName        string `json:"class_name" bun:"class_name"`
	HomeworkTitle    string `json:"homework_title" bun:"homework_title"`
	StudentFullName  string `json:"student_full_name" bun:"student_full_name"`
	StudentPhone     string `json:"student_phone" bun:"student_phone"`
}

type QuizAnswerDTO struct {
	QuestionID       string  `json:"question_id" bun:"question_id"`
	QuestionPrompt   string  `json:"question_prompt" bun:"question_prompt"`
	SelectedOptionID *string `json:"selected_option_id" bun:"selected_option_id"`
	SelectedLabel    *string `json:"selected_label" bun:"selected_label"`
	IsCorrect        bool    `json:"is_correct" bun:"is_correct"`
	Points           int     `json:"points" bun:"points"`
}

type QuizQuestionInput struct {
	ID          *string           `json:"id" validate:"omitempty,uuid"`
	Prompt      string            `json:"prompt" validate:"required,min=2,max=1000"`
	OrderNumber int               `json:"order_number" validate:"omitempty,min=1,max=1000"`
	Points      int               `json:"points" validate:"omitempty,min=0,max=1000"`
	Options     []QuizOptionInput `json:"options" validate:"omitempty,dive"`
}

type QuizOptionInput struct {
	ID          *string `json:"id" validate:"omitempty,uuid"`
	Label       string  `json:"label" validate:"required,min=1,max=500"`
	IsCorrect   bool    `json:"is_correct"`
	OrderNumber int     `json:"order_number" validate:"omitempty,min=1,max=1000"`
}

type QuizAnswerInput struct {
	QuestionID       string  `json:"question_id" validate:"required,uuid"`
	SelectedOptionID *string `json:"selected_option_id" validate:"omitempty,uuid"`
}

type lessonContext struct {
	LessonID       string `bun:"lesson_id"`
	OrganizationID string `bun:"organization_id"`
	CourseID       string `bun:"course_id"`
}
