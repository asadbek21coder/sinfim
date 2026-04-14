# Auth ve Access Notlari

Bu dosya henuz teknik dokuman degil. Amac: LMS'in okul/organization bazli calisma modelini, rollerini ve access kurallarini netlestirmek.

## Temel Model

Platform multi-tenant calismali.

Her okul, egitim merkezi veya bireysel ogretmen platformda kendi organization/brand profiline sahip olur. MVP'de bu organization'i public self-register ile kendisi acmaz; superadmin olusturur ve owner'a access verir. Kayit olan okul sahibi, sistem icinde kendi okulunu "kendi platformu gibi" kullanabilmeli: kurslarini, sinif/gruplarini, videolarini, materyallerini, mentorlarini, testlerini, odevlerini, ogrencilerini ve access ayarlarini kontrol eder.

Tenant siniri kritik: bir organization icindeki data baska organization tarafindan gorulmemeli veya yonetilmemeli.

## Roller

### Platform Admin

- Tum platformu yonetir.
- Organization'lari gorur ve gerekirse destek amacli mudahele eder.
- MVP'de cok kucuk tutulabilir.

### Organization Owner

- Okul/egitim merkezi/brend sahibi.
- Organization profilini yonetir.
- Kurs, sinif/grup, mentor, ogrenci, materyal, test, odev ve access ayarlarini kontrol eder.
- Odeme disarida alinmis olsa bile platform icinde odeme/access tasdigi yapar.

### Teacher

- Kurs icerigi uretir ve yonetir.
- Ders, video, PDF, odev, test olusturabilir.
- Yetkisi organization owner tarafindan sinirlandirilabilir.

### Mentor

- Bir veya birden fazla sinif/gruptan sorumlu olabilir.
- Kendi sorumlu oldugu sinif/gruplardaki ogrencilerin odevlerini gorur.
- Yazili, fotograf/dosya ve oral/audio odevleri kontrol eder.
- Puan, durum ve feedback girebilir.

### Student

- Sadece kayitli oldugu sinif/grup ve access verilen ders/materyalleri gorur.
- Ders izler, PDF/materyal indirir veya gorur.
- Odev teslim eder: yazili cevap, dosya/fotograf veya oral/audio message.
- Test/quiz cozer ve sonucunu gorur.

### Lead / Potential Student

- Henuz gercek ogrenci degildir.
- Public okul/kurs sayfasindan telefon raqami veya basvuru formu birakir.
- Organization icinde potansiyel musteri olarak gorunur.
- Owner/teacher/mentor tarafindan sonradan gercek ogrenciye cevrilebilir.
- Uygun kurs veya sinif/gruba atanabilir.

## Auth Sorulari

Bu kisim henuz karar bekliyor.

- Ogrenci girisi telefon raqami + password veya davet kodu ile mi olmali? Karar: MVP'de SMS OTP yok, dis servis bagimliligi azaltilecek.
- Email/password gerekli mi, yoksa ogretmen/mentorlar icin telefon + password, ogrenciler icin telefon + password/davet kodu daha mi dogru?
- Telegram login kullanilmali mi? Karar: MVP'de hayir, sonraki faza kalabilir.
- Organization owner kayit olurken okul/brend slug'i alacak mi? Karar: MVP'de evet, path slug kullanilacak. Ornek: `sinfim.uz/my-school`.
- Bir kullanici birden fazla organization'a uye olabilir mi?
- Mentor ayni anda iki farkli organization'da calisabilir mi?
- Owner, teacher, mentor ve student tek `users` tablosunda rol ile mi tutulmali, yoksa student ayri model mi olmali?
- Davet akisi nasil olmali: owner mentor/ogrenci eklerken invite link mi, telefon raqami ile direkt ekleme mi?
- Lead login olmadan basvuru birakabilecek mi, yoksa once kayit olmasi mi gerekecek?
- Lead'den student'a cevirme sirasinda telefon raqami dogrulamasi gerekecek mi?

## Auth Icin Baslangic Varsayimi

Bu henuz final karar degil, ama teknik dokumana gecmeden once tartisilacak baslangic modeli:

