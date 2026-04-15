# UC: add-student

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `add-student` |
| Method | `POST` |
| Path | `/api/v1/classroom/add-student` |
| Actor | Owner / Teacher / Mentor |
| Modul | `classroom` |

## Amac

Owner/teacher/mentor kabul edilen ogrenciyi class/group'a manuel olarak ekler. Gerekirse student user olusturulur ve ilk giris icin gecici password/kod uretilir.

## Request

```json
{
  "organization_id": "uuid",
  "class_id": "uuid",
  "full_name": "Vali Aliyev",
  "phone_number": "+998901234567",
  "temporary_password": "TempPass123",
  "access_status": "pending",
  "payment_status": "unknown"
}
```

## Validasyon

| Field | Kural |
|-------|-------|
| `organization_id` | required |
| `class_id` | required |
| `full_name` | required, min 2 |
| `phone_number` | required |
| `temporary_password` | optional; yoksa sistem uretir |
| `access_status` | `pending`, `active`, `blocked` |
| `payment_status` | `unknown`, `pending`, `confirmed`, `rejected` |

## Business Rules

1. Actor ilgili organization'da `OWNER`, `TEACHER` veya class'a atanmis `MENTOR` olmali.
2. Class organization icinde mevcut olmali.
3. Telefon raqami normalize edilir.
4. User yoksa `auth` portal ile student user olusturulur, `must_change_password = true`.
5. User varsa organization membership kontrol edilir; yoksa `STUDENT` membership eklenir.
6. Student class'a daha once eklenmisse `STUDENT_ALREADY_ENROLLED` doner.
7. Enrollment olusturulur.
8. Access grant olusturulur. MVP'de class-level access kullanilir.

## Response 200

```json
{
  "student": {
    "id": "uuid",
    "full_name": "Vali Aliyev",
    "phone_number": "+998901234567",
    "must_change_password": true
  },
  "enrollment": {
    "id": "uuid",
    "class_id": "uuid",
    "status": "active"
  },
  "access": {
    "access_status": "pending",
    "payment_status": "unknown"
  },
  "temporary_password_generated": true
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Actor bu class'a ogrenci ekleyemez |
| `CLASS_NOT_FOUND` | 404 | Class bulunamadi |
| `STUDENT_ALREADY_ENROLLED` | 409 | Ogrenci zaten class'ta |
| `VALIDATION_ERROR` | 422 | Request formati hatali |

## Side Effects

- Gerekirse `auth.users` ve `auth.user_memberships` olusur.
- `classroom.enrollments` kaydi olusur.
- `classroom.access_grants` kaydi olusur.

## Test Senaryolari

1. Owner yeni telefon raqami ile student olusturup class'a ekler.
2. Mentor sadece atanmis oldugu class'a student ekleyebilir.
3. Ayni student ikinci kez eklenirse `STUDENT_ALREADY_ENROLLED` doner.
4. Existing user icin yeni user olusturulmaz, membership/enrollment eklenir.
5. Access status default pending olur.
