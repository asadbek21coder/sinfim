# Sinfim.uz Frontend

Vue 3 + TypeScript + Vite + Pinia + Tailwind web application for Sinfim.uz.

Sinfim.uz is a desktop-first multi-tenant online school platform for schools, education centers, teachers, mentors, and course-selling teams.

## Stack

- Vue 3 Composition API with `<script setup>`
- TypeScript strict mode
- Vite
- Pinia
- Vue Router 4
- Axios with token storage and refresh handling
- Tailwind CSS v3

## Local Development

From the repository root, the preferred local workflow is:

```bash
make local-up
```

Frontend-only development:

```bash
cp .env.example .env
npm install
npm run dev
```

Default local URLs:

- Frontend: `http://localhost:5173`
- Backend: `http://localhost:9876`
- Backend health: `http://localhost:9876/health`

## Current Step 0 Structure

```text
src/
  api/              API clients and health check
  stores/           Pinia stores
  types/            Shared TypeScript DTOs
  router/           Route map and future auth guard
  layouts/          Public, auth, school app, and student layouts
  components/       Reusable UI and shell components
  views/            Route views
```

Important route shells:

- Public: `/`, `/enter`, `/apply-school`, `/demo`
- Auth: `/auth/login`
- School app: `/app/dashboard`, `/app/courses`, `/app/classes`, `/app/leads`
- Student learning: `/learn/dashboard`, `/learn/lessons/:lessonId`

## Auth Direction

MVP auth is phone number + password/temporary password.

No SMS OTP, no Telegram login, and no SSO/PKCE in the MVP.

## Verification

```bash
npm run build
```
