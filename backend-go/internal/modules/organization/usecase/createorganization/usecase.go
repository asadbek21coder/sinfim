package createorganization

import (
	"context"
	"regexp"
	"strings"

	authrbac "go-enterprise-blueprint/internal/modules/auth/domain/rbac"
	authuser "go-enterprise-blueprint/internal/modules/auth/domain/user"
	authpg "go-enterprise-blueprint/internal/modules/auth/infra/postgres"
	"go-enterprise-blueprint/internal/modules/organization/domain"
	"go-enterprise-blueprint/internal/modules/organization/domain/membership"
	"go-enterprise-blueprint/internal/modules/organization/domain/org"
	orgpg "go-enterprise-blueprint/internal/modules/organization/infra/postgres"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/google/uuid"
	"github.com/rise-and-shine/pkg/hasher"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type OwnerRequest struct {
	FullName          string `json:"full_name" validate:"required,min=2,max=120"`
	PhoneNumber       string `json:"phone_number" validate:"required,min=7,max=32"`
	TemporaryPassword string `json:"temporary_password" validate:"required,min=8,max=120" mask:"true"`
}

type Request struct {
	Name         string       `json:"name" validate:"required,min=2,max=160"`
	Slug         string       `json:"slug" validate:"required,min=2,max=120"`
	Description  *string      `json:"description" validate:"omitempty,max=2000"`
	LogoURL      *string      `json:"logo_url" validate:"omitempty,max=1000"`
	Category     *string      `json:"category" validate:"omitempty,max=80"`
	ContactPhone *string      `json:"contact_phone" validate:"omitempty,max=32"`
	TelegramURL  *string      `json:"telegram_url" validate:"omitempty,max=1000"`
	IsDemo       bool         `json:"is_demo"`
	Owner        OwnerRequest `json:"owner" validate:"required"`
}

type OrganizationDTO struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Slug         string  `json:"slug"`
	Description  *string `json:"description"`
	LogoURL      *string `json:"logo_url"`
	PublicStatus string  `json:"public_status"`
	IsDemo       bool    `json:"is_demo"`
}

type OwnerDTO struct {
	ID                 string `json:"id"`
	FullName           string `json:"full_name"`
	PhoneNumber        string `json:"phone_number"`
	Role               string `json:"role"`
	MustChangePassword bool   `json:"must_change_password"`
}

