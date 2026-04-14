# UC: login

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `login` |
| Method | `POST` |
| Path | `/api/v1/auth/login` |
| Actor | Public |
| Modul | `auth` |

## Amac

Kullanici telefon raqami + password ile sisteme girer, session olusturulur ve kullanicinin organization membership bilgileri doner.

MVP'de SMS OTP ve Telegram login yoktur.

## Request

```json
{
  "phone_number": "+998901234567",
  "password": "string"
}
```

### Validasyon

| Field | Kural |
|-------|-------|
| `phone_number` | required, normalize edilebilir telefon formatinda olmali |
| `password` | required, min 6 karakter |

## Business Rules

1. Telefon raqami normalize edilerek aranir.
2. Kullanici yoksa `INCORRECT_CREDENTIALS` doner. Guvenlik icin `USER_NOT_FOUND` donulmez.
3. Kullanici pasifse `USER_INACTIVE` doner.
4. Password hash dogrulanir.
5. Password yanlissa `INCORRECT_CREDENTIALS` doner.
6. Kullanici icin aktif membership'ler yuklenir.
7. Access token ve refresh token olusturulur.
8. `must_change_password = true` ise login yine basarili olur, fakat response'ta bu flag doner. Frontend yeni password ekranina yonlendirir.
9. Membership listesinde `PLATFORM_ADMIN` role varsa `organization_id` null olabilir.

## Response 200

```json
{
  "access_token": "jwt_access_token",
  "refresh_token": "jwt_refresh_token",
  "must_change_password": false,
  "user": {
    "id": "uuid",
    "full_name": "Ali Valiyev",
    "phone_number": "+998901234567",
    "is_active": true
  },
  "memberships": [
    {
      "organization_id": "uuid",
      "organization_name": "My School",
      "organization_slug": "my-school",
      "role": "OWNER"
    }
  ]
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `INCORRECT_CREDENTIALS` | 422 | Telefon yoki password notogri |
| `USER_INACTIVE` | 422 | Kullanici pasif |
| `VALIDATION_ERROR` | 422 | Request formati hatali |

## Side Effects

- `auth.sessions` tablosuna yeni session kaydi eklenir.
- `auth.users.last_active_at` guncellenebilir.

## Test Senaryolari

1. Dogru telefon + password ile login basarili olur.
2. Yanlis password ile `INCORRECT_CREDENTIALS` doner.
3. Olmayan telefon ile `INCORRECT_CREDENTIALS` doner.
4. Pasif kullanici ile `USER_INACTIVE` doner.
5. `must_change_password = true` kullanici login olur ve response'ta `must_change_password: true` doner.
6. Platform admin membership'i `organization_id = null` ile donebilir.
7. Owner membership'i organization bilgisiyle doner.
