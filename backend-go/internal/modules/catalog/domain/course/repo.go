package course

import "context"

type Repo interface {
	Create(ctx context.Context, item *Course) (*Course, error)
	Update(ctx context.Context, item *Course) (*Course, error)
	Get(ctx context.Context, filter Filter) (*Course, error)
	List(ctx context.Context, filter Filter) ([]Course, error)
}
