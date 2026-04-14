# UX Dokumani

Bu dosya Sinfim projesinin ekran ve akislarini netlestirmek icin calisma taslagidir.

Kural: Bu dosya simdilik final UI karari degil. Once ekranlar ve akislar yaziyla netlesecek, sonra teknik dokumana gecilecek.

## 1. Genel Urun Bilgisi

### Urun Adi

Sinfim.uz

### Tek Cumle

Sinfim.uz, okul, egitim merkezi, ogretmen veya kendi brendiyle kurs satan egitim ekipleri icin kurs, sinif/grup, video ders, materyal, odev, test, mentor kontrolu, lead, access ve ogrenci ilerlemesini tek yerde yoneten online okul platformudur.

### Tasarim Tonu

Karar: guvenilir, sistematik, modern ve egitim odakli.

Notlar:

- Ciddi ve guvenilir hissettirmeli.
- Okul/egitim merkezi kendi platformu gibi kullanabilmeli.
- Ogrenci tarafinda mobil web kullanimi rahat olmali.
- Public landing sayfasi sade ama guzel olmali; platformdan faydalanan okul logolari yavas hareket eden bir logo bandi/slider ile gosterilebilir.
- "Nasil kullanilir?" bolumu klavuz gibi olmali: okul kurulur, kurs acilir, sinif/grup olusturulur, materyal/odev eklenir, ogrenci access alir.

### Renk ve Tipografi Karari

Karar: Stitch AI tarafindan uretilen `docs/sinfim-design/scholar_slate/DESIGN.md` tasarim sistemi, Sinfim.uz MVP icin ana gorsel referans olarak alinacak.

Tasarim yonu:

- Desktop-first web SaaS: ana operasyon ekranlari 1440px genislikte browser uygulamasi gibi dusunulecek.
- Responsive destek olacak, fakat dashboard ve operasyon ekranlarinda ilk oncelik bilgisayar kullanimi.
- Genel his: ciddi, guvenilir, modern akademik ve operasyonel.
- Ana renk: deep navy `#041632`.
- Aksiyon rengi: forest green `#2c694e`.
- Ana yuzey rengi: cool white `#f7f9fb`.
- Tipografi: Manrope headline/display icin, Inter UI/body icin.
- Dashboard ana navigasyon: 260px koyu lacivert sidebar.
- Component stili: Vue blueprint ile uyumlu Tailwind token sistemi; Stitch HTML'leri birebir kopyalanmayacak, gorsel referans ve component hedefi olarak kullanilacak.
- Landing tarafinda temiz bir egitim/SaaS hissi; dashboard tarafinda daha operasyonel, hizli ve tablo/kart odakli UI.

Not: `docs/sinfim-design` ciktisinda hala placeholder marka ve icerikler var. Implementation'a gecmeden once `Architectural Academic`, `Sterling`, `platform.uz`, `lms.com`, SMS copy'si ve tasarim/physics gibi ornek icerikler Sinfim.uz kararlariyla temizlenecek. Marka/domain sabitleri icin `docs/projects/lms/docs/brand-constants.md` esas alinacak.

## 1.1 Public Landing Karari

Public ana sayfa olacak.

Amaci:

- Platformun ne oldugunu anlatmak.
- Kullanan okul/brand logolarini gostermek.
- Nasil kullanildigini adim adim anlatmak.
- Dogru giris noktasini sunmak: mevcut okuluna gir, demo okulunu gez, platforma katilma talebi birak.

Not: MVP'de herkes kendi kendine okul olusturmayacak. Okul/organization olusturma superadmin kontrolunde olacak. Public taraftaki "platforma katil" aksiyonu direkt okul yaratmaz; talep/basvuru olusturur.

## 2. Kullanici Akislari

### Akis 0: Public Landing ve Entry Point

Baslangic: Ziyaretci platform ana sayfasina girer.

1. Platform hakkinda kisa ve net aciklama gorur.
2. Platformu kullanan okul/brand logolarini hareketli logo bandinda gorur.
3. "Nasil ishlaydi?" klavuz bolumunu gorur.
4. Entry point butonuna tiklar.
5. Secenekleri gorur: mevcut okuluma gir, demo okulunu gez, platforma katilma talebi birak, ogrenci olarak okul ara/gir.

Bitis: Ziyaretci dogru akisa yonlendirilir.

### Akis 1: Superadmin Okul/Brand Olusturur

Baslangic: Platform superadmin developer/admin dashboard'a girer.

