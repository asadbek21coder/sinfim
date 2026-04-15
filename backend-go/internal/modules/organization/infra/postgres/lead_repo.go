package postgres

import (
	"context"
	"time"

	"go-enterprise-blueprint/internal/modules/organization/domain/lead"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type leadRepo struct {
	db bun.IDB
}

func NewLeadRepo(db bun.IDB) lead.Repo {
	return &leadRepo{db: db}
}

func (r *leadRepo) Create(ctx context.Context, item *lead.Lead) (*lead.Lead, error) {
	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now
	if item.Status == "" {
		item.Status = lead.StatusNew
	}
	if item.Source == "" {
		item.Source = lead.SourcePublicSchoolPage
	}
	_, err := r.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return item, nil
}

func (r *leadRepo) List(ctx context.Context, filter lead.Filter) ([]lead.Lead, error) {
	items := make([]lead.Lead, 0)
	q := r.db.NewSelect().Model(&items).Order("created_at DESC")
	if filter.ID != nil {
		q = q.Where("id = ?", *filter.ID)
	}
	if filter.OrganizationID != nil {
		q = q.Where("organization_id = ?", *filter.OrganizationID)
	}
	if filter.Status != nil {
		q = q.Where("status = ?", *filter.Status)
	}
	if filter.Limit > 0 {
		q = q.Limit(filter.Limit)
	}
	if err := q.Scan(ctx); err != nil {
		return nil, errx.Wrap(err)
	}
	return items, nil
}

func (r *leadRepo) UpdateStatus(ctx context.Context, id string, status string) (*lead.Lead, error) {
	item := &lead.Lead{ID: id}
	res, err := r.db.NewUpdate().
		Model(item).
		Set("status = ?", status).
		Set("updated_at = ?", time.Now()).
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
		return nil, errx.New("lead not found", errx.WithType(errx.T_NotFound), errx.WithCode(lead.CodeLeadNotFound))
	}
	return item, nil
}
