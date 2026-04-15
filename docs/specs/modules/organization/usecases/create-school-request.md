# UC: create-school-request

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `create-school-request` |
| Method | `POST` |
| Path | `/api/v1/organization/create-school-request` |
| Actor | Public |
| Modul | `organization` |

## Amac

Okul/egitim merkezi acmak isteyen ziyaretci public landing sayfasindan talep birakir. Bu talep direkt organization olusturmaz; superadmin tarafindan incelenir.

## Request

```json
{
  "full_name": "Ali Valiyev",
  "phone_number": "+998901234567",
  "school_name": "My School",
  "category": "language",
  "student_count": 120,
  "note": "Telegram orqali kurs sotamiz"
}
```

## Validasyon

| Field | Kural |
|-------|-------|
| `full_name` | required, min 2 |
| `phone_number` | required, telefon formatida |
| `school_name` | required, min 2 |
| `category` | optional |
| `student_count` | optional, >= 0 |
| `note` | optional, max 2000 |

## Business Rules

1. Public endpoint, login gerektirmez.
2. Telefon raqami normalize edilir.
3. Ayni telefon + ayni school_name icin cok yakin zamanda acik talep varsa duplicate olarak isaretlenebilir.
4. Talep status `new` olarak kaydedilir.
5. Superadmin listesinde gorunur.
6. Bu UC organization veya user olusturmaz.

## Response 200

```json
{
  "id": "uuid",
  "status": "new",
  "message": "Arizangiz qabul qilindi. Platform administratori siz bilan bog'lanadi."
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `VALIDATION_ERROR` | 422 | Request formati hatali |
| `SCHOOL_REQUEST_DUPLICATE` | 409 | Ayni talep zaten beklemede |

## Side Effects

- `organization.school_requests` kaydi olusur.

## Test Senaryolari

1. Valid request ile school request olusur.
2. Eksik telefon ile `VALIDATION_ERROR` doner.
3. Eksik school name ile `VALIDATION_ERROR` doner.
4. Duplicate acik talep varsa `SCHOOL_REQUEST_DUPLICATE` doner veya mevcut request bilgisi doner.
5. UC organization olusturmaz.
