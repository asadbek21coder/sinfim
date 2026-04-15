-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS catalog;

CREATE TABLE IF NOT EXISTS catalog.courses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    title TEXT NOT NULL,
    slug TEXT NOT NULL,
    description TEXT,
    category TEXT,
    level TEXT,
    status TEXT NOT NULL DEFAULT 'draft',
    public_status TEXT NOT NULL DEFAULT 'draft',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT courses_status_check CHECK (status IN ('draft', 'active', 'archived')),
    CONSTRAINT courses_public_status_check CHECK (public_status IN ('draft', 'public', 'hidden')),
    CONSTRAINT courses_organization_slug_unique UNIQUE (organization_id, slug)
);

CREATE INDEX IF NOT EXISTS idx_courses_organization_created_at
    ON catalog.courses (organization_id, created_at DESC);

CREATE INDEX IF NOT EXISTS idx_courses_organization_public_status
    ON catalog.courses (organization_id, public_status);

ALTER TABLE catalog.courses
    ADD CONSTRAINT fk_courses_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS catalog.courses DROP CONSTRAINT IF EXISTS fk_courses_organization;
DROP TABLE IF EXISTS catalog.courses;
DROP SCHEMA IF EXISTS catalog;
-- +goose StatementEnd
