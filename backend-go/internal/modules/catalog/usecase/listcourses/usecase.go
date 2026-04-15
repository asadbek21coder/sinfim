package listcourses

import (
	"context"

	"go-enterprise-blueprint/internal/modules/catalog/domain"
	"go-enterprise-blueprint/internal/modules/catalog/domain/course"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/shared"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	OrganizationID string  `query:"organization_id" validate:"required,uuid"`
	PublicStatus   *string `query:"public_status" validate:"omitempty,oneof=draft public hidden"`
	Limit          int     `query:"limit" validate:"omitempty,min=1,max=200"`
}

type Response struct {
	Items []course.Course `json:"items"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "list-courses" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	if err := shared.EnsureCourseReadAccess(ctx, uc.domainContainer, in.OrganizationID); err != nil {
		return nil, errx.Wrap(err)
	}
	limit := in.Limit
	if limit == 0 {
		limit = 100
	}
	items, err := uc.domainContainer.CourseRepo().List(ctx, course.Filter{OrganizationID: &in.OrganizationID, PublicStatus: in.PublicStatus, Limit: limit})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{Items: items}, nil
}
