# UC: create-organization

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `create-organization` |
| Method | `POST` |
| Path | `/api/v1/organization/create-organization` |
| Actor | Platform Admin |
| Modul | `organization` |

## Amac

Platform superadmin yangi okul/brand organization olusturur, owner kullanicisini baglar ve owner'a organization access verir.

MVP'de public kullanici kendi kendine organization olusturamaz.

## Request

```json
{
  "name": "My School",
  "slug": "my-school",
  "description": "Online kurs markazi",
  "logo_url": "https://cdn.example.com/logo.png",
  "category": "language",
  "contact_phone": "+998901234567",
  "telegram_url": "https://t.me/myschool",
  "is_demo": false,
  "owner": {
    "full_name": "Ali Valiyev",
    "phone_number": "+998901112233",
    "temporary_password": "TempPass123"
  }
}
```

### Validasyon

| Field | Kural |
|-------|-------|
| `name` | required, min 2 |
| `slug` | required, lowercase kebab-case, unique |
| `owner.full_name` | required |
| `owner.phone_number` | required |
| `owner.temporary_password` | required, min 8 |
| `logo_url` | optional |
| `is_demo` | optional, default false |

## Business Rules

1. Actor `PLATFORM_ADMIN` role'e sahip olmali.
2. Slug normalize edilir ve unique kontrol edilir.
3. Slug alinmissa `SLUG_ALREADY_TAKEN` doner.
4. Organization kaydi olusturulur.
5. Owner telefon raqami ile mevcut user aranir.
6. Owner user yoksa `auth` portal uzerinden user olusturulur, password hash'lenir, `must_change_password = true` set edilir.
7. Owner user varsa yeni organization membership'i eklenir.
8. Owner'a `OWNER` role ile membership verilir.
9. Response organization ve owner bilgisi doner.

## Response 200

```json
{
  "organization": {
    "id": "uuid",
    "name": "My School",
    "slug": "my-school",
    "description": "Online kurs markazi",
    "logo_url": "https://cdn.example.com/logo.png",
    "public_status": "draft",
    "is_demo": false
  },
  "owner": {
    "id": "uuid",
    "full_name": "Ali Valiyev",
    "phone_number": "+998901112233",
    "role": "OWNER",
    "must_change_password": true
  }
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Platform admin degil |
| `SLUG_ALREADY_TAKEN` | 409 | Slug daha once alinmis |
| `VALIDATION_ERROR` | 422 | Request formati hatali |
| `OWNER_ALREADY_MEMBER` | 409 | Owner bu organization'a zaten bagli |

## Side Effects

- `organization.organizations` kaydi olusur.
- `auth.users` owner yoksa olusur.
- `auth.user_memberships` icinde owner membership olusur.

## Test Senaryolari

1. Platform admin valid request ile organization olusturur.
2. Slug alinmissa `SLUG_ALREADY_TAKEN` doner.
3. Platform admin olmayan kullanici `FORBIDDEN` alir.
4. Owner user yoksa yeni user + membership olusur.
5. Owner user varsa sadece membership eklenir.
6. `is_demo = true` ile demo organization olusur.
7. Invalid slug ile `VALIDATION_ERROR` doner.
