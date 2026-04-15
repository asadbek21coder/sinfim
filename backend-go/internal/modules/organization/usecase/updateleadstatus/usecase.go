package updateleadstatus

import (
	"context"

	"go-enterprise-blueprint/internal/modules/organization/domain"
	"go-enterprise-blueprint/internal/modules/organization/domain/lead"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	ID     string `json:"id" validate:"required,uuid"`
	Status string `json:"status" validate:"required,oneof=new contacted converted archived"`
}

type Response struct {
	Item lead.Lead `json:"item"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "update-lead-status" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	existing, err := uc.domainContainer.LeadRepo().List(ctx, lead.Filter{ID: &in.ID, Limit: 1})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if len(existing) == 0 {
		return nil, errx.New("lead not found", errx.WithType(errx.T_NotFound), errx.WithCode(lead.CodeLeadNotFound))
	}
	if err := uc.ensureAccess(ctx, existing[0].OrganizationID); err != nil {
		return nil, errx.Wrap(err)
	}
	updated, updateErr := uc.domainContainer.LeadRepo().UpdateStatus(ctx, in.ID, in.Status)
	if updateErr != nil {
		return nil, errx.Wrap(updateErr)
	}
	return &Response{Item: *updated}, nil
}

func (uc *usecase) ensureAccess(ctx context.Context, organizationID string) error {
	userCtx := auth.MustUserContext(ctx)
	if auth.HasPermission(userCtx, auth.PermissionUserRead) {
		return nil
	}
	memberships, err := uc.domainContainer.MembershipRepo().ListByOrganization(ctx, organizationID)
	if err != nil {
		return errx.Wrap(err)
	}
	for _, item := range memberships {
		if item.UserID == userCtx.UserID {
			return nil
		}
	}
	return errx.New("organization membership required", errx.WithType(errx.T_Forbidden), errx.WithCode(auth.CodeForbidden))
}
