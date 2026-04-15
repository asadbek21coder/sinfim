package getcoursedetail

import (
	"context"

	"go-enterprise-blueprint/internal/modules/catalog/domain"
	"go-enterprise-blueprint/internal/modules/catalog/domain/course"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/shared"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	ID string `query:"id" validate:"required,uuid"`
}

type Response struct {
	Item course.Course `json:"item"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "get-course-detail" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	item, err := uc.domainContainer.CourseRepo().Get(ctx, course.Filter{ID: &in.ID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if accessErr := shared.EnsureCourseReadAccess(ctx, uc.domainContainer, item.OrganizationID); accessErr != nil {
		return nil, errx.Wrap(accessErr)
	}
	return &Response{Item: *item}, nil
}
