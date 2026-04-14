# Session Context

## Current task

Organize the Sinfim product idea using the project-builder and ai-blueprint workflow before creating the actual implementation repo.

## Goal

Create a clear project planning package for a multi-tenant LMS: startup scope, auth/access model, roles, MVP boundary, and next UX/technical documentation steps.

## Current status

- `docs/startup-idea.md` exists and contains the first LMS product scope.
- `docs/auth-access-notes.md` exists and captures the multi-tenant organization, role, lead, student access, mentor, and homework model.
- `docs/ux-doc.md` exists as a first UX skeleton with flows, routes, screen list, and components.
- Course/Class UX model is decided: course is reusable content; class/group is the live cohort operation.
- Ders Editor, Student Dashboard, Student Ders Detay, and Odev Kontrol UX details are drafted.
- Organization Setup/Edit, Public Okul/Kurs Lead Formu, Demo Okul, and Login/first-login UX details are drafted.
- `docs/tech-doc.md` exists as the first technical contract draft with modules, data model, API use cases, portals, and frontend API files.
- First UC docs exist for `auth/login`, `auth/change-initial-password`, `organization/create-organization`, and `organization/get-public-school-page`.
- Additional UC docs exist for `organization/create-school-request`, `catalog/create-course`, `catalog/create-lesson`, `classroom/create-class`, `classroom/add-student`, and `classroom/update-access`.
- Homework and learning UC docs exist for `submit-homework`, `review-submission`, `submit-quiz`, `get-student-dashboard`, `get-lesson-detail`, and `mark-lesson-completed`.
- Stitch AI master UI prompt exists at `prompts/ui/stitch-master-prompt.md`.
- Stitch AI design output was added under `docs/sinfim-design/` and reviewed. The visual direction is accepted as desktop-first web SaaS reference, with cleanup needed for placeholder brand/content and SMS copy.
- `docs/ui-design-review.md` exists and maps canonical design screens to UX screens.
- `docs/brand-constants.md` exists and fixes the canonical platform name/domain as `Sinfim.uz` / `sinfim.uz`.
- AI blueprint files were copied from `../blueprints/ai-blueprint` into this `projects/lms` workspace.
- The actual application repo has not been created yet.

## Important decisions

- Product is not only for language courses; target users include schools, education centers, individual teachers, mentors, and teams selling course videos.
- Platform must be multi-tenant: each school/brand/organization uses it like its own platform.
- MVP excludes automatic payment integration, but includes manual payment/access confirmation.
- MVP video strategy uses Telegram channel stream references.
- Mentor can be responsible for one or more classes/groups.
- Homework types include written answer, file/photo upload, quiz/test, and oral/audio message.
- Student acquisition supports both manual registration by admin/mentor and public lead submission from school/course pages.
- Product/platform name and domain: `Sinfim.uz` / `sinfim.uz`. Short `Sinfim` is allowed only as natural text shorthand.
- UI design direction: deep navy `#041632`, forest green `#2c694e`, cool surface `#f7f9fb`, Manrope headlines, Inter UI/body, 260px desktop sidebar.
- MVP organization URL model is path slug, e.g. `sinfim.uz/my-school`; subdomains can come later.
- MVP should avoid external auth dependencies: no SMS OTP and no Telegram login at first; use phone number identity with password/invite code.
- MVP organization creation is superadmin-controlled; public users cannot create schools directly. They can submit a school request or use demo school.
- Technical MVP decisions: keep `homework` and `learning` as separate modules; use storage abstraction with local dev and S3-compatible production target; use class-level access for MVP; keep Telegram stream references; keep quiz answer storage minimal.

## Blockers

- No blocker for starting the technical document.

## Open questions

- Exact student first-login flow can still be refined: admin-created temporary password, invite link, or temporary code.
## Exact next step

Use `docs/ui-design-review.md` and `docs/brand-constants.md` to normalize the Stitch output into Vue components, or start project setup.
