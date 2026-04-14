# UC: get-student-dashboard

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `get-student-dashboard` |
| Method | `GET` |
| Path | `/api/v1/learning/get-student-dashboard` |
| Actor | Student |
| Modul | `learning` |

## Amac

Student login olduktan sonra bugunku derslerini, bekleyen odevlerini, son feedback'ini ve progress durumunu gorur.

## Request Query

```text
organization_id=uuid
class_id=uuid
```

`class_id` optional olabilir; student birden fazla class'a enrolled ise frontend secim yaptirabilir.

## Business Rules

1. Actor `STUDENT` role'e sahip olmali.
2. Student organization membership'e sahip olmali.
3. Student class'a enrolled olmali.
4. Access pending/blocked ise dashboard locked state doner.
5. Lesson availability class `start_date`, `lesson_cadence`, lesson `publish_day/order_number` ile hesaplanir.
6. Bugun acik olan veya tamamlanmamis son dersler doner.
7. Bekleyen homework submissions ve son feedback doner.

## Response 200

```json
{
  "student": {
    "id": "uuid",
    "full_name": "Vali Aliyev"
  },
  "organization": {
    "id": "uuid",
    "name": "My School",
    "slug": "my-school",
    "logo_url": "https://cdn.example.com/logo.png"
  },
  "class": {
    "id": "uuid",
    "name": "Russian A1 - May 2026",
    "access_status": "active"
  },
  "progress": {
    "completed_lessons": 3,
    "total_lessons": 20,
    "percentage": 15
  },
  "today_lessons": [
    {
      "lesson_id": "uuid",
      "title": "Lesson 4",
      "status": "available",
      "has_video": true,
      "has_materials": true,
      "has_homework": true,
      "due_at": "2026-04-16T10:00:00Z"
    }
  ],
  "pending_homework_count": 2,
  "latest_feedback": {
    "lesson_title": "Lesson 3",
    "status": "approved",
    "score": 90,
    "feedback": "Yaxshi"
  }
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Actor student degil veya organization access yok |
| `CLASS_NOT_FOUND` | 404 | Class bulunamadi |
| `ENROLLMENT_NOT_FOUND` | 404 | Student class'a kayitli degil |

## Side Effects

- Yok. Read-only use case.

## Test Senaryolari

1. Active student dashboard data alir.
2. Pending access student locked dashboard state alir.
3. Enrolled olmayan class icin `ENROLLMENT_NOT_FOUND` doner.
4. Lesson schedule'e gore sadece acik dersler available doner.
5. Latest feedback varsa response'ta doner.
