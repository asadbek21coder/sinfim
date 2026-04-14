# UC: submit-quiz

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `submit-quiz` |
| Method | `POST` |
| Path | `/api/v1/homework/submit-quiz` |
| Actor | Student |
| Modul | `homework` |

## Amac

Student single-choice veya multiple-choice test/quiz cevaplarini teslim eder ve otomatik puan alir.

## Request

```json
{
  "organization_id": "uuid",
  "class_id": "uuid",
  "homework_task_id": "uuid",
  "answers": [
    {
      "question_id": "uuid",
      "selected_option_ids": ["uuid"]
    }
  ]
}
```

## Validasyon

| Field | Kural |
|-------|-------|
| `organization_id` | required |
| `class_id` | required |
| `homework_task_id` | required |
| `answers` | required, min 1 |
| `question_id` | required |
| `selected_option_ids` | required |

## Business Rules

1. Actor `STUDENT` role'e sahip olmali.
2. Student class'a enrolled olmali.
3. Student class-level access active olmali.
4. Homework task type `quiz` bo'lishi kerak.
5. Lesson student icin acilmis olmali.
6. Sistem correct option'lara gore score hesaplar.
7. MVP'de attempt summary saklanir: score, max_score, submitted_at.
8. Per-answer detay tablosu MVP disinda kalabilir.

## Response 200

```json
{
  "attempt_id": "uuid",
  "score": 8,
  "max_score": 10,
  "percentage": 80,
  "submitted_at": "2026-04-14T10:00:00Z"
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Student bu quiz'e access qilolmaydi |
| `HOMEWORK_NOT_FOUND` | 404 | Quiz task bulunamadi |
| `ACCESS_DENIED` | 403 | Access active degil |
| `LESSON_LOCKED` | 403 | Ders henuz acilmadi |
| `QUESTION_NOT_FOUND` | 404 | Question bulunamadi |
| `INVALID_QUIZ_ANSWER` | 422 | Answer formati hatali |

## Side Effects

- `homework.quiz_attempts` kaydi olusur.
- Student dashboard/lesson detail quiz sonucunu gosterir.

## Test Senaryolari

1. Student quiz submit eder ve score hesaplanir.
2. Access pending ise `ACCESS_DENIED` doner.
3. Quiz olmayan homework task icin `HOMEWORK_NOT_FOUND` veya validation error doner.
4. Invalid question_id icin `QUESTION_NOT_FOUND` doner.
5. Response score ve max_score doner.