- Superadmin platformda okul/brend organization'i olusturur.
- Owner kendi organization'inin slug'ini belirler.
- Owner slug ve okul bilgilerini sonradan duzenleyebilir, fakat ilk organization creation superadmin kontrolundedir.
- MVP'de public okul/kurs sayfalari path slug ile calisir: `sinfim.uz/{school-slug}`.
- Subdomain modeli (`{school-slug}.sinfim.uz`) MVP disinda kalir; ileride eklenebilir.
- Owner teacher ve mentor kullanicilarini davet eder.
- Mentor bir veya birden fazla sinif/gruba atanabilir.
- Ogrenci sinif/gruba owner/teacher/mentor tarafindan manuel eklenebilir.
- Ogrenci veya potansiyel ogrenci public okul/kurs sayfasindan basvuru birakabilir.
- Basvuru birakan kisi once `lead` olarak kaydedilir.
- Owner/teacher/mentor lead'i sonradan `student` haline getirir ve uygun sinif/gruba ekler.
- Owner/teacher/mentor icin telefon raqami + password tercih edilir.
- Student icin telefon raqami + password veya tek kullanimlik davet kodu tercih edilir.
- SMS OTP MVP'de yok; cunku dis servis entegrasyonu ve operasyonal bagimlilik getirir.
- Telegram login daha sonra dusunulebilir; MVP'de auth karmasikligini artirabilir.

## Ogrenci Katilim Akislari

### Akis 1: Manuel Kayit

1. Owner/teacher/mentor ogrencinin adini ve telefon raqamini girer.
2. Ogrenci organization icinde student olarak olusturulur.
3. Ogrenci bir sinif/gruba eklenir.
4. Odeme/access durumu manuel tasdik edilir.
5. Ogrenciye ilk giris icin davet linki veya gecici kod verilir.
6. Ogrenci telefon raqami + kendi belirledigi password ile devam eder.

### Akis 2: Lead Basvurusu

1. Potansiyel ogrenci public okul/kurs sayfasina girer.
2. Isim, telefon raqami ve ilgilendigi kursu birakir.
3. Sistem bu kisiyi organization icinde lead olarak kaydeder.
4. Okul admini veya mentor lead listesinden kisiyi gorur.
5. Konusma/odeme/onay sonrasi lead student'a cevrilir.
6. Student uygun sinif/gruba eklenir ve access verilir.

## Access Kurallari

- MVP'de odeme entegrasyonu yok.
- Odeme disarida alinacak.
- Platformda owner veya yetkili kisi ogrencinin odeme/access durumunu manuel tasdik edecek.
- Access sinif/grup bazinda verilebilir.
- Access ders/materyal seviyesinde kilitlenebilir.
- Dersler yayin tarihi veya sira kuraliyla acilabilir.

## Organization URL Modeli

MVP karari: path slug.

Ornek:

- Okul ana sayfasi: `sinfim.uz/my-school`
- Kurs sayfasi: `sinfim.uz/my-school/courses/russian-a1`
- Lead basvuru sayfasi: `sinfim.uz/my-school/apply` veya kurs detayindan basvuru formu

Subdomain modeli (`my-school.sinfim.uz`) simdilik out-of-scope. Bunun nedeni: DNS, SSL, tenant resolution ve deploy karmasikligini MVP'de artirmamak.

## Odev Tipleri

- Yazili cevap: ogrenci text olarak cevap verir.
- Dosya/fotograf: ogrenci defter fotografi veya dosya yukler.
- Test/quiz: opsiyonlu sorular, otomatik puanlama hedeflenir.
- Oral/audio message: ogrenci sesli cevap yukler; mentor manuel dinleyip feedback verir.

## MVP Icin Tavsiye

Auth'u aceleye getirmemek lazim. Ilk teknik dokumanda tenant izolasyonu, rol bazli yetki ve access kontrolu ana modul olarak ele alinmali.

Ilk olasi karar:

- Owner/Teacher/Mentor: telefon raqami + password.
- Student: telefon raqami + password veya invitation link/gecici kod ile ilk giris.
- Lead: login yok; public formdan isim + telefon raqami + ilgilendigi kurs.
- Her route/API request organization context'i tasimali.
- Her course/class/student/homework kaydi organization'a bagli olmali.
