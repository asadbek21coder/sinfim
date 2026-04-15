package updateaccess

import (
	"context"
	"time"

	"go-enterprise-blueprint/internal/modules/classroom/domain"
	"go-enterprise-blueprint/internal/modules/classroom/domain/accessgrant"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classgroup"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/shared"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	ClassID       string  `json:"class_id" validate:"required,uuid"`
	StudentUserID string  `json:"student_user_id" validate:"required"`
	AccessStatus  string  `json:"access_status" validate:"required,oneof=pending active paused blocked"`
	PaymentStatus string  `json:"payment_status" validate:"required,oneof=unknown pending confirmed rejected"`
	Note          *string `json:"note" validate:"omitempty,max=2000"`
}

type Response struct {
	Item accessgrant.AccessGrant `json:"item"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(dc *domain.Container) UseCase { return &usecase{dc: dc} }

type usecase struct{ dc *domain.Container }

func (uc *usecase) OperationID() string { return "update-access" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	classItem, err := uc.dc.ClassRepo().Get(ctx, classgroup.Filter{ID: &in.ClassID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if accessErr := shared.EnsureClassOperate(ctx, uc.dc, classItem); accessErr != nil {
		return nil, errx.Wrap(accessErr)
	}
	_, err = uc.dc.EnrollmentRepo().Get(ctx, classItem.ID, in.StudentUserID)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	item := &accessgrant.AccessGrant{OrganizationID: classItem.OrganizationID, ClassID: classItem.ID, StudentUserID: in.StudentUserID, AccessStatus: in.AccessStatus, PaymentStatus: in.PaymentStatus, Note: shared.TrimPtr(in.Note)}
	if in.AccessStatus == accessgrant.AccessActive {
		now := time.Now()
		actor := auth.MustUserContext(ctx).UserID
		item.GrantedAt = &now
		item.GrantedBy = &actor
	}
	updated, err := uc.dc.AccessRepo().Upsert(ctx, item)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{Item: *updated}, nil
}
