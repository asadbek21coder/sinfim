package listclasses

import (
	"context"

	"go-enterprise-blueprint/internal/modules/classroom/domain"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classgroup"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/shared"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	OrganizationID string  `query:"organization_id" validate:"required,uuid"`
	CourseID       *string `query:"course_id" validate:"omitempty,uuid"`
	Limit          int     `query:"limit" validate:"omitempty,min=1,max=200"`
}

type Response struct {
	Items []classgroup.Summary `json:"items"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(dc *domain.Container) UseCase { return &usecase{dc: dc} }

type usecase struct{ dc *domain.Container }

func (uc *usecase) OperationID() string { return "list-classes" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	if err := shared.EnsureClassWrite(ctx, uc.dc, in.OrganizationID); err != nil {
		return nil, errx.Wrap(err)
	}
	limit := in.Limit
	if limit == 0 {
		limit = 100
	}
	items, err := uc.dc.ClassRepo().List(ctx, classgroup.Filter{OrganizationID: &in.OrganizationID, CourseID: in.CourseID, Limit: limit})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{Items: items}, nil
}
