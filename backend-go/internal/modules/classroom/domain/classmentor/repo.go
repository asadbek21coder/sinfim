package classmentor

import "context"

type Repo interface {
	Assign(ctx context.Context, item *ClassMentor) (*ClassMentor, error)
	IsAssigned(ctx context.Context, classID string, mentorUserID string) (bool, error)
	ListByClass(ctx context.Context, classID string) ([]MentorDTO, error)
}
