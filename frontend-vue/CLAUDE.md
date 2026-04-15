# CLAUDE

## Skills

For any frontend work (implementation, review, modification, planning), load first:

1. **`frontend-guidelines`** — project architecture, conventions, and development rules

Do not improvise on patterns, naming, or structure — follow what the skill prescribes.

---

## Agents

Use specialized agents (`.claude/agents/`) instead of generic subagents. Each agent has focused instructions, restricted tools, and auto-loaded skills.

| Agent | Role | Tools |
|---|---|---|
| **`vue-architect`** | Plans feature implementation | Read, Glob, Grep *(read-only)* |
| **`vue-coder`** | Implements types, API, views, components, routes | Read, Write, Edit, Bash, Glob, Grep |
| **`vue-reviewer`** | Reviews code against guidelines, runs build | Read, Glob, Grep, Bash *(no edit)* |

### How to delegate

```
Task tool → subagent_type: "vue-architect" / "vue-coder" / "vue-reviewer"
```

### Pipeline

For features and pages, follow this pipeline:

1. **vue-architect** — plan the feature (read-only, for complex features)
2. **vue-coder** — implement types, API, view, route, sidebar entry
3. **vue-reviewer** — review against guidelines, run `npm run build`
4. Fix CRITICAL issues from reviewer, re-review if needed

For simple features (one view, one API file), skip the architect and go directly to vue-coder.

---

## Project

Vue 3 + TypeScript SPA connecting to the Sinfim.uz Go backend. Phone number + password auth. Pinia state. Tailwind CSS.

**Backend contract:** GET for queries, POST for mutations. No PATCH/PUT/DELETE.
**Auth:** Sinfim.uz MVP uses phone number + password, plus first-login temporary password/access code. No SMS OTP, no Telegram login, no SSO/PKCE in MVP. Replace the blueprint auth screens/API with this model when adapting the frontend.

---

## Verification

```bash
npm run build    # TypeScript compile + Vite build — must pass, zero errors
```

**NEVER deliver code that hasn't passed `npm run build` with zero TypeScript errors.**

---

## Workflow

Use `/generate-feature <domain>` to scaffold a new feature with all required files.
Use `/review` or `/review <domain>` to run a review.
Use `/enhance-docs` to add new rules or patterns to the guidelines.

For complex features, use `TaskCreate` to break work into steps and `Task` to run agents.

---

## File Structure

```
src/
├── api/              # All HTTP calls — one file per domain
├── stores/           # Pinia stores — one per domain, only for shared state
├── types/            # TypeScript types — mirror backend DTOs
├── router/           # Routes + navigation guards
├── layouts/          # Public, auth, school app, and student layouts
├── components/
│   ├── ui/           # Generic reusable components
│   └── {domain}/     # Domain-specific components
├── views/            # One view per route
│   └── auth/         # LoginView phone/password flow
└── assets/main.css   # Tailwind + component class definitions
```

---

## Rules

- Always `<script setup lang="ts">` — no Options API, no `defineComponent`
- API calls only through `src/api/*.ts` — never raw axios in views or components
- Never `any` type — all refs and params explicitly typed
- No inline `style=""` — Tailwind utilities and component classes only
- No `console.log` in committed code
- No direct localStorage access outside `client.ts`
- Roles as `Role` type constants — no magic strings
- Route views always lazy imported: `() => import('@/views/...')`

---

## Customisation Checklist (start of every new project)

1. `src/types/auth.ts` — update `Role` union type to match backend
2. `src/main.ts` — update `roleHomeMap` (role → home route)
3. `src/router/index.ts` — extend the Step 0 route shells with real feature routes
4. `src/layouts/AppLayout.vue` — add `<SidebarItem>` entries
5. `src/layouts/AuthLayout.vue` — keep phone/password auth copy aligned with Sinfim.uz
6. `index.html` — set `<title>`
7. `tailwind.config.js` — update brand colors if needed
8. `vite.config.ts` — update `VITE_API_PROXY_TARGET` if backend port differs from 9876
9. `package.json` — rename `"name"` field
10. `CLAUDE.md` — update project description to reflect actual project context
