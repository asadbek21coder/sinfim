# UC: create-lesson

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `create-lesson` |
| Method | `POST` |
| Path | `/api/v1/catalog/create-lesson` |
| Actor | Owner / Teacher |
| Modul | `catalog` |

## Amac

Course icinde yeni lesson olusturmak. Lesson video, material ve homework/test ile zenginlestirilebilir; bu UC ilk lesson kaydini olusturur.

## Request

```json
{
  "organization_id": "uuid",
  "course_id": "uuid",
  "title": "Lesson 1: Alphabet",
  "description": "Rus alifbosi bilan tanishuv",
  "order_number": 1,
  "estimated_minutes": 25,
  "publish_day": 1
}
```

## Validasyon

| Field | Kural |
|-------|-------|
| `organization_id` | required |
| `course_id` | required |
| `title` | required, min 2 |
| `order_number` | required, > 0, course icinde unique |
| `estimated_minutes` | optional, > 0 |
| `publish_day` | optional, >= 1 |

## Business Rules

1. Actor ilgili organization'da `OWNER` veya `TEACHER` role'e sahip olmali.
2. Course organization icinde mevcut olmali.
3. `order_number` course icinde unique olmali.
4. Lesson varsayilan status `draft` olur.
5. Video/material/homework bu UC ile zorunlu degildir; ders sonradan `update-lesson` ile zenginlestirilir.

## Response 200

```json
{
  "id": "uuid",
  "organization_id": "uuid",
  "course_id": "uuid",
  "title": "Lesson 1: Alphabet",
  "description": "Rus alifbosi bilan tanishuv",
  "order_number": 1,
  "estimated_minutes": 25,
  "publish_day": 1,
  "status": "draft"
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `FORBIDDEN` | 403 | Organization'da yetki yok |
| `COURSE_NOT_FOUND` | 404 | Course bulunamadi |
| `LESSON_ORDER_ALREADY_EXISTS` | 409 | Bu sira numarasi kullaniliyor |
| `VALIDATION_ERROR` | 422 | Request formati hatali |

## Side Effects

- `catalog.lessons` kaydi olusur.

## Test Senaryolari

1. Owner course icinde lesson olusturur.
2. Teacher course icinde lesson olusturur.
3. Course baska organization'a aitse `COURSE_NOT_FOUND` veya `FORBIDDEN` doner.
4. Duplicate order number ile `LESSON_ORDER_ALREADY_EXISTS` doner.
5. Video/material olmadan lesson olusabilir.
