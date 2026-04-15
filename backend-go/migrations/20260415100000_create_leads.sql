-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS lead;

CREATE TABLE IF NOT EXISTS lead.leads (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    full_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    note TEXT,
    source TEXT NOT NULL DEFAULT 'public_school_page',
    status TEXT NOT NULL DEFAULT 'new',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT leads_status_check CHECK (status IN ('new', 'contacted', 'converted', 'archived'))
);

CREATE INDEX IF NOT EXISTS idx_leads_organization_status_created_at
    ON lead.leads (organization_id, status, created_at DESC);

ALTER TABLE lead.leads
    ADD CONSTRAINT fk_leads_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS lead.leads DROP CONSTRAINT IF EXISTS fk_leads_organization;
DROP TABLE IF EXISTS lead.leads;
DROP SCHEMA IF EXISTS lead;
-- +goose StatementEnd
