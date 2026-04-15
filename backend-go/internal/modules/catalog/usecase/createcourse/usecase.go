package createcourse

import (
	"context"
	"regexp"
	"strings"

	"go-enterprise-blueprint/internal/modules/catalog/domain"
	"go-enterprise-blueprint/internal/modules/catalog/domain/course"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/shared"
	"go-enterprise-blueprint/internal/modules/organization/domain/org"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	OrganizationID string  `json:"organization_id" validate:"required,uuid"`
	Title          string  `json:"title" validate:"required,min=2,max=160"`
	Slug           string  `json:"slug" validate:"required,min=2,max=120"`
	Description    *string `json:"description" validate:"omitempty,max=2000"`
	Category       *string `json:"category" validate:"omitempty,max=80"`
	Level          *string `json:"level" validate:"omitempty,max=80"`
	PublicStatus   string  `json:"public_status" validate:"omitempty,oneof=draft public hidden"`
}

type Response struct {
	Item course.Course `json:"item"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "create-course" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	if err := shared.EnsureCourseWriteAccess(ctx, uc.domainContainer, in.OrganizationID); err != nil {
		return nil, errx.Wrap(err)
	}
	organizations, err := uc.domainContainer.OrganizationRepo().List(ctx, org.Filter{ID: &in.OrganizationID, Limit: 1})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if len(organizations) == 0 {
		return nil, errx.New("organization not found", errx.WithType(errx.T_NotFound), errx.WithCode(org.CodeOrganizationNotFound))
	}

	slug := normalizeSlug(in.Slug)
	if !isValidSlug(slug) {
		return nil, errx.New("slug must be lowercase kebab-case", errx.WithType(errx.T_Validation), errx.WithCode("VALIDATION_ERROR"))
	}
	publicStatus := in.PublicStatus
	if publicStatus == "" {
		publicStatus = course.PublicStatusDraft
	}
	item, createErr := uc.domainContainer.CourseRepo().Create(ctx, &course.Course{
		OrganizationID: in.OrganizationID,
		Title:          strings.TrimSpace(in.Title),
		Slug:           slug,
		Description:    shared.TrimPtr(in.Description),
		Category:       shared.TrimPtr(in.Category),
		Level:          shared.TrimPtr(in.Level),
		Status:         course.StatusDraft,
		PublicStatus:   publicStatus,
	})
	if createErr != nil {
		return nil, errx.WrapWithTypeOnCodes(createErr, errx.T_Conflict, course.CodeCourseSlugAlreadyUsed)
	}
	return &Response{Item: *item}, nil
}

func normalizeSlug(value string) string {
	return strings.Trim(strings.ToLower(strings.TrimSpace(value)), "-")
}

func isValidSlug(value string) bool {
	return regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`).MatchString(value)
}
