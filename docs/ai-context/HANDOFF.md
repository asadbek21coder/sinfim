# Agent Handoff

Use this file when switching between Claude and Codex or when a session is about to end.

## Task

Prepare the Sinfim.uz monorepo workspace for implementation.

## Current status

The workspace now has product docs under `docs/`, backend blueprint code under `backend-go/`, and frontend blueprint code under `frontend-vue/`. Shared AI handoff context is centralized under `docs/ai-context/`.

## Files touched

- `AGENTS.md`
- `CLAUDE.md`
- `.claude/`
- `.codex/`
- `docs/ai-context/SESSION.md`
- `docs/ai-context/HANDOFF.md`
- `docs/ai-context/WORKLOG.md`
- `docs/product/startup-idea.md`
- `docs/product/auth-access-notes.md`
- `docs/product/ux-doc.md`
- `docs/product/tech-doc.md`
- `docs/product/brand-constants.md`
- `docs/product/ui-design-review.md`
- `docs/design/sinfim-design/README.md`
- `backend-go/AGENTS.md`
- `frontend-vue/AGENTS.md`
- Removed duplicate/template AI folders from `ai/`, `docs/projects/lms/`, `backend-go/docs/ai-context/`, and `frontend-vue/docs/ai-context/`.

## Important decisions

- Sinfim.uz targets schools, education centers, individual teachers, mentors, and course-selling teams, not only language courses.
- Product/platform name and domain are fixed as `Sinfim.uz` / `sinfim.uz`. Short `Sinfim` is allowed only as natural text shorthand.
- Stitch AI design output exists under `docs/design/sinfim-design/`. It is accepted as a desktop-first web SaaS visual reference, but should not be copied directly into production code.
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

- Centralized planning AI context under shared root `docs/ai-context/`.
- Updated backend/frontend agent references to use `../docs/ai-context/`.
- Deleted duplicate/template AI folders, keeping root shared context plus backend/frontend stack-specific agent helpers.

## Open issues

- Decide exact student first-login flow: admin-created temporary password, invite link, or temporary code.
- Decide test/quiz auto-grading rules for MVP.
- Clean Stitch placeholders before implementation: Architectural Academic, Sterling, platform.uz, lms.com, SMS copy, physics/design demo content.
- Adapt backend/frontend blueprint code and README files to Sinfim.uz.
- `frontend-vue` was converted from a gitlink/submodule-style index entry into normal tracked files. `frontend-vue/node_modules` and `frontend-vue/dist` remain ignored by `frontend-vue/.gitignore`.

## Exact next step

Adapt backend/frontend blueprint configuration and docs to Sinfim.uz.
