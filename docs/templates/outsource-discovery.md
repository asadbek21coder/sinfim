# Outsource Discovery Rehberi

**Ne zaman kullanılır:** Client'tan proje teklifi geldi, anlaşmadan önce.
**Süre:** Discovery çağrısı (1–2 saat) + yazma (1 saat)
**Çıktı:** Net scope, kabul kriterleri, milestone takvimi, teklif hazır.

> Belirsiz proje = Sonsuz revize = Ücretsiz çalışma.
> Bu rehber seni belirsizlikten korumak için var.

---

## Bölüm 1: Discovery Çağrısında Sorulacak Sorular

### 1. Ne İstiyorlar?

```
- Bu projeyi neden yapıyorsunuz? (asıl hedef nedir?)
- Kimler kullanacak? (son kullanıcı kimdir?)
- Bu proje şu an var mı? Sıfırdan mı başlıyoruz, yoksa devam mı?
- Varsa mevcut sistemle entegrasyon olacak mı?
- Referans aldıkları bir ürün / site var mı?
```

### 2. Kapsam

```
- Proje tamamlandığında ne görmek istiyorlar?
  (her birini listele, sonra "başka?" diye sor — biter diye düşündüklerinde daha fazla çıkar)
- "Nice to have" listesi var mı? (MVP sonrası için)
- Hangi platformlar? (web / mobil / her ikisi)
- Admin paneli gerekiyor mu?
- Entegrasyon ihtiyaçları? (ödeme, SMS, e-posta, harita, AI...)
```

### 3. Teknik Kısıtlar

```
- Tercih ettikleri bir teknoloji var mı? (zorunlu mu, tercih mi?)
- Hosting nerede olmalı? (kendi sunucuları, cloud, önemsemiyorlar)
- Mevcut altyapıyla uyumluluk var mı?
- Kaynak kodu teslim edilecek mi? (IP kime kalıyor?)
- Güvenlik / compliance gereksinimi? (KVKK, GDPR, sektöre özel)
```

### 4. Zaman ve Bütçe

```
- Ne zaman lazım? (gerçek tarih, "mümkün olan en kısa sürede" değil)
- Bu tarihin arkasında ne var? (etkinlik, yatırımcı sunumu, sözleşme...)
- Bütçe aralığı? (vermek istemeyebilirler, ama sormak zorundayız)
- Daha önce başka biriyle çalıştılar mı? Ne oldu?
```

### 5. Süreç ve İletişim

```
- Proje boyunca tek iletişim noktası kim?
- Karar alma süreci nasıl? (onay için kaç kişi?)
- Haftalık demo / check-in uygun mu?
- İçerik ve varlıkları kim sağlayacak? (metin, logo, fotoğraf)
```

---

## Bölüm 2: Discovery Çıktısı — Doldur

### Proje Özeti

> **Client:** ___________________________
> **Proje Adı:** ________________________
> **Görüşme Tarihi:** ___________________
> **Tahmini Başlangıç:** ________________
> **İstenen Bitiş:** ____________________

**Projenin Amacı (1 paragraf):**
> ___________________________

**Son Kullanıcı:**
> ___________________________

---

### Kapsam Analizi

**Kesin olarak teslim edilecekler:**
- 
- 
- 

**Kesin olarak teslim edilmeyecekler (scope dışı):**
- 
- 
- 

**Belirsiz kalan alanlar (netleştirilmesi gereken):**
- 
- 

**Entegrasyonlar:**
| Servis | Amaç | Kim Sağlıyor |
|--------|------|-------------|
| | | |

---

### Kabul Kriterleri

> Her teslim edilecek feature için "bu ne zaman bitti sayılır?"
> Muğlak olanı kabul etme: "güzel görünsün" → kabul kriteri değil.

| Feature | Kabul Kriteri |
|---------|--------------|
| | |
| | |
| | |

---

### Risk Analizi

| Risk | Olasılık | Etki | Önlem |
|------|---------|------|-------|
| Kapsam genişlemesi | Yüksek | Yüksek | Scope freeze + CR süreci |
| İçerik gecikmesi | Orta | Orta | İçerik takvime bağlama |
| Teknik bilinmezlik | | | Discovery sprint |
| Onay gecikmeleri | | | Haftalık demo zorunluluğu |

