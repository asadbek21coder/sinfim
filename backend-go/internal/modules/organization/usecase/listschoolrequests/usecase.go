package listschoolrequests

import (
	"context"

	"go-enterprise-blueprint/internal/modules/organization/domain"
	"go-enterprise-blueprint/internal/modules/organization/domain/schoolrequest"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	Status *string `query:"status" validate:"omitempty,oneof=new contacted approved rejected"`
	Limit  int     `query:"limit" validate:"omitempty,min=1,max=200"`
}

type Response struct {
	Items []schoolrequest.SchoolRequest `json:"items"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "list-school-requests" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	limit := in.Limit
	if limit == 0 {
		limit = 100
	}
	items, err := uc.domainContainer.SchoolRequestRepo().List(ctx, schoolrequest.Filter{Status: in.Status, Limit: limit})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{Items: items}, nil
}
