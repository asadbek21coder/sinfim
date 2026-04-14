# Startup Fikir ve Problem Tanimi

## 1. Tek Cumle Tanimi

Sinfim.uz, Ozbekistan'da okul, egitim merkezi, ogretmen veya kendi brendiyle online kurs satan egitim ekipleri icin kurs, sinif/grup, video ders, PDF materyal, odev, mentor kontrolu, test, access ve ogrenci ilerlemesini tek yerde yoneten online okul platformudur.

## 2. Problem Statement

Ozbekistan'da online kurs veren okullar, egitim merkezleri, ogretmenler ve egitim ekipleri, kurs satisindan sonra ogrencileri Telegram kanal/gruplarina alip hazir video dersleri, PDF materyalleri ve odevleri gun asiri manuel gonderiyor.

Mevcut Telegram tabanli cozumler kurs yonetimi icin yetersiz kaliyor: video player rahat degil, materyaller buyuk dosya olarak dagiliyor, dersleri sirali acma manuel yapiliyor, odev toplama ve geri bildirim mentorlar icin daginik ilerliyor, ogrencinin progress ve basarisi olculemiyor.

Biz bu akisi online okul/sinif mantigina cevirerek egitim ekiplerinin kendi brendi altinda kurs, sinif/grup, ders, materyal access, odev, test, mentor feedback ve ilerleme takibini sistematik sekilde yonetmesini sagliyoruz.

## 3. Hedef Kullanici

### Persona 1: Okul / Egitim Merkezi Sahibi

- Kim: Kendi online okulunu, egitim merkezini veya brendini yoneten kisi.
- Gunluk rutini: Ders videosu hazirlar, yeni grup acar, odeme alan ogrencileri Telegram kanalina ekler, materyalleri sirayla yayinlar.
- Problem ani: Birden fazla grup olunca hangi ogrenci hangi dersi acti, odevi kim yapti, mentor kime cevap verdi takip etmek zorlasir.
- Su an ne yapiyor: Telegram kanal/grubu, Google Drive, PDF, Excel/Sheets, manuel mesajlasma ve mentor takibi kullaniyor.
- Basari kriteri: Kendi brendi altinda kurslar acmak, siniflar/gruplar yonetmek, odemeye gore access vermek, dersleri takvimlemek, odev/test sonuclarini ve mentorlarin isini takip etmek kolaylasir.

### Persona 2: Ana Ogretmen / Kurs Ureticisi

- Kim: Kendi basina calisan veya okul icinde ders icerigi ureteren ogretmen.
- Gunluk rutini: Ders plani yapar, videolari hazirlar, PDF materyal ve egzersizleri olusturur.
- Problem ani: Icerik hazir olsa bile bunu siniflara duzenli dagitmak, odevleri toplamak ve ogrencinin nerede kaldigini gormek zorlasir.
- Su an ne yapiyor: Telegram kanalina video/PDF yollar, odevleri mentor veya kendisi DM uzerinden kontrol eder.
- Basari kriteri: Bir kurs yapisini bir kez kurar, sonra farkli siniflara/gruplara planli sekilde uygular.

### Persona 3: Mentor / Cirak Ogretmen

- Kim: Ana ogretmenin altinda calisan, odev kontrol eden ve haftalik pratik voice chat yapan mentor.
- Problem ani: Ogrenciler odev defteri fotografini Telegram'dan gonderir; cevaplar kaybolur, kime donuldugu takip edilemez.
- Su an ne yapiyor: Telegram DM veya grup mesajlariyla odev alir, tek tek inceler, cevap yazar/ses atar.
- Basari kriteri: Kendi sinifindaki odevleri tek listede gorur, kontrol eder, hazir feedback sablonlari ve puanlama ile hizli doner.

### Persona 4: Ogrenci

- Kim: Online veya hybrid kurs egitimi alan, telefondan ders izleyen ogrenci.
- Problem ani: Telegram'da video, PDF, eski mesajlar ve odev cevaplari karisir; hangi ders sirada, ne teslim edilmeli net olmaz.
- Su an ne yapiyor: Kanalda mesaj arar, PDF indirir, deftere odev yapip fotograf atar.
- Basari kriteri: Bugunku dersini, materyalini, odevini ve sonucunu tek yerde gorur.

## 4. Varsayimlar

| # | Varsayim | Kritik? | Nasil test ederim? | Durum |
|---|----------|---------|-------------------|-------|
| 1 | Telegram tabanli kurs yoneten egitim ekipleri bu daginikligi ciddi problem olarak goruyor. | Evet | 5-10 okul/egitim merkezi/ogretmenle problem interview yap. | Senin ilk dogrulamana gore: Evet |
| 2 | Ogretmenler mevcut Telegram akisini tamamen birakmasa bile ders/odev/progress icin ayrica platform kullanmaya razi olur. | Evet | Clickable prototype veya basit demo goster. | Senin ilk dogrulamana gore: Evet |
| 3 | Ogrenciler Telegram yerine/yaninda web veya mobil web platformuna girmeyi kabul eder. | Evet | Bir pilot grupta 1 haftalik test yap. | Senin ilk dogrulamana gore: Evet |
| 4 | Mentorlar odev kontrolunu platformdan yaparsa zaman kazanir. | Evet | 1 mentorla eski akisa karsi yeni akisi sure olarak karsilastir. | Senin ilk dogrulamana gore: Evet |
| 5 | Ilk MVP'de full payment entegrasyonu olmadan manuel odeme + platform icinde access kontrol yeterli olur. | Evet | Ilk pilotta odemeyi disarida al, platformda odeme/access durumunu manuel tasdik et. | Karar: Evet |

