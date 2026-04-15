package schoolrequest

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	StatusNew       = "new"
	StatusContacted = "contacted"
	StatusApproved  = "approved"
	StatusRejected  = "rejected"

	CodeSchoolRequestNotFound  = "SCHOOL_REQUEST_NOT_FOUND"
	CodeSchoolRequestDuplicate = "SCHOOL_REQUEST_DUPLICATE"
)

type SchoolRequest struct {
	bun.BaseModel `bun:"table:organization.school_requests,alias:sr"`

	ID           string    `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	FullName     string    `json:"full_name" bun:"full_name,notnull"`
	PhoneNumber  string    `json:"phone_number" bun:"phone_number,notnull"`
	SchoolName   string    `json:"school_name" bun:"school_name,notnull"`
	Category     *string   `json:"category" bun:"category"`
	StudentCount *int      `json:"student_count" bun:"student_count"`
	Note         *string   `json:"note" bun:"note"`
	Status       string    `json:"status" bun:"status,notnull"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at" bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Filter struct {
	ID     *string
	Status *string
	Limit  int
}
