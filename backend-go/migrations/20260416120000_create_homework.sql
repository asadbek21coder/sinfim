-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS homework;

CREATE TABLE IF NOT EXISTS homework.definitions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    course_id UUID NOT NULL,
    lesson_id UUID NOT NULL,
    title TEXT NOT NULL,
    instructions TEXT,
    submission_type TEXT NOT NULL DEFAULT 'text',
    status TEXT NOT NULL DEFAULT 'draft',
    max_score INTEGER NOT NULL DEFAULT 100,
    due_days_after_publish INTEGER,
    allow_resubmission BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT homework_definitions_submission_type_check CHECK (submission_type IN ('text', 'file', 'audio', 'quiz')),
    CONSTRAINT homework_definitions_status_check CHECK (status IN ('draft', 'published', 'archived')),
    CONSTRAINT homework_definitions_max_score_check CHECK (max_score >= 0),
    CONSTRAINT homework_definitions_due_days_check CHECK (due_days_after_publish IS NULL OR due_days_after_publish >= 0),
    CONSTRAINT homework_definitions_lesson_unique UNIQUE (lesson_id)
);

CREATE INDEX IF NOT EXISTS idx_homework_definitions_lesson
    ON homework.definitions (lesson_id);

ALTER TABLE homework.definitions
    ADD CONSTRAINT fk_homework_definitions_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;

ALTER TABLE homework.definitions
    ADD CONSTRAINT fk_homework_definitions_course
    FOREIGN KEY (course_id) REFERENCES catalog.courses (id) ON DELETE CASCADE;

ALTER TABLE homework.definitions
    ADD CONSTRAINT fk_homework_definitions_lesson
    FOREIGN KEY (lesson_id) REFERENCES catalog.lessons (id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS homework.quiz_questions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    definition_id UUID NOT NULL,
    prompt TEXT NOT NULL,
    order_number INTEGER NOT NULL DEFAULT 1,
    points INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT homework_quiz_questions_order_check CHECK (order_number >= 1),
    CONSTRAINT homework_quiz_questions_points_check CHECK (points >= 0)
);

CREATE INDEX IF NOT EXISTS idx_homework_quiz_questions_definition_order
    ON homework.quiz_questions (definition_id, order_number ASC);

