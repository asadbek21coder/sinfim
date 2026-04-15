# UI Design Review

Bu dokuman repo kokundeki `docs/design/sinfim-design` klasorundeki Stitch AI ciktisini Sinfim.uz planina baglamak icin hazirlandi.

## Genel Karar

`docs/design/sinfim-design` ciktisi MVP icin kullanilabilir bir gorsel yon veriyor. Ikinci prompt sonrasi ekranlar artik mobil app degil, desktop-first web SaaS dashboard mantigina daha yakin.

Bu ciktinin dogru kullanimi:

- HTML dosyalari birebir production koda kopyalanmayacak.
- `screen.png` ve `code.html` dosyalari gorsel/component referans olarak kullanilacak.
- Vue 3 + TypeScript implementasyonunda layout, renk, spacing, tablo ve form davranislari proje component sistemine uygun yeniden kurulacak.
- Ana referans tasarim sistemi: `docs/design/sinfim-design/scholar_slate/DESIGN.md`.

Marka icin canonical kaynak: `docs/product/brand-constants.md`.

## Tasarim Sistemi

Stitch ciktisindaki ana tasarim dili:

- Yonetim ekranlari icin desktop-first 1440px canvas.
- Premium akademik/SaaS hissi.
- Koyu lacivert sabit sidebar.
- Ferah bosluk, buyuk basliklar, soft yuzey katmanlari.
- Dashboard ve data ekranlarinda guvenilir operasyon hissi.

Temel tokenlar:

| Alan | Deger |
|------|-------|
| Primary | `#041632` deep navy |
| Secondary | `#2c694e` forest green |
| Surface | `#f7f9fb` cool white |
| Headline font | Manrope |
| Body/UI font | Inter |
| Sidebar | 260px, primary background |
| Primary button | Secondary/green gradient, 8px radius |
| Metric card | White surface + left green accent bar |

Not: Tasarim sisteminde "no divider lines" deniyor, fakat uretilen bazi HTML ekranlarinda subtle border/divider kullanilmis. Implementation'da bu kural esnek uygulanabilir: tablo okunurlugu icin cok hafif border gerekirse kullanilabilir, ama genel ekran hissi cizgiyle bogulmamali.

## Canonical Screen Secimi

Ayni ekranin `_1` ve `_2` versiyonu varsa `_2` genelde Sinfim.uz markasina daha yakin. Implementation referansi olarak asagidaki secim kullanilacak:

| UX Ekrani | Tasarim Kaynagi | Karar |
|-----------|-----------------|-------|
| Platform Landing | `platform_landing_page_2` | `_2` kullan, marka/copy temizle |
| Entry Point | `entry_point_page_2` | `_2` kullan, tekrar eden footer/copy temizle |
| School Request | `school_request_page_2` | `_2` kullan, placeholder okul adlarini temizle |
| Login / First Password | `login_first_password_flow_2` | `_2` kullan, SMS copy'sini kaldir |
| Superadmin Organization Create | `superadmin_organization_create_2` | `_2` kullan, title/placeholder temizle |
| Organization Settings | `organization_settings` | Kullanilabilir |
| Owner Dashboard | `owner_dashboard` | Layout iyi, marka ve okul icerigi temizlenmeli |
| Course Detail | `course_detail` | Layout iyi, course icerigi Sinfim.uz'a uyarlanacak |
| Class/Group Detail | `class_group_detail` | Layout iyi, brand ve kisi/icerik placeholder'lari temizlenmeli |
| Lesson Editor | `lesson_editor` | MVP kararlarina yakin; Telegram ref alani iyi |
| Homework Review Inbox | `homework_review_inbox` | Split-pane model iyi; ders/ornek icerikler temizlenmeli |
| Student Dashboard | `student_dashboard` | Web portal hissi iyi; oyunlastirma/icerik abartisi sadelestirilebilir |
| Student Lesson Detail | `student_lesson_detail` | Kullanilabilir; homework tipleriyle kontrol edilmeli |
| Public School Page | `public_school_page` | Kullanilabilir; generic okul copy'si temizlenmeli |
| Public Course Page | `public_course_page` | Kullanilabilir; Architectural placeholder'lari temizlenmeli |

## Visual Screenshot Notes

Gorsel screenshot incelemesinden frontend icin ozel notlar:

