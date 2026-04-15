package enrollment

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	StatusActive  = "active"
	StatusRemoved = "removed"

	CodeEnrollmentNotFound     = "ENROLLMENT_NOT_FOUND"
	CodeStudentAlreadyEnrolled = "STUDENT_ALREADY_ENROLLED"
)

type Enrollment struct {
	bun.BaseModel `bun:"table:classroom.enrollments,alias:e"`

	ID             string    `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrganizationID string    `json:"organization_id" bun:"organization_id,type:uuid,notnull"`
	ClassID        string    `json:"class_id" bun:"class_id,type:uuid,notnull"`
	StudentUserID  string    `json:"student_user_id" bun:"student_user_id,notnull"`
	Status         string    `json:"status" bun:"status,notnull"`
	EnrolledAt     time.Time `json:"enrolled_at" bun:"enrolled_at,nullzero,notnull,default:current_timestamp"`
	CreatedAt      time.Time `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `json:"updated_at" bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type StudentDTO struct {
	EnrollmentID  string     `json:"enrollment_id" bun:"enrollment_id"`
	StudentUserID string     `json:"student_user_id" bun:"student_user_id"`
	FullName      string     `json:"full_name" bun:"full_name"`
	PhoneNumber   string     `json:"phone_number" bun:"phone_number"`
	Status        string     `json:"status" bun:"status"`
	AccessStatus  string     `json:"access_status" bun:"access_status"`
	PaymentStatus string     `json:"payment_status" bun:"payment_status"`
	Note          *string    `json:"note" bun:"note"`
	EnrolledAt    time.Time  `json:"enrolled_at" bun:"enrolled_at"`
	GrantedAt     *time.Time `json:"granted_at" bun:"granted_at"`
}
