-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS learning;

CREATE TABLE IF NOT EXISTS learning.lesson_completions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    class_id UUID NOT NULL,
    lesson_id UUID NOT NULL,
    student_user_id VARCHAR NOT NULL,
    completed_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT lesson_completions_class_lesson_student_unique UNIQUE (class_id, lesson_id, student_user_id)
);

CREATE INDEX IF NOT EXISTS idx_lesson_completions_student_class
    ON learning.lesson_completions (student_user_id, class_id);

ALTER TABLE learning.lesson_completions
    ADD CONSTRAINT fk_lesson_completions_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;

ALTER TABLE learning.lesson_completions
    ADD CONSTRAINT fk_lesson_completions_class
    FOREIGN KEY (class_id) REFERENCES classroom.classes (id) ON DELETE CASCADE;

ALTER TABLE learning.lesson_completions
    ADD CONSTRAINT fk_lesson_completions_lesson
    FOREIGN KEY (lesson_id) REFERENCES catalog.lessons (id) ON DELETE CASCADE;

ALTER TABLE learning.lesson_completions
    ADD CONSTRAINT fk_lesson_completions_student
    FOREIGN KEY (student_user_id) REFERENCES auth.users (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS learning.lesson_completions DROP CONSTRAINT IF EXISTS fk_lesson_completions_student;
ALTER TABLE IF EXISTS learning.lesson_completions DROP CONSTRAINT IF EXISTS fk_lesson_completions_lesson;
ALTER TABLE IF EXISTS learning.lesson_completions DROP CONSTRAINT IF EXISTS fk_lesson_completions_class;
ALTER TABLE IF EXISTS learning.lesson_completions DROP CONSTRAINT IF EXISTS fk_lesson_completions_organization;
DROP TABLE IF EXISTS learning.lesson_completions;
DROP SCHEMA IF EXISTS learning;
-- +goose StatementEnd
