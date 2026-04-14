# CLAUDE.md — [Proje Adı]

Bu dosya her Claude oturumunun başında otomatik okunur.
**Her alanda doldur, boş bırakma.**

---

## Proje Özeti

> **Proje Adı:** ___________________________
> **Tür:** Startup / Outsource / Hybrid
> **Amaç (1 cümle):** ___________________________
> **Hedef Kullanıcı:** ___________________________
> **Kuzey Yıldızı Metriği:** ___________________________

---

## Stack

| Katman | Teknoloji |
|--------|-----------|
| Backend | Go — `go-enterprise-blueprint` |
| HTTP | GET (query) + POST (mutation) only — REST değil |
| ORM | bun + repogen |
| Migrations | goose |
| Error | errx (`github.com/code19m/errx`) |
| Frontend | Vue 3 + TypeScript — `vue-blueprint-web` |
| State | Pinia |
| Router | Vue Router 4 |
| HTTP client | Axios (`src/api/client.ts`) |
| Styling | Tailwind CSS v3 + predefined component classes (`main.css`) |
| Icons | Material Symbols (Google Fonts CDN) |
| Auth | SSO / PKCE (OneID veya OAuth2) |
| Database | PostgreSQL (`timestamptz` for all timestamps) |
| Deploy | [Railway / Fly.io / VPS] |

---

## Blueprint Referansları

```
blueprints/
├── go-enterprise-blueprint-main/   ← Go backend'in kaynak yapısı
│   ├── CLAUDE.md                   ← Go pipeline ve agent kuralları
│   └── .claude/skills/backend-guidelines/SKILL.md  ← TÜM Go kuralları
└── vue-blueprint-web/
    └── CLAUDE.md                   ← Vue kuralları (bölüm 4 mandatory)
```

Go kodu yazarken her zaman `backend-guidelines` skill'ini yükle.

---

## Go Backend — Mimari

### Klasör Yapısı

```
backend/
├── cmd/                        ← entry point (Cobra CLI)
├── config/                     ← YAML config per environment
├── internal/
│   ├── app/                    ← bootstrap ve lifecycle
│   ├── modules/
│   │   └── {module}/
│   │       ├── module.go       ← init ve wiring
│   │       ├── domain/         ← entity, repo interface, container
│   │       ├── usecase/        ← one package per UC
│   │       │   └── {domain}/{operation}/usecase.go
│   │       ├── pblc/           ← reusable business logic
│   │       ├── infra/          ← postgres/, redis/ implementations
│   │       │   └── postgres/
│   │       ├── ctrl/           ← http/, cli/, consumer/
│   │       │   └── http/
│   │       └── embassy/        ← portal implementation
│   └── portal/                 ← cross-module interfaces
├── pkg/                        ← shared packages
├── migrations/                 ← goose migration SQL files
└── tests/
    ├── state/                  ← Given*, Get* helpers
    └── system/                 ← use case system tests
```

### Katman Kuralları

| Katman | Yer | Ne yapar |
|--------|-----|---------|
| Domain | `domain/` | Entity, value object, repo interface |
| Use Case | `usecase/` | Business logic, transaction sınırları |
| PBLC | `pblc/` | UC'ler arası paylaşılan business logic |
| Infra | `infra/` | Repo implementation, external client |
| Controller | `ctrl/` | HTTP handler — sadece parse + delegate |

- Business logic sadece UC veya PBLC'de
- Controller'da business logic kesinlikle yok
- SQL sadece infra/repository'de
- Modüller arası iletişim sadece Portal üzerinden

### API Tasarımı

- **Sadece GET** (query) ve **POST** (mutation) — PUT/PATCH/DELETE yok
- URL: `api/v1/{module}/{operation-id}` (kebab-case)
- GET: query param, POST: JSON body
- Path parameter yok — ID'ler body veya query'de

**Response formatları:**

