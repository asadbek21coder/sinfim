package org

import "context"

type Repo interface {
	Create(ctx context.Context, organization *Organization) (*Organization, error)
	Update(ctx context.Context, organization *Organization) (*Organization, error)
	GetBySlug(ctx context.Context, slug string) (*Organization, error)
	List(ctx context.Context, filter Filter) ([]Organization, error)
}
