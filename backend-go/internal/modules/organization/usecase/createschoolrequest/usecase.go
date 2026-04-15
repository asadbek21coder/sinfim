package createschoolrequest

import (
	"context"
	"regexp"
	"strings"

	"go-enterprise-blueprint/internal/modules/organization/domain"
	"go-enterprise-blueprint/internal/modules/organization/domain/schoolrequest"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	FullName     string  `json:"full_name" validate:"required,min=2,max=120"`
	PhoneNumber  string  `json:"phone_number" validate:"required,min=7,max=32"`
	SchoolName   string  `json:"school_name" validate:"required,min=2,max=160"`
	Category     *string `json:"category" validate:"omitempty,max=80"`
	StudentCount *int    `json:"student_count" validate:"omitempty,min=0"`
	Note         *string `json:"note" validate:"omitempty,max=2000"`
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

func (uc *usecase) OperationID() string { return "create-school-request" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	phone := normalizePhone(in.PhoneNumber)
	if !looksLikePhone(phone) {
		return nil, errx.New(
			"phone number format is invalid",
			errx.WithType(errx.T_Validation),
			errx.WithCode("VALIDATION_ERROR"),
		)
	}

	schoolName := strings.TrimSpace(in.SchoolName)
	existing, err := uc.domainContainer.SchoolRequestRepo().FindOpenDuplicate(ctx, phone, schoolName)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if existing != nil {
		return &Response{ID: existing.ID, Status: existing.Status, Message: "Bu ariza allaqachon qabul qilingan. Tez orada siz bilan bog'lanamiz."}, nil
	}

	created, err := uc.domainContainer.SchoolRequestRepo().Create(ctx, &schoolrequest.SchoolRequest{
		FullName:     strings.TrimSpace(in.FullName),
		PhoneNumber:  phone,
		SchoolName:   schoolName,
		Category:     trimPtr(in.Category),
		StudentCount: in.StudentCount,
		Note:         trimPtr(in.Note),
		Status:       schoolrequest.StatusNew,
	})
	if err != nil {
		return nil, errx.Wrap(err)
	}

	return &Response{ID: created.ID, Status: created.Status, Message: "Arizangiz qabul qilindi. Platform administratori siz bilan bog'lanadi."}, nil
}

func normalizePhone(value string) string {
	replacer := strings.NewReplacer(" ", "", "-", "", "(", "", ")", "")
	return replacer.Replace(strings.TrimSpace(value))
}

func looksLikePhone(value string) bool {
	return regexp.MustCompile(`^\+?[0-9]{7,15}$`).MatchString(value)
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