```json
// Liste (non-paginated)
{ "content": [] }

// Paginated liste
{ "page_number": 1, "page_size": 20, "count": 150, "content": [] }

// Hata
{ "trace_id": "...", "error": { "code": "...", "message": "...", "cause": "..." } }
```

Header: `X-Trace-ID` her response'ta.

### Error Handling Kuralları

```go
// DOĞRU — ayrı satırda call ve check
err := doSomething()
if err != nil {
    return errx.Wrap(err)
}

// YANLIŞ — inline if-assignment
if err := doSomething(); err != nil { ... }
```

Final return'de inline wrap kullanılabilir:
```go
return role, errx.WrapWithTypeOnCodes(err, errx.T_Conflict, rbac.CodeRoleNameConflict)
```

### Verification — Her zaman çalıştır

```bash
make fmt          # önce format
make lint         # linting
make test         # unit tests (pkg/)
make test-system  # system tests (use case coverage)
```

**Bu üçü geçmeden kod teslim edilmez.**

### Agents (Go Blueprint)

| Agent | Ne zaman |
|-------|---------|
| `architect` | Implementation planla (read-only) |
| `system-analyst` | UC dokümanı, ERD yaz |
| `go-coder` | Kod implement et |
| `go-tester` | Test yaz |
| `reviewer` | Guideline'a karşı review |

Pipeline sırası: architect → system-analyst → go-coder + go-tester → reviewer

### Commands (Go Blueprint)

- `/generate-module` — yeni module scaffold
- `/review-module` — module review
- `/review-full` — full codebase review
- `/enhance-docs` — doc iyileştir

---

## Vue Frontend — Kurallar

### Klasör Yapısı

```
frontend/src/
├── api/
│   ├── client.ts          ← axios instance + token storage + 401 auto-refresh
│   ├── auth.ts            ← SSO login-url, token exchange, me, logout
│   └── {domain}.ts        ← one file per domain
├── stores/
│   ├── auth.ts            ← Pinia auth store
│   └── {domain}.ts
├── types/
│   ├── common.ts          ← ApiResponse<T>, PageResponse<T>, ErrorDetail
│   ├── auth.ts            ← Role, UserDto, AuthResponse
│   └── {domain}.ts        ← domain types (mirror backend DTOs)
├── router/
│   └── index.ts           ← routes + beforeEach guard
├── layouts/
│   ├── AppLayout.vue      ← sidebar + topbar (authenticated pages)
│   └── AuthLayout.vue     ← centered card (login, callback)
├── components/
│   ├── ui/                ← generic reusable components
│   └── {domain}/          ← domain-specific components
└── views/
    ├── auth/
    │   ├── LoginView.vue
    │   └── CallbackView.vue
    └── {domain}/
        └── {Domain}View.vue
```

### Mandatory Kurallar (ihlal edilemez)

```typescript
// Her component'ta
<script setup lang="ts">

// Tip her zaman explicit — any yok
const users = ref<UserDto[]>([])
async function fetchUser(id: string): Promise<UserDto> { ... }

// API çağrısı sadece src/api/*.ts üzerinden
const res = await usersApi.list({ page: 0, size: 20 })  // DOĞRU
await axios.get('/api/v1/users')                          // YANLIŞ

// Route meta
{ meta: { requiresAuth: true, roles: ['ADMIN'], layout: 'app' } }

// Rol kontrolü
if (auth.hasRole(['ADMIN'])) { ... }  // DOĞRU
if (auth.user?.role === 'admin') { }  // YANLIŞ
```

### API Dosyası Şablonu

```typescript
// src/api/{domain}.ts
import client from './client'
import type { ApiResponse, PageResponse } from '@/types/common'
import type { ItemDto } from '@/types/{domain}'

export const {domain}Api = {
  list: (params: { page?: number; size?: number }) =>
    client.get<ApiResponse<PageResponse<ItemDto>>>('/{module}/list-items', { params }),
  create: (body: CreateItemDto) =>
    client.post<ApiResponse<ItemDto>>('/{module}/create-item', body),
  // NOT: update = POST, delete = POST — REST değil
  update: (body: UpdateItemDto) =>
    client.post<ApiResponse<ItemDto>>('/{module}/update-item', body),
  delete: (body: { id: string }) =>
    client.post<ApiResponse<null>>('/{module}/delete-item', body),
}
```

