# UC: get-public-school-page

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `get-public-school-page` |
| Method | `GET` |
| Path | `/api/v1/organization/get-public-school-page` |
| Actor | Public |
| Modul | `organization` |

## Amac

Public ziyaretci `sinfim.uz/{school-slug}` sayfasini actiginda okul/brand bilgisi, public kurs listesi ve lead formu icin gerekli minimal data doner.

## Request Query

```text
slug=my-school
```

### Validasyon

| Field | Kural |
|-------|-------|
| `slug` | required, lowercase kebab-case |

## Business Rules

1. Slug ile organization aranir.
2. Organization yoksa `ORGANIZATION_NOT_FOUND` doner.
3. Organization public status `published` degilse `ORGANIZATION_NOT_PUBLIC` doner. Demo school icin bu kural konfigurasyona gore esnetilebilir.
4. Public kurs listesi `catalog` portal/API read model uzerinden alinabilir.
5. Response sadece public alanlari doner; owner/admin/internal alanlar donmez.
6. Lead form submit icin gerekli `organization_id` frontend'e acik UUID olarak donebilir veya backend sadece slug ile lead alabilir. MVP karari: response'ta `organization_id` doner.

## Response 200

```json
{
  "organization": {
    "id": "uuid",
    "name": "My School",
    "slug": "my-school",
    "description": "Online kurs markazi",
    "logo_url": "https://cdn.example.com/logo.png",
    "category": "language",
    "contact_phone": "+998901234567",
    "telegram_url": "https://t.me/myschool",
    "is_demo": false
  },
  "courses": [
    {
      "id": "uuid",
      "title": "Russian A1",
      "slug": "russian-a1",
      "description": "Boshlangich rus tili kursi",
      "category": "language",
      "level": "A1"
    }
  ],
  "lead_form": {
    "enabled": true,
    "required_fields": ["full_name", "phone_number"]
  }
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `ORGANIZATION_NOT_FOUND` | 404 | Slug boyicha okul topilmadi |
| `ORGANIZATION_NOT_PUBLIC` | 404 | Okul public emas veya draft holatda |
| `VALIDATION_ERROR` | 422 | Slug formati hatali |

## Side Effects

- Yok. Read-only use case.

## Test Senaryolari

1. Published organization slug ile public data doner.
2. Draft organization icin `ORGANIZATION_NOT_PUBLIC` doner.
3. Olmayan slug icin `ORGANIZATION_NOT_FOUND` doner.
4. Public response owner membership veya internal admin data dondurmez.
5. Public kurs listesinde sadece public/published kurslar doner.
