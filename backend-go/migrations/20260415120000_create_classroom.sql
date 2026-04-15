-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS classroom;

CREATE TABLE IF NOT EXISTS classroom.classes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    course_id UUID NOT NULL,
    name TEXT NOT NULL,
    start_date DATE,
    lesson_cadence TEXT NOT NULL DEFAULT 'every_other_day',
    status TEXT NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT classes_lesson_cadence_check CHECK (lesson_cadence IN ('daily', 'every_other_day', 'weekly_3', 'manual')),
    CONSTRAINT classes_status_check CHECK (status IN ('active', 'paused', 'archived'))
);

CREATE INDEX IF NOT EXISTS idx_classes_organization_course_created_at
    ON classroom.classes (organization_id, course_id, created_at DESC);

ALTER TABLE classroom.classes
    ADD CONSTRAINT fk_classes_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;

ALTER TABLE classroom.classes
    ADD CONSTRAINT fk_classes_course
    FOREIGN KEY (course_id) REFERENCES catalog.courses (id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS classroom.class_mentors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    class_id UUID NOT NULL,
    mentor_user_id VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT class_mentors_class_user_unique UNIQUE (class_id, mentor_user_id)
);

CREATE INDEX IF NOT EXISTS idx_class_mentors_mentor_user_id
    ON classroom.class_mentors (mentor_user_id);

ALTER TABLE classroom.class_mentors
    ADD CONSTRAINT fk_class_mentors_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;

ALTER TABLE classroom.class_mentors
    ADD CONSTRAINT fk_class_mentors_class
    FOREIGN KEY (class_id) REFERENCES classroom.classes (id) ON DELETE CASCADE;

ALTER TABLE classroom.class_mentors
    ADD CONSTRAINT fk_class_mentors_user
    FOREIGN KEY (mentor_user_id) REFERENCES auth.users (id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS classroom.enrollments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    class_id UUID NOT NULL,
    student_user_id VARCHAR NOT NULL,
    status TEXT NOT NULL DEFAULT 'active',
    enrolled_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT enrollments_status_check CHECK (status IN ('active', 'removed')),
    CONSTRAINT enrollments_class_student_unique UNIQUE (class_id, student_user_id)
);

CREATE INDEX IF NOT EXISTS idx_enrollments_student_user_id
    ON classroom.enrollments (student_user_id);

ALTER TABLE classroom.enrollments
    ADD CONSTRAINT fk_enrollments_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;

ALTER TABLE classroom.enrollments
    ADD CONSTRAINT fk_enrollments_class
    FOREIGN KEY (class_id) REFERENCES classroom.classes (id) ON DELETE CASCADE;

ALTER TABLE classroom.enrollments
    ADD CONSTRAINT fk_enrollments_user
    FOREIGN KEY (student_user_id) REFERENCES auth.users (id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS classroom.access_grants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    class_id UUID NOT NULL,
    student_user_id VARCHAR NOT NULL,
    access_status TEXT NOT NULL DEFAULT 'pending',
    payment_status TEXT NOT NULL DEFAULT 'unknown',
    note TEXT,
    granted_by VARCHAR,
    granted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT access_grants_access_status_check CHECK (access_status IN ('pending', 'active', 'paused', 'blocked')),
    CONSTRAINT access_grants_payment_status_check CHECK (payment_status IN ('unknown', 'pending', 'confirmed', 'rejected')),
    CONSTRAINT access_grants_class_student_unique UNIQUE (class_id, student_user_id)
);

ALTER TABLE classroom.access_grants
    ADD CONSTRAINT fk_access_grants_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;

ALTER TABLE classroom.access_grants
    ADD CONSTRAINT fk_access_grants_class
    FOREIGN KEY (class_id) REFERENCES classroom.classes (id) ON DELETE CASCADE;

ALTER TABLE classroom.access_grants
    ADD CONSTRAINT fk_access_grants_student_user
    FOREIGN KEY (student_user_id) REFERENCES auth.users (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS classroom.access_grants DROP CONSTRAINT IF EXISTS fk_access_grants_student_user;
ALTER TABLE IF EXISTS classroom.access_grants DROP CONSTRAINT IF EXISTS fk_access_grants_class;
ALTER TABLE IF EXISTS classroom.access_grants DROP CONSTRAINT IF EXISTS fk_access_grants_organization;
DROP TABLE IF EXISTS classroom.access_grants;
ALTER TABLE IF EXISTS classroom.enrollments DROP CONSTRAINT IF EXISTS fk_enrollments_user;
ALTER TABLE IF EXISTS classroom.enrollments DROP CONSTRAINT IF EXISTS fk_enrollments_class;
ALTER TABLE IF EXISTS classroom.enrollments DROP CONSTRAINT IF EXISTS fk_enrollments_organization;
DROP TABLE IF EXISTS classroom.enrollments;
ALTER TABLE IF EXISTS classroom.class_mentors DROP CONSTRAINT IF EXISTS fk_class_mentors_user;
ALTER TABLE IF EXISTS classroom.class_mentors DROP CONSTRAINT IF EXISTS fk_class_mentors_class;
ALTER TABLE IF EXISTS classroom.class_mentors DROP CONSTRAINT IF EXISTS fk_class_mentors_organization;
DROP TABLE IF EXISTS classroom.class_mentors;
ALTER TABLE IF EXISTS classroom.classes DROP CONSTRAINT IF EXISTS fk_classes_course;
ALTER TABLE IF EXISTS classroom.classes DROP CONSTRAINT IF EXISTS fk_classes_organization;
DROP TABLE IF EXISTS classroom.classes;
DROP SCHEMA IF EXISTS classroom;
-- +goose StatementEnd
