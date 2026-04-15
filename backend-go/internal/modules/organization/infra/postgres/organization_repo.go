package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/modules/organization/domain/org"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type organizationRepo struct {
	db bun.IDB
}

func NewOrganizationRepo(db bun.IDB) org.Repo {
	return &organizationRepo{db: db}
}

func (r *organizationRepo) Create(ctx context.Context, organization *org.Organization) (*org.Organization, error) {
	now := time.Now()
	organization.CreatedAt = now
	organization.UpdatedAt = now
	if organization.PublicStatus == "" {
		organization.PublicStatus = org.PublicStatusDraft
	}
	_, err := r.db.NewInsert().Model(organization).Returning("*").Exec(ctx)
	if err != nil {
		return nil, errx.WrapWithTypeOnCodes(err, errx.T_Conflict, org.CodeSlugAlreadyTaken)
	}
	return organization, nil
}

func (r *organizationRepo) Update(ctx context.Context, organization *org.Organization) (*org.Organization, error) {
	organization.UpdatedAt = time.Now()
	res, err := r.db.NewUpdate().Model(organization).
		Column("name", "description", "logo_url", "category", "contact_phone", "telegram_url", "public_status", "is_demo", "updated_at").
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	rows, rowsErr := res.RowsAffected()
	if rowsErr != nil {
		return nil, errx.Wrap(rowsErr)
	}
	if rows == 0 {
		return nil, errx.New("organization not found", errx.WithType(errx.T_NotFound), errx.WithCode(org.CodeOrganizationNotFound))
	}
	return organization, nil
}

func (r *organizationRepo) GetBySlug(ctx context.Context, slug string) (*org.Organization, error) {
	organization := new(org.Organization)
	err := r.db.NewSelect().Model(organization).Where("slug = ?", slug).Limit(1).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errx.New("organization not found", errx.WithType(errx.T_NotFound), errx.WithCode(org.CodeOrganizationNotFound))
	}
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return organization, nil
}

func (r *organizationRepo) List(ctx context.Context, filter org.Filter) ([]org.Organization, error) {
	organizations := make([]org.Organization, 0)
	q := r.db.NewSelect().Model(&organizations).Order("created_at DESC")
	if filter.ID != nil {
		q = q.Where("id = ?", *filter.ID)
	}
	if filter.Slug != nil {
		q = q.Where("slug = ?", *filter.Slug)
	}
	if filter.Limit > 0 {
		q = q.Limit(filter.Limit)
	}
	if err := q.Scan(ctx); err != nil {
		return nil, errx.Wrap(err)
	}
	return organizations, nil
}
