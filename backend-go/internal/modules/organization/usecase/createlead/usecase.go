package createlead

import (
	"context"
	"regexp"
	"strings"

	"go-enterprise-blueprint/internal/modules/organization/domain"
	"go-enterprise-blueprint/internal/modules/organization/domain/lead"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	OrganizationID string  `json:"organization_id" validate:"required,uuid"`
	FullName       string  `json:"full_name" validate:"required,min=2,max=120"`
	PhoneNumber    string  `json:"phone_number" validate:"required,min=7,max=32"`
	Note           *string `json:"note" validate:"omitempty,max=2000"`
}

type Response struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "create-lead" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	phone := normalizePhone(in.PhoneNumber)
	if !regexp.MustCompile(`^\+?[0-9]{7,15}$`).MatchString(phone) {
		return nil, errx.New("phone number format is invalid", errx.WithType(errx.T_Validation), errx.WithCode("VALIDATION_ERROR"))
	}
	created, err := uc.domainContainer.LeadRepo().Create(ctx, &lead.Lead{
		OrganizationID: in.OrganizationID,
		FullName:       strings.TrimSpace(in.FullName),
		PhoneNumber:    phone,
		Note:           trimPtr(in.Note),
		Source:         lead.SourcePublicSchoolPage,
		Status:         lead.StatusNew,
	})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{ID: created.ID, Status: created.Status, Message: "Arizangiz qabul qilindi. Tez orada siz bilan bog'lanamiz."}, nil
}

func normalizePhone(value string) string {
	replacer := strings.NewReplacer(" ", "", "-", "", "(", "", ")", "")
	return replacer.Replace(strings.TrimSpace(value))
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
