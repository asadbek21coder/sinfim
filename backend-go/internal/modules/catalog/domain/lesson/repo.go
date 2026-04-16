package lesson

import "context"

type Repo interface {
	Create(ctx context.Context, item *Lesson) (*Lesson, error)
	Update(ctx context.Context, item *Lesson) (*Lesson, error)
	Get(ctx context.Context, filter Filter) (*Lesson, error)
	List(ctx context.Context, filter Filter) ([]Summary, error)
}
