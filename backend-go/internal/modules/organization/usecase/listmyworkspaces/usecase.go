package listmyworkspaces

import (
	"context"

	"go-enterprise-blueprint/internal/modules/organization/domain"
	"go-enterprise-blueprint/internal/modules/organization/domain/org"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct{}

type WorkspaceItem struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Slug         string  `json:"slug"`
	Description  *string `json:"description"`
	LogoURL      *string `json:"logo_url"`
	Category     *string `json:"category"`
	ContactPhone *string `json:"contact_phone"`
	TelegramURL  *string `json:"telegram_url"`
	PublicStatus string  `json:"public_status"`
	IsDemo       bool    `json:"is_demo"`
	Role         string  `json:"role"`
}

type Response struct {
	Items []WorkspaceItem `json:"items"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "list-my-workspaces" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	userCtx := auth.MustUserContext(ctx)
	memberships, err := uc.domainContainer.MembershipRepo().ListByUser(ctx, userCtx.UserID)
	if err != nil {
		return nil, errx.Wrap(err)
	}

	items := make([]WorkspaceItem, 0, len(memberships))
	for _, membership := range memberships {
		organizations, orgErr := uc.domainContainer.OrganizationRepo().List(ctx, org.Filter{ID: &membership.OrganizationID, Limit: 1})
		if orgErr != nil {
			return nil, errx.Wrap(orgErr)
		}
		if len(organizations) == 0 {
			continue
		}
		organization := organizations[0]
		items = append(items, WorkspaceItem{
			ID:           organization.ID,
			Name:         organization.Name,
			Slug:         organization.Slug,
			Description:  organization.Description,
			LogoURL:      organization.LogoURL,
			Category:     organization.Category,
			ContactPhone: organization.ContactPhone,
			TelegramURL:  organization.TelegramURL,
			PublicStatus: organization.PublicStatus,
			IsDemo:       organization.IsDemo,
			Role:         membership.Role,
		})
	}

	return &Response{Items: items}, nil
}
