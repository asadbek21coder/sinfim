package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/modules/organization/domain/membership"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type membershipRepo struct {
	db bun.IDB
}

func NewMembershipRepo(db bun.IDB) membership.Repo {
	return &membershipRepo{db: db}
}

func (r *membershipRepo) Create(ctx context.Context, m *membership.Membership) (*membership.Membership, error) {
	now := time.Now()
	m.CreatedAt = now
	m.UpdatedAt = now
	m.IsActive = true
	_, err := r.db.NewInsert().Model(m).Returning("*").Exec(ctx)
	if err != nil {
		return nil, errx.WrapWithTypeOnCodes(err, errx.T_Conflict, membership.CodeOwnerAlreadyMember)
	}
	return m, nil
}

func (r *membershipRepo) Exists(ctx context.Context, userID string, organizationID string, role string) (bool, error) {
	m := new(membership.Membership)
	err := r.db.NewSelect().Model(m).
		Where("user_id = ?", userID).
		Where("organization_id = ?", organizationID).
		Where("role = ?", role).
		Limit(1).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, errx.Wrap(err)
	}
	return true, nil
}

func (r *membershipRepo) ListByUser(ctx context.Context, userID string) ([]membership.Membership, error) {
	memberships := make([]membership.Membership, 0)
	err := r.db.NewSelect().Model(&memberships).
		Where("user_id = ?", userID).
		Where("is_active = true").
		Order("created_at ASC").
		Scan(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return memberships, nil
}

func (r *membershipRepo) ListByOrganization(ctx context.Context, organizationID string) ([]membership.Membership, error) {
	memberships := make([]membership.Membership, 0)
	err := r.db.NewSelect().Model(&memberships).
		Where("organization_id = ?", organizationID).
		Where("is_active = true").
		Order("created_at ASC").
		Scan(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return memberships, nil
}
