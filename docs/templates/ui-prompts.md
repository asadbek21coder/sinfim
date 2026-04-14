# UI Prompt Üretme Rehberi

**Ne zaman kullanılır:** UX dokümanı ve tech-doc tamamlandı, şimdi Vue component'larını üretiyorsun.
**Stack:** Vue 3 + TypeScript + Pinia + Tailwind CSS v3 (`vue-blueprint-web`)
**Araç:** Claude Code (proje klasöründe CLAUDE.md hazır olmalı)

> Prompt yazmadan önce CLAUDE.md projeye kopyalanmış ve doldurulmuş olmalı.
> Claude projeyi zaten biliyor — her seferinde stack'i açıklamana gerek yok.

---

## Vue Blueprint'in Önemli Farklılıkları

Bu blueprint'te bazı şeyler alışılmışın dışında — prompta bunları yansıt:

| Konu | Vue Blueprint Kuralı |
|------|---------------------|
| API method | GET (query) + POST (mutation) — PUT/PATCH/DELETE yok |
| URL formatı | `/api/v1/{module}/{operation-id}` kebab-case |
| Pagination | `page` 0-based, response'ta `content` + `totalElements` |
| Auth | SSO/PKCE — basit form login değil |
| CSS | Predefined class'lar (`btn-primary`, `form-input`, `card` vb.) |
| Icons | Material Symbols — `<span class="material-symbols-outlined">` |
| Stil | Inline `style=""` yok, yeni CSS dosyası yok |
| Tip | `any` yok, her şey explicit tipli |

---

## Master Context Bloğu

Claude'a her büyük prompt'tan önce kısa bir context ver:

```
Proje CLAUDE.md'si mevcut. Ek bağlam:
- Şu an yapıyorum: [ne yapıyorsun]
- İlgili API: [hangi endpoint'ler — operation ID ve path]
- İlgili tip: [hangi DTO/interface]
```

---

## Prompt Şablonu: Yeni Ekran (View)

```
## Görev
src/views/{domain}/{EkranAdi}View.vue dosyasını oluştur.

## Bu Ekranın Amacı
[Tek cümle]

## Layout
AppLayout içinde / AuthLayout içinde / Tam sayfa

## İçerik (Yukarıdan Aşağı)
1. [PageHeader — başlık + sağda aksiyon butonu varsa]
2. [filter-bar — arama input + filtre dropdown]
3. [data-table / form / kart listesi]
4. [modal / drawer — varsa]

## Aksiyonlar
- [Buton adı] → [ne olur / nereye gider]
- [Form submit] → [hangi API çağrısı]

## State'ler
- Loading: skeleton / spinner
- Boş (no data): "[Mesaj]" + "[Aksiyon butonu]"
- Hata: error mesajı + retry

## API
- `src/api/{domain}.ts → {domainApi}.{method}(params)`
- Request: { page?: number; size?: number; search?: string }
- Response: PageResponse<{ItemDto}> → content + totalElements

## Store (gerekiyorsa)
- useAuthStore() → [hangi alan kullanılacak]

## Route
- Path: /{path}
- meta: { requiresAuth: true, roles: ['{ROLE}'], layout: 'app' }

## Dikkat
- btn-primary, form-input, card, data-table, page-header class'larını kullan
- Material Symbols ile ikon: <span class="material-symbols-outlined">icon_name</span>
- Pagination: page ref(0), totalElements ref(0), size=20
- Error: err?.response?.data?.error?.message ?? 'Beklenmeyen hata'
```

---

## Prompt Örnekleri

### Liste Ekranı

```
## Görev
src/views/users/UsersView.vue oluştur.

## Bu Ekranın Amacı
Admin kullanıcı listesini göster, ara, yeni kullanıcı oluştur.

## Layout
AppLayout içinde.

## İçerik
1. PageHeader: "Kullanıcılar" başlık, sağda "Yeni Kullanıcı" btn-primary
2. filter-bar: arama input (300ms debounce) + "Durum" select (Tümü / Aktif / Pasif)
3. data-table:
   - Kolonlar: Ad, E-posta, Durum (chip), Oluşturulma, İşlemler
   - İşlemler: "Düzenle" (router-link) + "Sil" (onay modal)
4. Silme onay modal: "Bu kullanıcıyı silmek istediğine emin misin?" + İptal + btn-danger

## State'ler
- Loading: 5 satır skeleton row
- Boş: "Henüz kullanıcı yok" + "İlk kullanıcıyı oluştur" btn-primary
- Hata: alert + "Tekrar Dene" butonu

## API
- GET list: usersApi.list({ page, size: 20, search, isActive })
  Response: PageResponse<UserDto> → content, totalElements
- POST delete: usersApi.delete({ id }) → onay sonrası

## Route
Path: /users, meta: { requiresAuth: true, roles: ['ADMIN'], layout: 'app' }

## Dikkat
- Silme başarıysa listeyi yenile
- Arama 300ms debounce
- Status chip: chip-approved (aktif) / chip-rejected (pasif)
```

