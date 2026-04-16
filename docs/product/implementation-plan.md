# Sinfim.uz Implementation Plan

Bu dokuman Sinfim.uz MVP'sini frontend ve backend birlikte ilerletecek vertical-slice plana boler.

Ana kural: Her step sonunda kullanici browser uzerinden anlamli bir akis test edebilmeli. Sadece backend endpoint'i veya sadece statik UI bitmis sayilmaz.

## Calisma Prensipleri

- Her step kendi icinde backend, frontend, seed/demo data, manuel test ve dokuman guncellemesi icerir.
- Ilk asamada odeme entegrasyonu, SMS OTP, Telegram login, video transcoding ve mobil native app yoktur.
- MVP auth: telefon raqami + password/gecici password.
- Organization modeli multi-tenant kalacak; her okul kendi workspace'i gibi kullanacak.
- URL modeli: `sinfim.uz/{school-slug}`.
- UI desktop-first web uygulamasi olarak kurulacak; responsive destek olacak ama ana operasyon bilgisayar icindir.
- Stitch ciktisi birebir kopyalanmayacak; `docs/design/sinfim-design` gorsel referans olarak kullanilacak.
- Platform adi her yerde `Sinfim.uz`, domain `sinfim.uz`.

## Step 0 - Local Foundation ve Proje Iskeleti

Durum: frontend route/layout iskeleti kuruldu; local Docker/health standardi `docs/product/local-development.md` icinde sabitlendi.

Amac: Developer ortaminda tek komutla calisan temel uygulama iskeleti.

Backend:

- Docker local config'in netlestirilmesi.
- Backend health endpoint kontrolu.
- Mevcut blueprint modullerinin Sinfim.uz icin hangi kisimlarinin korunacagini netlestirme.
- Migration ve seed komutlari icin standart yol belirleme.

Frontend:

- Vue app route yapisini Sinfim.uz ekran listesine gore duzenleme.
- Public layout, app layout, student layout ayrimini kurma.
- Design token ve temel component kararlarini baslatma.
- API client base URL ve auth token storage kararini sabitleme.

Manual test:

- `make local-up` ile servisler kalkar.
- `http://localhost:5173` frontend acilir.
- `http://localhost:9876/health` backend health doner.
- Frontend API proxy backend'e ulasir.

Bitis kriteri:

- Local stack tekrar kurulabilir durumdadir.
- Bos ama dogru route/layout iskeleti vardir.

## Step 1 - Public Landing, Entry Point ve School Request

Durum: public okul arizasi backend'e kaydediliyor; platform admin preview ekranindan arizalar listelenip status degistirilebiliyor. Admin endpoint'leri Step 2 auth tamamlanana kadar gecici olarak acik tutuldu.

Amac: Ziyaretci Sinfim.uz'u anlar ve okul acmak icin talep birakabilir.

Backend:

- `organization.school_requests` migration'i. Tamamlandi.
- School request create/list/update status use case'leri. Tamamlandi.
- Public endpoint: okul acma talebi olusturma. Tamamlandi: `POST /api/v1/organization/create-school-request`.
- Superadmin/admin endpoint: talepleri listeleme ve status degistirme. Preview tamamlandi: `GET /api/v1/organization/list-school-requests`, `POST /api/v1/organization/update-school-request-status`.

Frontend:

- Desktop-first landing page.
- Entry point page: mevcut okuluma gir, demo okul, platforma katilma talebi.
- School request formu. Tamamlandi: `/apply-school`.
- Basarili talep, validation hata ve loading state'leri. Tamamlandi.
- Platform admin preview ekrani. Tamamlandi: `/admin/school-requests`.

Manual test:

- Ana sayfadan platform aciklamasi gorulur.
- Entry point'ten school request formuna gidilir.
- Zorunlu alanlar bosken hata gorulur.
- Gecerli form gonderilince basari mesaji gorulur.
- Backend'de talep kaydi olusur.

Bitis kriteri:

- Public visitor ilk temas akisini tamamlayabilir.
- Talep kaydi backend'de kalici tutulur.
- Admin preview ekranindan talep status'u degistirilebilir.

## Step 2 - Auth, First Login ve Session

Durum: telefon/password login calisiyor; `get-me`, refresh/logout/change-password endpoint'leri frontend ile uyumlu hale getirildi. Protected app/admin route guard eklendi. Local seed platform admin: `+998900000001` / `admin12345`.