type Response struct {
	Organization OrganizationDTO `json:"organization"`
	Owner        OwnerDTO        `json:"owner"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(db *bun.DB, domainContainer *domain.Container, hashingCost int) UseCase {
	return &usecase{db: db, domainContainer: domainContainer, hashingCost: hashingCost}
}

type usecase struct {
	db              *bun.DB
	domainContainer *domain.Container
	hashingCost     int
}

func (uc *usecase) OperationID() string { return "create-organization" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	if err := uc.ensurePlatformAdmin(ctx); err != nil {
		return nil, errx.Wrap(err)
	}

	slug := normalizeSlug(in.Slug)
	if !isValidSlug(slug) {
		return nil, errx.New("slug must be lowercase kebab-case", errx.WithType(errx.T_Validation), errx.WithCode("VALIDATION_ERROR"))
	}
	ownerPhone := normalizePhone(in.Owner.PhoneNumber)
	if !looksLikePhone(ownerPhone) {
		return nil, errx.New("owner phone number format is invalid", errx.WithType(errx.T_Validation), errx.WithCode("VALIDATION_ERROR"))
	}

	var response *Response
	err := uc.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		organizationRepo := orgpg.NewOrganizationRepo(tx)
		membershipRepo := orgpg.NewMembershipRepo(tx)
		userRepo := authpg.NewUserRepo(tx)
		roleRepo := authpg.NewRoleRepo(tx)
		userRoleRepo := authpg.NewUserRoleRepo(tx)

		_, existingErr := organizationRepo.GetBySlug(ctx, slug)
		if existingErr == nil {
			return errx.New("organization slug is already taken", errx.WithType(errx.T_Conflict), errx.WithCode(org.CodeSlugAlreadyTaken))
		}
		if !errx.IsCodeIn(existingErr, org.CodeOrganizationNotFound) {
			return errx.Wrap(existingErr)
		}

		organization, createErr := organizationRepo.Create(ctx, &org.Organization{
			Name:         strings.TrimSpace(in.Name),
			Slug:         slug,
			Description:  trimPtr(in.Description),
			LogoURL:      trimPtr(in.LogoURL),
			Category:     trimPtr(in.Category),
			ContactPhone: trimPhonePtr(in.ContactPhone),
			TelegramURL:  trimPtr(in.TelegramURL),
			PublicStatus: org.PublicStatusDraft,
			IsDemo:       in.IsDemo,
		})
		if createErr != nil {
			return errx.WrapWithTypeOnCodes(createErr, errx.T_Conflict, org.CodeSlugAlreadyTaken)
		}

		owner, ownerErr := userRepo.Get(ctx, authuser.Filter{PhoneNumber: &ownerPhone})
		if errx.IsCodeIn(ownerErr, authuser.CodeUserNotFound) {
			passwordHash, hashErr := hasher.Hash(in.Owner.TemporaryPassword, hasher.WithCost(uc.hashingCost))
			if hashErr != nil {
				return errx.Wrap(hashErr)
			}
			owner, ownerErr = userRepo.Create(ctx, &authuser.User{
				ID:                 uuid.NewString(),
				Username:           &ownerPhone,
				PhoneNumber:        &ownerPhone,
				FullName:           stringPtr(strings.TrimSpace(in.Owner.FullName)),
				PasswordHash:       &passwordHash,
				IsActive:           true,
				MustChangePassword: true,
			})
		}
		if ownerErr != nil {
			return errx.Wrap(ownerErr)
		}

		ownerRoleName := membership.RoleOwner
		ownerRole, roleErr := roleRepo.Get(ctx, authrbac.RoleFilter{Name: &ownerRoleName})
		if roleErr != nil {
			return errx.Wrap(roleErr)
		}

		assigned, assignedErr := userRoleRepo.List(ctx, authrbac.UserRoleFilter{UserID: &owner.ID, RoleID: &ownerRole.ID})
		if assignedErr != nil {
			return errx.Wrap(assignedErr)
		}
		if len(assigned) == 0 {
			_, roleCreateErr := userRoleRepo.Create(ctx, &authrbac.UserRole{UserID: owner.ID, RoleID: ownerRole.ID})
			if roleCreateErr != nil {
				return errx.Wrap(roleCreateErr)
			}
		}

		exists, existsErr := membershipRepo.Exists(ctx, owner.ID, organization.ID, membership.RoleOwner)
		if existsErr != nil {
			return errx.Wrap(existsErr)
		}
		if exists {
			return errx.New("owner is already member", errx.WithType(errx.T_Conflict), errx.WithCode(membership.CodeOwnerAlreadyMember))
		}
		_, membershipErr := membershipRepo.Create(ctx, &membership.Membership{
			UserID:         owner.ID,
			OrganizationID: organization.ID,
			Role:           membership.RoleOwner,
			IsActive:       true,
		})
		if membershipErr != nil {
			return errx.WrapWithTypeOnCodes(membershipErr, errx.T_Conflict, membership.CodeOwnerAlreadyMember)
		}

		response = &Response{
			Organization: OrganizationDTO{
				ID:           organization.ID,
				Name:         organization.Name,
				Slug:         organization.Slug,
				Description:  organization.Description,
				LogoURL:      organization.LogoURL,
				PublicStatus: organization.PublicStatus,
				IsDemo:       organization.IsDemo,
			},
			Owner: OwnerDTO{
				ID:                 owner.ID,
				FullName:           deref(owner.FullName),
				PhoneNumber:        ownerPhone,
				Role:               membership.RoleOwner,
				MustChangePassword: owner.MustChangePassword,
			},
		}
		return nil
	})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return response, nil
}

func (uc *usecase) ensurePlatformAdmin(ctx context.Context) error {
	userCtx := auth.MustUserContext(ctx)
	if auth.HasPermission(userCtx, auth.PermissionUserManage) {
		return nil
	}
	return errx.New("platform admin permission required", errx.WithType(errx.T_Forbidden), errx.WithCode(auth.CodeForbidden))
}

func normalizeSlug(value string) string {
	return strings.Trim(strings.ToLower(strings.TrimSpace(value)), "-")
}

func isValidSlug(value string) bool {
	return regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`).MatchString(value)
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

func trimPhonePtr(value *string) *string {
	if value == nil {
		return nil
	}
	trimmed := normalizePhone(*value)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}

func stringPtr(value string) *string {
	return &value
}

func deref(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
