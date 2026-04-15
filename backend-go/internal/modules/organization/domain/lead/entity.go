package lead

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	StatusNew       = "new"
	StatusContacted = "contacted"
	StatusConverted = "converted"
	StatusArchived  = "archived"

	SourcePublicSchoolPage = "public_school_page"

	CodeLeadNotFound = "LEAD_NOT_FOUND"
)

type Lead struct {
	bun.BaseModel `bun:"table:lead.leads,alias:l"`

	ID             string    `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrganizationID string    `json:"organization_id" bun:"organization_id,type:uuid,notnull"`
	FullName       string    `json:"full_name" bun:"full_name,notnull"`
	PhoneNumber    string    `json:"phone_number" bun:"phone_number,notnull"`
	Note           *string   `json:"note" bun:"note"`
	Source         string    `json:"source" bun:"source,notnull"`
	Status         string    `json:"status" bun:"status,notnull"`
	CreatedAt      time.Time `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `json:"updated_at" bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Filter struct {
	ID             *string
	OrganizationID *string
	Status         *string
	Limit          int
}