Amac: Owner, mentor, teacher ve student telefon/password ile sisteme girebilir.

Backend:

- Blueprint auth modulu Sinfim.uz kararlarina uyarlanir. Tamamlandi: `phone_number`, `full_name`, `must_change_password` alanlari eklendi.
- Telefon raqami + password login endpoint'i. Tamamlandi: `POST /api/v1/auth/admin-login`.
- Refresh/logout/change password endpoint'leri. Tamamlandi.
- `must_change_password` first-login flow. Backend ve frontend ekran tamamlandi.
- Role yapisi: `PLATFORM_ADMIN`, `OWNER`, `TEACHER`, `MENTOR`, `STUDENT`. Tamamlandi.
- Seed: bir platform admin. Tamamlandi. Demo owner Step 3 organization create ile eklenecek.
- Step 1 admin school request endpoint'leri token/permission ile korundu.

Frontend:

- Login ekranlari. Tamamlandi: `/auth/login`.
- First password change ekranlari. Tamamlandi: `/auth/change-password`.
- Auth guard. Tamamlandi.
- Role bazli ilk yonlendirme. Baslangic seviyesi tamamlandi.
- Session kaybi/401 davranisi. Baslangic seviyesi tamamlandi.

Manual test:

- Seed admin login olur: `+998900000001` / `admin12345`.
- Yanlis password hata verir.
- `must_change_password=true` olan kullanici password degistirme ekranina duser.
- Password degistikten sonra dashboard'a yonlenir.
- Token yokken protected route'a girilmez.

Bitis kriteri:

- Temel guvenli giris akisi calisir.
- SMS/OTP/Telegram login kullanilmaz.
- School request admin ekranlari login olmadan acilmaz.

## Step 3 - Organization Create ve Owner Workspace

Durum: platform admin yeni organization olusturabiliyor, owner user yoksa gecici password ile olusturuluyor, `OWNER` role ve `auth.user_memberships` kaydi veriliyor. Owner temporary password ile login oldugunda `mustChangePassword=true` donuyor. Owner kendi workspace'ini gorebiliyor ve Organization Settings ekranindan temel okul bilgilerini guncelleyebiliyor.

Amac: Platform admin yeni okul/brand olusturur, owner kendi okul workspace'ine girer.

Backend:

- `organization.organizations` migration'i. Tamamlandi.
- `auth.user_memberships` migration'i. Tamamlandi.
- Organization create use case'i. Tamamlandi: `POST /api/v1/organization/create-organization`.
- My workspaces use case'i. Tamamlandi: `GET /api/v1/organization/list-my-workspaces`.
- Organization update use case'i. Tamamlandi: `POST /api/v1/organization/update-organization`.
- Slug unique kontrolu. Tamamlandi: duplicate slug `409` doner.
- Owner kullanicisi olusturma veya mevcut kullaniciyi owner yapma. Tamamlandi.
- Membership kaydi. Tamamlandi. Owner update yetkisi membership ile kontrol ediliyor; route bazli tenant enforcement Step 4-6 icinde genisletilecek.

Frontend:

- Superadmin organization create ekrani. Tamamlandi: `/admin/organizations/new`.
- Organization settings ekranlari. Tamamlandi: `/app/settings/organization`.
- Owner dashboard shell. Mevcut app dashboard kullaniliyor.
- Sidebar/topbar icinde aktif okul bilgisi. Ilk workspace adini gosteriyor; multi-workspace selector sonraki tenant-context adiminda eklenecek.

Manual test:

- Superadmin yeni okul olusturur.
- Slug conflict olursa hata gorur.
- Owner telefon/gecici password ile login olur.
- Owner `mustChangePassword=true` ile password degistirme ekranina yonlendirilir.
- Owner kendi workspace listesini gorur.
- Organization settings'te isim, description, logo URL, category, contact, Telegram URL ve public status gorulur/duzenlenir.

Bitis kriteri:

- Multi-tenant workspace mantigi ilk kez gercek calisir.
- Organization, owner user, owner role ve membership birlikte olusur.

## Step 4 - Public School Page ve Lead Capture

Amac: Her okulun public sayfasi olur; potansiyel ogrenci lead birakir.

