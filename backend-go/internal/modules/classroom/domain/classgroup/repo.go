package classgroup

import "context"

type Repo interface {
	Create(ctx context.Context, item *Class) (*Class, error)
	Update(ctx context.Context, item *Class) (*Class, error)
	Get(ctx context.Context, filter Filter) (*Class, error)
	List(ctx context.Context, filter Filter) ([]Summary, error)
}
