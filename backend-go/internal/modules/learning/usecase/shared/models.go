package shared

import "time"

const (
	CodeAccessDenied       = "ACCESS_DENIED"
	CodeEnrollmentNotFound = "ENROLLMENT_NOT_FOUND"
	CodeLessonLocked       = "LESSON_LOCKED"
	CodeLessonNotFound     = "LESSON_NOT_FOUND"
)

type ClassContext struct {
	OrganizationID   string     `bun:"organization_id" json:"organization_id"`
	Organization     string     `bun:"organization" json:"organization"`
	OrganizationSlug string     `bun:"organization_slug" json:"organization_slug"`
	LogoURL          *string    `bun:"logo_url" json:"logo_url"`
	ClassID          string     `bun:"class_id" json:"class_id"`
	ClassName        string     `bun:"class_name" json:"class_name"`
	StartDate        *time.Time `bun:"start_date" json:"start_date"`
	LessonCadence    string     `bun:"lesson_cadence" json:"lesson_cadence"`
	CourseID         string     `bun:"course_id" json:"course_id"`
	CourseTitle      string     `bun:"course_title" json:"course_title"`
	AccessStatus     string     `bun:"access_status" json:"access_status"`
	PaymentStatus    string     `bun:"payment_status" json:"payment_status"`
}

type LessonRow struct {
	ID             string     `bun:"id" json:"id"`
	OrganizationID string     `bun:"organization_id" json:"organization_id"`
	CourseID       string     `bun:"course_id" json:"course_id"`
	Title          string     `bun:"title" json:"title"`
	Description    *string    `bun:"description" json:"description"`
	OrderNumber    int        `bun:"order_number" json:"order_number"`
	PublishDay     int        `bun:"publish_day" json:"publish_day"`
	Status         string     `bun:"status" json:"status"`
	HasVideo       bool       `bun:"has_video" json:"has_video"`
	MaterialCount  int        `bun:"material_count" json:"material_count"`
	CompletedAt    *time.Time `bun:"completed_at" json:"completed_at"`
}

type VideoDTO struct {
	Provider          string  `json:"provider" bun:"provider"`
	StreamRef         *string `json:"stream_ref" bun:"stream_ref"`
	TelegramChannelID *string `json:"telegram_channel_id" bun:"telegram_channel_id"`
	TelegramMessageID *string `json:"telegram_message_id" bun:"telegram_message_id"`
	EmbedURL          *string `json:"embed_url" bun:"embed_url"`
	DurationSeconds   *int    `json:"duration_seconds" bun:"duration_seconds"`
}

type MaterialDTO struct {
	ID           string  `json:"id" bun:"id"`
	Title        string  `json:"title" bun:"title"`
	MaterialType string  `json:"material_type" bun:"material_type"`
	SourceType   string  `json:"source_type" bun:"source_type"`
	URL          *string `json:"url" bun:"url"`
	FileID       *string `json:"file_id" bun:"file_id"`
	OrderNumber  int     `json:"order_number" bun:"order_number"`
}

type Availability struct {
	Status      string     `json:"status"`
	AvailableAt *time.Time `json:"available_at"`
}

func ComputeAvailability(startDate *time.Time, cadence string, publishDay int, lessonStatus string, now time.Time) Availability {
	if lessonStatus != "published" {
		return Availability{Status: "locked"}
	}
	if startDate == nil {
		return Availability{Status: "available"}
	}
	stepDays := 1
	switch cadence {
	case "every_other_day", "weekly_3":
		stepDays = 2
	case "manual":
		stepDays = 0
	}
	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, now.Location())
	availableAt := start.AddDate(0, 0, (publishDay-1)*stepDays)
	if stepDays == 0 {
		availableAt = start
	}
	if now.Before(availableAt) {
		return Availability{Status: "locked", AvailableAt: &availableAt}
	}
	return Availability{Status: "available", AvailableAt: &availableAt}
}