1. Yeni organization olusturur.
2. Okul/brand adini girer.
3. Description ve logo ekler.
4. Slug secer: `sinfim.uz/{school-slug}`.
5. Okul owner kisinin telefon raqamini girer.
6. Owner icin gecici password/kod olusturur.
7. Owner role ve organization access verilir.
8. Owner login bilgisiyle kendi okul dashboard'una girer.

Bitis: Okul workspace'i hazirdir, owner kendi okulunu yonetebilir.

### Akis 2: Owner Kurs ve Sinif/Grup Olusturur

Baslangic: Owner dashboard'dadir.

1. Kurs olusturur.
2. Kurs icin temel bilgileri girer.
3. Kurs icinde sinif/grup olusturur.
4. Sinif/grup icin baslangic tarihi, ders yayin ritmi ve access ayarlarini belirler.
5. Mentor atama adimina gecer.

Bitis: Ilk sinif/grup hazirdir.

### Akis 3: Mentor Eklenir ve Sinifa Atanir

Baslangic: Owner sinif/grup detayindadir.

1. Mentorun isim ve telefon raqamini girer.
2. Mentor davet edilir veya gecici giris bilgisi olusturulur.
3. Mentor bir veya birden fazla sinif/gruba atanir.

Bitis: Mentor kendi sorumlu oldugu sinif/gruplari gorebilir.

### Akis 4: Ogrenci Manuel Kayit Edilir

Baslangic: Owner/mentor sinif/grup detayindadir.

1. Ogrencinin isim ve telefon raqamini girer.
2. Ogrenci sinif/gruba eklenir.
3. Odeme/access durumu manuel tasdik edilir.
4. Ogrenciye gecici kod/password veya davet linki verilir.
5. Ogrenci ilk giriste kendi passwordini belirler.

Bitis: Ogrenci access verilen ders/materyalleri gorebilir.

### Akis 5: Public Okul/Kurs Sayfasindan Lead Gelir

Baslangic: Potansiyel ogrenci `sinfim.uz/{school-slug}` veya kurs sayfasina girer.

1. Okul/kurs hakkinda temel bilgi gorur.
2. Isim, telefon raqami ve ilgilendigi kursu birakir.
3. Sistem lead kaydi olusturur.
4. Owner/mentor lead listesinden kisiyi gorur.
5. Konusma/odeme/onay sonrasi lead student'a cevrilir.
6. Student uygun sinif/gruba eklenir.

Bitis: Lead gercek ogrenciye donusebilir.

### Akis 6: Ders, Materyal ve Odev Yayinlanir

Baslangic: Teacher/owner kurs detayindadir.

1. Ders olusturur.
2. Telegram stream video referansi ekler.
3. PDF/materyal ekler.
4. Odev tipi secer: yazili, dosya/fotograf, test/quiz, oral/audio.
5. Dersin yayin sirasini veya yayin tarihini belirler.

Bitis: Ders ilgili sinif/grup icin yayinlanmaya hazirdir.

### Akis 7: Student Ders Izler ve Odev Teslim Eder

Baslangic: Student kendi dashboard'una girer.

1. Bugunku dersleri gorur.
2. Video dersi izler.
3. PDF/materyali acar.
4. Odev tipine gore cevap verir.
5. Teslim durumunu gorur.

Bitis: Odev mentor kontrolune duser veya test sonucu otomatik hesaplanir.

### Akis 8: Mentor Odev Kontrol Eder

Baslangic: Mentor dashboard'dadir.

1. Sorumlu oldugu sinif/gruplardaki bekleyen odevleri gorur.
2. Odev detayini acar.
3. Yazili/fotograf/audio cevabi inceler.
4. Durum, puan ve feedback girer.
5. Student sonucu gorur.

Bitis: Odev kontrol edilmis olur.

### Akis 9: Owner Progress ve Access Takip Eder

Baslangic: Owner dashboard'dadir.

1. Kurs/sinif/grup listesini gorur.
2. Progress, odev teslimleri ve access durumlarini inceler.
3. Gerekirse ogrencinin access durumunu degistirir.
4. Mentor performansini ve bekleyen odevleri gorur.

Bitis: Okul operasyonu tek yerden takip edilir.

## 3. Edge Case'ler

| Durum | Ekran / Davranis |
|-------|------------------|
| Yanlis sifre | Login ekraninda hata mesaji |
| Gecici kod/password gecersiz | Ilk giris ekraninda yeniden kod isteme veya adminle iletisime gecme mesaji |
| Slug daha once alinmis | Organization setup ekraninda alternatif slug onerisi |
| Access yok | Student ders sayfasinda "Bu derse erisiminiz yok" mesaji |
| Ders henuz yayinlanmadi | Student ders listesinde kilitli/yayin tarihi bilgisi |
| Odev gec teslim edildi | Odev detayinda gec teslim etiketi |
| Mentor yetkisiz sinifa girmeye calisti | Yetki yok mesaji |
| Telegram video referansi calismiyor | Video alaninda hata + owner/teacher icin kontrol uyarisi |

