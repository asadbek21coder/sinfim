# CLAUDE

This is the Sinfim.uz monorepo workspace.

Read `AGENTS.md` first for workspace routing and canonical product decisions.

## Shared Context

Use the single shared context folder:

- `docs/ai-context/HANDOFF.md`
- `docs/ai-context/SESSION.md`
- `docs/ai-context/WORKLOG.md`

Do not create duplicate AI context folders inside subprojects unless there is a specific reason. Backend and frontend agents should refer back to the shared root docs.

## Where To Work

- Product and planning docs: `docs/`
- Backend implementation: `backend-go/`
- Frontend implementation: `frontend-vue/`

## Important Product Docs

- `docs/product/brand-constants.md`
- `docs/product/startup-idea.md`
- `docs/product/auth-access-notes.md`
- `docs/product/ux-doc.md`
- `docs/product/tech-doc.md`
- `docs/product/ui-design-review.md`

## Design Output

Raw Stitch AI output is in `docs/design/sinfim-design/`.

Read `docs/design/sinfim-design/README.md` before using it. The raw HTML/screenshots contain placeholder copy and must not be copied directly into production.

## Subproject Agent Instructions

When working in `backend-go/`, follow `backend-go/CLAUDE.md`.

When working in `frontend-vue/`, follow `frontend-vue/CLAUDE.md`.