ALTER TABLE homework.quiz_questions
    ADD CONSTRAINT fk_homework_quiz_questions_definition
    FOREIGN KEY (definition_id) REFERENCES homework.definitions (id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS homework.quiz_options (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    question_id UUID NOT NULL,
    label TEXT NOT NULL,
    is_correct BOOLEAN NOT NULL DEFAULT false,
    order_number INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT homework_quiz_options_order_check CHECK (order_number >= 1)
);

CREATE INDEX IF NOT EXISTS idx_homework_quiz_options_question_order
    ON homework.quiz_options (question_id, order_number ASC);

ALTER TABLE homework.quiz_options
    ADD CONSTRAINT fk_homework_quiz_options_question
    FOREIGN KEY (question_id) REFERENCES homework.quiz_questions (id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS homework.submissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    definition_id UUID NOT NULL,
    lesson_id UUID NOT NULL,
    class_id UUID NOT NULL,
    student_user_id VARCHAR NOT NULL,
    submission_type TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'submitted',
    attempt_number INTEGER NOT NULL DEFAULT 1,
    text_answer TEXT,
    file_url TEXT,
    audio_url TEXT,
    score INTEGER,
    max_score INTEGER NOT NULL DEFAULT 100,
    auto_scored BOOLEAN NOT NULL DEFAULT false,
    submitted_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    reviewed_at TIMESTAMPTZ,
    reviewer_user_id VARCHAR,
    feedback TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT homework_submissions_submission_type_check CHECK (submission_type IN ('text', 'file', 'audio', 'quiz')),
    CONSTRAINT homework_submissions_status_check CHECK (status IN ('submitted', 'reviewed', 'needs_revision')),
    CONSTRAINT homework_submissions_attempt_check CHECK (attempt_number >= 1),
    CONSTRAINT homework_submissions_score_check CHECK (score IS NULL OR score >= 0),
    CONSTRAINT homework_submissions_class_student_definition_unique UNIQUE (class_id, student_user_id, definition_id)
);

CREATE INDEX IF NOT EXISTS idx_homework_submissions_definition_status
    ON homework.submissions (definition_id, status, submitted_at DESC);

CREATE INDEX IF NOT EXISTS idx_homework_submissions_student
    ON homework.submissions (student_user_id, submitted_at DESC);

ALTER TABLE homework.submissions
    ADD CONSTRAINT fk_homework_submissions_organization
    FOREIGN KEY (organization_id) REFERENCES organization.organizations (id) ON DELETE CASCADE;

ALTER TABLE homework.submissions
    ADD CONSTRAINT fk_homework_submissions_definition
    FOREIGN KEY (definition_id) REFERENCES homework.definitions (id) ON DELETE CASCADE;

ALTER TABLE homework.submissions
    ADD CONSTRAINT fk_homework_submissions_lesson
    FOREIGN KEY (lesson_id) REFERENCES catalog.lessons (id) ON DELETE CASCADE;

ALTER TABLE homework.submissions
    ADD CONSTRAINT fk_homework_submissions_class
    FOREIGN KEY (class_id) REFERENCES classroom.classes (id) ON DELETE CASCADE;

ALTER TABLE homework.submissions
    ADD CONSTRAINT fk_homework_submissions_student_user
    FOREIGN KEY (student_user_id) REFERENCES auth.users (id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS homework.quiz_answers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    submission_id UUID NOT NULL,
    question_id UUID NOT NULL,
    selected_option_id UUID,
    is_correct BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT homework_quiz_answers_submission_question_unique UNIQUE (submission_id, question_id)
);

ALTER TABLE homework.quiz_answers
    ADD CONSTRAINT fk_homework_quiz_answers_submission
    FOREIGN KEY (submission_id) REFERENCES homework.submissions (id) ON DELETE CASCADE;

ALTER TABLE homework.quiz_answers
    ADD CONSTRAINT fk_homework_quiz_answers_question
    FOREIGN KEY (question_id) REFERENCES homework.quiz_questions (id) ON DELETE CASCADE;

ALTER TABLE homework.quiz_answers
    ADD CONSTRAINT fk_homework_quiz_answers_option
    FOREIGN KEY (selected_option_id) REFERENCES homework.quiz_options (id) ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS homework.quiz_answers DROP CONSTRAINT IF EXISTS fk_homework_quiz_answers_option;
ALTER TABLE IF EXISTS homework.quiz_answers DROP CONSTRAINT IF EXISTS fk_homework_quiz_answers_question;
ALTER TABLE IF EXISTS homework.quiz_answers DROP CONSTRAINT IF EXISTS fk_homework_quiz_answers_submission;
DROP TABLE IF EXISTS homework.quiz_answers;
ALTER TABLE IF EXISTS homework.submissions DROP CONSTRAINT IF EXISTS fk_homework_submissions_student_user;
ALTER TABLE IF EXISTS homework.submissions DROP CONSTRAINT IF EXISTS fk_homework_submissions_class;
ALTER TABLE IF EXISTS homework.submissions DROP CONSTRAINT IF EXISTS fk_homework_submissions_lesson;
ALTER TABLE IF EXISTS homework.submissions DROP CONSTRAINT IF EXISTS fk_homework_submissions_definition;
ALTER TABLE IF EXISTS homework.submissions DROP CONSTRAINT IF EXISTS fk_homework_submissions_organization;
DROP TABLE IF EXISTS homework.submissions;
ALTER TABLE IF EXISTS homework.quiz_options DROP CONSTRAINT IF EXISTS fk_homework_quiz_options_question;
DROP TABLE IF EXISTS homework.quiz_options;
ALTER TABLE IF EXISTS homework.quiz_questions DROP CONSTRAINT IF EXISTS fk_homework_quiz_questions_definition;
DROP TABLE IF EXISTS homework.quiz_questions;
ALTER TABLE IF EXISTS homework.definitions DROP CONSTRAINT IF EXISTS fk_homework_definitions_lesson;
ALTER TABLE IF EXISTS homework.definitions DROP CONSTRAINT IF EXISTS fk_homework_definitions_course;
ALTER TABLE IF EXISTS homework.definitions DROP CONSTRAINT IF EXISTS fk_homework_definitions_organization;
DROP TABLE IF EXISTS homework.definitions;
DROP SCHEMA IF EXISTS homework;
-- +goose StatementEnd
