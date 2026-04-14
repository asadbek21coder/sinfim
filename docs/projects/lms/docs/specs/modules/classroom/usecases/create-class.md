# UC: create-class

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `create-class` |
| Method | `POST` |
| Path | `/api/v1/classroom/create-class` |
| Actor | Owner / Teacher |
| Modul | `classroom` |

## Amac

Bir course icin gercek ogrenci cohort'u olan class/group olusturmak.

## Request

```json
{
  "organization_id": "uuid",
  "course_id": "uuid",
  "name": "Russian A1 - May 2026",
  "start_date": "2026-05-01",
  "lesson_cadence": "every_other_day",
  "mentor_user_ids": ["uuid"]
}
```

## Validasyon

| Field | Kural |
|-------|-------|
| `organization_id` | required |
| `course_id` | required |
| `name` | required, min 2 |
| `start_date` | optional |
| `lesson_cadence` | `daily`, `every_other_day`, `weekly_3`, `manual` |
| `mentor_user_ids` | optional |

## Business Rules

1. Actor ilgili organization'da `OWNER` veya `TEACHER` role'e sahip olmali.
2. Course organization icinde mevcut olmali.
3. Class varsayilan status `active` olur.
4. Mentor atanacaksa mentor user ilgili organization'da `MENTOR` role'e sahip olmali.
5. Class olusturulunca ogrenci otomatik eklenmez.
6. Class-level access modeli kullanilir.

## Response 200

```json
{
  "id": "uuid",
  "organization_id": "uuid",
  "course_id": "uuid",
  "name": "Russian A1 - May 2026",
  "start_date": "2026-05-01",
  "lesson_cadence": "every_other_day",
  "status": "active",
  "mentor_count": 1,
  "student_count": 0
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Yetki yok |
| `COURSE_NOT_FOUND` | 404 | Course bulunamadi |
| `MENTOR_NOT_FOUND` | 404 | Mentor bulunamadi |
| `MENTOR_ROLE_REQUIRED` | 422 | User mentor role'e sahip degil |
| `VALIDATION_ERROR` | 422 | Request formati hatali |

## Side Effects

- `classroom.classes` kaydi olusur.
- Mentor atanirsa `classroom.class_mentors` kaydi olusur.

## Test Senaryolari

1. Owner valid request ile class olusturur.
2. Teacher valid request ile class olusturur.
3. Mentor class olusturamaz, `FORBIDDEN` doner.
4. Mentor role'u olmayan user mentor olarak atanamaz.
5. Class ogrencisiz olusabilir.