Kirmizi cizgi: Eger varsayim #1 ve #2 yanlissa bu fikri yeniden konumlandirmak gerekir.

## 5. Deger Onerisi

Okul, egitim merkezi veya kendi brendiyle kurs satan ogretmen, Sinfim sayesinde kurslarini ve siniflarini tek yerde yonetir; video dersleri sirayla acar, PDF ve odevleri duzenli dagitir, odemeye gore access verir, mentor feedback surecini takip eder ve ogrenci ilerlemesini olcer.

Telegram + Drive + Excel kombinasyonundan farkli olarak bu platform kursun asil operasyonunu tek yerde toplar.

## 6. Kuzey Yildizi Metrigi

Secilen metrik: Haftalik tamamlanan ders + teslim edilen odev sayisi.

Neden: Platformun asil degeri, ogretmenin kurs akisini duzenlemesi ve ogrencinin ders/odev dongusune devam etmesi.

Mevcut deger: 0

3 aylik hedef: 3 pilot okul/ogretmen, toplam 5 aktif sinif/grup, haftalik en az 100 ders tamamlama veya odev teslimi.

Nasil olcegim: Her ders goruntuleme/tamamlama ve her odev teslimini event olarak kaydet.

## 7. Odeme / Kullanim Niyeti Testi

- [ ] Problem interview: Telegram uzerinden kurs satan 10 okul/egitim merkezi/ogretmenle konus.
- [ ] Demo talebi: En az 3 okul veya ogretmene basit akisi gosterip "pilot grup acar misin?" diye sor.
- [ ] Pilot: 1 gercek kurs grubunu MVP'de yonet.
- [ ] Odeme niyeti: Ogretmenlere aylik platform ucreti veya ogrenci basina ucret modelini sor.

Sonuc: Ilk icgoruya gore problem ciddi; ogrenciler icin rahatsizlik ve ekipler icin operasyon zorlugu var. Disarida gercek interview ile dogrulanacak.

Karar: Simdilik devam. MVP'de odeme entegrasyonu disarida kalacak, fakat platform icinde odeme/access durumunu manuel tasdik etme ve materyal access kontrolu olacak.

## 8. MVP Siniri

### In-scope

- Ogretmen kaydi ve girisi.
- Okul / egitim merkezi / ogretmen icin organization veya brand profili olusturma.
- MVP'de organization olusturmanin superadmin kontrolunde olmasi; herkesin kendi kendine okul acmamasi.
- Organization icinde kurs olusturma.
- Kurs icinde sinif/grup olusturma.
- Sinif/grup icin ogrenci ekleme.
- Okul admini veya mentor tarafindan kabul edilen ogrenciyi manuel kayit etme.
- Public okul sayfasindan gelen potansiyel ogrenci/lead basvurularini toplama.
- Lead'i sonradan gercek ogrenciye cevirme ve uygun sinif/gruba ekleme.
- Organization icinde mentor ekleme ve mentoru bir veya birden fazla sinif/gruba atama.
- Ogrencinin odeme/access durumunu manuel tasdik etme.
- Odemeye/access durumuna gore ders ve materyal erisimi verme/kapatma.
- Dersleri sira ve yayin tarihiyle olusturma.
- Telegram channel uzerinden video stream referansi ekleme.
- PDF/materyal ekleme.
- Odev olusturma: yazili cevap, test/quiz veya oral/audio message tipi.
- Ogrencinin odev teslim etmesi: fotograf veya dosya yukleme.
- Mentorun odev kontrol etmesi, durum/puan/yorum vermesi.
- Basit opsiyonlu test/quiz.
- Ogrenci panelinde bugunku ders, materyal, odev ve sonuc gorunumu.
- Ogretmen panelinde grup progress ve teslim durumlari.

### Out-of-scope

- Otomatik odeme entegrasyonu; odeme disarida alinacak.
- Tam mobil uygulama; ilk versiyon mobil uyumlu web olabilir.
- Canli ders/voice chat altyapisi; ilk versiyonda Zoom/Telegram linki yeterli olabilir.
- AI ile otomatik odev kontrolu.
- Sertifika sistemi.
- Marketplace: ogrencilerin kurs kesfetmesi ve satin almasi.
- Cok gelismis video hosting/transcoding altyapisi. Ilk versiyonda Telegram stream yaklasimi denenir.
- Gamification, leaderboard, rozetler.
- Gelismis public landing + marketing funnel. Basit public okul/kurs lead formu MVP icinde kalacak.

