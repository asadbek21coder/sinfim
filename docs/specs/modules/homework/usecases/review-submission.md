# UC: review-submission

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `review-submission` |
| Method | `POST` |
| Path | `/api/v1/homework/review-submission` |
| Actor | Mentor / Owner / Teacher |
| Modul | `homework` |

## Amac

Mentor yoki yetkili owner/teacher student odev teslimini kontrol eder, feedback ve durum girer.

## Request

```json
{
  "organization_id": "uuid",
  "submission_id": "uuid",
  "status": "approved",
  "score": 85,
  "feedback": "Yaxshi, lekin 3-mashqda xatolar bor."
}
```

## Validasyon

| Field | Kural |
|-------|-------|
| `organization_id` | required |
| `submission_id` | required |
| `status` | `approved`, `needs_revision`, `rejected`, `score_only` |
| `score` | optional, 0-100 |
| `feedback` | optional ama `needs_revision` icin required |

## Business Rules

1. Actor `OWNER`, `TEACHER` veya ilgili class'a atanmis `MENTOR` olmali.
2. Submission organization icinde mevcut olmali.
3. Mentor sadece atanmis oldugu class submission'larini review edebilir.
4. Status `needs_revision` ise feedback required.
5. Review sonrasi `reviewed_by`, `reviewed_at`, `status`, `score`, `feedback` set edilir.
6. Student ders detayinda feedback'i gorur.

## Response 200

```json
{
  "id": "uuid",
  "status": "approved",
  "score": 85,
  "feedback": "Yaxshi, lekin 3-mashqda xatolar bor.",
  "reviewed_by": "uuid",
  "reviewed_at": "2026-04-14T10:00:00Z"
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Actor bu submission'i review edemez |
| `SUBMISSION_NOT_FOUND` | 404 | Submission bulunamadi |
| `FEEDBACK_REQUIRED` | 422 | Needs revision icin feedback gerekli |
| `VALIDATION_ERROR` | 422 | Request formati hatali |

## Side Effects

- `homework.homework_submissions` review alanlari guncellenir.
- Student tarafinda feedback/sonuc gorunur.

## Test Senaryolari

1. Atanmis mentor submission review eder.
2. Atanmamis mentor `FORBIDDEN` alir.
3. Owner herhangi organization submission'ini review eder.
4. `needs_revision` feedback olmadan `FEEDBACK_REQUIRED` doner.
5. Review sonrasi student lesson detail'da reviewed state gorur.
