# UX Dokümantasyon Rehberi

**Ne zaman kullanılır:** Problem tanımı tamamlandı, şimdi ürünü ekranlara dönüştürüyorsun.
**Süre:** 2–4 saat
**Çıktı:** User flow, ekran listesi, her ekranın state'leri ve içeriği.

> UX dokümanı wireframe çizmek değildir. Ekranları, akışları ve içerikleri **yazıyla** tarif etmektir.
> Bu doküman hem Claude'un UI üretmesi için context, hem de senin kafandaki ürünü netleştirmek için kullanılır.

---

## UX Dokümanı Neden Önemli?

Claude'a "login sayfası yap" dersen genel bir şey alırsın.
Claude'a bu dokümanı verirsen ürününe özel, tutarlı bir UI alırsın.

İyi bir UX dokümanı şunu sağlar:
- Claude doğru component'ları üretir (ilk seferinde)
- Eksik state'ler (boş sayfa, hata, yükleniyor) unutulmaz
- Ekranlar arası tutarlılık korunur

---

## Bölüm 1: Genel Ürün Bilgisi

### Ürün Adı ve Tek Cümle
> ___________________________

### Tasarım Tonu
Birkaç kelimeyle: minimal / canlı / kurumsal / teknik / sıcak / profesyonel

> ___________________________

### Renk ve Tipografi Kararı
- **Primary color:** (hex veya "koyu mavi, marka rengi")
- **Accent color:**
- **Background:**
- **Font:** (serif / sans-serif, örn: Inter, Geist)
- **Component stili:** shadcn-vue / Tailwind custom / Element Plus / Naive UI

> Karar vermekte zorlanıyorsan: **shadcn-vue + Tailwind**, dark mode opsiyonel.

---

## Bölüm 2: Kullanıcı Akışları (User Flows)

### Nasıl Yazılır?

Her akış için:
1. Başlangıç noktası (kullanıcı nerede?)
2. Her adım (ne görüyor, ne yapıyor)
3. Bitiş noktası (nereye ulaştı, ne değer elde etti)

**Örnek format:**
```
[Akış Adı]
Başlangıç: Kullanıcı uygulamayı ilk kez açıyor
→ Landing / onboarding sayfası görünüyor
→ "Kayıt ol" butonuna tıklıyor
→ Email + şifre formu
→ Email doğrulama gönderildi mesajı
→ Email'e tıklıyor → Dashboard
Bitiş: Kullanıcı ilk değer anına 1 adım uzakta
```

---

### Akış 1: Happy Path (İlk Değer Anı)

> Yeni bir kullanıcı, ürüne girip core value'ya ilk kez ulaşana kadar ne yaşıyor?

**Başlangıç:**
1. 
2. 
3. 
4. 
**Bitiş (değer anı):**

---

### Akış 2: Geri Dönen Kullanıcı

> Kullanıcı 2. kez giriyor. Nereye gidiyor, ne yapıyor?

1. 
2. 
3. 

---

### Akış 3: [İkinci Önemli Flow]

> ___________________________

1. 
2. 
3. 

---

### Edge Case'ler

| Durum | Ekran / Davranış |
|-------|-----------------|
| Yanlış şifre girdi | |
| Session süresi doldu | |
| İnternet bağlantısı yok | |
| İzin yokken erişmeye çalıştı | |
| Quota / limit aşıldı | |

---

## Bölüm 3: Ekran Listesi

Her ekran için: Ad, Amaç, Hangi akışta kullanılıyor.

| # | Ekran Adı | Amaç | Route / Path |
|---|-----------|------|-------------|
| 1 | | | |
| 2 | | | |
| 3 | | | |
| 4 | | | |
| 5 | | | |

---

## Bölüm 4: Ekran Detayları

Her ekran için ayrı bir bölüm doldur. Bu bölümler **UI prompt'larının ham maddesidir.**

---

### Ekran: [Ekran Adı]

**Amaç:** Bu ekranda kullanıcı ne yapmak istiyor?

**Layout:**
- [ ] Tam sayfa (auth/landing)
- [ ] Sidebar + içerik
- [ ] Üst bar + içerik
- [ ] Modal / drawer
- [ ] Mobil first

**İçerik (Yukarıdan Aşağı):**
1. (başlık, açıklama, buton...)
2. 
3. 

**Aksiyonlar (kullanıcı ne yapabilir):**
- Buton: [ad] → [ne olur]
- Form field: [ad] → [validasyon kuralı]
- Link: [ad] → [nereye gider]

**State'ler:**

| State | Görünüm / Davranış |
|-------|------------------|
| Loading | Skeleton loader / spinner |
| Boş (no data) | "Henüz ... yok, ilk ... oluştur" |
| Hata | Hata mesajı + retry butonu |
| Başarı | Success toast / redirect |

**Data (bu ekranda ne gösteriliyor):**
```
{
  field: tip,
  field: tip
}
```

---

### Ekran: [Ekran Adı]

*(Yukarıdaki yapıyı tekrarla)*

---

## Bölüm 5: Navigation Yapısı

### Web (Vue Router)

```
/                    → Landing / Home
/auth/login          → Giriş
/auth/register       → Kayıt
/dashboard           → Ana panel (auth gerekli)
/[kaynak]            → Liste
/[kaynak]/:id        → Detay
/[kaynak]/new        → Oluşturma
/settings            → Ayarlar
/settings/profile    → Profil
```

**Benim route yapım:**
```
/

```

### Auth Guard
- Auth gerektiren route'lar: 
- Auth gerektirmeyen route'lar: 

---

## Bölüm 6: Component Kataloğu

Tüm sayfalarda kullanılacak tekrar eden component'lar.

| Component | Nerede Kullanılıyor | Props | Notlar |
|-----------|-------------------|-------|--------|
| `AppButton` | Her yerde | variant, size, loading | |
| `AppInput` | Formlar | label, error, type | |
| `DataTable` | Liste ekranları | columns, rows, loading | |
| `EmptyState` | Boş listeler | icon, title, action | |
| `PageHeader` | Her içerik sayfası | title, subtitle, action | |
| | | | |

---

## ✓ Çıkış Kriteri

- [ ] Happy path akışı yazıldı
- [ ] Tüm ekranlar listelendi
- [ ] Her ekranın 3 state'i tanımlandı (loading, empty, error)
- [ ] Navigation yapısı belirlendi
- [ ] Tekrar eden component'lar listelendi

**Sonraki adım:** `templates/tech-doc.md`
