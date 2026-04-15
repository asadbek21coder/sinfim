package addstudent

import (
	"context"
	"strings"
	"time"

	authrbac "go-enterprise-blueprint/internal/modules/auth/domain/rbac"
	authuser "go-enterprise-blueprint/internal/modules/auth/domain/user"
	authpg "go-enterprise-blueprint/internal/modules/auth/infra/postgres"
	"go-enterprise-blueprint/internal/modules/classroom/domain"
	"go-enterprise-blueprint/internal/modules/classroom/domain/accessgrant"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classgroup"
	"go-enterprise-blueprint/internal/modules/classroom/domain/enrollment"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/shared"
	"go-enterprise-blueprint/internal/modules/organization/domain/membership"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/google/uuid"
	"github.com/rise-and-shine/pkg/hasher"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

type Request struct {
	ClassID           string  `json:"class_id" validate:"required,uuid"`
	FullName          string  `json:"full_name" validate:"required,min=2,max=120"`
	PhoneNumber       string  `json:"phone_number" validate:"required,min=7,max=32"`
	TemporaryPassword *string `json:"temporary_password" validate:"omitempty,min=8,max=120" mask:"true"`
	AccessStatus      string  `json:"access_status" validate:"omitempty,oneof=pending active paused blocked"`
	PaymentStatus     string  `json:"payment_status" validate:"omitempty,oneof=unknown pending confirmed rejected"`
	Note              *string `json:"note" validate:"omitempty,max=2000"`
}

type StudentDTO struct {
	ID                 string `json:"id"`
	FullName           string `json:"full_name"`
	PhoneNumber        string `json:"phone_number"`
	MustChangePassword bool   `json:"must_change_password"`
}

type Response struct {
	Student                    StudentDTO              `json:"student"`
	Enrollment                 enrollment.Enrollment   `json:"enrollment"`
	Access                     accessgrant.AccessGrant `json:"access"`
	TemporaryPasswordGenerated bool                    `json:"temporary_password_generated"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(db *bun.DB, dc *domain.Container, hashingCost int) UseCase {
	return &usecase{db: db, dc: dc, hashingCost: hashingCost}
}

type usecase struct {
	db          *bun.DB
	dc          *domain.Container
	hashingCost int
}

func (uc *usecase) OperationID() string { return "add-student" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	classItem, err := uc.dc.ClassRepo().Get(ctx, classgroup.Filter{ID: &in.ClassID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if accessErr := shared.EnsureClassOperate(ctx, uc.dc, classItem); accessErr != nil {
		return nil, errx.Wrap(accessErr)
	}

	phone := normalizePhone(in.PhoneNumber)
	userRepo := authpg.NewUserRepo(uc.db)
	roleRepo := authpg.NewRoleRepo(uc.db)
	userRoleRepo := authpg.NewUserRoleRepo(uc.db)
	student, created, tempPassword, err := uc.ensureStudentUser(ctx, userRepo, phone, strings.TrimSpace(in.FullName), in.TemporaryPassword)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if err := ensureGlobalRole(ctx, roleRepo, userRoleRepo, student.ID, membership.RoleStudent); err != nil {
		return nil, errx.Wrap(err)
	}
	if err := uc.ensureMembership(ctx, student.ID, classItem.OrganizationID); err != nil {
		return nil, errx.Wrap(err)
	}

	enrolled, err := uc.dc.EnrollmentRepo().Create(ctx, &enrollment.Enrollment{OrganizationID: classItem.OrganizationID, ClassID: classItem.ID, StudentUserID: student.ID, Status: enrollment.StatusActive})
	if err != nil {
		return nil, errx.WrapWithTypeOnCodes(err, errx.T_Conflict, enrollment.CodeStudentAlreadyEnrolled)
	}
	accessStatus := in.AccessStatus
	if accessStatus == "" {
		accessStatus = accessgrant.AccessPending
	}
	paymentStatus := in.PaymentStatus
	if paymentStatus == "" {
		paymentStatus = accessgrant.PaymentUnknown
	}
	grant := &accessgrant.AccessGrant{OrganizationID: classItem.OrganizationID, ClassID: classItem.ID, StudentUserID: student.ID, AccessStatus: accessStatus, PaymentStatus: paymentStatus, Note: shared.TrimPtr(in.Note)}
	if accessStatus == accessgrant.AccessActive {
		now := time.Now()
		actor := auth.MustUserContext(ctx).UserID
		grant.GrantedAt = &now
		grant.GrantedBy = &actor
	}
	grant, err = uc.dc.AccessRepo().Upsert(ctx, grant)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{Student: StudentDTO{ID: student.ID, FullName: deref(student.FullName), PhoneNumber: phone, MustChangePassword: student.MustChangePassword}, Enrollment: *enrolled, Access: *grant, TemporaryPasswordGenerated: created && tempPassword != ""}, nil
}

func (uc *usecase) ensureStudentUser(ctx context.Context, userRepo authuser.Repo, phone string, fullName string, temporaryPassword *string) (*authuser.User, bool, string, error) {
	student, err := userRepo.Get(ctx, authuser.Filter{PhoneNumber: &phone})
	if err == nil {
		return student, false, "", nil
	}
	if !errx.IsCodeIn(err, authuser.CodeUserNotFound) {
		return nil, false, "", errx.Wrap(err)
	}
	password := "TempPass123"
	generated := ""
	if temporaryPassword != nil && strings.TrimSpace(*temporaryPassword) != "" {
		password = *temporaryPassword
	} else {
		generated = password
	}
	hash, hashErr := hasher.Hash(password, hasher.WithCost(uc.hashingCost))
	if hashErr != nil {
		return nil, false, "", errx.Wrap(hashErr)
	}
	student, createErr := userRepo.Create(ctx, &authuser.User{ID: uuid.NewString(), Username: &phone, PhoneNumber: &phone, FullName: &fullName, PasswordHash: &hash, IsActive: true, MustChangePassword: true})
	return student, true, generated, createErr
}

func (uc *usecase) ensureMembership(ctx context.Context, userID string, organizationID string) error {
	ok, err := uc.dc.MembershipRepo().Exists(ctx, userID, organizationID, membership.RoleStudent)
	if err != nil {
		return errx.Wrap(err)
	}
	if ok {
		return nil
	}
	_, err = uc.dc.MembershipRepo().Create(ctx, &membership.Membership{UserID: userID, OrganizationID: organizationID, Role: membership.RoleStudent, IsActive: true})
	return errx.Wrap(err)
}

func ensureGlobalRole(ctx context.Context, roleRepo authrbac.RoleRepo, userRoleRepo authrbac.UserRoleRepo, userID string, roleName string) error {
	role, err := roleRepo.Get(ctx, authrbac.RoleFilter{Name: &roleName})
	if err != nil {
		return errx.Wrap(err)
	}
	assigned, err := userRoleRepo.List(ctx, authrbac.UserRoleFilter{UserID: &userID, RoleID: &role.ID})
	if err != nil {
		return errx.Wrap(err)
	}
	if len(assigned) > 0 {
		return nil
	}
	_, err = userRoleRepo.Create(ctx, &authrbac.UserRole{UserID: userID, RoleID: role.ID})
	return errx.Wrap(err)
}

func normalizePhone(value string) string {
	return strings.ReplaceAll(strings.TrimSpace(value), " ", "")
}
func deref(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
