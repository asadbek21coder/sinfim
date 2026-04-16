package listlessons

import (
	"context"

	"go-enterprise-blueprint/internal/modules/catalog/domain"
	"go-enterprise-blueprint/internal/modules/catalog/domain/course"
	"go-enterprise-blueprint/internal/modules/catalog/domain/lesson"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/shared"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	CourseID string  `query:"course_id" validate:"required,uuid"`
	Status   *string `query:"status" validate:"omitempty,oneof=draft published archived"`
	Limit    int     `query:"limit" validate:"omitempty,min=1,max=1000"`
}

type Response struct {
	Items []lesson.Summary `json:"items"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(dc *domain.Container) UseCase { return &usecase{dc: dc} }

type usecase struct{ dc *domain.Container }

func (uc *usecase) OperationID() string { return "list-lessons" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	courseItem, err := uc.dc.CourseRepo().Get(ctx, course.Filter{ID: &in.CourseID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if accessErr := shared.EnsureCourseReadAccess(ctx, uc.dc, courseItem.OrganizationID); accessErr != nil {
		return nil, errx.Wrap(accessErr)
	}
	limit := in.Limit
	if limit == 0 {
		limit = 200
	}
	items, listErr := uc.dc.LessonRepo().List(ctx, lesson.Filter{CourseID: &in.CourseID, Status: in.Status, Limit: limit})
	if listErr != nil {
		return nil, errx.Wrap(listErr)
	}
	return &Response{Items: items}, nil
}
