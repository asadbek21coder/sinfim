package adminlogin

import (
	"context"
	"strings"
	"time"

	"go-enterprise-blueprint/internal/modules/auth/domain"
	"go-enterprise-blueprint/internal/modules/auth/domain/rbac"
	"go-enterprise-blueprint/internal/modules/auth/domain/session"
	"go-enterprise-blueprint/internal/modules/auth/domain/uow"
	"go-enterprise-blueprint/internal/modules/auth/domain/user"
	"go-enterprise-blueprint/internal/portal"
	"go-enterprise-blueprint/internal/portal/audit"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/hasher"
	"github.com/rise-and-shine/pkg/meta"
	"github.com/rise-and-shine/pkg/sorter"
	"github.com/rise-and-shine/pkg/token"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/samber/lo"
)

var (
	errIncorrectCreds = errx.New(
		"username or password is incorrect",
		errx.WithType(errx.T_Validation),
		errx.WithCode(user.CodeIncorrectCreds),
	)
)

type Request struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password"     validate:"required" mask:"true"`
}

type Response struct {
	AccessToken           string  `json:"accessToken"`
	AccessTokenExpiresAt  string  `json:"accessTokenExpiresAt"`
	RefreshToken          string  `json:"refreshToken"`
	RefreshTokenExpiresAt string  `json:"refreshTokenExpiresAt"`
	TokenType             string  `json:"tokenType"`
	ExpiresIn             int64   `json:"expiresIn"`
	User                  UserDTO `json:"user"`
}

type UserDTO struct {
	ID                 string  `json:"id"`
	PhoneNumber        string  `json:"phoneNumber"`
	FullName           string  `json:"fullName"`
	Role               string  `json:"role"`
	OrganizationID     *string `json:"organizationId"`
	IsActive           bool    `json:"isActive"`
	MustChangePassword bool    `json:"mustChangePassword"`
	CreatedAt          string  `json:"createdAt"`
}

// UseCase implements "admin-login" user action.
type UseCase = ucdef.UserAction[*Request, *Response]

func New(
	domainContainer *domain.Container,
	portalContainer *portal.Container,
	accessTokenTTL,
	refreshTokenTTL time.Duration,
	maxActiveSessions int,
) UseCase {
	return &usecase{
		domainContainer:   domainContainer,
		portalContainer:   portalContainer,
		accessTokenTTL:    accessTokenTTL,
		refreshTokenTTL:   refreshTokenTTL,
		maxActiveSessions: maxActiveSessions,
	}
}

type usecase struct {
	domainContainer *domain.Container
	portalContainer *portal.Container

	accessTokenTTL    time.Duration
	refreshTokenTTL   time.Duration
	maxActiveSessions int
}