## 4. Ekran Listesi

| # | Ekran Adi | Amac | Route / Path | Durum |
|---|-----------|------|--------------|-------|
| 1 | Platform Landing | Platform tanitimi, logo bandi, rehber ve entry point | `/` | Detay taslagi var |
| 2 | Entry Point | Mevcut okul, demo okul, platform talebi, ogrenci girisi secimi | `/enter` | Detay taslagi var |
| 3 | Login | Owner/teacher/mentor/student girisi | `/auth/login` | Detay taslagi var |
| 4 | Platforma Katilma Talebi | Okul acmak isteyen kisi basvuru birakir | `/apply-school` | Detay taslagi var |
| 5 | Superadmin Organization Create | Okul/brend olusturma, owner atama | `/admin/organizations/new` | Detay taslagi var |
| 6 | Organization Setup/Edit | Okul/brend adi, description, logo ve slug yonetimi | `/app/settings/organization` | Detay taslagi var |
| 7 | Owner Dashboard | Okul operasyon ozeti | `/app/dashboard` | Detay taslagi var |
| 8 | Kurs Listesi | Organization kurslarini yonetme | `/app/courses` | Detay bekliyor |
| 9 | Kurs Detay | Dersler, siniflar, materyaller | `/app/courses/:courseId` | Detay taslagi var |
| 10 | Sinif/Grup Detay | Ogrenci, mentor, access ve progress | `/app/classes/:classId` | Detay taslagi var |
| 11 | Mentor Listesi | Mentor ekleme ve atama | `/app/mentors` | Detay bekliyor |
| 12 | Student Listesi | Ogrenci kayit ve access takibi | `/app/students` | Detay bekliyor |
| 13 | Lead Listesi | Potansiyel ogrencileri takip etme | `/app/leads` | Detay bekliyor |
| 14 | Ders Editor | Video, materyal, odev yayinlama | `/app/lessons/:lessonId/edit` | Detay taslagi var |
| 15 | Odev Kontrol | Mentorun odev kontrol etmesi | `/app/homework/review` | Detay taslagi var |
| 16 | Student Dashboard | Ogrencinin bugunku ders ve odevleri | `/learn/dashboard` | Detay taslagi var |
| 17 | Student Ders Detay | Video, materyal ve odev teslimi | `/learn/lessons/:lessonId` | Detay taslagi var |
| 18 | Public Okul Sayfasi | Okul/kurs tanitimi ve lead formu | `/{schoolSlug}` | Detay bekliyor |
| 19 | Public Kurs Sayfasi | Kurs tanitimi ve basvuru | `/{schoolSlug}/courses/:courseSlug` | Detay bekliyor |
| 20 | Demo Okul | Fake okul deneyimi | `/demo` veya `/demo-school` | Detay taslagi var |

## 5. Navigation Yapisi

### Public

```text
/
/enter
/apply-school
/demo
/{schoolSlug}
/{schoolSlug}/courses/:courseSlug
/{schoolSlug}/apply
```

### Auth

```text
/auth/login
```

Not: Public self-register ile okul acma MVP'de yok. Organization superadmin tarafindan olusturulur.

### Superadmin

```text
/admin/organizations
/admin/organizations/new
/admin/organizations/:organizationId
```

### Owner / Teacher / Mentor App

```text
/app/dashboard
/app/settings/organization
/app/courses
/app/courses/:courseId
/app/classes/:classId
/app/mentors
/app/students
/app/leads
/app/lessons/:lessonId/edit
/app/homework/review
```

### Student

```text
/learn/dashboard
/learn/lessons/:lessonId
```

## 6. Auth Guard

Auth gerektirmeyen route'lar:

- `/`
- `/enter`
- `/apply-school`
- `/demo`
- `/{schoolSlug}`
- `/{schoolSlug}/courses/:courseSlug`
- `/{schoolSlug}/apply`
- `/auth/login`

Auth gerektiren route'lar:

- `/admin/*`
- `/app/*`
- `/learn/*`

Rol bazli route notlari:

- Platform superadmin: organization olusturma ve owner atama ekranlari.
- Organization owner: tum organization ekranlari.
- Teacher: kurs/ders/materyal ekranlari, yetkisine gore sinif ekranlari.
- Mentor: atanmis sinif/gruplar, odev kontrol ekranlari.
- Student: sadece kendi `learn` ekranlari.

