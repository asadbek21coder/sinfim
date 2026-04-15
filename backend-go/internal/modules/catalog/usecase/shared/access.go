package shared

import (
	"context"
	"strings"

	"go-enterprise-blueprint/internal/modules/organization/domain/membership"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
)

type domainAccess interface {
	MembershipRepo() membership.Repo
}

func EnsureCourseWriteAccess(ctx context.Context, domainContainer domainAccess, organizationID string) error {
	userCtx := auth.MustUserContext(ctx)
	if auth.HasPermission(userCtx, auth.PermissionUserManage) {
		return nil
	}
	for _, role := range []string{membership.RoleOwner, membership.RoleTeacher} {
		allowed, err := domainContainer.MembershipRepo().Exists(ctx, userCtx.UserID, organizationID, role)
		if err != nil {
			return errx.Wrap(err)
		}
		if allowed {
			return nil
		}
	}
	return errx.New("owner or teacher membership required", errx.WithType(errx.T_Forbidden), errx.WithCode(auth.CodeForbidden))
}

func EnsureCourseReadAccess(ctx context.Context, domainContainer domainAccess, organizationID string) error {
	userCtx := auth.MustUserContext(ctx)
	if auth.HasPermission(userCtx, auth.PermissionUserRead) {
		return nil
	}
	memberships, err := domainContainer.MembershipRepo().ListByOrganization(ctx, organizationID)
	if err != nil {
		return errx.Wrap(err)
	}
	for _, item := range memberships {
		if item.UserID == userCtx.UserID {
			return nil
		}
	}
	return errx.New("organization membership required", errx.WithType(errx.T_Forbidden), errx.WithCode(auth.CodeForbidden))
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
