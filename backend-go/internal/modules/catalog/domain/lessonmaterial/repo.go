package lessonmaterial

import "context"

type Repo interface {
	ReplaceByLesson(ctx context.Context, organizationID string, lessonID string, items []ReplaceItem) ([]LessonMaterial, error)
	ListByLesson(ctx context.Context, lessonID string) ([]LessonMaterial, error)
}
