# Brand Constants

Bu dokuman frontend ve backend blueprint kopyalandiktan sonra marka/domain kararlarinda karisiklik olmamasi icin sabit referans olarak tutulur.

## Canonical Brand

| Alan | Karar |
|------|-------|
| Platform adi | `Sinfim.uz` |
| Domain | `sinfim.uz` |
| Kisa urun adi | `Sinfim` sadece metin icinde dogal kullanim icin kabul edilebilir |
| Public organization URL | `sinfim.uz/{school-slug}` |
| Public course URL | `sinfim.uz/{school-slug}/courses/{course-slug}` |
| Public school apply URL | `sinfim.uz/{school-slug}/apply` |
| Subdomain modeli | MVP disi: `{school-slug}.sinfim.uz` |

## Frontend Copy Rules

- Navbar/logo alaninda varsayilan marka: `Sinfim.uz`.
- Browser title ornegi: `Sinfim.uz - Online School Platform`.
- App shell sidebar logo metni: `Sinfim.uz`.
- Footer ornegi: `(c) 2026 Sinfim.uz. All rights reserved.`.
- Organization preview alanlarinda slug prefix'i: `sinfim.uz/`.
- Demo veya placeholder domain olarak `platform.uz`, `lms.com`, `sterling-edu.lms.com` kullanilmayacak.

## Forbidden Placeholder Copy

Stitch AI ciktisindan gelen asagidaki metinler production UI'a tasinmayacak:

- `Architectural Academic`
- `Architectural Academic LMS`
- `Scholarly Canvas`
- `Sterling Institute of Design`
- `Arthur Sterling`
- `Marcus Sterling`
- `platform.uz`
- `sterling-edu.lms.com`
- SMS ile temporary code yonlendirmesi

## Auth Copy

MVP auth kopyasi su kararla uyumlu olmali:

- Telefon raqami + password.
- Ilk giriste admin/owner/mentor tarafindan verilen gecici password veya access code.
- SMS OTP yok.
- Telegram login yok.

Dogru ilk-giris yardim metni ornegi:

`First time here? Use the temporary password or access code given by your school admin. You will set your own password after login.`

Yanlis metin ornegi:

`Check your registered SMS for the temporary access code.`
