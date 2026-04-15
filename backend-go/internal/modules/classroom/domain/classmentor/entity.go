package classmentor

import (
	"time"

	"github.com/uptrace/bun"
)

const CodeMentorAlreadyAssigned = "MENTOR_ALREADY_ASSIGNED"

type ClassMentor struct {
	bun.BaseModel `bun:"table:classroom.class_mentors,alias:cm"`

	ID             string    `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrganizationID string    `json:"organization_id" bun:"organization_id,type:uuid,notnull"`
	ClassID        string    `json:"class_id" bun:"class_id,type:uuid,notnull"`
	MentorUserID   string    `json:"mentor_user_id" bun:"mentor_user_id,notnull"`
	CreatedAt      time.Time `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `json:"updated_at" bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type MentorDTO struct {
	ID          string `json:"id" bun:"id"`
	UserID      string `json:"user_id" bun:"user_id"`
	FullName    string `json:"full_name" bun:"full_name"`
	PhoneNumber string `json:"phone_number" bun:"phone_number"`
}