Durum: tamamlandi. Public sayfa sadece `public_status=public` veya `is_demo=true` organization'lar icin acilir. Lead'ler `lead.leads` tablosunda organization'a bagli saklanir.

Backend:

- Public organization get by slug endpoint'i. Tamamlandi: `GET /api/v1/organization/get-public-school-page?slug={schoolSlug}`.
- `lead.leads` migration'i. Tamamlandi.
- Lead create/list/update status use case'leri. Tamamlandi.
- Lead kaydini organization'a baglama. Tamamlandi.

Frontend:

- `/{schoolSlug}` public school page. Tamamlandi.
- Public lead formu. Tamamlandi.
- Owner lead listesi. Tamamlandi: `/app/leads`.
- Lead status: new, contacted, converted, archived. Tamamlandi.

Manual test:

- `/{schoolSlug}` public okul sayfasi acilir.
- Var olmayan slug 404/empty state verir.
- Lead formundan isim/telefon gonderilir.
- Owner dashboard'da lead gorulur.
- Owner lead status degistirir.

Bitis kriteri:

- Reklamdan gelen ogrenci adayini platform icine almak mumkundur.

## Step 5 - Course Management

Amac: Owner/teacher okul icinde kurs olusturur ve public/private durumunu yonetir.

Durum: tamamlandi. Course `catalog` modulu altinda ayri bounded context olarak eklendi. Public course page sadece okul public/demo ve course `public_status=public` ise acilir.

Backend:

- `catalog.courses` migration'i. Tamamlandi.
- Course create/update/list/detail use case'leri. Tamamlandi.
- Public course page use case'i. Tamamlandi: `GET /api/v1/catalog/get-public-course-page?school_slug={schoolSlug}&course_slug={courseSlug}`.
- Organization scope ve role permission kontrolleri. Tamamlandi: create/update Owner veya Teacher, list/detail organization member veya platform admin.
- Course slug unique kontrolu organization icinde. Tamamlandi.

Frontend:

- Course list ekranı. Tamamlandi: `/app/courses`.
- Course create/edit formu. Tamamlandi.
- Course detail ana ekranı. Tamamlandi: `/app/courses/:courseId`.
- Public course page baslangici. Tamamlandi: `/{schoolSlug}/courses/{courseSlug}`.

Manual test:

- Owner kurs olusturur.
- Kurs listede gorulur.
- Kurs draft/public olarak degistirilir.
- Public course page sadece public kurs icin acilir.
- Baska okulun kursuna app icinden erisim engellenir.

Bitis kriteri:

- Kurs Sinfim.uz icinde ana content container olarak calisir.

## Step 6 - Class/Group Management ve Access

Amac: Kursa bagli sinif/grup acilir, ogrenci access/payment durumu manuel yonetilir.

Durum: tamamlandi. Class/group modeli `classroom` modulu altinda eklendi. Owner/teacher kursa bagli sinif acabilir, sinifa ogrenci ekleyebilir ve ogrencinin manual payment/access durumunu yonetebilir. Mentor assignment API hazir; MVP UI bu stepte class detail icinde ogrenci/access operasyonuna odaklandi.

Backend:

- `classroom.classes`, `classroom.enrollments`, `classroom.class_mentors`, `classroom.access_grants` migration'lari. Tamamlandi.
- Class create/list/detail use case'leri. Tamamlandi.
- Student add use case'i. Tamamlandi: yeni student user gerekirse telefon + temporary password ile olusur.
- Mentor assign use case'i. Tamamlandi: API seviyesinde hazir.
- Access/payment status update: pending, active, paused, blocked. Tamamlandi.
- Organization scope ve role permission kontrolleri. Tamamlandi: owner/teacher operate edebilir, mentor atandigi sinifi gorebilir.

Frontend:

- Class/group list ve create formu. Tamamlandi: `/app/classes`.
- Course detail icinden class create/list. Tamamlandi: `/app/courses/:courseId`.
- Class/group detail ekranı. Tamamlandi: `/app/classes/:classId`.
- Student ekleme formu. Tamamlandi.
- Access/payment durum kontrolleri. Tamamlandi.
- Mentor atama UI. API hazir; detayli UI sonraki admin/mentor polish adimina birakildi.

Manual test:

- Kurs detayindan yeni sinif acilir.
- Sinifa mentor API ile atanir.
- Ogrenci telefon/isim ile manuel eklenir.
- Ogrencinin access durumu active/blocked yapilir.
- Mentor sadece atandigi sinifi gorur.

