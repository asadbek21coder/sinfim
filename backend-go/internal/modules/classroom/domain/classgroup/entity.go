package classgroup

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	CadenceDaily         = "daily"
	CadenceEveryOtherDay = "every_other_day"
	CadenceWeekly3       = "weekly_3"
	CadenceManual        = "manual"

	StatusActive   = "active"
	StatusPaused   = "paused"
	StatusArchived = "archived"

	CodeClassNotFound = "CLASS_NOT_FOUND"
)

type Class struct {
	bun.BaseModel `bun:"table:classroom.classes,alias:cl"`

	ID             string     `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrganizationID string     `json:"organization_id" bun:"organization_id,type:uuid,notnull"`
	CourseID       string     `json:"course_id" bun:"course_id,type:uuid,notnull"`
	Name           string     `json:"name" bun:"name,notnull"`
	StartDate      *time.Time `json:"start_date" bun:"start_date"`
	LessonCadence  string     `json:"lesson_cadence" bun:"lesson_cadence,notnull"`
	Status         string     `json:"status" bun:"status,notnull"`
	CreatedAt      time.Time  `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `json:"updated_at" bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Summary struct {
	Class
	CourseTitle  string `json:"course_title" bun:"course_title"`
	MentorCount  int    `json:"mentor_count" bun:"mentor_count"`
	StudentCount int    `json:"student_count" bun:"student_count"`
}

type Filter struct {
	ID             *string
	OrganizationID *string
	CourseID       *string
	MentorUserID   *string
	Limit          int
}