## 7. Component Katalogu

| Component | Nerede Kullaniliyor | Notlar |
|-----------|---------------------|--------|
| `PageHeader` | App ekranlari | Baslik, aciklama, primary action |
| `DataTable` | Kurs, sinif, ogrenci, lead, odev listeleri | Loading/empty/error gerekli |
| `EmptyState` | Bos kurs/sinif/lead/odev listeleri | Ilk aksiyona yonlendirmeli |
| `AccessBadge` | Student, ders, materyal listeleri | Access var/yok/beklemede |
| `ProgressBar` | Dashboard ve sinif detay | Ogrenci/grup progress |
| `LessonCard` | Student dashboard | Ders durumu ve yayin tarihi |
| `HomeworkSubmissionCard` | Mentor odev kontrol | Yazili/dosya/audio/test durumlari |
| `LeadForm` | Public okul/kurs sayfasi | Isim, telefon, kurs |
| `InvitePanel` | Student/mentor ekleme | Gecici kod/link bilgisi |
| `LogoMarquee` | Platform landing | Okul/brand logolari yavas hareket eder |
| `EntryChoiceCard` | Entry point | Mevcut okul, demo, talep, ogrenci secenekleri |

## 8. Ekran Detay Taslaklari

## 8.0 Ana UX Model Karari

Course ve Class/Group ayrimi net:

- Course = reusable content package. Kursun dersleri, videolari, PDF/materiallari, odevleri, testleri ve public kurs bilgisi burada yonetilir.
- Class/Group = live cohort. Gercek ogrenciler, mentorlar, access/payment durumu, ders acilma takvimi, progress ve odev teslimleri burada yonetilir.

Ornek:

- Course: `Russian A1`
- Class/Group: `Russian A1 - May 2026`, `Russian A1 - VIP`, `Russian A1 - Weekend`

Bu ayrim onemli: ayni kurs icerigi farkli gruplara farkli baslangic tarihi, mentor, ogrenci listesi ve access durumu ile uygulanabilir.

### Ekran: Platform Landing

Amac: Ziyaretciye platformu anlatmak ve dogru giris noktasina yonlendirmek.

Layout:

- Full public page.
- Hero bolumu.
- Logo marquee / hareketli okul logolari.
- "Nasil ishlaydi?" klavuz bolumu.
- Fayda bolumleri.
- Entry point CTA.

Icerik:

1. Kisa baslik: online okul operasyonunu tek yerde yonet.
2. Alt aciklama: kurs, sinif, mentor, odev, test, access ve lead takibi.
3. CTA: "Platforma kirish" veya "Boshlash".
4. Kullanan okul/brand logolari.
5. Nasil kullanilir adimlari: okul olusturulur, kurs acilir, sinif/grup kurulur, materyal/odev yayinlanir, ogrenci progress takip edilir.
6. Demo okul linki.

Aksiyonlar:

- CTA: Platforma kirish -> `/enter`
- Demo okul -> `/demo`
- Platforma katilma talebi -> `/apply-school`

### Ekran: Entry Point

Amac: Kullanici tipini ve niyetini ayirmak.

Secenekler:

- Mevcut okuluma girmek istiyorum -> okul slug veya telefon login akisi
- Ogrenciyim -> okul/kurs sayfasina git veya login
- Demo okulunu gezmek istiyorum -> `/demo`
- Platformu okuluma kurmak istiyorum -> `/apply-school`

### Ekran: Platforma Katilma Talebi

Amac: Okul acmak isteyen kisiden basvuru almak; direkt organization olusturmamak.

Alanlar:

- Isim
- Telefon raqami
- Okul/brand adi
- Kurs kategorisi
- Kac ogrenci/grup var?
- Kisa not

Sonuc:

- Talep superadmin tarafinda gorulur.
- Superadmin uygun gorurse organization olusturur ve owner access verir.

### Ekran: Superadmin Organization Create

Amac: Superadmin'in yeni okul/organization olusturmasi ve owner atamasi.

Alanlar:

- Okul/brand adi
- Description
- Logo
- Slug: `sinfim.uz/{school-slug}`
- Owner ismi
- Owner telefon raqami
- Owner gecici password/kod

Aksiyonlar:

- Organization olustur
- Owner access ver
- Demo/fake okul olarak isaretle

### Ekran: Owner Dashboard

Amac: Okul sahibinin tum okul operasyonunu tek yerde gormesi.

Icerik:

