package accessgrant

import "context"

type Repo interface {
	Upsert(ctx context.Context, item *AccessGrant) (*AccessGrant, error)
}
