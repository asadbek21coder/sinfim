package getclassdetail

import (
	"context"

	"go-enterprise-blueprint/internal/modules/classroom/domain"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classgroup"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classmentor"
	"go-enterprise-blueprint/internal/modules/classroom/domain/enrollment"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/shared"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	ID string `query:"id" validate:"required,uuid"`
}

type Response struct {
	Item     classgroup.Class        `json:"item"`
	Mentors  []classmentor.MentorDTO `json:"mentors"`
	Students []enrollment.StudentDTO `json:"students"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(dc *domain.Container) UseCase { return &usecase{dc: dc} }

type usecase struct{ dc *domain.Container }

func (uc *usecase) OperationID() string { return "get-class-detail" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	item, err := uc.dc.ClassRepo().Get(ctx, classgroup.Filter{ID: &in.ID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if accessErr := shared.EnsureClassOperate(ctx, uc.dc, item); accessErr != nil {
		return nil, errx.Wrap(accessErr)
	}
	mentors, err := uc.dc.MentorRepo().ListByClass(ctx, item.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	students, err := uc.dc.EnrollmentRepo().ListStudents(ctx, item.ID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{Item: *item, Mentors: mentors, Students: students}, nil
}
