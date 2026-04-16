package usecase

import (
	"context"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type ListReviewSubmissionsRequest struct {
	OrganizationID *string `query:"organization_id" validate:"omitempty,uuid"`
	ClassID        *string `query:"class_id" validate:"omitempty,uuid"`
	Status         *string `query:"status" validate:"omitempty,oneof=submitted reviewed needs_revision all"`
	Limit          int     `query:"limit" validate:"omitempty,min=1,max=200"`
}

type ListReviewSubmissionsResponse struct {
	Items []ReviewSubmissionSummaryDTO `json:"items"`
}

type ListReviewSubmissionsUseCase = ucdef.UserAction[*ListReviewSubmissionsRequest, *ListReviewSubmissionsResponse]

func NewListReviewSubmissions(db *bun.DB) ListReviewSubmissionsUseCase {
	return &listReviewSubmissions{db: db}
}

type listReviewSubmissions struct{ db *bun.DB }

func (uc *listReviewSubmissions) OperationID() string { return "list-homework-review-submissions" }

func (uc *listReviewSubmissions) Execute(ctx context.Context, in *ListReviewSubmissionsRequest) (*ListReviewSubmissionsResponse, error) {
	limit := in.Limit
	if limit == 0 {
		limit = 100
	}
	items := make([]ReviewSubmissionSummaryDTO, 0)
	query := uc.db.NewSelect().TableExpr("homework.submissions AS hs").
		ColumnExpr("hs.id, hs.organization_id, hs.definition_id, hs.lesson_id, hs.class_id, hs.student_user_id, hs.submission_type, hs.status, hs.attempt_number, hs.text_answer, hs.file_url, hs.audio_url, hs.score, hs.max_score, hs.auto_scored, hs.submitted_at, hs.reviewed_at, hs.reviewer_user_id, hs.feedback").
		ColumnExpr("o.name AS organization_name, c.title AS course_title, l.title AS lesson_title, cl.name AS class_name, hd.title AS homework_title, u.full_name AS student_full_name, u.phone_number AS student_phone").
		Join("JOIN homework.definitions AS hd ON hd.id = hs.definition_id").
		Join("JOIN catalog.lessons AS l ON l.id = hs.lesson_id").
		Join("JOIN catalog.courses AS c ON c.id = hd.course_id").
		Join("JOIN classroom.classes AS cl ON cl.id = hs.class_id").
		Join("JOIN organization.organizations AS o ON o.id = hs.organization_id").
		Join("JOIN auth.users AS u ON u.id = hs.student_user_id").
		Order("hs.submitted_at DESC").
		Limit(limit)
	query = applyReviewAccessFilter(ctx, query, uc.db)
	if in.OrganizationID != nil {
		query = query.Where("hs.organization_id = ?", *in.OrganizationID)
	}
	if in.ClassID != nil {
		query = query.Where("hs.class_id = ?", *in.ClassID)
	}
	if in.Status == nil || *in.Status == "" {
		query = query.Where("hs.status = 'submitted'")
	} else if *in.Status != "all" {
		query = query.Where("hs.status = ?", *in.Status)
	}
	if err := query.Scan(ctx, &items); err != nil {
		return nil, errx.Wrap(err)
	}
	return &ListReviewSubmissionsResponse{Items: items}, nil
}
