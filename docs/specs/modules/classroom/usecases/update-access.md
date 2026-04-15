# UC: update-access

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `update-access` |
| Method | `POST` |
| Path | `/api/v1/classroom/update-access` |
| Actor | Owner / Teacher / Mentor |
| Modul | `classroom` |

## Amac

Ogrencinin class/group icindeki access ve payment status'unu manuel guncellemek.

MVP'de access class-level olarak yonetilir.

## Request

```json
{
  "organization_id": "uuid",
  "class_id": "uuid",
  "student_user_id": "uuid",
  "access_status": "active",
  "payment_status": "confirmed",
  "note": "IBAN orqali to'lov tasdiqlandi"
}
```

## Validasyon

| Field | Kural |
|-------|-------|
| `organization_id` | required |
| `class_id` | required |
| `student_user_id` | required |
| `access_status` | `pending`, `active`, `blocked` |
| `payment_status` | `unknown`, `pending`, `confirmed`, `rejected` |
| `note` | optional, max 2000 |

## Business Rules

1. Actor ilgili organization'da `OWNER`, `TEACHER` veya class'a atanmis `MENTOR` olmali.
2. Class organization icinde mevcut olmali.
3. Student class'a enrolled olmali.
4. Access grant yoksa olusturulur; varsa guncellenir.
5. `access_status = active` oldugunda `granted_by` ve `granted_at` set edilir.
6. `access_status = blocked` oldugunda student learning ekranlarinda class dersleri kilitlenir.

## Response 200

```json
{
  "class_id": "uuid",
  "student_user_id": "uuid",
  "access_status": "active",
  "payment_status": "confirmed",
  "note": "IBAN orqali to'lov tasdiqlandi",
  "granted_by": "uuid",
  "granted_at": "2026-04-14T10:00:00Z"
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Actor access guncelleyemez |
| `CLASS_NOT_FOUND` | 404 | Class bulunamadi |
| `ENROLLMENT_NOT_FOUND` | 404 | Student class'a kayitli degil |
| `VALIDATION_ERROR` | 422 | Request formati hatali |

## Side Effects

- `classroom.access_grants` kaydi olusturulur veya guncellenir.
- Learning ekranlarinda access durumu degisir.

## Test Senaryolari

1. Owner student access'i active yapar.
2. Mentor sadece atanmis class icin access gunceller.
3. Class'a kayitli olmayan student icin `ENROLLMENT_NOT_FOUND` doner.
4. Blocked access sonrasi student lesson detail'da locked state gorur.
5. Active access icin `granted_by` ve `granted_at` set edilir.
