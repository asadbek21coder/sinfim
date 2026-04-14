# UC: mark-lesson-completed

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `mark-lesson-completed` |
| Method | `POST` |
| Path | `/api/v1/learning/mark-lesson-completed` |
| Actor | Student |
| Modul | `learning` |

## Amac

Student dersi tamamlandi olarak isaretler. Bu progress hesaplamasinda kullanilir.

## Request

```json
{
  "organization_id": "uuid",
  "class_id": "uuid",
  "lesson_id": "uuid"
}
```

## Business Rules

1. Actor `STUDENT` role'e sahip olmali.
2. Student class'a enrolled olmali.
3. Student class-level access active olmali.
4. Lesson class course'una ait olmali.
5. Lesson schedule'e gore available olmali.
6. Daha once completed ise idempotent davranir; basarili response doner.
7. MVP data modelinde lesson completion tablosu gerekiyorsa `learning.lesson_completions` eklenir. Bu tablo tech-doc'a eklenmelidir.

## Response 200

```json
{
  "lesson_id": "uuid",
  "class_id": "uuid",
  "student_user_id": "uuid",
  "completed": true,
  "completed_at": "2026-04-14T10:00:00Z"
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Student bu lesson uchun yetkili emas |
| `LESSON_NOT_FOUND` | 404 | Lesson bulunamadi |
| `ACCESS_DENIED` | 403 | Access active degil |
| `LESSON_LOCKED` | 403 | Ders henuz acilmadi |
| `ENROLLMENT_NOT_FOUND` | 404 | Student class'a kayitli degil |

## Side Effects

- `learning.lesson_completions` kaydi olusur veya mevcut kayit korunur.

## Test Senaryolari

1. Student available lesson'i completed isaretler.
2. Ayni lesson ikinci kez completed isaretlenirse idempotent basarili doner.
3. Locked lesson icin `LESSON_LOCKED` doner.
4. Access pending ise `ACCESS_DENIED` doner.
5. Progress percentage sonraki dashboard response'ta guncellenir.
