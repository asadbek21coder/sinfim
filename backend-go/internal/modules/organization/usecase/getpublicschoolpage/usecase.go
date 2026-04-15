package getpublicschoolpage

import (
	"context"
	"regexp"
	"strings"

	"go-enterprise-blueprint/internal/modules/organization/domain"
	"go-enterprise-blueprint/internal/modules/organization/domain/org"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	Slug string `query:"slug" validate:"required,min=2,max=120"`
}

type OrganizationDTO struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Slug         string  `json:"slug"`
	Description  *string `json:"description"`
	LogoURL      *string `json:"logo_url"`
	Category     *string `json:"category"`
	ContactPhone *string `json:"contact_phone"`
	TelegramURL  *string `json:"telegram_url"`
	IsDemo       bool    `json:"is_demo"`
}

type LeadFormDTO struct {
	Enabled        bool     `json:"enabled"`
	RequiredFields []string `json:"required_fields"`
}

type Response struct {
	Organization OrganizationDTO `json:"organization"`
	Courses      []any           `json:"courses"`
	LeadForm     LeadFormDTO     `json:"lead_form"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "get-public-school-page" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	slug := strings.TrimSpace(strings.ToLower(in.Slug))
	if !regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`).MatchString(slug) {
		return nil, errx.New("slug must be lowercase kebab-case", errx.WithType(errx.T_Validation), errx.WithCode("VALIDATION_ERROR"))
	}

	organization, err := uc.domainContainer.OrganizationRepo().GetBySlug(ctx, slug)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if organization.PublicStatus != org.PublicStatusPublic && !organization.IsDemo {
		return nil, errx.New("organization is not public", errx.WithType(errx.T_NotFound), errx.WithCode(org.CodeOrganizationNotPublic))
	}

	return &Response{
		Organization: OrganizationDTO{
			ID:           organization.ID,
			Name:         organization.Name,
			Slug:         organization.Slug,
			Description:  organization.Description,
			LogoURL:      organization.LogoURL,
			Category:     organization.Category,
			ContactPhone: organization.ContactPhone,
			TelegramURL:  organization.TelegramURL,
			IsDemo:       organization.IsDemo,
		},
		Courses: []any{},
		LeadForm: LeadFormDTO{
			Enabled:        true,
			RequiredFields: []string{"full_name", "phone_number"},
		},
	}, nil
}
