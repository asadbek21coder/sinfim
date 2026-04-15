package shared

import (
	"context"
	"strings"

	"go-enterprise-blueprint/internal/modules/classroom/domain/classgroup"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classmentor"
	"go-enterprise-blueprint/internal/modules/organization/domain/membership"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
)

type AccessContainer interface {
	ClassRepo() classgroup.Repo
	MentorRepo() classmentor.Repo
	MembershipRepo() membership.Repo
}

func EnsureClassWrite(ctx context.Context, dc AccessContainer, organizationID string) error {
	userCtx := auth.MustUserContext(ctx)
	if auth.HasPermission(userCtx, auth.PermissionUserManage) {
		return nil
	}
	for _, role := range []string{membership.RoleOwner, membership.RoleTeacher} {
		ok, err := dc.MembershipRepo().Exists(ctx, userCtx.UserID, organizationID, role)
		if err != nil {
			return errx.Wrap(err)
		}
		if ok {
			return nil
		}
	}
	return errx.New("owner or teacher membership required", errx.WithType(errx.T_Forbidden), errx.WithCode(auth.CodeForbidden))
}

func EnsureClassOperate(ctx context.Context, dc AccessContainer, classItem *classgroup.Class) error {
	userCtx := auth.MustUserContext(ctx)
	if auth.HasPermission(userCtx, auth.PermissionUserRead) {
		return nil
	}
	for _, role := range []string{membership.RoleOwner, membership.RoleTeacher} {
		ok, err := dc.MembershipRepo().Exists(ctx, userCtx.UserID, classItem.OrganizationID, role)
		if err != nil {
			return errx.Wrap(err)
		}
		if ok {
			return nil
		}
	}
	assigned, err := dc.MentorRepo().IsAssigned(ctx, classItem.ID, userCtx.UserID)
	if err != nil {
		return errx.Wrap(err)
	}
	if assigned {
		return nil
	}
	return errx.New("class access required", errx.WithType(errx.T_Forbidden), errx.WithCode(auth.CodeForbidden))
}

func TrimPtr(value *string) *string {
	if value == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}
