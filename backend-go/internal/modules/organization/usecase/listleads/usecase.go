package listleads

import (
	"context"

	"go-enterprise-blueprint/internal/modules/organization/domain"
	"go-enterprise-blueprint/internal/modules/organization/domain/lead"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	OrganizationID string  `query:"organization_id" validate:"required,uuid"`
	Status         *string `query:"status" validate:"omitempty,oneof=new contacted converted archived"`
	Limit          int     `query:"limit" validate:"omitempty,min=1,max=200"`
}

type Response struct {
	Items []lead.Lead `json:"items"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "list-leads" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	if err := uc.ensureAccess(ctx, in.OrganizationID); err != nil {
		return nil, errx.Wrap(err)
	}
	limit := in.Limit
	if limit == 0 {
		limit = 100
	}
	items, err := uc.domainContainer.LeadRepo().List(ctx, lead.Filter{OrganizationID: &in.OrganizationID, Status: in.Status, Limit: limit})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{Items: items}, nil
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
