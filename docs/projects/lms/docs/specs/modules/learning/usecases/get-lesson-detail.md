# UC: get-lesson-detail

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `get-lesson-detail` |
| Method | `GET` |
| Path | `/api/v1/learning/get-lesson-detail` |
| Actor | Student |
| Modul | `learning` |

## Amac

Student tek dersin video, material, homework ve submission/feedback durumunu gorur.

## Request Query

```text
organization_id=uuid
class_id=uuid
lesson_id=uuid
```

## Business Rules

1. Actor `STUDENT` role'e sahip olmali.
2. Student class'a enrolled olmali.
3. Student class-level access active olmali.
4. Lesson class'in course'una ait olmali.
5. Lesson schedule'e gore acik degilse `LESSON_LOCKED` doner veya locked response doner. MVP karari: locked response 200 ile donsun, frontend kilitli state gosterir.
6. Video stream reference ve materials doner.
7. Homework task ve mevcut submission/quiz attempt status doner.

## Response 200

```json
{
  "lesson": {
    "id": "uuid",
    "title": "Lesson 4",
    "description": "Rus tili mashqlari",
    "status": "available",
    "order_number": 4
  },
  "video": {
    "provider": "telegram",
    "stream_ref": "telegram-stream-ref",
    "duration_seconds": 1500
  },
  "materials": [
    {
      "id": "uuid",
      "title": "Lesson PDF",
      "material_type": "pdf",
      "file_url": "https://cdn.example.com/file.pdf"
    }
  ],
  "homework": {
    "id": "uuid",
    "homework_type": "text",
    "instructions": "Daftarga mashqni bajaring",
    "submission": {
      "status": "submitted",
      "score": null,
      "feedback": null
    }
  }
}
```

### Locked Response 200

```json
{
  "lesson": {
    "id": "uuid",
    "title": "Lesson 8",
    "status": "locked",
    "available_at": "2026-04-20T00:00:00Z"
  },
  "video": null,
  "materials": [],
  "homework": null
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Student bu dersga access qilolmaydi |
| `LESSON_NOT_FOUND` | 404 | Lesson bulunamadi |
| `CLASS_NOT_FOUND` | 404 | Class bulunamadi |
| `ACCESS_DENIED` | 403 | Class-level access active degil |
| `ENROLLMENT_NOT_FOUND` | 404 | Student class'a kayitli degil |

## Side Effects

- Yok. Read-only use case.
- Lesson view tracking daha sonra eklenebilir.

## Test Senaryolari

1. Active access student available lesson detail alir.
2. Locked lesson 200 ile locked state doner.
3. Access pending ise `ACCESS_DENIED` doner.
4. Lesson baska course'a aitse `LESSON_NOT_FOUND` veya `FORBIDDEN` doner.
5. Existing submission varsa response'ta submission status doner.
