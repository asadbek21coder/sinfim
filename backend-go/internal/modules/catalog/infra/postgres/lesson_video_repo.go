package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-enterprise-blueprint/internal/modules/catalog/domain/lessonvideo"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type lessonVideoRepo struct{ db bun.IDB }

func NewLessonVideoRepo(db bun.IDB) lessonvideo.Repo { return &lessonVideoRepo{db: db} }

func (r *lessonVideoRepo) Upsert(ctx context.Context, item *lessonvideo.LessonVideo) (*lessonvideo.LessonVideo, error) {
	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now
	if item.Provider == "" {
		item.Provider = lessonvideo.ProviderTelegram
	}
	_, err := r.db.NewInsert().Model(item).
		On("CONFLICT (lesson_id) DO UPDATE").
		Set("provider = EXCLUDED.provider").
		Set("stream_ref = EXCLUDED.stream_ref").
		Set("telegram_channel_id = EXCLUDED.telegram_channel_id").
		Set("telegram_message_id = EXCLUDED.telegram_message_id").
		Set("embed_url = EXCLUDED.embed_url").
		Set("duration_seconds = EXCLUDED.duration_seconds").
		Set("updated_at = EXCLUDED.updated_at").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return item, nil
}

func (r *lessonVideoRepo) GetByLesson(ctx context.Context, lessonID string) (*lessonvideo.LessonVideo, error) {
	item := new(lessonvideo.LessonVideo)
	err := r.db.NewSelect().Model(item).Where("lesson_id = ?", lessonID).Limit(1).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return item, nil
}

func (r *lessonVideoRepo) DeleteByLesson(ctx context.Context, lessonID string) error {
	_, err := r.db.NewDelete().Model((*lessonvideo.LessonVideo)(nil)).Where("lesson_id = ?", lessonID).Exec(ctx)
	return errx.Wrap(err)
}
