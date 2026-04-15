package lead

import "context"

type Repo interface {
	Create(ctx context.Context, lead *Lead) (*Lead, error)
	List(ctx context.Context, filter Filter) ([]Lead, error)
	UpdateStatus(ctx context.Context, id string, status string) (*Lead, error)
}
