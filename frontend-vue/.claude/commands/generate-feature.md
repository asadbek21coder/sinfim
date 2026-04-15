# Generate Feature

Scaffold a new frontend feature with all required files and wiring.

The feature/domain name is provided as: $ARGUMENTS

## Steps

1. **Validate input.** The domain name must be provided. It should be lowercase, singular (e.g., `user`, `product`, `invoice`).

2. **Ask the user** for:
   - What this feature does (brief description)
   - Which backend API endpoints exist for it (paths and HTTP methods)
   - What roles can access it (e.g., `['ADMIN']` or `['USER']`)
   - Whether it needs a sidebar navigation link (and if so, what icon and label)
   - Whether it needs a Pinia store (only if state is shared across multiple views)

3. **Create types file** — `src/types/{domain}.ts`:
   - DTO interface mirroring backend response
   - Request interfaces (Create, Update) if mutations exist
   - Use only types from `src/types/common.ts` for shared contracts

4. **Create API file** — `src/api/{domain}.ts`:
   - One exported const object with all API calls for this domain
   - Match HTTP methods to the backend contract (GET for queries, POST for mutations)
   - Use `ApiResponse<T>` and `PageResponse<T>` from `src/types/common.ts`

5. **Create Pinia store** (only if user said yes) — `src/stores/{domain}.ts`:
   - Composition API style (`defineStore('name', () => { ... })`)
   - Loading and error refs included
   - Expose only what's needed by consuming views

6. **Create view** — `src/views/{Domain}View.vue`:
   - `<script setup lang="ts">` — always
   - `loading`, `error`, and data refs
   - `load()` function with try/catch/finally
   - `onMounted(() => load())`
   - Template with loading state, error state, and content
   - Pagination if the list endpoint returns `PageResponse<T>`

7. **Add route** to `src/router/index.ts`:
   - Lazy import
   - Full meta: `{ requiresAuth: true, roles: [...], layout: 'app' }`
   - Place logically among existing routes

8. **Add sidebar entry** (if user said yes) to `src/layouts/AppLayout.vue`:
   - Add `<SidebarItem icon="material-icon-name" label="Label" to="/path" />`
   - Place in logical order among existing items

## File Templates

### src/types/{domain}.ts

```typescript
export interface {Domain}Dto {
  id: string
  // mirror backend DTO fields exactly
  createdAt: string
  updatedAt: string
}

export interface Create{Domain}Dto {
  // request body fields
}

export interface Update{Domain}Dto {
  // request body fields
}
```

### src/api/{domain}.ts

```typescript
import client from './client'
import type { ApiResponse, PageResponse } from '@/types/common'
import type { {Domain}Dto, Create{Domain}Dto } from '@/types/{domain}'

export const {domain}Api = {
  list: (params: { page?: number; size?: number }) =>
    client.get<ApiResponse<PageResponse<{Domain}Dto>>>('/{domain}', { params }),
  getById: (id: string) =>
    client.get<ApiResponse<{Domain}Dto>>(`/{domain}/${id}`),
  create: (body: Create{Domain}Dto) =>
    client.post<ApiResponse<{Domain}Dto>>('/{domain}', body),
}
```

### src/views/{Domain}View.vue

```vue
<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { {domain}Api } from '@/api/{domain}'
import type { {Domain}Dto } from '@/types/{domain}'

const items = ref<{Domain}Dto[]>([])
const loading = ref(false)
const error = ref<string | null>(null)
const page = ref(0)
const total = ref(0)

async function load() {
  loading.value = true
  error.value = null
  try {
    const res = await {domain}Api.list({ page: page.value, size: 20 })
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
</script>

<template>
  <div>
    <div class="page-header">
      <h1 class="page-title">{Domain}</h1>
    </div>

    <div v-if="loading" class="text-center py-8 text-gray-400">Loading...</div>
    <div v-else-if="error" class="text-center py-8 text-red-500">{{ error }}</div>
    <div v-else class="card">
      <!-- table or content -->
    </div>
  </div>
</template>
```

### Route entry (in src/router/index.ts)

```typescript
{
  path: '/{domain}',
  component: () => import('@/views/{Domain}View.vue'),
  meta: { requiresAuth: true, roles: ['ADMIN'], layout: 'app' },
},
```

### Sidebar entry (in AppLayout.vue)

```html
<SidebarItem icon="icon_name" label="{Domain}" to="/{domain}" />
```

## After Scaffolding

Run `npm run build` to confirm zero TypeScript errors before considering the feature done.