Bitis kriteri:

- Telegram'daki grup operasyonunun platformdaki temel karsiligi olusur.

## Step 7 - Lesson, Video Reference ve Materials

Amac: Kurs icine dersler, Telegram stream referansi ve PDF/material eklenir.

Durum: tamamlandi. `go test ./...`, `npm run build` ve local API smoke gecti. MVP'de material upload yerine URL/file metadata saklanir; MinIO/filevault upload sonraki dosya polish adimina birakildi.

Backend:

- `catalog.lessons`, `catalog.lesson_videos`, `catalog.lesson_materials` migration'lari. Tamamlandi.
- Lesson create/update/list/detail use case'leri. Tamamlandi.
- Telegram stream ref saklama. Tamamlandi: provider, stream ref, channel/message id, embed URL metadata.
- File/material metadata saklama. Tamamlandi: URL veya future filevault id metadata.
- Publish day/order kurallari. Tamamlandi: order number ve publish day alanlari eklendi; student availability hesaplamasi Step 8'de uygulanacak.

Frontend:

- Course detail icinde lesson list/create. Tamamlandi: `/app/courses/:courseId`.
- Lesson editor ekranı. Tamamlandi: `/app/lessons/:lessonId/edit`.
- Ders listesi, siralama ve status UI. Tamamlandi.
- Video reference formu. Tamamlandi.
- Material/PDF link metadata ekleme UI. Tamamlandi.
- Publish day veya yayin sirasi alanlari. Tamamlandi.

Manual test:

- Course detail'de yeni ders olusturulur.
- Telegram stream referansi girilir.
- PDF/material eklenir.
- Ders draft/published yapilir.
- Class start date'e gore student tarafinda kilitli/acik gorunmesi Step 8'de learning endpoint ile uygulanir.

Bitis kriteri:

- Onceden hazirlanmis video ve material dagitimi platform icinde modellenir.

## Step 8 - Student Learning Experience

Amac: Ogrenci kendi dashboard'unda aktif dersleri gorur, ders detayini acip izler.

Durum: tamamlandi. `go test ./...`, `npm run build` ve local API smoke gecti. Student dashboard, lesson detail, publish/access kilidi ve lesson completion MVP akisi calisir durumda.

Backend:

- Learning dashboard read model endpoint'i. Tamamlandi: `GET /api/v1/learning/get-student-dashboard`.
- Lesson detail student endpoint'i. Tamamlandi: `GET /api/v1/learning/get-lesson-detail`.
- Access check: class membership + active access + publish rule. Tamamlandi.
- Lesson completion endpoint'i. Tamamlandi: `POST /api/v1/learning/mark-lesson-completed`.

Frontend:

- Student dashboard. Tamamlandi: `/learn/dashboard`.
- Student lesson detail. Tamamlandi: `/learn/lessons/{lessonId}`.
- Video player container. Tamamlandi: Telegram/embed URL metadata ile MVP container.
- Material list/download/open action. Tamamlandi.
- Lesson completed action ve progress state. Tamamlandi.

Manual test:

- Student login olur.
- Sadece kendi aktif sinifindaki dersleri gorur.
- Kilitli ders acilmaz.
- Acik ders detayinda video/material gorulur.
- Dersi tamamlandi isaretleyince progress degisir.

Bitis kriteri:

- Ogrenci Telegram yerine Sinfim.uz'da ders takip etmeye baslar.

## Step 9 - Homework Definitions ve Submission

Amac: Teacher/owner ders icin odev tanimlar; ogrenci yazili/dosya/audio/quiz teslim eder.

Durum: tamamlandi. `go test ./...`, `npm run build` ve local API smoke gecti. Lesson editor icinde homework definition; student lesson detail icinde text/file/audio URL ve quiz submission MVP akisi calisir durumda. Mentor review Step 10'da kalir.

Backend:

- `homework.homework_definitions`, `homework.submissions`, quiz tablolarinin MVP migration'i. Tamamlandi: `homework.definitions`, `quiz_questions`, `quiz_options`, `submissions`, `quiz_answers`.
- Homework create/update/list use case'leri. Tamamlandi: `POST /api/v1/homework/save-definition`, `GET /api/v1/homework/get-lesson-homework`.
- Submission create use case'i. Tamamlandi: `POST /api/v1/homework/submit-homework`.
- Submission type: text, file/photo, audio, quiz. Tamamlandi: file/photo/audio MVP'de URL olarak saklanir.
- Quiz automatic scoring MVP. Tamamlandi: correct option points toplanir, quiz submission `reviewed` ve `auto_scored=true` olur.

