package course

import "time"

import "github.com/uptrace/bun"

const (
	StatusDraft    = "draft"
	StatusActive   = "active"
	StatusArchived = "archived"

	PublicStatusDraft  = "draft"
	PublicStatusPublic = "public"
	PublicStatusHidden = "hidden"

	CodeCourseNotFound        = "COURSE_NOT_FOUND"
	CodeCourseSlugAlreadyUsed = "COURSE_SLUG_ALREADY_TAKEN"
)

type Course struct {
	bun.BaseModel `bun:"table:catalog.courses,alias:c"`

	ID             string    `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	OrganizationID string    `bun:"organization_id,type:uuid,notnull" json:"organization_id"`
	Title          string    `bun:"title,notnull" json:"title"`
	Slug           string    `bun:"slug,notnull" json:"slug"`
	Description    *string   `bun:"description" json:"description"`
	Category       *string   `bun:"category" json:"category"`
	Level          *string   `bun:"level" json:"level"`
	Status         string    `bun:"status,notnull" json:"status"`
	PublicStatus   string    `bun:"public_status,notnull" json:"public_status"`
	CreatedAt      time.Time `bun:"created_at,notnull" json:"created_at"`
	UpdatedAt      time.Time `bun:"updated_at,notnull" json:"updated_at"`
}

type Filter struct {
	ID             *string
	OrganizationID *string
	Slug           *string
	PublicStatus   *string
	Limit          int
}