func (uc *usecase) OperationID() string { return "admin-login" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	phoneNumber := normalizePhone(in.PhoneNumber)

	// Find user by phone number. Username remains as a compatibility fallback for old blueprint data.
	u, err := uc.domainContainer.UserRepo().Get(ctx, user.Filter{
		PhoneNumber: &phoneNumber,
	})
	if errx.IsCodeIn(err, user.CodeUserNotFound) {
		u, err = uc.domainContainer.UserRepo().Get(ctx, user.Filter{Username: &phoneNumber})
	}
	if errx.IsCodeIn(err, user.CodeUserNotFound) {
		return nil, errx.Wrap(errIncorrectCreds, errx.WithDetails(errx.D{"cause": "phone_number"}))
	}
	if err != nil {
		return nil, errx.Wrap(err)
	}

	ctx = context.WithValue(ctx, meta.ActorType, auth.ActorTypeUser)
	ctx = context.WithValue(ctx, meta.ActorID, u.ID)

	// Check if user is active
	if !u.IsActive {
		return nil, errx.Wrap(errIncorrectCreds, errx.WithDetails(errx.D{"cause": "is_active"}))
	}

	// Verify password hash
	if u.PasswordHash == nil {
		return nil, errx.Wrap(errIncorrectCreds, errx.WithDetails(errx.D{"cause": "password"}))
	}
	ok := hasher.Compare(in.Password, *u.PasswordHash)
	if !ok {
		return nil, errx.Wrap(errIncorrectCreds, errx.WithDetails(errx.D{"cause": "password"}))
	}

	// Start UOW
	uow, err := uc.domainContainer.UOWFactory().NewUOW(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	defer uow.DiscardUnapplied()

	// Enforce max active sessions limit (delete least recently used sessions if exceeded)
	err = uc.deleteExceededSessions(ctx, u, uow)
	if err != nil {
		return nil, errx.Wrap(err)
	}

	// Create session record with tokens and meta info (IP, user_agent)
	s, err := uow.Session().Create(ctx, &session.Session{
		UserID:                u.ID,
		AccessToken:           token.NewOpaqueToken(),
		AccessTokenExpiresAt:  time.Now().Add(uc.accessTokenTTL),
		RefreshToken:          token.NewOpaqueToken(),
		RefreshTokenExpiresAt: time.Now().Add(uc.refreshTokenTTL),
		IPAddress:             meta.Find(ctx, meta.IPAddress),
		UserAgent:             meta.Find(ctx, meta.UserAgent),
		LastUsedAt:            time.Now(),
	})
	if err != nil {
		return nil, errx.Wrap(err)
	}

	// Update user's last_active_at timestamp
	u.LastActiveAt = lo.ToPtr(time.Now())
	_, err = uow.User().Update(ctx, u)
	if err != nil {
		return nil, errx.Wrap(err)
	}

	// Record audit log
	err = uc.portalContainer.Audit().Log(uow.Lend(), audit.Action{
		Module: auth.ModuleName, OperationID: uc.OperationID(), Payload: in,
	})
	if err != nil {
		return nil, errx.Wrap(err)
	}

	// Apply UOW
	err = uow.ApplyChanges()
	if err != nil {
		return nil, errx.Wrap(err)
	}

	userDTO, err := uc.buildUserDTO(ctx, u)
	if err != nil {
		return nil, errx.Wrap(err)
	}

	return &Response{
		AccessToken:           s.AccessToken,
		AccessTokenExpiresAt:  s.AccessTokenExpiresAt.Format(time.RFC3339),
		RefreshToken:          s.RefreshToken,
		RefreshTokenExpiresAt: s.RefreshTokenExpiresAt.Format(time.RFC3339),
		TokenType:             "Bearer",
		ExpiresIn:             int64(uc.accessTokenTTL.Seconds()),
		User:                  userDTO,
	}, nil
}

func (uc *usecase) buildUserDTO(ctx context.Context, u *user.User) (UserDTO, error) {
	role := "STUDENT"
	userRoles, err := uc.domainContainer.UserRoleRepo().List(ctx, rbac.UserRoleFilter{UserID: &u.ID})
	if err != nil {
		return UserDTO{}, errx.Wrap(err)
	}
	if len(userRoles) > 0 {
		roleIDs := make([]int64, 0, len(userRoles))
		for _, userRole := range userRoles {
			roleIDs = append(roleIDs, userRole.RoleID)
		}
		roles, roleErr := uc.domainContainer.RoleRepo().List(ctx, rbac.RoleFilter{IDs: roleIDs})
		if roleErr != nil {
			return UserDTO{}, errx.Wrap(roleErr)
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

	return UserDTO{
		ID:                 u.ID,
		PhoneNumber:        phoneNumber,
		FullName:           fullName,
		Role:               role,
		OrganizationID:     nil,
		IsActive:           u.IsActive,
		MustChangePassword: u.MustChangePassword,
		CreatedAt:          u.CreatedAt.Format(time.RFC3339),
	}, nil
}

func normalizePhone(value string) string {
	replacer := strings.NewReplacer(" ", "", "-", "", "(", "", ")", "")
	return replacer.Replace(value)
}

func (uc *usecase) deleteExceededSessions(ctx context.Context, u *user.User, uow uow.UnitOfWork) error {
	activeSessions, err := uow.Session().List(ctx, session.Filter{
		UserID:            &u.ID,
		OrderByLastUsedAt: lo.ToPtr(sorter.Asc),
	})
	if err != nil {
		return errx.Wrap(err)
	}
	sessionsToDelete := len(activeSessions) - uc.maxActiveSessions + 1

	if sessionsToDelete <= 0 {
		return nil
	}

	err = uow.Session().BulkDelete(ctx, activeSessions[:sessionsToDelete])
	return errx.Wrap(err)
}
