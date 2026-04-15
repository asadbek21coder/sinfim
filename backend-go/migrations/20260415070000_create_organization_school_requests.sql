-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS organization;

CREATE TABLE IF NOT EXISTS organization.school_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    school_name TEXT NOT NULL,
    category TEXT,
    student_count INTEGER,
    note TEXT,
    status TEXT NOT NULL DEFAULT 'new',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT school_requests_status_check CHECK (status IN ('new', 'contacted', 'approved', 'rejected')),
    CONSTRAINT school_requests_student_count_check CHECK (student_count IS NULL OR student_count >= 0)
);

CREATE INDEX IF NOT EXISTS idx_school_requests_status_created_at
    ON organization.school_requests (status, created_at DESC);

CREATE UNIQUE INDEX IF NOT EXISTS idx_school_requests_open_phone_school_unique
    ON organization.school_requests (phone_number, lower(school_name))
    WHERE status IN ('new', 'contacted');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS organization.school_requests;
DROP SCHEMA IF EXISTS organization;
-- +goose StatementEnd