Frontend:

- Lesson editor icinde homework block. Tamamlandi.
- Student lesson detail icinde homework submission UI. Tamamlandi.
- Text answer, file/photo upload placeholder veya filevault upload. Tamamlandi: URL placeholder ile.
- Audio answer icin MVP kararina gore file upload UI. Tamamlandi: URL placeholder ile.
- Quiz UI. Tamamlandi.

Manual test:

- Teacher ders icin yazili odev olusturur.
- Student cevap gonderir.
- Teacher quiz olusturur.
- Student quiz submit eder ve otomatik score gorur.
- Tekrar teslim kurali dogru calisir.

Bitis kriteri:

- Telegram'daki defter fotografi/cevap gonderme operasyonu platforma tasinir.

## Step 10 - Mentor Homework Review Inbox

Amac: Mentor bekleyen odevleri gorur, feedback ve puan verir.

Durum: tamamlandi. `go test ./...`, `npm run build` ve local API smoke gecti. Review queue, submission detail, feedback/score kaydi ve student result display calisir durumda.

Backend:

- Mentor review list endpoint'i. Tamamlandi: `GET /api/v1/homework/list-review-submissions`.
- Submission detail endpoint'i. Tamamlandi: `GET /api/v1/homework/get-review-submission`.
- Review submit endpoint'i: status, score, feedback. Tamamlandi: `POST /api/v1/homework/review-submission`.
- Mentor permission: sadece atandigi class/group submissions. Tamamlandi; owner/teacher organization icinde, mentor sadece assigned class icinde gorur.
- Student result endpoint'i. Tamamlandi: `GET /api/v1/homework/get-student-homework` feedback/score/review status doner.

Frontend:

- Homework review inbox. Tamamlandi: `/app/homework/review`.
- Submission detail paneli. Tamamlandi.
- Feedback, score, approve/revision state. Tamamlandi.
- Student tarafinda review sonucunu gosterme. Tamamlandi.

Manual test:

- Mentor login olur.
- Sadece kendi siniflarindan pending submission gorur.
- Submission'a feedback ve score girer.
- Student sonucunu gorur.
- Yetkisiz mentor submission'a erisemez.

Bitis kriteri:

- Mentor operasyonu platformdaki ana is akisi haline gelir.

## Step 11 - Owner Operational Dashboard

Amac: Okul sahibi kurs, sinif, access, lead, progress ve bekleyen odev durumunu tek yerden gorur.

Durum: tamamlandi. `go test ./...`, `npm run build` ve local API smoke gecti. Owner dashboard aggregate endpoint'i ve gercek dashboard UI calisir durumda.

Backend:

- Owner dashboard aggregate endpoint'i. Tamamlandi: `GET /api/v1/organization/get-owner-dashboard`.
- Metrics: active courses, classes, students, leads, pending homework, access confirmations. Tamamlandi.
- Recent activity read model'i. Tamamlandi: lead, homework submission ve access activity union read model.

Frontend:

- Owner dashboard'u Stitch referansina gore desktop-first uygulama. Tamamlandi: `/app/dashboard`.
- Quick actions. Tamamlandi.
- Pending homework/access confirmation kartlari. Tamamlandi.
- Recent activity. Tamamlandi.
- Course/class progress summary. Tamamlandi.

Manual test:

- Owner login olur.
- Dashboard metric'leri seed/gercek data ile uyumlu gorur.
- Quick action'lar dogru ekranlara goturur.
- Pending homework ve lead sayilari aksiyon sonrasi degisir.

Bitis kriteri:

- Owner icin platformun degeri ilk bakista gorunur hale gelir.

## Step 12 - Demo School ve Public Trial Experience

Amac: Potansiyel musteri platformu gercek kayit yapmadan deneyebilir.

Durum: tamamlandi. Public `GET /api/v1/organization/get-demo-access` endpoint'i demo organization/course/class/lesson/material/homework ve owner/mentor/student hesaplarini idempotent sekilde hazirlar. Demo data read-only degil; MVP karari olarak resetlenebilir seed mantigi kullanilir ve `/demo` ekraninda uyari gosterilir.

