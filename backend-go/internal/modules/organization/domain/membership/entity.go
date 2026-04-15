package membership

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	RoleOwner   = "OWNER"
	RoleTeacher = "TEACHER"
	RoleMentor  = "MENTOR"
	RoleStudent = "STUDENT"

	CodeOwnerAlreadyMember = "OWNER_ALREADY_MEMBER"
)

type Membership struct {
	bun.BaseModel `bun:"table:auth.user_memberships,alias:um"`

	ID             int64     `json:"id" bun:"id,pk,autoincrement"`
	UserID         string    `json:"user_id" bun:"user_id,notnull"`
	OrganizationID string    `json:"organization_id" bun:"organization_id,type:uuid,notnull"`
	Role           string    `json:"role" bun:"role,notnull"`
	IsActive       bool      `json:"is_active" bun:"is_active,notnull"`
	CreatedAt      time.Time `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `json:"updated_at" bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
