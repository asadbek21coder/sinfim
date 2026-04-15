---
name: frontend-guidelines
description: Use when implementing any frontend feature, writing components, creating API files, managing stores, handling routing, writing styles, or making any code changes to the Vue blueprint.
---

# Frontend Guidelines

## Overview

Reference guide for all frontend development conventions in the Vue blueprint. Covers architecture, file structure, coding standards, component patterns, API contracts, state management, routing, styling, and verification. Every rule is mandatory — do not improvise.

**Core principle:** One pattern per concern. No raw axios in components. No inline styles. No `any` types. Every feature follows the same shape: types → api → route → view.

---

## Architecture

### File Structure

```
src/
├── api/              # One file per domain — all HTTP calls live here
├── stores/           # Pinia stores — one per domain if state is needed
├── types/            # TypeScript types — mirror backend DTOs exactly
├── router/           # Routes + navigation guards
├── layouts/          # AppLayout (authenticated), AuthLayout (login/callback)
├── components/
│   ├── ui/           # Generic reusable components (SidebarItem, etc.)
│   └── {domain}/     # Domain-specific components, grouped by feature
├── views/            # One view per route
│   └── auth/         # LoginView, CallbackView
└── assets/main.css   # Tailwind directives + component class definitions
```

### Layer Rules

| Layer | Responsibility | What goes here |
|---|---|---|
| `types/` | Shape of data | Interfaces, enums, DTOs mirroring backend |
| `api/` | HTTP calls | Typed axios calls, nothing else |
| `stores/` | Shared app state | Pinia stores for cross-component state |
| `views/` | Page logic | Fetch data, bind to components, handle user actions |
| `components/` | Reusable UI | Dumb components (take props, emit events) |

**Rules:**
- Components do NOT call axios directly — always through `src/api/*.ts`
- Views orchestrate: they call api, update local refs, handle errors
- Stores hold only shared state — don't put view-local state in stores
- Types mirror backend response DTOs exactly — no field renaming

---

## Tech Stack (do NOT change core libs without asking)

| Layer | Library | Version |
|---|---|---|
| Framework | Vue 3 Composition API | ^3.5 |
| Build | Vite | ^5.4 |
| Language | TypeScript strict | ~5.5 |
| State | Pinia | ^2.2 |
| Routing | Vue Router 4 | ^4.4 |
| HTTP | Axios | ^1.7 |
| Styling | Tailwind CSS v3 | ^3.4 |
| Icons | Material Symbols (Google Fonts CDN) | — |

---

## Coding Standards

### Component syntax — always `<script setup lang="ts">`

```vue
<script setup lang="ts">
// always lang="ts", always setup — no exceptions
</script>
```

No Options API. No `defineComponent`. No `export default {}`.

### Types — explicit, never `any`

```typescript
// Correct
const users = ref<UserDto[]>([])
async function fetchUser(id: string): Promise<UserDto> { ... }

// Wrong
const users = ref([])
async function fetchUser(id: any) { ... }
```

### Refs — typed and initialized

```typescript
const items = ref<ItemDto[]>([])
const loading = ref(false)
const error = ref<string | null>(null)
const total = ref(0)
const page = ref(0)
```

### Error handling

```typescript
try {
  await someApi.doThing()
} catch (err: unknown) {
  error.value = (err as any)?.response?.data?.error?.message ?? 'Unexpected error'
}
```

Always use `error` ref. Never leave errors silent. Never `console.log` — use `console.warn` only for real warnings.

### No magic strings for roles

```typescript
// Correct
if (auth.hasRole(['ADMIN'])) { ... }

// Wrong
if (auth.user?.role === 'admin') { ... }
```

---

## API Layer

### One file per domain

```typescript
// src/api/users.ts
import client from './client'
import type { ApiResponse, PageResponse } from '@/types/common'
import type { UserDto, CreateUserDto } from '@/types/user'

export const usersApi = {
  list: (params: { page?: number; size?: number }) =>
    client.get<ApiResponse<PageResponse<UserDto>>>('/users', { params }),
  getById: (id: string) =>
    client.get<ApiResponse<UserDto>>(`/users/${id}`),
  create: (body: CreateUserDto) =>
    client.post<ApiResponse<UserDto>>('/users', body),
  update: (id: string, body: UpdateUserDto) =>
    client.patch<ApiResponse<UserDto>>(`/users/${id}`, body),
  delete: (id: string) =>
    client.delete<ApiResponse<null>>(`/users/${id}`),
}
```

### API response types

- Single item: `ApiResponse<T>` — `res.data.data`
- Paginated list: `ApiResponse<PageResponse<T>>` — `res.data.data!.content`, `res.data.data!.totalElements`
- Error: `res.data.error?.message`

### Go backend URL contract

The Go backend uses `GET` for queries and `POST` for mutations. No PATCH/PUT/DELETE in the Go API.
Map accordingly:

```typescript
list: (params) => client.get('/resource', { params })       // GET  — query
create: (body) => client.post('/resource', body)            // POST — mutation
update: (id, body) => client.post('/resource/update', body) // POST — mutation (no PATCH)
delete: (id) => client.post('/resource/delete', { id })     // POST — mutation (no DELETE)
```

**Only deviate if the backend explicitly uses REST verbs.**

---

## Types

### Mirror backend DTOs exactly

