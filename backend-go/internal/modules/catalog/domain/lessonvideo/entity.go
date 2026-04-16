package lessonvideo

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	ProviderTelegram = "telegram"
	ProviderExternal = "external"
)

type LessonVideo struct {
	bun.BaseModel `bun:"table:catalog.lesson_videos,alias:lv"`

	ID                string    `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	OrganizationID    string    `bun:"organization_id,type:uuid,notnull" json:"organization_id"`
	LessonID          string    `bun:"lesson_id,type:uuid,notnull" json:"lesson_id"`
	Provider          string    `bun:"provider,notnull" json:"provider"`
	StreamRef         *string   `bun:"stream_ref" json:"stream_ref"`
	TelegramChannelID *string   `bun:"telegram_channel_id" json:"telegram_channel_id"`
	TelegramMessageID *string   `bun:"telegram_message_id" json:"telegram_message_id"`
	EmbedURL          *string   `bun:"embed_url" json:"embed_url"`
	DurationSeconds   *int      `bun:"duration_seconds" json:"duration_seconds"`
	CreatedAt         time.Time `bun:"created_at,notnull" json:"created_at"`
	UpdatedAt         time.Time `bun:"updated_at,notnull" json:"updated_at"`
}
