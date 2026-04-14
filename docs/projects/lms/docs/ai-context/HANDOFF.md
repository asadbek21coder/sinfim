# Agent Handoff

Use this file when switching between Claude and Codex or when a session is about to end.

## Task

Prepare the Sinfim idea for implementation using project-builder docs and ai-blueprint workflow files.

## Current status

Planning workspace is in progress under `projects/lms`. The implementation repo has not been created yet. Startup scope, auth/access notes, UX docs, and first technical contract draft exist.

## Files touched

- `projects/lms/docs/startup-idea.md`
- `projects/lms/docs/auth-access-notes.md`
- `projects/lms/docs/ux-doc.md`
- `projects/lms/docs/tech-doc.md`
- `projects/lms/docs/brand-constants.md`
- `projects/lms/docs/ui-design-review.md`
- `projects/lms/docs/specs/modules/auth/usecases/login.md`
- `projects/lms/docs/specs/modules/auth/usecases/change-initial-password.md`
- `projects/lms/docs/specs/modules/organization/usecases/create-organization.md`
- `projects/lms/docs/specs/modules/organization/usecases/get-public-school-page.md`
- `projects/lms/docs/specs/modules/organization/usecases/create-school-request.md`
- `projects/lms/docs/specs/modules/catalog/usecases/create-course.md`
- `projects/lms/docs/specs/modules/catalog/usecases/create-lesson.md`
- `projects/lms/docs/specs/modules/classroom/usecases/create-class.md`
- `projects/lms/docs/specs/modules/classroom/usecases/add-student.md`
- `projects/lms/docs/specs/modules/classroom/usecases/update-access.md`
- `projects/lms/docs/specs/modules/homework/usecases/submit-homework.md`
- `projects/lms/docs/specs/modules/homework/usecases/review-submission.md`
- `projects/lms/docs/specs/modules/homework/usecases/submit-quiz.md`
- `projects/lms/docs/specs/modules/learning/usecases/get-student-dashboard.md`
- `projects/lms/docs/specs/modules/learning/usecases/get-lesson-detail.md`
- `projects/lms/docs/specs/modules/learning/usecases/mark-lesson-completed.md`
- `projects/lms/prompts/ui/stitch-master-prompt.md`
- `projects/lms/AGENTS.md`
- `projects/lms/docs/ai-context/SESSION.md`
- `projects/lms/docs/ai-context/HANDOFF.md`
- `projects/lms/docs/ai-context/WORKLOG.md`
- `projects/lms/.claude/`
- `projects/lms/.codex/`
- `projects/lms/AI-BLUEPRINT.md`

## Important decisions

- Sinfim targets schools, education centers, individual teachers, mentors, and course-selling teams, not only language courses.
- Product/platform name and domain are fixed as `Sinfim.uz` / `sinfim.uz`. Short `Sinfim` is allowed only as natural text shorthand.
- Stitch AI design output exists under `docs/sinfim-design/`. It is accepted as a desktop-first web SaaS visual reference, but should not be copied directly into production code.
- UI direction: deep navy `#041632`, forest green `#2c694e`, cool surface `#f7f9fb`, Manrope headlines, Inter UI/body, 260px sidebar.
- Multi-tenant organization boundary is mandatory.
- Organization owner manages courses, classes/groups, mentors, students, videos, materials, tests, homework, and access.
- Payment stays outside the MVP, but access/payment status is manually confirmed inside the platform.
- Telegram channel stream references are the initial video strategy.
- Leads and students are separate: public school/course pages can collect leads; admins/mentors can convert leads to students.
- Homework supports written, file/photo, quiz/test, and oral/audio message types.
- Course/Class UX model is decided: course is reusable content package; class/group is live cohort operation.
- Ders Editor, Student Dashboard, Student Ders Detay, and Odev Kontrol UX details are drafted.
- Organization Setup/Edit, Public Okul/Kurs Lead Formu, Demo Okul, and Login/first-login UX details are drafted.
- MVP organization URL model is path slug: `sinfim.uz/{school-slug}`. Subdomain model is postponed.
- MVP avoids external auth dependencies: no SMS OTP and no Telegram login at first. Use phone number identity with password/invite code.
- MVP organization creation is superadmin-controlled. Public visitors can submit school requests or use a demo/fake school, but cannot create a real school directly.
- Technical MVP decisions: keep `homework` and `learning` separate; use storage abstraction with local dev and S3-compatible production target; use class-level access for MVP; Telegram stream references only; minimal quiz answer storage.

## Commands run

- `find ../blueprints/ai-blueprint ...` -> initial relative path did not exist.
- `find .. -maxdepth 3 ...` -> found `../blueprints/ai-blueprint`.
- `cp -R ../blueprints/ai-blueprint/... projects/lms/...` -> copied AI blueprint files into LMS planning workspace.

## Open issues

- Decide exact student first-login flow: admin-created temporary password, invite link, or temporary code.
- Decide test/quiz auto-grading rules for MVP.
- Clean Stitch placeholders before implementation: Architectural Academic, Sterling, platform.uz, lms.com, SMS copy, physics/design demo content.

## Exact next step

Use `projects/lms/docs/ui-design-review.md` and `projects/lms/docs/brand-constants.md` to normalize Stitch screens into implementation-ready Vue component/layout decisions, or start project setup.
