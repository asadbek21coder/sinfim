package updatecourse

import (
	"context"
	"strings"

	"go-enterprise-blueprint/internal/modules/catalog/domain"
	"go-enterprise-blueprint/internal/modules/catalog/domain/course"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/shared"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	ID           string  `json:"id" validate:"required,uuid"`
	Title        string  `json:"title" validate:"required,min=2,max=160"`
	Description  *string `json:"description" validate:"omitempty,max=2000"`
	Category     *string `json:"category" validate:"omitempty,max=80"`
	Level        *string `json:"level" validate:"omitempty,max=80"`
	Status       string  `json:"status" validate:"required,oneof=draft active archived"`
	PublicStatus string  `json:"public_status" validate:"required,oneof=draft public hidden"`
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

func (uc *usecase) OperationID() string { return "update-course" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	item, err := uc.domainContainer.CourseRepo().Get(ctx, course.Filter{ID: &in.ID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if accessErr := shared.EnsureCourseWriteAccess(ctx, uc.domainContainer, item.OrganizationID); accessErr != nil {
		return nil, errx.Wrap(accessErr)
	}
	item.Title = strings.TrimSpace(in.Title)
	item.Description = shared.TrimPtr(in.Description)
	item.Category = shared.TrimPtr(in.Category)
	item.Level = shared.TrimPtr(in.Level)
	item.Status = in.Status
	item.PublicStatus = in.PublicStatus
	updated, updateErr := uc.domainContainer.CourseRepo().Update(ctx, item)
	if updateErr != nil {
		return nil, errx.Wrap(updateErr)
	}
	return &Response{Item: *updated}, nil
}
