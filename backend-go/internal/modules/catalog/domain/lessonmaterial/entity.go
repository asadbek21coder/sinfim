package lessonmaterial

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	TypePDF   = "pdf"
	TypeImage = "image"
	TypeDoc   = "doc"
	TypeLink  = "link"
	TypeOther = "other"

	SourceURL       = "url"
	SourceFilevault = "filevault"
)

type LessonMaterial struct {
	bun.BaseModel `bun:"table:catalog.lesson_materials,alias:lm"`

	ID             string    `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	OrganizationID string    `bun:"organization_id,type:uuid,notnull" json:"organization_id"`
	LessonID       string    `bun:"lesson_id,type:uuid,notnull" json:"lesson_id"`
	Title          string    `bun:"title,notnull" json:"title"`
	MaterialType   string    `bun:"material_type,notnull" json:"material_type"`
	SourceType     string    `bun:"source_type,notnull" json:"source_type"`
	URL            *string   `bun:"url" json:"url"`
	FileID         *string   `bun:"file_id,type:uuid" json:"file_id"`
	OrderNumber    int       `bun:"order_number,notnull" json:"order_number"`
	CreatedAt      time.Time `bun:"created_at,notnull" json:"created_at"`
	UpdatedAt      time.Time `bun:"updated_at,notnull" json:"updated_at"`
}

type ReplaceItem struct {
	Title        string  `json:"title"`
	MaterialType string  `json:"material_type"`
	SourceType   string  `json:"source_type"`
	URL          *string `json:"url"`
	FileID       *string `json:"file_id"`
	OrderNumber  int     `json:"order_number"`
}
