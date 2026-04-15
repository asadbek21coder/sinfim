# UC: change-initial-password

## Operation

| Alan | Deger |
|------|-------|
| Operation ID | `change-initial-password` |
| Method | `POST` |
| Path | `/api/v1/auth/change-initial-password` |
| Actor | Authenticated |
| Modul | `auth` |

## Amac

Gecici password/kod ile ilk kez giren kullanici kendi kalici password'unu belirler.

## Request

```json
{
  "new_password": "new-secure-password",
  "new_password_confirm": "new-secure-password"
}
```

### Validasyon

| Field | Kural |
|-------|-------|
| `new_password` | required, min 8 karakter |
| `new_password_confirm` | required, `new_password` ile ayni olmali |

## Business Rules

1. Kullanici authenticated olmali.
2. Kullanici aktif olmali.
3. `new_password` ve `new_password_confirm` ayni olmali.
4. Yeni password mevcut gecici password ile ayni olmamali.
5. Yeni password hash'lenerek saklanir.
6. `must_change_password` false yapilir.
7. Guvenlik icin mevcut refresh/session token'lari revoke edilebilir; MVP'de mevcut session kalabilir, ama yeni token donmek daha temizdir.

## Response 200

```json
{
  "access_token": "new_jwt_access_token",
  "refresh_token": "new_jwt_refresh_token",
  "must_change_password": false,
  "user": {
    "id": "uuid",
    "full_name": "Ali Valiyev",
    "phone_number": "+998901234567",
    "is_active": true
  }
}
```

## Errors

| Code | HTTP | Sebep |
|------|------|-------|
| `UNAUTHORIZED` | 401 | Token yok/gecersiz |
| `USER_INACTIVE` | 422 | Kullanici pasif |
| `PASSWORD_CONFIRM_MISMATCH` | 422 | Password confirm eslesmiyor |
| `PASSWORD_TOO_WEAK` | 422 | Password minimum kurallari saglamiyor |
| `PASSWORD_SAME_AS_TEMPORARY` | 422 | Yeni password gecici password ile ayni |

## Side Effects

- `auth.users.password_hash` guncellenir.
- `auth.users.must_change_password` false olur.
- Opsiyonel: eski session/refresh token'lar revoke edilir.

## Test Senaryolari

1. Gecerli kullanici yeni password belirler ve `must_change_password` false olur.
2. Confirm eslesmezse `PASSWORD_CONFIRM_MISMATCH` doner.
3. Kisa/zayif password ile `PASSWORD_TOO_WEAK` doner.
4. Pasif kullanici ile `USER_INACTIVE` doner.
5. Token yoksa `UNAUTHORIZED` doner.
6. Yeni password ile tekrar login basarili olur.
