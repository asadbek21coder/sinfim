package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/modules/classroom/domain/classgroup"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type classRepo struct{ db bun.IDB }

func NewClassRepo(db bun.IDB) classgroup.Repo { return &classRepo{db: db} }

func (r *classRepo) Create(ctx context.Context, item *classgroup.Class) (*classgroup.Class, error) {
	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now
	if item.LessonCadence == "" {
		item.LessonCadence = classgroup.CadenceEveryOtherDay
	}
	if item.Status == "" {
		item.Status = classgroup.StatusActive
	}
	_, err := r.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return item, nil
}

func (r *classRepo) Update(ctx context.Context, item *classgroup.Class) (*classgroup.Class, error) {
	item.UpdatedAt = time.Now()
	res, err := r.db.NewUpdate().Model(item).Column("name", "start_date", "lesson_cadence", "status", "updated_at").WherePK().Returning("*").Exec(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	rows, rowsErr := res.RowsAffected()
	if rowsErr != nil {
		return nil, errx.Wrap(rowsErr)
	}
	if rows == 0 {
		return nil, errx.New("class not found", errx.WithType(errx.T_NotFound), errx.WithCode(classgroup.CodeClassNotFound))
	}
	return item, nil
}

func (r *classRepo) Get(ctx context.Context, filter classgroup.Filter) (*classgroup.Class, error) {
	item := new(classgroup.Class)
	q := r.db.NewSelect().Model(item).Limit(1)
	applyClassFilter(q, filter)
	err := q.Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errx.New("class not found", errx.WithType(errx.T_NotFound), errx.WithCode(classgroup.CodeClassNotFound))
	}
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return item, nil
}

func (r *classRepo) List(ctx context.Context, filter classgroup.Filter) ([]classgroup.Summary, error) {
	items := make([]classgroup.Summary, 0)
	q := r.db.NewSelect().TableExpr("classroom.classes AS cl").
		ColumnExpr("cl.*").
		ColumnExpr("co.title AS course_title").
		ColumnExpr("COUNT(DISTINCT cm.id) AS mentor_count").
		ColumnExpr("COUNT(DISTINCT e.id) AS student_count").
		Join("JOIN catalog.courses AS co ON co.id = cl.course_id").
		Join("LEFT JOIN classroom.class_mentors AS cm ON cm.class_id = cl.id").
		Join("LEFT JOIN classroom.enrollments AS e ON e.class_id = cl.id AND e.status = ?", "active").
		Group("cl.id", "co.title").
		Order("cl.created_at DESC")
	applyClassFilter(q, filter)
	if filter.MentorUserID != nil {
		q = q.Where("cm.mentor_user_id = ?", *filter.MentorUserID)
	}
	if filter.Limit > 0 {
		q = q.Limit(filter.Limit)
	}
	if err := q.Scan(ctx, &items); err != nil {
		return nil, errx.Wrap(err)
	}
	return items, nil
}

func applyClassFilter(q *bun.SelectQuery, filter classgroup.Filter) {
	if filter.ID != nil {
		q.Where("cl.id = ?", *filter.ID)
	}
	if filter.OrganizationID != nil {
		q.Where("cl.organization_id = ?", *filter.OrganizationID)
	}
	if filter.CourseID != nil {
		q.Where("cl.course_id = ?", *filter.CourseID)
	}
}
