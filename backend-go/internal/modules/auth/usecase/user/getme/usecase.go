package getme

import (
	"context"
	"time"

	"go-enterprise-blueprint/internal/modules/auth/domain"
	"go-enterprise-blueprint/internal/modules/auth/domain/rbac"
	"go-enterprise-blueprint/internal/modules/auth/domain/user"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct{}

type Response struct {
	ID                 string  `json:"id"`
	PhoneNumber        string  `json:"phoneNumber"`
	FullName           string  `json:"fullName"`
	Role               string  `json:"role"`
	OrganizationID     *string `json:"organizationId"`
	IsActive           bool    `json:"isActive"`
	MustChangePassword bool    `json:"mustChangePassword"`
	CreatedAt          string  `json:"createdAt"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(domainContainer *domain.Container) UseCase {
	return &usecase{domainContainer: domainContainer}
}

type usecase struct {
	domainContainer *domain.Container
}

func (uc *usecase) OperationID() string { return "get-me" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	userCtx := auth.MustUserContext(ctx)
	u, err := uc.domainContainer.UserRepo().Get(ctx, user.Filter{ID: &userCtx.UserID})
	if err != nil {
		return nil, errx.WrapWithTypeOnCodes(err, errx.T_NotFound, user.CodeUserNotFound)
	}

	role := "STUDENT"
	userRoles, err := uc.domainContainer.UserRoleRepo().List(ctx, rbac.UserRoleFilter{UserID: &u.ID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if len(userRoles) > 0 {
		roleIDs := make([]int64, 0, len(userRoles))
		for _, userRole := range userRoles {
			roleIDs = append(roleIDs, userRole.RoleID)
		}
		roles, roleErr := uc.domainContainer.RoleRepo().List(ctx, rbac.RoleFilter{IDs: roleIDs})
		if roleErr != nil {
			return nil, errx.Wrap(roleErr)
		}
		if len(roles) > 0 {
			role = roles[0].Name
		}
	}

	phoneNumber := ""
	if u.PhoneNumber != nil {
		phoneNumber = *u.PhoneNumber
	} else if u.Username != nil {
		phoneNumber = *u.Username
	}
	fullName := ""
	if u.FullName != nil {
		fullName = *u.FullName
	}
	createdAt := ""
	if !u.CreatedAt.IsZero() {
		createdAt = u.CreatedAt.Format(time.RFC3339)
	}

	return &Response{
		ID:                 u.ID,
		PhoneNumber:        phoneNumber,
		FullName:           fullName,
		Role:               role,
		OrganizationID:     nil,
		IsActive:           u.IsActive,
		MustChangePassword: u.MustChangePassword,
		CreatedAt:          createdAt,
	}, nil
}
