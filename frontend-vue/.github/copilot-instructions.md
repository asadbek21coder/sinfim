# GitHub Copilot Instructions — Vue Blueprint Web

## Stack
Vue 3 Composition API · TypeScript (strict) · Vite · Pinia · Vue Router 4 · Axios · Tailwind CSS v3

## Non-negotiable rules

- Always use `<script setup lang="ts">` — never Options API
- Never use `any` type — always explicit generics
- All HTTP calls go through `src/api/*.ts` — never raw axios in components
- One API file per domain: `usersApi`, `ordersApi`, etc., all using the typed `client` from `src/api/client.ts`
- Styling: Tailwind classes only — use the predefined classes in `src/assets/main.css` (`btn-primary`, `form-input`, `card`, `data-table`, `page-header`, etc.)
- State management: Pinia only — no Vuex, no local component-level global state

## Response shape from backend

All responses follow:
```ts
{ success: boolean; data?: T; error?: { code: string; message: string } }
```
Paginated responses: `data` is `{ content: T[]; page: number; size: number; totalElements: number; totalPages: number }`

## Error handling pattern

```ts
try {
  const res = await someApi.doThing()
  // use res.data.data!
} catch (err: any) {
  errorMsg.value = err?.response?.data?.error?.message ?? 'Unexpected error'
}
```

## Pagination pattern

```ts
const page = ref(0)
const items = ref<T[]>([])
const total = ref(0)
const loading = ref(false)

async function load() {
  loading.value = true
  try {
    const res = await api.list({ page: page.value, size: 20 })
    items.value = res.data.data!.content
    total.value = res.data.data!.totalElements
  } finally {
    loading.value = false
  }
}
```

## New feature steps (always follow this order)

1. `src/types/{domain}.ts` — DTO types mirroring backend
2. `src/api/{domain}.ts` — typed API calls
3. Route in `src/router/index.ts` with `meta: { requiresAuth: true, roles: [...] }`
4. `<SidebarItem>` in `AppLayout.vue` if it needs nav
5. `src/views/{Domain}View.vue` — page component

## Auth

Do NOT modify the auth flow in `src/api/client.ts`, `src/stores/auth.ts`, or the PKCE logic in `LoginView.vue` / `CallbackView.vue`. The token refresh interceptor is already handled.
