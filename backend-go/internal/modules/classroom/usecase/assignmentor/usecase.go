package assignmentor

import (
	"context"

	"go-enterprise-blueprint/internal/modules/classroom/domain"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classgroup"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classmentor"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/shared"
	"go-enterprise-blueprint/internal/modules/organization/domain/membership"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	ClassID      string `json:"class_id" validate:"required,uuid"`
	MentorUserID string `json:"mentor_user_id" validate:"required"`
}

type Response struct {
	Item classmentor.ClassMentor `json:"item"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(dc *domain.Container) UseCase { return &usecase{dc: dc} }

type usecase struct{ dc *domain.Container }

func (uc *usecase) OperationID() string { return "assign-mentor" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	classItem, err := uc.dc.ClassRepo().Get(ctx, classgroup.Filter{ID: &in.ClassID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if err := shared.EnsureClassWrite(ctx, uc.dc, classItem.OrganizationID); err != nil {
		return nil, errx.Wrap(err)
	}
	ok, err := uc.dc.MembershipRepo().Exists(ctx, in.MentorUserID, classItem.OrganizationID, membership.RoleMentor)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if !ok {
		return nil, errx.New("mentor role required", errx.WithType(errx.T_Validation), errx.WithCode("MENTOR_ROLE_REQUIRED"))
	}
	item, err := uc.dc.MentorRepo().Assign(ctx, &classmentor.ClassMentor{OrganizationID: classItem.OrganizationID, ClassID: classItem.ID, MentorUserID: in.MentorUserID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{Item: *item}, nil
}
