# Teknik Dokümantasyon Rehberi

**Ne zaman kullanılır:** UX dokümanı tamamlandı, şimdi teknik kararları netleştiriyorsun.
**Süre:** 1–2 saat
**Stack:** Go enterprise blueprint (backend) + Vue blueprint web (frontend)
**Çıktı:** Modül listesi, data model, API contract, UC listesi.

> Bu doküman Claude'un kod yazarken referans alacağı teknik sözleşmedir.
> UC dokümanı yazılmadan implementation başlamaz — bu kural ihlal edilemez.

---

## Proje Bilgisi

> **Proje:** ___________________________
> **Backend:** Go — `go-enterprise-blueprint-main`
> **Frontend:** Vue.js 3 — `vue-blueprint-web`
> **Database:** PostgreSQL (`timestamptz` her timestamp için)
> **Auth:** SSO / PKCE (OneID veya OAuth2)
> **Deploy:** [Railway / Fly.io / VPS]

---

## Bölüm 1: Modül Listesi

Go backend modüler mimari kullanıyor. Her modül bağımsız.
Modüller sadece Portal interface'leri üzerinden haberleşir.

| Modül | Sorumluluğu | Bağlı Modüller (portal üzerinden) |
|-------|-------------|----------------------------------|
| `auth` | Kimlik doğrulama, session yönetimi | — |
| | | |
| | | |

---

## Bölüm 2: Veri Modeli (ERD)

### Kurallar
- Her tablo: `id UUID PRIMARY KEY DEFAULT gen_random_uuid()`, `created_at timestamptz`, `updated_at timestamptz`
- Soft delete gerekiyorsa: `deleted_at timestamptz NULL`
- Timestamp tipi: her zaman `timestamptz` — `timestamp` değil
- Cross-module join yok — her modül kendi datasına sahip
- Migration adları: `{module}_init_schema`, `{module}_{aciklama}`

### Modül: `auth`

```sql
-- migrations/{timestamp}_auth_init_schema.sql
-- +goose Up

CREATE TABLE auth.admins (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username        VARCHAR(255) UNIQUE NOT NULL,
    password_hash   VARCHAR(255) NOT NULL,
    is_active       BOOLEAN DEFAULT true NOT NULL,
    last_active_at  timestamptz,
    created_at      timestamptz DEFAULT NOW() NOT NULL,
    updated_at      timestamptz DEFAULT NOW() NOT NULL
);

CREATE TABLE auth.sessions (
    id                      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    admin_id                UUID NOT NULL REFERENCES auth.admins(id),
    access_token            VARCHAR(512) NOT NULL,
    refresh_token           VARCHAR(512) NOT NULL,
    access_token_expires_at  timestamptz NOT NULL,
    refresh_token_expires_at timestamptz NOT NULL,
    created_at              timestamptz DEFAULT NOW() NOT NULL
);

CREATE INDEX idx_sessions_admin_id ON auth.sessions(admin_id);
CREATE INDEX idx_sessions_refresh_token ON auth.sessions(refresh_token);

-- +goose Down
DROP TABLE IF EXISTS auth.sessions;
DROP TABLE IF EXISTS auth.admins;
```

### Modül: `[module_adi]`

```sql
-- migrations/{timestamp}_{module}_init_schema.sql
-- +goose Up

CREATE TABLE {module}.{table} (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- alanlar
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS {module}.{table};
```

### ERD (Mermaid)

```mermaid
erDiagram
    admins {
        uuid id PK
        varchar username
        varchar password_hash
        boolean is_active
        timestamptz last_active_at
        timestamptz created_at
        timestamptz updated_at
    }
    sessions {
        uuid id PK
        uuid admin_id FK
        varchar access_token
        varchar refresh_token
        timestamptz refresh_token_expires_at
        timestamptz created_at
    }
    admins ||--o{ sessions : "has"
```

---

## Bölüm 3: API Contract

### Önemli Kural: REST Değil

| ✓ | ✗ |
|---|---|
| GET /api/v1/{module}/list-items | GET /api/v1/items |
| POST /api/v1/{module}/create-item | POST /api/v1/items |
| POST /api/v1/{module}/update-item | PUT /api/v1/items/:id |
| POST /api/v1/{module}/delete-item | DELETE /api/v1/items/:id |

- **GET** = query (sadece okuma) → query param ile
- **POST** = mutation (yazma) → JSON body ile
- **Path parameter yok** — ID'ler body veya query'de
- Operation ID = use case adı kebab-case

### Response Formatları