---

### Teklif Temeli

**Tahmini süre:** ___ hafta / ___ gün
**Ücret yapısı:**
- [ ] Sabit fiyat (Fixed price) — kapsam netse
- [ ] Zaman + malzeme (T&M) — kapsam belirsizse
- [ ] Milestone bazlı ödeme (önerilir)

**Milestone Planı:**
| Milestone | Teslim | Ödeme |
|-----------|--------|-------|
| Discovery + Prototip | Hft ___ | %___ |
| Alpha (core feature) | Hft ___ | %___ |
| Beta + UAT | Hft ___ | %___ |
| Final teslim | Hft ___ | %___ |

**Revize hakkı:** ___ tur (sonrası saatlik ücret)
**Destek süresi:** ___ hafta (lansman sonrası, kapsam dışı bug'lar için)

---

## Bölüm 3: Tuzak Soruları — Kritik Kontroller

Discovery biter bitmez bu listeyi gözden geçir:

### Kapsam Tuzakları
- [ ] "Admin paneli" beklentisi var mı? → Kapsama al veya çıkar, yaz.
- [ ] "Çoklu dil" gerekiyor mu?
- [ ] "Mobil uygulama" da mı isteniyor, sadece web mi?
- [ ] "İkinci bir platform" için ileride entegrasyon planı var mı?
- [ ] Kullanıcı sayısı beklentisi nedir? (100 / 10.000 / 1M — mimariyi etkiler)

### İçerik Tuzakları
- [ ] Metinleri kim yazacak? (sen mi, client mi)
- [ ] Görseller / ikon'lar kim sağlayacak?
- [ ] Logo ve marka rehberi var mı?

### Teknik Tuzakları
- [ ] Üçüncü taraf API key'leri kim temin edecek?
- [ ] Domain ve hosting kim kuracak?
- [ ] SSL sertifikası?
- [ ] E-posta altyapısı? (transactional email)

### Süreç Tuzakları
- [ ] Onay verecek kaç kişi var? (tek kişi değilse revize riski artar)
- [ ] Client'ın teknik bilgisi var mı? (UAT nasıl yapılacak?)
- [ ] Sözleşme gerekiyor mu? (genellikle evet)

---

## Bölüm 4: Standart Sözleşme Maddeleri (Checklist)

Teklifini göndermeden önce şunların sözleşmede olduğunu kontrol et:

- [ ] Net kapsam (in-scope / out-of-scope)
- [ ] Kabul kriterleri veya referansı
- [ ] Milestone takvimi ve ödeme zamanlaması
- [ ] Revize hakkı: kaç tur, neyi kapsıyor
- [ ] Scope dışı talepler için Change Request prosedürü
- [ ] Kaynak kodu ve IP devir koşulları
- [ ] Destek ve garanti süresi + kapsamı
- [ ] İletişim kanalı ve yanıt süresi
- [ ] Projenin iptal/askıya alınması durumunda ne olur

---

## Bölüm 5: Buradan Sonra

Discovery tamamlandı, kapsam netleşti.

**Sıradaki adım:** Ortak geliştirme yoluna katıl:
1. `templates/ux-doc.md` — ekranları ve akışları yaz
2. `templates/tech-doc.md` — teknik kararları al
3. `new-project-template/CLAUDE.md` — projeye kopyala, doldur
4. `templates/ui-prompts.md` — UI üret
5. `templates/dev-kickoff.md` — geliştirmeye başla

---

## Haftalık Milestone Yönetimi (Proje Boyunca)

Her hafta başında client'a gönder:

```
Hafta [N] özeti:
✓ Tamamlanan: [...]
→ Bu hafta: [...]
⚠ Bekleyen / Engel: [...]
```

Her demo'dan sonra yazılı onay al:

```
[Feature adı] demo'sunu [tarih]'de gösterdim.
Client: [onay / revize talebi]
Revize talebi scope içinde mi: Evet / Hayır (CR gerekiyor)
```
