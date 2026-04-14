# UC: create-course

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `create-course` |
| Method | `POST` |
| Path | `/api/v1/catalog/create-course` |
| Actor | Owner / Teacher |
| Modul | `catalog` |

## Amac

Organization icinde tekrar kullanilabilir course content package olusturmak.

Course ogrenci operasyonu degildir; ogrenciler class/group uzerinden yonetilir.

## Request

```json
{
  "organization_id": "uuid",
  "title": "Russian A1",
  "slug": "russian-a1",
  "description": "Boshlangich rus tili kursi",
  "category": "language",
  "level": "A1",
  "public_status": "draft"
}
```

## Validasyon

| Field | Kural |
|-------|-------|
| `organization_id` | required |
| `title` | required, min 2 |
| `slug` | required, lowercase kebab-case, organization icinde unique |
| `description` | optional |
| `category` | optional |
| `level` | optional |
| `public_status` | `draft` veya `published` |

## Business Rules

1. Actor ilgili organization'da `OWNER` veya `TEACHER` role'e sahip olmali.
2. Organization aktif olmali.
3. Slug organization icinde unique olmali.
4. Course status varsayilan `draft` olur.
5. Public status varsayilan `draft` olur.
6. Course olusturulunca henuz lesson/class olusturulmaz.

## Response 200

```json
{
  "id": "uuid",
  "organization_id": "uuid",
  "title": "Russian A1",
  "slug": "russian-a1",
  "description": "Boshlangich rus tili kursi",
  "category": "language",
  "level": "A1",
  "status": "draft",
  "public_status": "draft"
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Organization'da yetki yok |
| `ORGANIZATION_NOT_FOUND` | 404 | Organization bulunamadi |
| `COURSE_SLUG_ALREADY_TAKEN` | 409 | Slug organization icinde kullaniliyor |
| `VALIDATION_ERROR` | 422 | Request formati hatali |

## Side Effects

- `catalog.courses` kaydi olusur.

## Test Senaryolari

1. Owner valid request ile course olusturur.
2. Teacher valid request ile course olusturur.
3. Mentor course olusturamaz, `FORBIDDEN` doner.
4. Duplicate slug ile `COURSE_SLUG_ALREADY_TAKEN` doner.
5. Course olusturuldugunda lesson/class otomatik olusmaz.