```json
// Başarılı liste (paginated)
{
  "page_number": 1,
  "page_size": 20,
  "count": 150,
  "content": [...]
}

// Başarılı liste (non-paginated)
{ "content": [...] }

// Başarılı tek kayıt (UC'e göre değişir, direkt entity de dönebilir)
{ "id": "...", "username": "..." }

// Hata
{
  "trace_id": "abc123",
  "error": {
    "code": "INCORRECT_CREDENTIALS",
    "message": "username or password is incorrect",
    "cause": "password"
  }
}
```

Header: `X-Trace-ID` her response'ta.

### Use Case Listesi

> Her UC için `docs/specs/modules/{module}/usecases/{domain}/{operation}.md` yazılmalı.

#### Modül: `auth`

| Operation ID | Method | Path | Actor | Açıklama |
|-------------|--------|------|-------|---------|
| `admin-login` | POST | `/api/v1/auth/admin-login` | Public | Giriş yap, session oluştur |
| `admin-logout` | POST | `/api/v1/auth/admin-logout` | Admin | Oturumu kapat |
| `refresh-token` | POST | `/api/v1/auth/refresh-token` | Public | Token yenile |
| `get-me` | GET | `/api/v1/auth/get-me` | Admin | Mevcut kullanıcıyı getir |

#### Modül: `[module]`

| Operation ID | Method | Path | Actor | Açıklama |
|-------------|--------|------|-------|---------|
| | | | | |

### Request / Response Örnekleri

#### POST /api/v1/auth/admin-login

```json
// Request body
{
  "username": "string", // required, min=3
  "password": "string"  // required, min=8
}

// Response 200
{
  "access_token": "eyJ...",
  "refresh_token": "eyJ...",
  "admin": {
    "id": "uuid",
    "username": "john",
    "is_active": true
  }
}

// Response 422 — validasyon hatası
{
  "trace_id": "...",
  "error": {
    "code": "INCORRECT_CREDENTIALS",
    "message": "username or password is incorrect"
  }
}
```

---

## Bölüm 4: Error Code Kataloğu

| Modül | Code | Tip | Açıklama |
|-------|------|-----|---------|
| auth | `ADMIN_NOT_FOUND` | NotFound | Admin bulunamadı |
| auth | `INCORRECT_CREDENTIALS` | Validation | Yanlış kullanıcı adı veya şifre |
| auth | `SESSION_NOT_FOUND` | NotFound | Session bulunamadı |
| auth | `ADMIN_INACTIVE` | Validation | Admin pasif durumdadır |
| | | | |

---

## Bölüm 5: Portal (Cross-Module) Arayüzleri

Modüller sadece burada tanımlanan interface'ler üzerinden haberleşir.

```
internal/portal/{module}/portal.go  ← interface tanımı
internal/modules/{module}/embassy/  ← implementation
```

| Portal | Sağlayan Modül | Metodlar |
|--------|---------------|---------|
| | | |

---

## Bölüm 6: Frontend API Dosyaları

Vue blueprint convention'ına göre — her domain için bir dosya.

```typescript
// src/api/auth.ts — blueprint'teki mevcut auth.ts'i extend et
// src/api/{domain}.ts — yeni domain'ler için

export const {domain}Api = {
  // GET — query operations
  list: (params: { page?: number; size?: number; search?: string }) =>
    client.get<ApiResponse<PageResponse<ItemDto>>>('/{module}/list-items', { params }),

  // POST — mutation operations
  create: (body: CreateItemDto) =>
    client.post<ApiResponse<ItemDto>>('/{module}/create-item', body),
  update: (body: UpdateItemDto) =>      // id body'de
    client.post<ApiResponse<ItemDto>>('/{module}/update-item', body),
  delete: (body: { id: string }) =>    // id body'de
    client.post<ApiResponse<null>>('/{module}/delete-item', body),
}
```

---

## Bölüm 7: Environment Variables

### Backend `config/`

```yaml
# config/development.yaml
server:
  port: 8080

database:
  url: "postgresql://dev:dev@localhost:5432/{proje}?sslmode=disable"

jwt:  # Eğer custom JWT kullanılıyorsa
  access_secret: changeme
  access_expire: 15m
  refresh_expire: 168h
```

### Frontend `.env.example`

```env
VITE_API_BASE_URL=/api/v1
```

---

## ✓ Çıkış Kriteri

- [ ] Modül listesi belirlendi
- [ ] Her modül için tablo SQL'leri yazıldı (`timestamptz` kullanıldı)
- [ ] ERD çizildi
- [ ] Tüm UC'ler listelendi (operation ID, method, path)
- [ ] En az 1 UC için request/response örneği yazıldı
- [ ] Error code kataloğu başlatıldı
- [ ] Frontend API dosya yapısı planlandı

**Sonraki adım:** `new-project-template/CLAUDE.md` → projeye kopyala ve doldur, ardından `templates/ui-prompts.md`