```typescript
// src/types/user.ts
export interface UserDto {
  id: string
  username: string
  email: string
  role: Role
  isActive: boolean
  createdAt: string  // ISO string — don't use Date unless formatting
}

export interface CreateUserDto {
  username: string
  email: string
  role: Role
  password: string
}
```

No renaming fields. No adding computed fields to DTOs. Use separate interfaces for request vs response bodies.

### Common types (never modify)

```typescript
// src/types/common.ts
interface ApiResponse<T> {
  data?: T
  error?: ErrorDetail
}

interface PageResponse<T> {
  content: T[]
  totalElements: number
  page: number
  size: number
}
```

---

## Routing

### Always declare meta

```typescript
{
  path: '/users',
  component: () => import('@/views/UsersView.vue'),
  meta: { requiresAuth: true, roles: ['ADMIN'], layout: 'app' },
}
```

- `requiresAuth: true` — enforces auth guard
- `roles: ['ADMIN']` — enforces role-based access (empty array = any authenticated user)
- `layout: 'app'` or `layout: 'auth'` — controls which layout wraps the view

### Lazy loading — always use dynamic import

```typescript
component: () => import('@/views/SomeView.vue')
// NOT: component: SomeView (static import)
```

---

## Pagination Pattern

```typescript
const page = ref(0)
const pageSize = ref(20)
const total = ref(0)
const items = ref<ItemDto[]>([])
const loading = ref(false)
const error = ref<string | null>(null)

async function load() {
  loading.value = true
  error.value = null
  try {
    const res = await itemsApi.list({ page: page.value, size: pageSize.value })
    items.value = res.data.data!.content
    total.value = res.data.data!.totalElements
  } catch (err: unknown) {
    error.value = (err as any)?.response?.data?.error?.message ?? 'Failed to load'
  } finally {
    loading.value = false
  }
}

onMounted(() => load())
watch(page, () => load())
```

---

## Stores (Pinia)

Use stores for state shared across multiple components or views. Do NOT use stores for view-local state.

```typescript
// src/stores/users.ts
import { defineStore } from 'pinia'
import { usersApi } from '@/api/users'
import type { UserDto } from '@/types/user'

export const useUsersStore = defineStore('users', () => {
  const list = ref<UserDto[]>([])
  const loading = ref(false)

  async function fetchAll() {
    loading.value = true
    try {
      const res = await usersApi.list({})
      list.value = res.data.data!.content
    } finally {
      loading.value = false
    }
  }

  return { list, loading, fetchAll }
})
```

Composition API style (not options object) — matches `<script setup>` patterns.

---

## Styling

### Use predefined component classes from `src/assets/main.css`

| Category | Classes |
|---|---|
| Buttons | `btn-primary`, `btn-secondary`, `btn-danger`, `btn-success`, `btn-icon` |
| Forms | `form-label`, `form-input`, `form-select`, `form-textarea`, `form-error` |
| Layout | `card`, `data-table`, `page-header`, `page-title`, `filter-bar` |
| Status chips | `chip-pending`, `chip-approved`, `chip-rejected`, `chip-info` |

Use Tailwind utilities for spacing, sizing, flex/grid. Combine with component classes:

```html
<button class="btn-primary w-full mt-4">Save</button>
<div class="card p-6 mt-4">...</div>
```

### Rules

- No `style=""` attributes on any element
- No new CSS files — add to `@layer components` in `main.css`
- No hardcoded colors — use Tailwind config colors
- Icons: use Material Symbols with `<span class="material-symbols-outlined">`

---

## Auth Flow

Sinfim.uz MVP auth uses phone number + password. First-time users may enter a temporary password or access code given by the school admin, then set their own password.

- No SMS OTP in MVP.
- No Telegram login in MVP.
- No SSO/PKCE in MVP.
- Token refresh stays centralized in `src/api/client.ts`. Do NOT add manual refresh logic anywhere else.
- Do NOT access localStorage directly outside `client.ts`.

---

## Feature Implementation Checklist

When adding a new feature:

- [ ] `src/types/{domain}.ts` — define DTOs mirroring backend
- [ ] `src/api/{domain}.ts` — typed API calls using client
- [ ] Route in `src/router/index.ts` with `meta.requiresAuth` and `meta.roles`
- [ ] `<SidebarItem>` in `AppLayout.vue` if it needs a nav link
- [ ] `src/views/{Domain}View.vue` — loading + error refs, never silent errors
- [ ] Paginated lists: use `page` + `totalElements` from `PageResponse<T>`
- [ ] If state is shared: `src/stores/{domain}.ts`

---

## What NOT to Do

- No Options API (`data()`, `methods:`, `computed:`)
- No raw `axios` calls outside `src/api/`
- No `any` type
- No inline `style=""` attributes
- No direct `localStorage` access outside `client.ts`
- No Vuex (use Pinia)
- No `console.log` in committed code
- No hardcoded API base URLs — always `import.meta.env.VITE_API_BASE_URL`
- No static component imports — always use lazy `() => import(...)`
- No bare array responses — always wrapped in `{ content: [] }`

---

## Verification

```bash
npm run build     # TypeScript compile + Vite build — must pass, zero type errors
npm run dev       # Dev server — confirm feature works end-to-end
```

- **Never deliver code that fails `npm run build`**
- TypeScript strict mode is on — no type assertions to bypass errors
- Fix type errors, don't cast around them with `as any`
