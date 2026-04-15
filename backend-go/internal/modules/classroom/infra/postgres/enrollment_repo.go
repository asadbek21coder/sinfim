package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/modules/classroom/domain/enrollment"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type enrollmentRepo struct{ db bun.IDB }

func NewEnrollmentRepo(db bun.IDB) enrollment.Repo { return &enrollmentRepo{db: db} }

func (r *enrollmentRepo) Create(ctx context.Context, item *enrollment.Enrollment) (*enrollment.Enrollment, error) {
	now := time.Now()
	item.EnrolledAt = now
	item.CreatedAt = now
	item.UpdatedAt = now
	if item.Status == "" {
		item.Status = enrollment.StatusActive
	}
	_, err := r.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	if err != nil {
		return nil, errx.WrapWithTypeOnCodes(err, errx.T_Conflict, enrollment.CodeStudentAlreadyEnrolled)
	}
	return item, nil
}

func (r *enrollmentRepo) Get(ctx context.Context, classID string, studentUserID string) (*enrollment.Enrollment, error) {
	item := new(enrollment.Enrollment)
	err := r.db.NewSelect().Model(item).Where("class_id = ?", classID).Where("student_user_id = ?", studentUserID).Limit(1).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errx.New("enrollment not found", errx.WithType(errx.T_NotFound), errx.WithCode(enrollment.CodeEnrollmentNotFound))
	}
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return item, nil
}

func (r *enrollmentRepo) ListStudents(ctx context.Context, classID string) ([]enrollment.StudentDTO, error) {
	items := make([]enrollment.StudentDTO, 0)
	err := r.db.NewSelect().TableExpr("classroom.enrollments AS e").
		ColumnExpr("e.id AS enrollment_id").
		ColumnExpr("e.student_user_id AS student_user_id").
		ColumnExpr("COALESCE(u.full_name, '') AS full_name").
		ColumnExpr("COALESCE(u.phone_number, '') AS phone_number").
		ColumnExpr("e.status AS status").
		ColumnExpr("COALESCE(ag.access_status, 'pending') AS access_status").
		ColumnExpr("COALESCE(ag.payment_status, 'unknown') AS payment_status").
		ColumnExpr("ag.note AS note").
		ColumnExpr("e.enrolled_at AS enrolled_at").
		ColumnExpr("ag.granted_at AS granted_at").
		Join("JOIN auth.users AS u ON u.id = e.student_user_id").
		Join("LEFT JOIN classroom.access_grants AS ag ON ag.class_id = e.class_id AND ag.student_user_id = e.student_user_id").
		Where("e.class_id = ?", classID).
		Order("e.created_at DESC").
		Scan(ctx, &items)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return items, nil
}
