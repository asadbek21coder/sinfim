package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/modules/organization/domain/schoolrequest"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type schoolRequestRepo struct {
	db bun.IDB
}

func NewSchoolRequestRepo(db bun.IDB) schoolrequest.Repo {
	return &schoolRequestRepo{db: db}
}

func (r *schoolRequestRepo) Create(ctx context.Context, request *schoolrequest.SchoolRequest) (*schoolrequest.SchoolRequest, error) {
	now := time.Now()
	request.CreatedAt = now
	request.UpdatedAt = now
	if request.Status == "" {
		request.Status = schoolrequest.StatusNew
	}

	_, err := r.db.NewInsert().Model(request).Returning("*").Exec(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return request, nil
}

func (r *schoolRequestRepo) FindOpenDuplicate(ctx context.Context, phoneNumber string, schoolName string) (*schoolrequest.SchoolRequest, error) {
	request := new(schoolrequest.SchoolRequest)
	err := r.db.NewSelect().
		Model(request).
		Where("phone_number = ?", phoneNumber).
		Where("lower(school_name) = lower(?)", schoolName).
		Where("status IN (?)", bun.In([]string{schoolrequest.StatusNew, schoolrequest.StatusContacted})).
		Order("created_at DESC").
		Limit(1).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return request, nil
}

func (r *schoolRequestRepo) List(ctx context.Context, filter schoolrequest.Filter) ([]schoolrequest.SchoolRequest, error) {
	requests := make([]schoolrequest.SchoolRequest, 0)
	q := r.db.NewSelect().Model(&requests).Order("created_at DESC")
	if filter.ID != nil {
		q = q.Where("id = ?", *filter.ID)
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
	return requests, nil
}

func (r *schoolRequestRepo) UpdateStatus(ctx context.Context, id string, status string) (*schoolrequest.SchoolRequest, error) {
	request := &schoolrequest.SchoolRequest{ID: id}
	res, err := r.db.NewUpdate().
		Model(request).
		Set("status = ?", status).
		Set("updated_at = ?", time.Now()).
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if rows == 0 {
		return nil, errx.New(
			"school request not found",
			errx.WithType(errx.T_NotFound),
			errx.WithCode(schoolrequest.CodeSchoolRequestNotFound),
		)
	}
	return request, nil
}