1. Okul adi, logo ve public sayfa linki.
2. Ozet kartlari: aktif kurs, aktif sinif/grup, ogrenci, lead, bekleyen odev, access bekleyen ogrenci.
3. Hizli aksiyonlar: kurs olustur, sinif olustur, mentor ekle, ogrenci ekle, leadleri gor.
4. Son aktiviteler: yeni lead, yeni odev teslimi, yeni ogrenci, access degisikligi.
5. Bekleyen isler: kontrol edilmemis odevler, access tasdigi bekleyenler, yayin tarihi yaklasan dersler.
6. Kurs/sinif progress listesi.

### Ekran: Kurs Detay

Amac: Owner/teacher'in tekrar kullanilabilir kurs icerigini yonetmesi.

Ana model:

- Kurs burada bir icerik paketi olarak dusunulur.
- Ogrenci operasyonu burada degil, sinif/grup ekraninda yonetilir.

MVP tablari:

- Overview
- Lessons
- Classes
- Settings

Overview icerigi:

1. Kurs adi, kategori, seviye ve yayin durumu.
2. Public kurs sayfasi linki: `sinfim.uz/{schoolSlug}/courses/{courseSlug}`.
3. Kisa aciklama ve hedef kitle.
4. Toplam ders sayisi.
5. Bu kursa bagli aktif sinif/grup sayisi.
6. Hizli aksiyonlar: ders ekle, sinif/grup olustur, public sayfayi gor.

Lessons tab icerigi:

1. Ders listesi.
2. Her ders icin sira numarasi, baslik, video durumu, materyal durumu, odev/test durumu.
3. Ders ekleme aksiyonu.
4. Ders sirasi degistirme.
5. Ders editor'e gecis.

Classes tab icerigi:

1. Bu kursa bagli sinif/grup listesi.
2. Grup adi, mentor, ogrenci sayisi, baslangic tarihi, aktif/pasif durumu.
3. Yeni sinif/grup olusturma aksiyonu.

Settings tab icerigi:

1. Kurs adi, description, kategori, seviye.
2. Public kurs sayfasi aktif/pasif.
3. Kurs slug.
4. Kurs status: draft / active / archived.

Aksiyonlar:

- Yeni ders ekle -> ders editor acilir.
- Yeni sinif/grup olustur -> class/group creation akisi.
- Public sayfayi gor -> public kurs sayfasi.
- Kursu arsivle -> aktif sinif/grup varsa uyar.

State'ler:

| State | Davranis |
|-------|----------|
| Loading | Kurs bilgisi ve ders listesi skeleton |
| Empty lessons | "Bu kursta henuz ders yok" + ders ekle CTA |
| Empty classes | "Bu kurs henuz hicbir sinifta kullanilmiyor" + sinif olustur CTA |
| Error | Kurs yuklenemedi + retry |

### Ekran: Sinif/Grup Detay

Amac: Owner/mentorun gercek grup operasyonunu yonetmesi.

Ana model:

- Sinif/grup, bir kursun gercek ogrencilerle calisan cohort halidir.
- Mentorlar, ogrenciler, access, progress ve odev teslimleri burada yonetilir.

MVP tablari:

- Overview
- Students
- Homework
- Access

Overview icerigi:

1. Grup adi: ornek `Russian A1 - May 2026`.
2. Bagli kurs.
3. Baslangic tarihi.
4. Ders acilma ritmi: har kuni / kun ora / haftada 3 marta.
5. Atanmis mentorlar.
6. Ogrenci sayisi.
7. Progress ozeti.
8. Bekleyen odev sayisi.
9. Access tasdigi bekleyen ogrenci sayisi.

Students tab icerigi:

1. Ogrenci listesi.
2. Ogrenci adi, telefon raqami, access durumu, progress, son aktivite.
3. Manuel ogrenci ekleme.
4. Lead'den ogrenciye cevirme.
5. Gecici kod/password uretme veya davet bilgisini gosterme.

Homework tab icerigi:

1. Bekleyen odev teslimleri.
2. Teslim tipi: yazili, fotograf/dosya, test/quiz, oral/audio.
3. Mentor atamasi.
4. Durum: bekliyor / kontrol edildi / tekrar isteniyor.
5. Odev kontrol ekranina gecis.

Access tab icerigi:

1. Ogrenci bazli access/payment durumu.
2. Manuel access ac/kapat.
3. Access sebebi veya notu.
4. Hangi ders/materyaller acik, hangileri kilitli.
5. Bulk action: secili ogrencilere access ver.

Aksiyonlar:

- Ogrenci ekle.
- Mentor ata.
- Lead'i student'a cevir.
- Access ver/kapat.
- Odevleri kontrol et.
- Ders takvimini duzenle.

