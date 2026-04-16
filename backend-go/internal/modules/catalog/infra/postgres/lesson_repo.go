package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/modules/catalog/domain/lesson"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type lessonRepo struct{ db bun.IDB }

func NewLessonRepo(db bun.IDB) lesson.Repo { return &lessonRepo{db: db} }

func (r *lessonRepo) Create(ctx context.Context, item *lesson.Lesson) (*lesson.Lesson, error) {
	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now
	if item.Status == "" {
		item.Status = lesson.StatusDraft
	}
	if item.OrderNumber == 0 {
		item.OrderNumber = 1
	}
	if item.PublishDay == 0 {
		item.PublishDay = item.OrderNumber
	}
	_, err := r.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	if err != nil {
		return nil, errx.WrapWithTypeOnCodes(err, errx.T_Conflict, lesson.CodeLessonOrderConflict)
	}
	return item, nil
}

func (r *lessonRepo) Update(ctx context.Context, item *lesson.Lesson) (*lesson.Lesson, error) {
	item.UpdatedAt = time.Now()
	res, err := r.db.NewUpdate().Model(item).
		Column("title", "description", "order_number", "publish_day", "status", "updated_at").
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, errx.WrapWithTypeOnCodes(err, errx.T_Conflict, lesson.CodeLessonOrderConflict)
	}
	rows, rowsErr := res.RowsAffected()
	if rowsErr != nil {
		return nil, errx.Wrap(rowsErr)
	}
	if rows == 0 {
		return nil, errx.New("lesson not found", errx.WithType(errx.T_NotFound), errx.WithCode(lesson.CodeLessonNotFound))
	}
	return item, nil
}

func (r *lessonRepo) Get(ctx context.Context, filter lesson.Filter) (*lesson.Lesson, error) {
	item := new(lesson.Lesson)
	q := r.db.NewSelect().Model(item).Limit(1)
	applyLessonFilter(q, filter)
	err := q.Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errx.New("lesson not found", errx.WithType(errx.T_NotFound), errx.WithCode(lesson.CodeLessonNotFound))
	}
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return item, nil
}

func (r *lessonRepo) List(ctx context.Context, filter lesson.Filter) ([]lesson.Summary, error) {
	items := make([]lesson.Summary, 0)
	q := r.db.NewSelect().TableExpr("catalog.lessons AS l").
		ColumnExpr("l.*").
		ColumnExpr("COUNT(DISTINCT lv.id) > 0 AS has_video").
		ColumnExpr("COUNT(DISTINCT lm.id) AS material_count").
		Join("LEFT JOIN catalog.lesson_videos AS lv ON lv.lesson_id = l.id").
		Join("LEFT JOIN catalog.lesson_materials AS lm ON lm.lesson_id = l.id").
		Group("l.id").
		Order("l.order_number ASC")
	applyLessonFilter(q, filter)
	if filter.Limit > 0 {
		q = q.Limit(filter.Limit)
	}
	if err := q.Scan(ctx, &items); err != nil {
		return nil, errx.Wrap(err)
	}
	return items, nil
}

func applyLessonFilter(q *bun.SelectQuery, filter lesson.Filter) {
	if filter.ID != nil {
		q.Where("l.id = ?", *filter.ID)
	}
	if filter.OrganizationID != nil {
		q.Where("l.organization_id = ?", *filter.OrganizationID)
	}
	if filter.CourseID != nil {
		q.Where("l.course_id = ?", *filter.CourseID)
	}
	if filter.Status != nil {
		q.Where("l.status = ?", *filter.Status)
	}
}
