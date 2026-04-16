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

type ReviewSubmissionRequest struct {
	ID       string  `json:"id" validate:"required,uuid"`
	Status   string  `json:"status" validate:"required,oneof=reviewed needs_revision"`
	Score    *int    `json:"score" validate:"omitempty,min=0,max=10000"`
	Feedback *string `json:"feedback" validate:"omitempty,max=4000"`
}

type ReviewSubmissionResponse struct {
	Item SubmissionDTO `json:"item"`
}

type ReviewSubmissionUseCase = ucdef.UserAction[*ReviewSubmissionRequest, *ReviewSubmissionResponse]

func NewReviewSubmission(db *bun.DB) ReviewSubmissionUseCase { return &reviewSubmission{db: db} }

type reviewSubmission struct{ db *bun.DB }

func (uc *reviewSubmission) OperationID() string { return "review-homework-submission" }

func (uc *reviewSubmission) Execute(ctx context.Context, in *ReviewSubmissionRequest) (*ReviewSubmissionResponse, error) {
	current, err := uc.load(ctx, in.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if err := ensureReviewAccess(ctx, uc.db, current.OrganizationID, current.ClassID); err != nil {
		return nil, errx.Wrap(err)
	}
	reviewerID := auth.MustUserContext(ctx).UserID
	now := time.Now()
	item := new(SubmissionDTO)
	err = uc.db.NewUpdate().TableExpr("homework.submissions AS hs").
		Set("status = ?", in.Status).
		Set("score = ?", in.Score).
		Set("feedback = ?", trimPtr(in.Feedback)).
		Set("reviewed_at = ?", now).
		Set("reviewer_user_id = ?", reviewerID).
		Set("updated_at = ?", now).
		Where("hs.id = ?", in.ID).
		Returning("hs.id, hs.organization_id, hs.definition_id, hs.lesson_id, hs.class_id, hs.student_user_id, hs.submission_type, hs.status, hs.attempt_number, hs.text_answer, hs.file_url, hs.audio_url, hs.score, hs.max_score, hs.auto_scored, hs.submitted_at, hs.reviewed_at, hs.reviewer_user_id, hs.feedback").
		Scan(ctx, item)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &ReviewSubmissionResponse{Item: *item}, nil
}

func (uc *reviewSubmission) load(ctx context.Context, id string) (*SubmissionDTO, error) {
	item := new(SubmissionDTO)
	err := uc.db.NewSelect().TableExpr("homework.submissions AS hs").
		ColumnExpr("hs.id, hs.organization_id, hs.definition_id, hs.lesson_id, hs.class_id, hs.student_user_id, hs.submission_type, hs.status, hs.attempt_number, hs.text_answer, hs.file_url, hs.audio_url, hs.score, hs.max_score, hs.auto_scored, hs.submitted_at, hs.reviewed_at, hs.reviewer_user_id, hs.feedback").
		Where("hs.id = ?", id).
		Limit(1).
		Scan(ctx, item)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errx.New("submission not found", errx.WithType(errx.T_NotFound), errx.WithCode(CodeDefinitionNotFound))
	}
	return item, err
}
