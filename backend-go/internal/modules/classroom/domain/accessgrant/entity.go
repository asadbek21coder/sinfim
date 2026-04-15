package accessgrant

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	AccessPending = "pending"
	AccessActive  = "active"
	AccessPaused  = "paused"
	AccessBlocked = "blocked"

	PaymentUnknown   = "unknown"
	PaymentPending   = "pending"
	PaymentConfirmed = "confirmed"
	PaymentRejected  = "rejected"
)

type AccessGrant struct {
	bun.BaseModel `bun:"table:classroom.access_grants,alias:ag"`

	ID             string     `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrganizationID string     `json:"organization_id" bun:"organization_id,type:uuid,notnull"`
	ClassID        string     `json:"class_id" bun:"class_id,type:uuid,notnull"`
	StudentUserID  string     `json:"student_user_id" bun:"student_user_id,notnull"`
	AccessStatus   string     `json:"access_status" bun:"access_status,notnull"`
	PaymentStatus  string     `json:"payment_status" bun:"payment_status,notnull"`
	Note           *string    `json:"note" bun:"note"`
	GrantedBy      *string    `json:"granted_by" bun:"granted_by"`
	GrantedAt      *time.Time `json:"granted_at" bun:"granted_at"`
	CreatedAt      time.Time  `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `json:"updated_at" bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
