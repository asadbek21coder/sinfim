package enrollment

import "context"

type Repo interface {
	Create(ctx context.Context, item *Enrollment) (*Enrollment, error)
	Get(ctx context.Context, classID string, studentUserID string) (*Enrollment, error)
	ListStudents(ctx context.Context, classID string) ([]StudentDTO, error)
}