---

### Form Ekranı (Create / Edit)

```
## Görev
src/views/users/UserFormView.vue oluştur.
Query param ?id varsa edit modu, yoksa create modu.

## Layout
AppLayout içinde.

## İçerik
1. PageHeader: "Yeni Kullanıcı" / "Kullanıcıyı Düzenle" (moda göre)
2. card içinde form:
   - Ad Soyad: form-input, required, min=2
   - E-posta: form-input type=email, required
   - Şifre: form-input type=password, required (create), opsiyonel (edit)
   - Durum: form-select, Aktif / Pasif (sadece edit modunda)
3. Alt: "İptal" btn-secondary (/users'a git) + "Kaydet" btn-primary (loading state ile)

## State'ler
- Edit modunda başlangıç: veriyi fetch et, skeleton göster
- Submit loading: buton spinner + disabled
- Hata (422): form-error altında her field'ın yanında
- Başarı: /users'a redirect

## API
- Fetch (edit): GET usersApi.getById({ id })
- Create: POST usersApi.create(body)
- Edit: POST usersApi.update(body) — id body'de

## Route
Path: /users/form, meta: { requiresAuth: true, roles: ['ADMIN'], layout: 'app' }
```

---

### API Dosyası

```
## Görev
src/api/users.ts dosyasını oluştur.

## Endpoint'ler (tech-doc'tan)
- GET /api/v1/users/list-users → params: { page?, size?, search?, isActive? }
  Response: ApiResponse<PageResponse<UserDto>>
- GET /api/v1/users/get-user → params: { id }
  Response: ApiResponse<UserDto>
- POST /api/v1/users/create-user → body: CreateUserDto
  Response: ApiResponse<UserDto>
- POST /api/v1/users/update-user → body: UpdateUserDto (id dahil)
  Response: ApiResponse<UserDto>
- POST /api/v1/users/delete-user → body: { id }
  Response: ApiResponse<null>

## Tipler
src/types/user.ts'e şunları ekle:
- UserDto: { id, username, email, isActive, createdAt }
- CreateUserDto: { username, email, password }
- UpdateUserDto: { id, username?, email?, password?, isActive? }

## Dikkat
- client'ı src/api/client.ts'den import et
- Response tipleri ApiResponse<T> ve PageResponse<T> common.ts'den import et
```

---

### Reusable Component

```
## Görev
src/components/ui/ConfirmModal.vue reusable component oluştur.

## Props
- isOpen: boolean
- title: string
- message: string
- confirmLabel?: string (default: "Onayla")
- confirmVariant?: 'danger' | 'primary' (default: 'danger')
- loading?: boolean (default: false)

## Emits
- confirm: () — kullanıcı onayladı
- cancel: () — iptal / dışarı tıklandı

## Görünüm
- Overlay backdrop (koyu)
- Ortalanmış beyaz card: başlık + mesaj + İptal (btn-secondary) + Onayla (btn-danger veya btn-primary)
- loading true ise buton spinner + disabled

## Kullanım
<ConfirmModal
  :is-open="showModal"
  title="Kullanıcıyı Sil"
  message="Bu kullanıcıyı silmek istediğine emin misin?"
  confirm-variant="danger"
  :loading="isDeleting"
  @confirm="handleDelete"
  @cancel="showModal = false"
/>
```

---

## Ekran Üretim Sırası

1. **Auth ekranları** — blueprint'te zaten mevcut, sadece customise et
2. **AppLayout** — sidebar item'ları ekle
3. **Dashboard** — önce mock data
4. **Ana resource CRUD** — liste → form sırası
5. **Reusable component'lar** — birden fazla yerde gerekince extract et
6. **Settings / Profile** — son

---

## Üretim Sonrası Kontrol

Her component için:
- [ ] `<script setup lang="ts">` kullanıyor
- [ ] `any` tipi yok
- [ ] Loading state var
- [ ] Empty state var
- [ ] Error state var (`err?.response?.data?.error?.message`)
- [ ] API çağrısı `src/api/` üzerinden
- [ ] `btn-primary`, `form-input`, `card` gibi predefined class kullanıyor
- [ ] Inline style yok
- [ ] Route meta'da `requiresAuth` ve `roles` var
- [ ] `console.log` yok

**Sonraki adım:** `templates/dev-kickoff.md`
