package updateschoolrequeststatus

import (
	"context"

	"go-enterprise-blueprint/internal/modules/organization/domain"
	"go-enterprise-blueprint/internal/modules/organization/domain/schoolrequest"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	ID     string `json:"id" validate:"required,uuid"`
	Status string `json:"status" validate:"required,oneof=new contacted approved rejected"`
}

type Response struct {
	Item schoolrequest.SchoolRequest `json:"item"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "update-school-request-status" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	updated, err := uc.domainContainer.SchoolRequestRepo().UpdateStatus(ctx, in.ID, in.Status)
	if err != nil {
		return nil, errx.WrapWithTypeOnCodes(err, errx.T_NotFound, schoolrequest.CodeSchoolRequestNotFound)
	}
	return &Response{Item: *updated}, nil
}