- `owner_dashboard/screen.png`: genel dashboard layout'u iyi; metric cards, quick actions ve course/class progress tablosu referans alinabilir. Fakat marka, owner isimleri, okul adi ve domain tamamen placeholder.
- `course_detail/screen.png`: course detail icin tab yapisi iyi: Overview, Lessons, Classes, Settings. Lessons tablosu MVP icin uygun; sadece icerik ve status adlari Sinfim.uz domainine uyarlanacak.
- `homework_review_inbox/screen.png`: mentor review icin en guclu ekranlardan biri. Sol submission listesi + orta student response + sag grading paneli korunabilir. Copy ve ornek dersler degisecek.
- `lesson_editor/screen.png`: Telegram stream source link, material upload, homework setup ve quiz bolumlerini ayni ekranda topluyor; MVP ders editor icin iyi referans. Ekran dar export edilmis, implementation'da 1440px web canvas'a daha ferah yayilmali.
- `student_lesson_detail/screen.png`: video player + material panel + homework submit paneli dogru model. Student ekrani da web app gibi duruyor, mobil app degil. Sidebar student icin fazla admin hissi verebilir; implementation'da student navigation sade tutulabilir.
- `platform_landing_page_2/screen.png`: marka `Sinfim.uz` olarak gelmis ama export 560px genislikte ve landing copy'sinde hala Ingilizce/akademik SaaS tonu var. Frontend'de bu ekran genis desktop landing olarak yeniden kurulacak, birebir dar layout alinmayacak.
- `public_school_page/screen.png` ve `public_course_page/screen.png`: public sayfa hiyerarsisi fikir olarak iyi, fakat screenshot export'lari dar ve uzun. Gercek web implementasyonunda 1200-1440px desktop layout ve responsive breakpoint'ler yeniden ele alinmali.
- `_1` versiyonlu ekranlar genelde eski marka icerir; sadece layout fikri icin bakilabilir. `_2` versiyonlar canonical'a daha yakindir.

Gorsel boyut siniflandirmasi:

| Tip | Ekranlar | Not |
|-----|----------|-----|
| Desktop app'e yakin | `owner_dashboard`, `class_group_detail`, `homework_review_inbox`, `student_dashboard`, `student_lesson_detail` | Ana uygulama layout referansi olarak kullanilabilir |
| Dikey/uzun app ekranlari | `course_detail`, `lesson_editor`, `organization_settings` | Icerik iyi, implementation'da genis canvas'a uyarlanacak |
| Dar public export | `platform_landing_page_2`, `public_school_page`, `public_course_page` | Birebir genislik alinmayacak; desktop-first landing olarak yeniden kurulacak |

## Temizlenmesi Gerekenler

Stitch ciktisi final Sinfim.uz ciktisi degil. Implementation oncesi su placeholder ve celiskiler temizlenecek:

- `Architectural Academic`
- `Scholarly Canvas`
- `Sterling Institute of Design`
- `Arthur Sterling`, `Marcus Sterling`
- `Advanced Architectural Design`
- `Physics`, `Quantum`, `Grade 10-A`, `Grade 11-A`
- `platform.uz`
- `sterling-edu.lms.com`
- `Architectural Academic LMS`
- SMS ile gecici kod metinleri

Auth kopyasi icin dogru MVP metni:

- Telefon raqami + password.
- Ilk giriste admin/owner/mentor tarafindan verilen gecici password veya access code.
- SMS OTP yok.
- Telegram login yok.

## Uygulama Notlari

Layout:

- App ekranlarinda ortak shell kullan: sidebar + topbar + content area.
- Public sayfalarda marketing shell kullan: header + hero + sectionlar + footer.
- Student ekranlari daha sade olabilir, ama mobil app gibi degil browser tabanli web portal hissinde kalmali.

Component adaylari:

- `AppShell`
- `PublicShell`
- `SidebarNav`
- `Topbar`
- `MetricCard`
- `ActionPanel`
- `DataTable`
- `StatusBadge`
- `AccessBadge`
- `ProgressBar`
- `SplitReviewLayout`
- `LessonEditorForm`
- `LeadForm`
- `FirstPasswordNotice`

Content dili:

- MVP'de urun dili sonradan kararlastirilabilir, fakat tasarim implementasyonunda Sinfim.uz ve Ozbekistan egitim pazari hissi korunmali.
- Placeholder dersler icin mimari/physics/design ornekleri yerine daha genel ornekler kullanilabilir: `Ingliz tili A1`, `Rus tili A1`, `Matematika tayyorlov`, `Frontend asoslari`.

## Sonuc

Tasarim yonu kabul edilebilir. En buyuk risk gorsel layout degil, ciktinin icinde kalan yanlis marka ve auth metinleri. Ilk implementation adiminda tasarim kodu birebir alinmadan once Sinfim.uz domain kararlarina gore temizlenmis component map cikartilmali.
