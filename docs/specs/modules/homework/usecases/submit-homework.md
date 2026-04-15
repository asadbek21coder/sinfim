# UC: submit-homework

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `submit-homework` |
| Method | `POST` |
| Path | `/api/v1/homework/submit-homework` |
| Actor | Student |
| Modul | `homework` |

## Amac

Student yazili, dosya/fotograf veya audio turundeki odevini teslim eder.

Quiz/test teslimi bu UC degil; `submit-quiz` ayridir.

## Request

```json
{
  "organization_id": "uuid",
  "class_id": "uuid",
  "homework_task_id": "uuid",
  "submission_type": "text",
  "text_answer": "Mening javobim...",
  "file_url": null,
  "audio_url": null
}
```

## Validasyon

| Field | Kural |
|-------|-------|
| `organization_id` | required |
| `class_id` | required |
| `homework_task_id` | required |
| `submission_type` | `text`, `file`, `audio` |
| `text_answer` | `submission_type=text` ise required |
| `file_url` | `submission_type=file` ise required |
| `audio_url` | `submission_type=audio` ise required |

## Business Rules

1. Actor `STUDENT` role'e sahip olmali.
2. Student ilgili class'a enrolled olmali.
3. Student'in class-level access'i active olmali.
4. Homework task ilgili lesson'a ait olmali.
5. Lesson ilgili course/class schedule icinde student icin acilmis olmali.
6. Homework type request submission type ile uyumlu olmali.
7. Daha once submission varsa ve status `reviewed` degilse update edilebilir.
8. Status varsayilan `submitted` olur.
9. Mentor inbox'ina duser.

## Response 200

```json
{
  "id": "uuid",
  "homework_task_id": "uuid",
  "class_id": "uuid",
  "student_user_id": "uuid",
  "submission_type": "text",
  "status": "submitted",
  "submitted_at": "2026-04-14T10:00:00Z"
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Student bu class/homework'e access qilolmaydi |
| `HOMEWORK_NOT_FOUND` | 404 | Homework task bulunamadi |
| `CLASS_NOT_FOUND` | 404 | Class bulunamadi |
| `ACCESS_DENIED` | 403 | Access active degil |
| `LESSON_LOCKED` | 403 | Ders henuz acilmadi |
| `SUBMISSION_ALREADY_REVIEWED` | 409 | Reviewed submission update edilemez |
| `VALIDATION_ERROR` | 422 | Request formati hatali |

## Side Effects

- `homework.homework_submissions` kaydi olusur veya guncellenir.
- Mentor homework inbox'inda yeni/bekleyen submission olarak gorunur.

## Test Senaryolari

1. Active access'e sahip student text homework teslim eder.
2. File homework icin file_url yoksa `VALIDATION_ERROR` doner.
3. Access pending ise `ACCESS_DENIED` doner.
4. Lesson locked ise `LESSON_LOCKED` doner.
5. Reviewed submission tekrar guncellenemez.
6. Submitted ama reviewed olmayan submission guncellenebilir.
