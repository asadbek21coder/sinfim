-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS organization.organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,
    description TEXT,
    logo_url TEXT,
    category TEXT,
    contact_phone TEXT,
    telegram_url TEXT,
    public_status TEXT NOT NULL DEFAULT 'draft',
    is_demo BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT organizations_public_status_check CHECK (public_status IN ('draft', 'public', 'hidden'))
);

CREATE INDEX IF NOT EXISTS idx_organizations_slug ON organization.organizations (slug);

CREATE TABLE IF NOT EXISTS auth.user_memberships (
    id BIGSERIAL PRIMARY KEY,
    user_id VARCHAR NOT NULL,
    organization_id UUID NOT NULL,
    role VARCHAR NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT user_memberships_role_check CHECK (role IN ('OWNER', 'TEACHER', 'MENTOR', 'STUDENT')),
    CONSTRAINT user_memberships_user_org_role_unique UNIQUE (user_id, organization_id, role)
);

CREATE INDEX IF NOT EXISTS idx_user_memberships_user_id ON auth.user_memberships (user_id);
CREATE INDEX IF NOT EXISTS idx_user_memberships_organization_id ON auth.user_memberships (organization_id);

ALTER TABLE auth.user_memberships
    ADD CONSTRAINT fk_user_memberships_user
    FOREIGN KEY (user_id) REFERENCES auth.users (id) ON DELETE CASCADE;

ALTER TABLE auth.user_memberships
    ADD CONSTRAINT fk_user_memberships_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS auth.user_memberships DROP CONSTRAINT IF EXISTS fk_user_memberships_organization;
ALTER TABLE IF EXISTS auth.user_memberships DROP CONSTRAINT IF EXISTS fk_user_memberships_user;
DROP TABLE IF EXISTS auth.user_memberships;
DROP TABLE IF EXISTS organization.organizations;
-- +goose StatementEnd