Backend:

- Demo organization seed. Tamamlandi: `demo-school` slug'i public demo school olarak olusur/guncellenir.
- Demo course/class/student/mentor/content seed. Tamamlandi: `Russian A1 Demo`, `Demo Group A`, owner/mentor/student hesaplari, lesson video/material ve text homework seed edilir.
- Demo read-only veya resetlenebilir data kurali. Tamamlandi: strict read-only yerine resetlenebilir/idempotent seed karari alindi; material duplicate temizlenir, class/course/lesson/access kayitlari tekrar kullanilir.
- Demo login veya demo direct preview endpoint karari. Tamamlandi: direct public endpoint credential ve route linklerini doner; login ekranina query ile prefill yapilir.

Frontend:

- Demo entry point. Tamamlandi: `/demo` backend'den demo access alir.
- Demo school public page. Tamamlandi: `/demo-school` mevcut public school sayfasini kullanir.
- Demo owner/student/mentor experience. Tamamlandi: `/demo` owner, student ve mentor hesaplarini gosterir; login ekranina phone/password prefill ve redirect verir.
- Demo mode banner ve read-only uyari state'i. Tamamlandi: `/demo` ekraninda resetlenebilir playground uyarisi var; global read-only enforcement MVP hardening'e birakildi.

Manual test:

- Landing'den demo okul acilir.
- Demo owner dashboard gezilir.
- Demo student lesson detail ve homework gezilir.
- Demo mentor review inbox login'i acilir.
- Demo data tekrar endpoint cagrisi ile resetlenebilir/idempotent hazirlanir.

Bitis kriteri:

- Satis/oneri icin gosterilebilir demo deneyimi vardir.

## Step 13 - MVP Hardening ve Beta Hazirlik

Amac: Ilk gercek okul/egitim merkeziyle denenebilecek kaliteye getirmek.

Backend:

- Permission audit.
- Tenant isolation testleri.
- Error response standardi.
- Basic rate limit veya brute-force koruma karari.
- Migration rollback kontrolu.
- Seed/admin command dokumani.

Frontend:

- Empty/loading/error states tamamlanir.
- Responsive kritik ekran kontrolu.
- Form validation tutarliligi.
- Navigation ve role guard testleri.
- UX copy temizligi: placeholder marka kalmaz.

Manual test:

- Tum ana rollerle smoke test yapilir: platform admin, owner, mentor, student.
- Baska tenant datasina erisim denenir ve engellenir.
- Browser refresh sonrasi session korunur.
- Empty state'ler ve hata state'leri anlaşılırdır.

Bitis kriteri:

- Ilk beta okulu icin MVP hazirdir.

## Onerilen Uygulama Sirasi

1. Step 0 - Local Foundation ve Proje Iskeleti
2. Step 1 - Public Landing, Entry Point ve School Request
3. Step 2 - Auth, First Login ve Session
4. Step 3 - Organization Create ve Owner Workspace
5. Step 4 - Public School Page ve Lead Capture
6. Step 5 - Course Management
7. Step 6 - Class/Group Management ve Access
8. Step 7 - Lesson, Video Reference ve Materials
9. Step 8 - Student Learning Experience
10. Step 9 - Homework Definitions ve Submission
11. Step 10 - Mentor Homework Review Inbox
12. Step 11 - Owner Operational Dashboard
13. Step 12 - Demo School ve Public Trial Experience
14. Step 13 - MVP Hardening ve Beta Hazirlik

Not: Step 11 owner dashboard gorsel olarak erken istenebilir, fakat backend metriklerinin anlamli olmasi icin course/class/homework/lead datasi olustuktan sonra tamamlanmasi daha dogrudur. Ilk dashboard shell Step 3'te acilir, gercek operasyonel dashboard Step 11'de olgunlasir.

## Her Step Icin Standart Checklist

- Backend migration eklendi.
- Backend use case ve permission kontrolleri eklendi.
- API endpoint ve request/response contract'i net.
- Frontend route ve ekran eklendi.
- Loading, empty, validation error ve success state'leri var.
- Manuel test senaryolari denendi.
- Ilgili dokuman veya worklog guncellendi.
- Bir sonraki step icin eksik kararlar not edildi.