State'ler:

| State | Davranis |
|-------|----------|
| Loading | Grup bilgisi, ogrenciler ve odevler skeleton |
| Empty students | "Bu grupta henuz ogrenci yok" + ogrenci ekle CTA |
| Empty homework | "Bekleyen odev yok" |
| Empty access pending | "Access tasdigi bekleyen ogrenci yok" |
| Error | Grup bilgisi yuklenemedi + retry |

### Ekran: Ders Editor

Amac: Owner/teacher'in kurs icindeki tek bir dersi olusturmasi veya duzenlemesi.

Best-practice MVP karari:

- Ders editor tek sayfa icinde bolumlere ayrilir.
- Zorunlu alanlar en basta, opsiyonel materyal/odev/test alanlari sonra gelir.
- Video upload MVP'de yok; Telegram stream referansi girilir.
- Odev tek ders icinde birden fazla tipte olabilir, ama MVP'de her ders icin bir ana odev seti yeterli.

Bolumler:

1. Basic Info
2. Video
3. Materials
4. Homework
5. Quiz/Test
6. Publish Rules

Basic Info:

- Ders basligi
- Kisa aciklama
- Ders sirasi
- Tahmini sure
- Status: draft / ready / archived

Video:

- Telegram channel/message referansi
- Video basligi
- Video sure bilgisi
- Video test/preview aksiyonu
- Hata durumunda owner/teacher icin "video referansini kontrol et" mesaji

Materials:

- PDF/material ekleme
- Material adi
- Material tipi: PDF / image / link / other
- Dosya boyutu ve indirme/goruntuleme bilgisi

Homework:

- Odev var/yok toggle
- Odev tipi: yazili cevap / dosya-fotograf / oral-audio
- Odev talimati
- Son teslim kurali: yayinlandiktan sonra X gun veya tarih secimi
- Mentor feedback gerekli mi?

Quiz/Test:

- Test var/yok toggle
- Soru listesi
- Soru tipi: single choice / multiple choice
- Dogru cevap
- Otomatik puanlama
- Gecme puani opsiyonel

Publish Rules:

- Ders siraya gore acilsin.
- Ders class/group baslangic tarihine gore gun N'de acilsin.
- Manuel yayinla opsiyonu.

Aksiyonlar:

- Taslak kaydet.
- Ready olarak isaretle.
- Preview et.
- Sil/arsivle.

State'ler:

| State | Davranis |
|-------|----------|
| Loading | Form skeleton |
| Empty material | Material yoksa bos state, ders yine kaydedilebilir |
| Empty homework | Odev yoksa "Bu derste odev yok" |
| Validation error | Eksik zorunlu alanlar inline gosterilir |
| Video error | Telegram referansi gecersiz/preview alinamadi mesaji |

### Ekran: Student Dashboard

Amac: Ogrencinin bugun ne yapmasi gerektigini hizli ve net gormesi.

Best-practice MVP karari:

- Student dashboard mobile-first olmali.
- Ogrenciye tablo degil, "bugunku isler" akisi gosterilmeli.
- Dersler kilitli/acik/tamamlandi durumlariyla sade kartlarda gorunmeli.

Icerik:

1. Selamlama ve aktif okul/kurs bilgisi.
2. Bugunku ders karti.
3. Bekleyen odevler.
4. Son feedback/sonuc.
5. Kurs progress bar.
6. Siradaki acilacak ders bilgisi.

Lesson card bilgileri:

- Ders adi
- Durum: available / locked / completed
- Video var mi?
- Material var mi?
- Odev var mi?
- Son teslim tarihi

Aksiyonlar:

- Derse devam et.
- Odev teslim et.
- Feedback'i gor.
- Kurs materyallerini gor.

State'ler:

| State | Davranis |
|-------|----------|
| Loading | Kart skeleton |
| No active course | "Henuz aktif kursunuz yok" + okul adminiyle iletisime gec mesajı |
| No lesson today | "Bugun yeni ders yok" + onceki derslere git |
| Access locked | Odeme/access tasdigi bekleniyor mesaji |
| Error | Dashboard yuklenemedi + retry |

### Ekran: Student Ders Detay

Amac: Ogrencinin tek derste video izlemesi, materyali acmasi ve odev teslim etmesi.

Best-practice MVP karari:

- Mobilde once video, sonra materyal, sonra odev akisi.
- Desktop'ta video ana alan, sag tarafta ders bilgisi ve odev paneli olabilir.
- Odev teslim formu dersle ayni sayfada olmali; ayri sayfaya atlamamali.

Icerik:

