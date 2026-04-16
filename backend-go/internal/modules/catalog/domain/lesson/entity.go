package lesson

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	StatusDraft     = "draft"
	StatusPublished = "published"
	StatusArchived  = "archived"

	CodeLessonNotFound      = "LESSON_NOT_FOUND"
	CodeLessonOrderConflict = "LESSON_ORDER_ALREADY_USED"
)

type Lesson struct {
	bun.BaseModel `bun:"table:catalog.lessons,alias:l"`

	ID             string    `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	OrganizationID string    `bun:"organization_id,type:uuid,notnull" json:"organization_id"`
	CourseID       string    `bun:"course_id,type:uuid,notnull" json:"course_id"`
	Title          string    `bun:"title,notnull" json:"title"`
	Description    *string   `bun:"description" json:"description"`
	OrderNumber    int       `bun:"order_number,notnull" json:"order_number"`
	PublishDay     int       `bun:"publish_day,notnull" json:"publish_day"`
	Status         string    `bun:"status,notnull" json:"status"`
	CreatedAt      time.Time `bun:"created_at,notnull" json:"created_at"`
	UpdatedAt      time.Time `bun:"updated_at,notnull" json:"updated_at"`
}

type Summary struct {
	Lesson
	HasVideo      bool `json:"has_video" bun:"has_video"`
	MaterialCount int  `json:"material_count" bun:"material_count"`
}

type Filter struct {
	ID             *string
	OrganizationID *string
	CourseID       *string
	Status         *string
	Limit          int
}