### Pagination Şablonu

```typescript
const page    = ref(0)          // 0-based
const total   = ref(0)
const items   = ref<ItemDto[]>([])
const loading = ref(false)

async function load() {
  loading.value = true
  try {
    const res  = await itemsApi.list({ page: page.value, size: 20 })
    items.value = res.data.data!.content
    total.value = res.data.data!.totalElements
  } finally {
    loading.value = false
  }
}
```

### CSS — Predefined Classes (main.css)

```
Buton:   btn-primary, btn-secondary, btn-danger, btn-success, btn-icon
Form:    form-label, form-input, form-select, form-textarea, form-error
Layout:  card, data-table, page-header, page-title, filter-bar
Status:  chip-pending, chip-approved, chip-rejected, chip-info
```

- Inline `style=""` attribute yok
- Yeni CSS dosyası yok — yeni class `main.css @layer components`'e eklenir
- `console.log` committed code'da yok

### Yeni Feature Checklist (Vue)

- [ ] `src/types/{domain}.ts` — backend DTO'yu mirror et
- [ ] `src/api/{domain}.ts` — typed API calls
- [ ] `src/router/index.ts`'e route ekle (meta.roles ile)
- [ ] `AppLayout.vue`'ya `<SidebarItem>` ekle (nav link gerekiyorsa)
- [ ] `src/views/{Domain}View.vue` oluştur
- [ ] `loading` + `error` ref'leri kullan, silent error state bırakma
- [ ] Paginated: `page` + `totalElements` kullan

### Customisation Checklist (Proje Başında)

- [ ] `src/types/auth.ts` — `Role` union type'ı backend'e göre güncelle
- [ ] `src/main.ts` — `roleHomeMap` güncelle
- [ ] `src/router/index.ts` — placeholder route'ları değiştir
- [ ] `src/layouts/AppLayout.vue` — `<SidebarItem>` girdilerini ekle
- [ ] `src/layouts/AuthLayout.vue` — `appName` ve `appSubtitle` set et
- [ ] `index.html` — `<title>` güncelle
- [ ] `tailwind.config.js` — brand renkleri güncelle
- [ ] `package.json` — `name` alanını güncelle

---

## Docs Yapısı (Her Modül İçin)

```
docs/specs/modules/{module}/
├── overview.md              ← module amacı, entity'ler
├── ERD.md                   ← Mermaid ERD
└── usecases/{domain}/{operation}.md
```

UC dokümanı implementation'dan önce yazılır. Bu kural ihlal edilemez.

---

## Proje Modülleri

> Modül ekledikçe güncelle

| Modül | Amaç | Durum |
|-------|------|-------|
| auth | Kimlik doğrulama, session | Başlamadı |
| | | |

---

## Mevcut Use Case'ler

> UC ekledikçe güncelle

| Module | UC | Operation ID | Durum |
|--------|----|-------------|-------|
| | | | |

---

## Şu An Ne Üzerinde Çalışıyoruz

> Her sprint başında güncelle — Claude'un odak noktası

**Sprint hedefi:** ___________________________

**Aktif feature / UC:** ___________________________

**Bu sprint'te biten:**
- [ ]

**Engel / Bilinmeyen:**
-

---

## Sık Kullanılan Komutlar

```bash
# Backend
cd backend
make run              # infra up + go run
make lint             # golangci-lint
make fmt              # format
make test             # unit tests
make test-system      # system tests
make migrate-create   # yeni migration oluştur
make migrate-up       # migration çalıştır
make migrate-down     # geri al

# Frontend
cd frontend
npm run dev           # geliştirme sunucusu
npm run type-check    # TypeScript kontrol
npm run lint          # ESLint
npm run build         # production build
```