1. Ders basligi ve progress durumu.
2. Video player alanı.
3. PDF/material listesi.
4. Odev talimati.
5. Odev teslim formu.
6. Teslim durumu ve mentor feedback.

Odev teslim tipleri:

- Yazili cevap: text area.
- Dosya/fotograf: upload alanı.
- Oral/audio: audio file upload veya ileride browser recorder.
- Quiz/test: soru listesi ve cevap secimi.

Aksiyonlar:

- Dersi tamamlandi olarak isaretle.
- Material ac/indir.
- Odev teslim et.
- Teslimi guncelle, eger mentor henuz kontrol etmediyse.
- Feedback'i gor.

State'ler:

| State | Davranis |
|-------|----------|
| Loading | Video ve ders icerigi skeleton |
| Locked | Ders henuz acilmadi veya access yok |
| No material | Material bolumu gizlenebilir veya bos state |
| No homework | "Bu derste odev yok" |
| Submitted | Teslim edildi, mentor kontrolu bekleniyor |
| Reviewed | Mentor feedback ve puan gosterilir |
| Needs revision | Mentor tekrar istemisse uyarı ve yeniden teslim aksiyonu |

### Ekran: Odev Kontrol

Amac: Mentorun sorumlu oldugu sinif/gruplardaki odev teslimlerini hizli ve duzenli kontrol etmesi.

Best-practice MVP karari:

- Mentor icin inbox mantigi kullanilsin.
- Sol tarafta filtreli teslim listesi, sagda secili teslim detayi iyi calisir.
- Mobilde liste ve detay ayri ekran gibi davranabilir.

Filtreler:

- Sinif/grup
- Kurs
- Ders
- Odev tipi
- Durum: bekliyor / kontrol edildi / tekrar isteniyor
- Gec teslimler

Liste bilgileri:

- Ogrenci adi
- Ders/odev adi
- Teslim tipi
- Teslim zamani
- Durum

Detay icerigi:

1. Ogrenci bilgisi.
2. Ders ve odev talimati.
3. Ogrencinin cevabi: text / dosya-fotograf / audio / quiz sonucu.
4. Mentor feedback formu.
5. Puan veya durum.
6. Hazir feedback sablonlari, MVP'de opsiyonel.

Feedback durumlari:

- Approved
- Needs revision
- Rejected
- Score only

Aksiyonlar:

- Feedback gonder.
- Puan ver.
- Tekrar iste.
- Sonraki teslimata gec.

State'ler:

| State | Davranis |
|-------|----------|
| Loading | Liste ve detay skeleton |
| Empty inbox | "Kontrol bekleyen odev yok" |
| No permission | Bu sinif/grup mentora atanmis degil |
| File preview error | Dosya/aciklama indirilebilir ama preview yok |
| Audio unavailable | Audio acilamadi + download fallback |

### Ekran: Organization Setup/Edit

Amac: Owner'in kendi okul/brand bilgisini duzenlemesi.

Best-practice MVP karari:

- Organization ilk kez superadmin tarafindan olusturulur.
- Owner bu ekranda okul bilgisini duzenler, ama gercek ownership/organization creation yetkisi superadmin'de kalir.
- Slug degisikligi MVP'de izinli olabilir ama uyarili olmalı; public link degisir.

Bolumler:

1. Brand Info
2. Public Page
3. Contact
4. Danger Zone

Brand Info:

- Okul/brand adi
- Description
- Logo upload
- Kategori veya egitim tipi

Public Page:

- Slug: `sinfim.uz/{school-slug}`
- Public sayfa aktif/pasif
- Public sayfa preview linki
- Logo ve description public sayfada nasil gorunecek preview

Contact:

- Telefon raqami
- Telegram username/link
- Support contact
- Location, opsiyonel

Danger Zone:

- Organization arsivle, MVP'de sadece superadmin.
- Slug degistir, uyarili.

Aksiyonlar:

- Bilgileri kaydet.
- Public sayfayi gor.
- Logo degistir.
- Slug degistir.

State'ler:

| State | Davranis |
|-------|----------|
| Loading | Form skeleton |
| Slug taken | Slug alinmis, alternatif oner |
| Save success | Basarili kayit toast |
| Validation error | Eksik/hatali alanlari inline goster |

### Ekran: Public Okul/Kurs Lead Formu

Amac: Potansiyel ogrencinin login olmadan okul veya kursa ilgi birakmasi.

Best-practice MVP karari:

- Lead form cok kisa olmali.
- Login gerektirmez.
- Form okul sayfasinda ve kurs sayfasinda ayni component olarak kullanilir.
- Kurs sayfasindan gelirse ilgilendigi kurs otomatik dolu gelir.

