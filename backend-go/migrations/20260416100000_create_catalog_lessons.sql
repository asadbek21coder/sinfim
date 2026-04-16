-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS catalog.lessons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    course_id UUID NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    order_number INTEGER NOT NULL DEFAULT 1,
    publish_day INTEGER NOT NULL DEFAULT 1,
    status TEXT NOT NULL DEFAULT 'draft',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT lessons_status_check CHECK (status IN ('draft', 'published', 'archived')),
    CONSTRAINT lessons_order_number_check CHECK (order_number >= 1),
    CONSTRAINT lessons_publish_day_check CHECK (publish_day >= 1),
    CONSTRAINT lessons_course_order_unique UNIQUE (course_id, order_number)
);

CREATE INDEX IF NOT EXISTS idx_lessons_course_order
    ON catalog.lessons (course_id, order_number ASC);

ALTER TABLE catalog.lessons
    ADD CONSTRAINT fk_lessons_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;

ALTER TABLE catalog.lessons
    ADD CONSTRAINT fk_lessons_course
    FOREIGN KEY (course_id) REFERENCES catalog.courses (id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS catalog.lesson_videos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    lesson_id UUID NOT NULL,
    provider TEXT NOT NULL DEFAULT 'telegram',
    stream_ref TEXT,
    telegram_channel_id TEXT,
    telegram_message_id TEXT,
    embed_url TEXT,
    duration_seconds INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT lesson_videos_provider_check CHECK (provider IN ('telegram', 'external')),
    CONSTRAINT lesson_videos_duration_check CHECK (duration_seconds IS NULL OR duration_seconds >= 0),
    CONSTRAINT lesson_videos_lesson_unique UNIQUE (lesson_id)
);

ALTER TABLE catalog.lesson_videos
    ADD CONSTRAINT fk_lesson_videos_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;

ALTER TABLE catalog.lesson_videos
    ADD CONSTRAINT fk_lesson_videos_lesson
    FOREIGN KEY (lesson_id) REFERENCES catalog.lessons (id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS catalog.lesson_materials (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    lesson_id UUID NOT NULL,
    title TEXT NOT NULL,
    material_type TEXT NOT NULL DEFAULT 'pdf',
    source_type TEXT NOT NULL DEFAULT 'url',
    url TEXT,
    file_id UUID,
    order_number INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT lesson_materials_material_type_check CHECK (material_type IN ('pdf', 'image', 'doc', 'link', 'other')),
    CONSTRAINT lesson_materials_source_type_check CHECK (source_type IN ('url', 'filevault')),
    CONSTRAINT lesson_materials_order_number_check CHECK (order_number >= 1)
);

CREATE INDEX IF NOT EXISTS idx_lesson_materials_lesson_order
    ON catalog.lesson_materials (lesson_id, order_number ASC);

ALTER TABLE catalog.lesson_materials
    ADD CONSTRAINT fk_lesson_materials_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;

ALTER TABLE catalog.lesson_materials
    ADD CONSTRAINT fk_lesson_materials_lesson
    FOREIGN KEY (lesson_id) REFERENCES catalog.lessons (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS catalog.lesson_materials DROP CONSTRAINT IF EXISTS fk_lesson_materials_lesson;
ALTER TABLE IF EXISTS catalog.lesson_materials DROP CONSTRAINT IF EXISTS fk_lesson_materials_organization;
DROP TABLE IF EXISTS catalog.lesson_materials;
ALTER TABLE IF EXISTS catalog.lesson_videos DROP CONSTRAINT IF EXISTS fk_lesson_videos_lesson;
ALTER TABLE IF EXISTS catalog.lesson_videos DROP CONSTRAINT IF EXISTS fk_lesson_videos_organization;
DROP TABLE IF EXISTS catalog.lesson_videos;
ALTER TABLE IF EXISTS catalog.lessons DROP CONSTRAINT IF EXISTS fk_lessons_course;
ALTER TABLE IF EXISTS catalog.lessons DROP CONSTRAINT IF EXISTS fk_lessons_organization;
DROP TABLE IF EXISTS catalog.lessons;
-- +goose StatementEnd
