package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/modules/classroom/domain/classmentor"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type mentorRepo struct{ db bun.IDB }

func NewMentorRepo(db bun.IDB) classmentor.Repo { return &mentorRepo{db: db} }

func (r *mentorRepo) Assign(ctx context.Context, item *classmentor.ClassMentor) (*classmentor.ClassMentor, error) {
	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now
	_, err := r.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	if err != nil {
		return nil, errx.WrapWithTypeOnCodes(err, errx.T_Conflict, classmentor.CodeMentorAlreadyAssigned)
	}
	return item, nil
}

func (r *mentorRepo) IsAssigned(ctx context.Context, classID string, mentorUserID string) (bool, error) {
	item := new(classmentor.ClassMentor)
	err := r.db.NewSelect().Model(item).Where("class_id = ?", classID).Where("mentor_user_id = ?", mentorUserID).Limit(1).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, errx.Wrap(err)
	}
	return true, nil
}

func (r *mentorRepo) ListByClass(ctx context.Context, classID string) ([]classmentor.MentorDTO, error) {
	items := make([]classmentor.MentorDTO, 0)
	err := r.db.NewSelect().TableExpr("classroom.class_mentors AS cm").
		ColumnExpr("cm.id AS id").
		ColumnExpr("u.id AS user_id").
		ColumnExpr("COALESCE(u.full_name, '') AS full_name").
		ColumnExpr("COALESCE(u.phone_number, '') AS phone_number").
		Join("JOIN auth.users AS u ON u.id = cm.mentor_user_id").
		Where("cm.class_id = ?", classID).
		Order("cm.created_at ASC").
		Scan(ctx, &items)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return items, nil
}
