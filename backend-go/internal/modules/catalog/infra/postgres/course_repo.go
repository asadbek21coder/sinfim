package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/modules/catalog/domain/course"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type courseRepo struct {
	db bun.IDB
}

func NewCourseRepo(db bun.IDB) course.Repo {
	return &courseRepo{db: db}
}

func (r *courseRepo) Create(ctx context.Context, item *course.Course) (*course.Course, error) {
	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now
	if item.Status == "" {
		item.Status = course.StatusDraft
	}
	if item.PublicStatus == "" {
		item.PublicStatus = course.PublicStatusDraft
	}
	_, err := r.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	if err != nil {
		return nil, errx.WrapWithTypeOnCodes(err, errx.T_Conflict, course.CodeCourseSlugAlreadyUsed)
	}
	return item, nil
}

func (r *courseRepo) Update(ctx context.Context, item *course.Course) (*course.Course, error) {
	item.UpdatedAt = time.Now()
	res, err := r.db.NewUpdate().Model(item).
		Column("title", "description", "category", "level", "status", "public_status", "updated_at").
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, errx.WrapWithTypeOnCodes(err, errx.T_Conflict, course.CodeCourseSlugAlreadyUsed)
	}
	rows, rowsErr := res.RowsAffected()
	if rowsErr != nil {
		return nil, errx.Wrap(rowsErr)
	}
	if rows == 0 {
		return nil, errx.New("course not found", errx.WithType(errx.T_NotFound), errx.WithCode(course.CodeCourseNotFound))
	}
	return item, nil
}

func (r *courseRepo) Get(ctx context.Context, filter course.Filter) (*course.Course, error) {
	item := new(course.Course)
	q := r.db.NewSelect().Model(item).Limit(1)
	applyFilter(q, filter)
	err := q.Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errx.New("course not found", errx.WithType(errx.T_NotFound), errx.WithCode(course.CodeCourseNotFound))
	}
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return item, nil
}

func (r *courseRepo) List(ctx context.Context, filter course.Filter) ([]course.Course, error) {
	items := make([]course.Course, 0)
	q := r.db.NewSelect().Model(&items).Order("created_at DESC")
	applyFilter(q, filter)
	if filter.Limit > 0 {
		q = q.Limit(filter.Limit)
	}
	if err := q.Scan(ctx); err != nil {
		return nil, errx.Wrap(err)
	}
	return items, nil
}

func applyFilter(q *bun.SelectQuery, filter course.Filter) {
	if filter.ID != nil {
		q.Where("id = ?", *filter.ID)
	}
	if filter.OrganizationID != nil {
		q.Where("organization_id = ?", *filter.OrganizationID)
	}
	if filter.Slug != nil {
		q.Where("slug = ?", *filter.Slug)
	}
	if filter.PublicStatus != nil {
		q.Where("public_status = ?", *filter.PublicStatus)
	}
}
