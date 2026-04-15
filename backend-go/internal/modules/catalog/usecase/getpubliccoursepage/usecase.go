package getpubliccoursepage

import (
	"context"

	"go-enterprise-blueprint/internal/modules/catalog/domain"
	"go-enterprise-blueprint/internal/modules/catalog/domain/course"
	"go-enterprise-blueprint/internal/modules/organization/domain/org"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	SchoolSlug string `query:"school_slug" validate:"required,min=2,max=120"`
	CourseSlug string `query:"course_slug" validate:"required,min=2,max=120"`
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
	Source         string   `json:"source"`
}

type Response struct {
	Organization OrganizationDTO `json:"organization"`
	Course       course.Course   `json:"course"`
	Lessons      []any           `json:"lessons"`
	LeadForm     LeadFormDTO     `json:"lead_form"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "get-public-course-page" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	organization, err := uc.domainContainer.OrganizationRepo().GetBySlug(ctx, in.SchoolSlug)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if organization.PublicStatus != org.PublicStatusPublic && !organization.IsDemo {
		return nil, errx.New("organization is not public", errx.WithType(errx.T_NotFound), errx.WithCode(org.CodeOrganizationNotPublic))
	}
	item, courseErr := uc.domainContainer.CourseRepo().Get(ctx, course.Filter{OrganizationID: &organization.ID, Slug: &in.CourseSlug})
	if courseErr != nil {
		return nil, errx.Wrap(courseErr)
	}
	if item.PublicStatus != course.PublicStatusPublic {
		return nil, errx.New("course is not public", errx.WithType(errx.T_NotFound), errx.WithCode(course.CodeCourseNotFound))
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
		Course:   *item,
		Lessons:  []any{},
		LeadForm: LeadFormDTO{Enabled: true, RequiredFields: []string{"full_name", "phone_number"}, Source: "public_course_page"},
	}, nil
}
