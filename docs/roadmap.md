# Yol Haritası: Fikirden Lansmana

AI agent'larla solo/iki kişilik geliştirme için A–Z rehber.
Startup fikri mi var? Outsource proje mi aldın? İkisi de burada.

**Ortağın için tek dosya:** `new-project-template/partner.md` — teknik detay yok, sadece sprint durumu, bekleyen kararlar ve scope değişiklikleri.

---

## İki Yol, Ortak Bitiş

```
STARTUP FİKRİ          OUTSOURCE PROJE
      ↓                       ↓
 startup-idea.md      outsource-discovery.md
      ↓                       ↓
      └──────────┬────────────┘
                 ↓
          ux-doc.md          (2–4 saat)
                 ↓
         tech-doc.md         (1–2 saat)
                 ↓
       CLAUDE.md kurulumu    (30 dk)
                 ↓
        ui-prompts.md        (30 dk – 1 saat)
                 ↓
        dev-kickoff.md       (1. gün)
                 ↓
           LANSMAN
```

---

## STARTUP YOLU — Adım Adım

### Adım 1 — Fikri Kağıda Dök
**Dosya:** `templates/startup-idea.md`
**Süre:** 1–2 saat
**Amaç:** Kafandaki fikri test edilebilir varsayımlara çevir.

Çıktı: Problem statement, 3 kritik varsayım, hedef kullanıcı, kuzey yıldızı.

Ne zaman geçersin: Problem statement'ı 3 dakikada yazabildiğinde.

---

### Adım 2 — UX'i Yaz
**Dosya:** `templates/ux-doc.md`
**Süre:** 2–4 saat
**Amaç:** Ürünü ekranlara ve akışlara dönüştür. Kod yazmadan.

Çıktı: Happy path, ekran listesi, her ekranın state'leri.

Ne zaman geçersin: Her ekranı ve kullanıcı akışını sözle anlatabildiğinde.

---

### Adım 3 — Teknik Doküman
**Dosya:** `templates/tech-doc.md`
**Süre:** 1–2 saat
**Amaç:** Veri modeli, API contract ve stack kararı.

Çıktı: Entity tabloları, endpoint listesi, Go + Vue klasör yapısı.

Ne zaman geçersin: "Hangi endpoint'i ne zaman yazarım?" sorusuna cevabın olduğunda.

---

### Adım 4 — AI Agent Context Hazırla
**Dosya:** `new-project-template/CLAUDE.md` → proje klasörüne kopyala
**Süre:** 30 dakika
**Amaç:** Claude her session'da projeyi sıfırdan anlasın.

```bash
cp -r blueprints/project-builder/new-project-template/ projects/proje-adi/
```

CLAUDE.md içindeki alanları doldur: proje özeti, stack, klasör yapısı, kurallar.

---

### Adım 5 — UI Üret
**Dosya:** `templates/ui-prompts.md`
**Süre:** 30 dk – 1 saat hazırlık, sonra iterasyon
**Amaç:** UX dokümanından prompt üret → Claude / v0 / Lovable ile UI al.

Çıktı: Her ekran için Vue component'ları.

---

### Adım 6 — Geliştirmeye Başla
**Dosya:** `templates/dev-kickoff.md`
**Süre:** 1. gün kurulum, sonra haftalık sprint
**Amaç:** golang-blueprint + vue-blueprint'i klonla, ilk vertical slice'ı çalıştır.

---

## OUTSOURCE YOLU — Adım Adım

### Adım 1 — Client Discovery
**Dosya:** `templates/outsource-discovery.md`
**Süre:** Discovery çağrısı (1–2 saat) + yazma (1 saat)
**Amaç:** Client'tan doğru soruları sor, scope'u kilitle, tuzakları önle.

Çıktı: SOW taslağı, kabul kriterleri, milestone takvimi.

### Adım 2 — Buradan ortak yola katıl
Scope netleştikten sonra Startup Yolu'nun 2. adımından devam et.
(UX → Tech Doc → CLAUDE.md → UI → Dev)

---

## Hızlı Referans Tablosu

| Dosya | Ne Zaman | Süre |
|-------|----------|------|
| `templates/startup-idea.md` | Startup fikri var | 1–2 saat |
| `templates/ux-doc.md` | Ekranları tarif etmem lazım | 2–4 saat |
| `templates/tech-doc.md` | Teknik kararları netleştir | 1–2 saat |
| `new-project-template/CLAUDE.md` | Projeye başlamadan önce | 30 dk |
| `templates/ui-prompts.md` | UI üretmek istiyorum | 30 dk–1 saat |
| `templates/dev-kickoff.md` | Kod yazmaya başlıyorum | 1. gün |
| `templates/outsource-discovery.md` | Client projesi alıyorum | Discovery çağrısı |

---

## Altın Kurallar

1. **Sırayı atlama.** UX yokken UI prompt yazmak zaman kaybı.
2. **CLAUDE.md'yi güncelle.** Her sprint sonunda "Şu an ne üzerinde çalışıyoruz" bölümünü güncelle.
3. **Vertical slice.** Bir feature'ı baştan sona bitir, sonra diğerine geç.
4. **Küçük PR, sık deploy.** "Büyük PR review'u" yoktur, sadece "küçük PR merge'ü" vardır.
5. **Doküman önce, kod sonra.** Claude'a "şu endpoint'i yaz" demeden önce tech-doc'ta o endpoint tanımlı mı kontrol et.
