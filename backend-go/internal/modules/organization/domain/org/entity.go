package org

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	PublicStatusDraft  = "draft"
	PublicStatusPublic = "public"
	PublicStatusHidden = "hidden"

	CodeOrganizationNotFound  = "ORGANIZATION_NOT_FOUND"
	CodeOrganizationNotPublic = "ORGANIZATION_NOT_PUBLIC"
	CodeSlugAlreadyTaken      = "SLUG_ALREADY_TAKEN"
)

type Organization struct {
	bun.BaseModel `bun:"table:organization.organizations,alias:o"`

	ID           string    `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name         string    `json:"name" bun:"name,notnull"`
	Slug         string    `json:"slug" bun:"slug,notnull"`
	Description  *string   `json:"description" bun:"description"`
	LogoURL      *string   `json:"logo_url" bun:"logo_url"`
	Category     *string   `json:"category" bun:"category"`
	ContactPhone *string   `json:"contact_phone" bun:"contact_phone"`
	TelegramURL  *string   `json:"telegram_url" bun:"telegram_url"`
	PublicStatus string    `json:"public_status" bun:"public_status,notnull"`
	IsDemo       bool      `json:"is_demo" bun:"is_demo,notnull"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at" bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Filter struct {
	ID    *string
	Slug  *string
	Limit int
}