Alanlar:

- Isim
- Telefon raqami
- Qaysi kursga qiziqyapsiz? (kurs sayfasinda prefilled)
- Kisa not, opsiyonel

Gizli metadata:

- Organization ID
- Course ID, eger kurs sayfasindan geldiyse
- Source page
- Created at

Aksiyonlar:

- Basvuru gonder.
- Telefon raqami format hatasini goster.
- Basarili gonderim sonrasi tesekkur mesaji.

Basarili mesaj:

- "Arizangiz qabul qilindi. Maktab administratori siz bilan bog'lanadi."

State'ler:

| State | Davranis |
|-------|----------|
| Loading | Submit buton loading |
| Success | Tesekkur mesaji ve tekrar basvuru yapma opsiyonu |
| Validation error | Telefon/isim eksik |
| Duplicate lead | "Bu raqam bilan ariza bor" mesaji, ama kullaniciyi korkutmadan |
| Error | Ariza yuborilmadi + retry |

### Ekran: Demo Okul

Amac: Ziyaretcinin platformu gercek okul yaratmadan deneyimlemesi.

Best-practice MVP karari:

- Demo okul read-only veya kontrollu sandbox olmali.
- Public kullanici demo icinde veri bozabilmemeli.
- Demo, owner dashboard ve student experience hissini gostermeli.

Demo entry:

- `/demo`
- "Owner dashboard demo" ve "Student view demo" secenekleri.

Owner demo icerigi:

- Fake okul adi ve logo.
- Fake kurs listesi.
- Fake sinif/grup progress.
- Fake bekleyen odevler.
- Fake lead listesi.
- Aksiyon butonlari calisiyormus gibi modal acabilir, ama gercek kayit yapmaz.

Student demo icerigi:

- Bugunku ders karti.
- Video preview placeholder.
- Material listesi.
- Odev teslim formu demo modunda.
- Mentor feedback ornegi.

Uyari:

- Demo mod banner: "Bu demo. Kiritilgan ma'lumotlar saqlanmaydi."

State'ler:

| State | Davranis |
|-------|----------|
| Loading | Demo data skeleton |
| Demo unavailable | Demo yuklenemedi + retry |
| Action disabled | "Demo rejimida bu amal saqlanmaydi" |

### Ekran: Login ve Ilk Giris Akisi

Amac: Owner/teacher/mentor/student kullanicilarinin dis servis bagimliligi olmadan sisteme girmesi.

Best-practice MVP karari:

- Tek login ekrani kullanilir: telefon raqami + password.
- SMS OTP yok.
- Telegram login yok.
- Ilk giris icin gecici password/kod desteklenir.

Login alanlari:

- Telefon raqami
- Password

Ilk giris akisi:

1. Superadmin veya owner/mentor kullanici icin gecici password/kod olusturur.
2. Bu bilgi kullaniciya manuel iletilir.
3. Kullanici telefon raqami + gecici password/kod ile girer.
4. Sistem yeni password belirleme ekranina yonlendirir.
5. Password belirlendikten sonra role gore dashboard'a gider.

Role gore redirect:

- Platform superadmin -> `/admin/organizations`
- Organization owner -> `/app/dashboard`
- Teacher -> `/app/courses`
- Mentor -> `/app/homework/review`
- Student -> `/learn/dashboard`

Aksiyonlar:

- Login ol.
- Ilk giriste password belirle.
- Password unutuldu: MVP'de "admin/mentor bilan bog'laning" mesaji.

State'ler:

| State | Davranis |
|-------|----------|
| Wrong credentials | Telefon veya password hatali |
| Temporary password | Yeni password belirleme ekranina yonlendir |
| No organization access | Kullaniciya hic organization access verilmemis |
| Inactive user | Admin bilan bog'laning mesaji |
| Loading | Submit loading |

## 9. UX Cikis Kriteri

- [x] Public landing ve entry point akisi yazildi.
- [x] Superadmin kontrollu okul olusturma akisi yazildi.
- [x] Course ve Class/Group ayrimi yazildi.
- [x] Kurs Detay ve Sinif/Grup Detay ekranlari taslaklandi.
- [x] Ders Editor, Student Dashboard, Student Ders Detay ve Odev Kontrol ekranlari taslaklandi.
- [x] Organization Setup/Edit, Lead Form, Demo Okul ve Login/ilk giris akisi taslaklandi.
- [ ] Kurs Listesi, Mentor Listesi, Student Listesi ve Lead Listesi icin detay gerekirse sonraki turda genisletilecek.

Sonraki adim: Teknik dokumana gec.
