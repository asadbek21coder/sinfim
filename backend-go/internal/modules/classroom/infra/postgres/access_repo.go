package postgres

import (
	"context"
	"time"

	"go-enterprise-blueprint/internal/modules/classroom/domain/accessgrant"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type accessRepo struct{ db bun.IDB }

func NewAccessRepo(db bun.IDB) accessgrant.Repo { return &accessRepo{db: db} }

func (r *accessRepo) Upsert(ctx context.Context, item *accessgrant.AccessGrant) (*accessgrant.AccessGrant, error) {
	now := time.Now()
	if item.CreatedAt.IsZero() {
		item.CreatedAt = now
	}
	item.UpdatedAt = now
	_, err := r.db.NewInsert().Model(item).
		On("CONFLICT (class_id, student_user_id) DO UPDATE").
		Set("access_status = EXCLUDED.access_status").
		Set("payment_status = EXCLUDED.payment_status").
		Set("note = EXCLUDED.note").
		Set("granted_by = EXCLUDED.granted_by").
		Set("granted_at = EXCLUDED.granted_at").
		Set("updated_at = EXCLUDED.updated_at").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return item, nil
}
