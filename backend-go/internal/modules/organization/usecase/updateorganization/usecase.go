package updateorganization

import (
	"context"
	"strings"

	"go-enterprise-blueprint/internal/modules/organization/domain"
	"go-enterprise-blueprint/internal/modules/organization/domain/membership"
	"go-enterprise-blueprint/internal/modules/organization/domain/org"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	ID           string  `json:"id" validate:"required,uuid"`
	Name         string  `json:"name" validate:"required,min=2,max=160"`
	Description  *string `json:"description" validate:"omitempty,max=2000"`
	LogoURL      *string `json:"logo_url" validate:"omitempty,max=1000"`
	Category     *string `json:"category" validate:"omitempty,max=80"`
	ContactPhone *string `json:"contact_phone" validate:"omitempty,max=32"`
	TelegramURL  *string `json:"telegram_url" validate:"omitempty,max=1000"`
	PublicStatus string  `json:"public_status" validate:"required,oneof=draft public hidden"`
	IsDemo       bool    `json:"is_demo"`
}

type Response struct {
	Item org.Organization `json:"item"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "update-organization" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	userCtx := auth.MustUserContext(ctx)
	if !auth.HasPermission(userCtx, auth.PermissionUserManage) {
		allowed, err := uc.domainContainer.MembershipRepo().Exists(ctx, userCtx.UserID, in.ID, membership.RoleOwner)
		if err != nil {
			return nil, errx.Wrap(err)
		}
		if !allowed {
			return nil, errx.New("owner membership required", errx.WithType(errx.T_Forbidden), errx.WithCode(auth.CodeForbidden))
		}
	}

	organizations, err := uc.domainContainer.OrganizationRepo().List(ctx, org.Filter{ID: &in.ID, Limit: 1})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if len(organizations) == 0 {
		return nil, errx.New("organization not found", errx.WithType(errx.T_NotFound), errx.WithCode(org.CodeOrganizationNotFound))
	}

	organization := organizations[0]
	organization.Name = strings.TrimSpace(in.Name)
	organization.Description = trimPtr(in.Description)
	organization.LogoURL = trimPtr(in.LogoURL)
	organization.Category = trimPtr(in.Category)
	organization.ContactPhone = trimPtr(in.ContactPhone)
	organization.TelegramURL = trimPtr(in.TelegramURL)
	organization.PublicStatus = in.PublicStatus
	organization.IsDemo = in.IsDemo

	updated, updateErr := uc.domainContainer.OrganizationRepo().Update(ctx, &organization)
	if updateErr != nil {
		return nil, errx.Wrap(updateErr)
	}

	return &Response{Item: *updated}, nil
}

func trimPtr(value *string) *string {
	if value == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}
