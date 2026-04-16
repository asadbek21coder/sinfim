package lessonvideo

import "context"

type Repo interface {
	Upsert(ctx context.Context, item *LessonVideo) (*LessonVideo, error)
	GetByLesson(ctx context.Context, lessonID string) (*LessonVideo, error)
	DeleteByLesson(ctx context.Context, lessonID string) error
}