## 9. Mevcut Rakipler / Alternatifler

| Cozum | Ne yapiyor | Neden yetersiz olabilir | Kullanici neden bizi secer |
|-------|------------|-------------------------|----------------------------|
| Telegram + kanal/grup | Duyuru, video/PDF dagitimi, mentor mesajlasmasi | Kurs progress, odev takibi, test, access kontrol, dashboard ve sistematik sinif yonetimi zayif | Mevcut akisi bozmadan egitim operasyonunu duzenler |
| Google Classroom | Sinif, materyal, odev ve puanlama | Kurs satisi/cohort/drip video akisi ve yerel online kurs modeli icin yeterince ticari olmayabilir | Ogretmenlerin sattigi grup bazli kurs akisi icin ozellesir |
| Moodle | Guclu klasik LMS | Kurulum, yonetim ve UX kucuk ogretmenler icin agir olabilir | Daha sade, Telegram'dan gelen ogretmene uygun olur |
| Thinkific / Teachable / Kajabi / Podia | Online kurs satisi, ders yayinlama, icerik dagitimi, bazi drip/community ozellikleri | Yerel odeme, mentorlu odev kontrolu ve Telegram benzeri cohort operasyonuna tam uymayabilir | Ozbekistan'daki grup + mentor + manuel odeme modeline gore tasarlanir |
| LearnWorlds / Mighty Networks | Kurs, community ve interaktif egitim ozellikleri | Genel platform; yerel dil, fiyat ve operasyon uyumu sorun olabilir | Daha hedefli, daha hafif ve yerel ihtiyaca yakin olur |

Not: Thinkific ve LearnWorlds gibi platformlarda dersleri zamanla acma (drip/drip-feed) ozelligi var. Bu yuzden farki sadece "dersleri gun gun aciyoruz" olarak degil, "Telegram'da satis yapan yerel okul/ogretmenin sinif + mentor + odev + feedback + access operasyonunu duzenliyoruz" diye konumlandirmak daha guclu.

Kontrol edilecek referanslar:

- Thinkific: drip schedule
- Teachable: lesson content ve quiz/open response bloklari
- LearnWorlds: drip-feed courses
- Mighty Networks: community, courses ve memberships

## Cikis Kriteri

- [ ] Problem statement 3 dakikada anlatilabiliyor.
- [ ] En az 5 ogretmenle problem interview yapildi.
- [x] MVP'de odeme entegrasyonu olup olmayacagina karar verildi: olmayacak, odeme disarida kalacak.
- [x] MVP'de odeme/access kontrolu olup olmayacagina karar verildi: olacak, manuel tasdik ve access verme olacak.
- [x] Video hosting stratejisi secildi: ilk MVP'de Telegram channel uzerinden stream referansi denenecek.
- [x] Mentor modeli belirlendi: bir mentor bir veya birden fazla sinif/gruptan sorumlu olabilir.
- [x] Odev tipleri belirlendi: yazili, test/quiz ve oral/audio message desteklenecek.
- [x] Ogrenci katilim modeli belirlendi: admin/mentor manuel kayit edebilir, ayrica public okul sayfasindan lead basvurusu alinabilir.
- [x] Organization URL modeli belirlendi: MVP'de `sinfim.uz/{school-slug}`, subdomain sonraya.
- [x] MVP auth prensibi belirlendi: SMS OTP ve Telegram login yok; dis servis bagimliligini azaltmak icin telefon raqami + password/davet kodu kullanilacak.
- [x] Okul/organization creation modeli belirlendi: MVP'de superadmin kontrollu, public self-service okul acma yok.
- [ ] Ilk pilot icin 1 ogretmen ve 1 gercek grup bulundu.
- [ ] MVP siniri onaylandi.

## Acik Sorular

- Ilk hedef kitle: her tur ders/kurs videosu satan okul, egitim merkezi, ogretmen ve mentor ekibi. Ilk pilotta yine de bir niche secilecek mi?
- Ogrenci login detayi: telefon raqami + password mu, yoksa ilk giris icin davet/gecici kod + sonra password mu?
- Odeme ilk versiyonda disarida kalacak; platformda odeme/access tasdigi olacak. Bunun ekrani nasil olmali?
- Lead toplama ekrani nasil olmali: public okul sayfasi, kurs sayfasi veya sadece basit basvuru formu?
- Video dosyalari ilk MVP'de Telegram channel uzerinden stream edilecek. Teknik olarak hangi metadata saklanacak?
- Auth ve tenant modeli nasil olmali: okul sahibi kendi okulunu platform icinde kendi platformu gibi kullanacak; her data organization siniri icinde kalmali.
- Odev kontrolu ilk versiyonda: yazili ve oral odevler manuel kontrol, test/quiz otomatik puanlama mi olacak?
- Haftalik voice chat ilk versiyonda Zoom/Telegram linki olarak verilecek.
