# Proje Başlangıç Listesi

> **Proje:** ___________________________
> **Başlangıç:** ________________________
> **Tür:** Startup / Outsource / Hybrid

Her yeni projede bu dosyayı yukarıdan aşağı işaretle.
İşaretlemeden sonraki adıma geçme.

**Görev sahibi:**
- 👤 Sen yapacaksın (AI yardım edemez veya etmemeli)
- 🤖 AI'a vereceksin (prompt hazır, sen sadece tetikleyeceksin)
- 👥 İkisi birlikte (sen yönlendirirsin, AI üretir, sen onaylarsın)

---

## AŞAMA 0 — Fikir / Kapsam

> Startup ise bu aşama, outsource ise → Aşama 1'e atla.

- [ ] 👤 `templates/startup-idea.md`'yi kopyala → `docs/startup-idea.md`
- [ ] 👤 Problem statement'ı yaz (3 dakikadan fazla sürerse dur, fikri olgunlaştır)
- [ ] 👤 5 kritik varsayımı listele
- [ ] 👤 En az 2 persona tanımla
- [ ] 👤 Kuzey yıldızı metriğini seç (bir tane, değiştirme)
- [ ] 👤 MVP sınırını çiz: in-scope / out-of-scope listesi
- [ ] 👤 Ortakla birlikte gözden geçir, onayını al
- [ ] 👤 `partner.md` → Alınan Kararlar: "MVP scope onaylandı"

**Çıkış kriteri:** Problem statement 3 dakikada yazıldı, scope yazılı ve onaylı.

---

## AŞAMA 1 — Client Discovery (Sadece Outsource)

- [ ] 👤 `templates/outsource-discovery.md`'yi kopyala → `docs/discovery.md`
- [ ] 👤 Discovery çağrısını yap — soruları sor, notları al
- [ ] 👥 Discovery çıktısını doldur (sen doldururken AI organize eder)
- [ ] 👤 Tuzak soruları listesini gözden geçir — eksik var mı?
- [ ] 👤 Scope'u client'a yazılı olarak gönder, onay al
- [ ] 👤 Milestone takvimini client ile mutabık kal
- [ ] 👤 `partner.md` → Alınan Kararlar: scope + tarihler
- [ ] 👤 Sözleşme checklist'ini gözden geçir

**Çıkış kriteri:** Yazılı onaylı scope ve milestone takvimi var.

---

## AŞAMA 2 — UX Dokümanı

- [ ] 👤 `templates/ux-doc.md`'yi kopyala → `docs/ux-doc.md`
- [ ] 👥 Happy path akışını yaz (sen tarif et, AI yapılandırır)
- [ ] 👤 Tüm ekranları listele
- [ ] 👥 Her ekranın içeriğini ve aksiyonlarını yaz
- [ ] 👤 Her ekran için 3 state'i tanımla: loading / empty / error
- [ ] 👥 Navigation yapısını (route'ları) belirle
- [ ] 👤 Ortakla birlikte gözden geçir — akışlar mantıklı mı?
- [ ] 👤 `partner.md` → Alınan Kararlar: "UX onaylandı"

**Çıkış kriteri:** Her ekranı ve akışı sözle anlatabiliyorsun.

---

## AŞAMA 3 — Teknik Doküman

- [ ] 👤 `templates/tech-doc.md`'yi kopyala → `docs/tech-doc.md`
- [ ] 👤 Modül listesini belirle
- [ ] 👥 Her modül için tablo SQL'lerini yaz
  ```
  AI'a ver: "Bu entity'leri Go enterprise blueprint kurallarına göre
  PostgreSQL tablosu olarak yaz: [entity listesi]"
  ```
- [ ] 👥 ERD'yi oluştur (Mermaid)
  ```
  AI'a ver: "Bu tabloları Mermaid ERD olarak çiz: [tablolar]"
  ```
- [ ] 👤 Tüm UC'leri (use case) listele — her UC için operation ID belirle
- [ ] 👥 Her UC için request/response örneği yaz
- [ ] 👤 Error code kataloğunu başlat
- [ ] 👤 `docs/tech-doc.md` tamamlandı mı kontrol et

**Çıkış kriteri:** Tüm UC'ler listelendi, tablolar SQL olarak tanımlandı.

---

## AŞAMA 4 — Proje Kurulumu

- [ ] 👤 Ana proje klasörünü oluştur
- [ ] 👤 `go-enterprise-blueprint-main`'i kopyala → `backend/`, `.git` sil
- [ ] 👤 `vue-blueprint-web`'i kopyala → `frontend/`, `.git` sil
- [ ] 👤 `new-project-template/CLAUDE.md`'yi kopyala → `./CLAUDE.md`
- [ ] 👤 `new-project-template/partner.md`'yi kopyala → `./partner.md`
- [ ] 👤 `git init` + ilk commit: `"chore: init project from blueprints"`
- [ ] 👤 `backend/.env.example` → `backend/.env`, değerleri doldur
- [ ] 👤 `frontend/.env.example` → `frontend/.env`, değerleri doldur
- [ ] 👤 `make infra-up` → Docker servisleri ayağa kaldır
- [ ] 👤 `make run` → backend çalışıyor mu?
- [ ] 👤 `npm install && npm run dev` → frontend çalışıyor mu?

**Vue Blueprint Customisation (bir kez):**
- [ ] 👤 `src/types/auth.ts` → `Role` union type backend'e göre güncelle
- [ ] 👤 `src/main.ts` → `roleHomeMap` güncelle
- [ ] 👤 `src/layouts/AuthLayout.vue` → `appName` + `appSubtitle`
- [ ] 👤 `index.html` → `<title>`
- [ ] 👤 `tailwind.config.js` → brand renkleri (varsa)
- [ ] 👤 `package.json` → `"name"` alanı

**Çıkış kriteri:** Backend ve frontend ayağa kalkıyor, boş blueprint çalışıyor.

---

## AŞAMA 5 — CLAUDE.md Doldur

- [ ] 👤 Proje adı, amaç, hedef kullanıcı, kuzey yıldızı
- [ ] 👤 Deploy hedefi
- [ ] 👤 Modüller tablosunu doldur
- [ ] 👤 "Şu An Ne Üzerinde Çalışıyoruz" → Sprint 1 hedefini yaz
- [ ] 👤 `partner.md` → Milestone'lar tablosunu doldur

**Çıkış kriteri:** CLAUDE.md'yi okuyan biri projeyi anlayabilir.

---

## AŞAMA 6 — İlk Modül: UC Dokümanları

> Her UC için bu bloğu tekrarla.

- [ ] 👤 `docs/specs/modules/{module}/overview.md` yaz (yoksa)
- [ ] 🤖 ERD'yi oluştur:
  ```
  Agent: system-analyst
  "Write ERD for {module} module based on these tables: [tablolar]"
  ```
- [ ] 🤖 UC dokümanını yaz:
  ```
  Agent: system-analyst
  "Write UC doc for {operation-id} based on tech-doc.md → [UC satırı]"
  ```
- [ ] 👤 UC dokümanını oku, mantıklı mı? Eksik var mı?
- [ ] 👤 Onayladıysan bir sonraki adıma geç

**Çıkış kriteri:** UC dokümanı yazılı ve onaylı — implementation başlayabilir.

---

## AŞAMA 7 — Backend Implementation

> Her UC için bu bloğu tekrarla. Sıra: migration → domain → infra → UC → ctrl.

- [ ] 🤖 Implementation planla:
  ```
  Agent: architect
  "Plan implementation of {operation-id} UC in {module} module."
  ```
- [ ] 🤖 Migration yaz:
  ```
  Agent: go-coder
  "Write goose migration for {module} init schema based on docs/tech-doc.md"
  ```
- [ ] 👤 Migration'ı gözden geçir, çalıştır: `make migrate-up`
- [ ] 🤖 Kod + test (paralel):
  ```
  Agent: go-coder  → "Implement {operation-id} use case"
  Agent: go-tester → "Write system tests for {operation-id} use case"
  ```
- [ ] 🤖 Review:
  ```
  Agent: reviewer
  "Review {operation-id} implementation against backend-guidelines skill."
  ```
- [ ] 👤 Review çıktısını oku, fix et
- [ ] 👤 Verification: `make fmt && make lint && make test && make test-system`
- [ ] 👤 Commit: `"feat({module}): implement {operation-id}"`

**Çıkış kriteri:** 4'lü verification geçiyor, commit yapıldı.

---

## AŞAMA 8 — UI Üretimi

> Her ekran için bu bloğu tekrarla.

- [ ] 👤 `templates/ui-prompts.md`'den ilgili şablonu seç
- [ ] 👤 `docs/ux-doc.md`'den o ekranın detaylarını kopyala
- [ ] 👥 API dosyasını üret:
  ```
  "src/api/{domain}.ts dosyasını yaz. [ui-prompts.md → API Dosyası şablonu]"
  ```
- [ ] 👤 Üretilen API dosyasını kontrol et — endpoint path'ler tech-doc ile uyuşuyor mu?
- [ ] 👥 View'ı üret:
  ```
  "[ui-prompts.md → Prompt Şablonu: Yeni Ekran]"
  ```
- [ ] 👤 Üretilen component'ı kontrol listesiyle doğrula:
  - `<script setup lang="ts">` var mı?
  - Loading / empty / error state var mı?
  - `any` tip var mı? (olmamalı)
  - Inline style var mı? (olmamalı)
  - Predefined class kullanıyor mu? (`btn-primary`, `form-input` vb.)
- [ ] 👤 Route'u `src/router/index.ts`'e ekle (meta.roles ile)
- [ ] 👤 Tarayıcıda aç, manuel test et
- [ ] 👤 Commit: `"feat({domain}): add {EkranAdi} view"`

**Çıkış kriteri:** Ekran tarayıcıda çalışıyor, manuel test edildi.

---

## AŞAMA 9 — Sprint Sonu Review

> Her sprint sonunda (her hafta cuma).

- [ ] 👤 `make test && make test-system` → hepsi geçiyor mu?
- [ ] 👤 `npm run type-check && npm run lint` → temiz mi?
- [ ] 👤 `partner.md` → "Haftalık Özet" güncelle
- [ ] 👤 `partner.md` → biten UC'leri işaretle
- [ ] 👤 `CLAUDE.md` → "Şu An Ne Üzerinde Çalışıyoruz" güncelle
- [ ] 👤 Ortakla 15 dk konuş: ne bitti, önümüzdeki hafta ne
- [ ] 👤 `partner.md` → "Senden Cevap Bekleyen Sorular" güncelle

---

## AŞAMA 10 — Lansman Öncesi

- [ ] 👤 Tüm UC'lerin system test'i var mı? `make test-system`
- [ ] 👤 Happy path uçtan uca manuel test edildi
- [ ] 👤 Error state'ler manuel test edildi
- [ ] 👤 Production `.env` değerleri set edildi
- [ ] 👤 Migration'lar production'da çalıştırıldı: `make migrate-up`
- [ ] 👤 Error monitoring aktif (Sentry vb.)
- [ ] 👤 Uptime monitoring aktif
- [ ] 🤖 Launch checklist dokümanı üret:
  ```
  "Projenin mevcut durumuna göre bir launch checklist yaz."
  ```
- [ ] 👤 Ortakla birlikte gözden geçir
- [ ] 👤 `partner.md` → Milestone: "Lansman" tarihini işaretle

---

## Hızlı Başvuru: Ne Zaman AI, Ne Zaman Sen?

| Görev | Kim |
|-------|-----|
| Problem statement yazmak | 👤 Sen — AI fikri bilemez |
| Kullanıcıyla konuşmak | 👤 Sen |
| Scope kararı vermek | 👤 Sen + ortak |
| Teknik karar vermek | 👤 Sen |
| SQL / ERD üretmek | 🤖 AI (sen onaylarsın) |
| UC dokümanı yazmak | 🤖 system-analyst (sen onaylarsın) |
| Implementation planlamak | 🤖 architect (sen onaylarsın) |
| Kod yazmak | 🤖 go-coder (sen review edersin) |
| Test yazmak | 🤖 go-tester (sen çalıştırırsın) |
| Code review | 🤖 reviewer (sen fix edersin) |
| Vue component üretmek | 🤖 AI (sen test edersin) |
| Manuel test yapmak | 👤 Sen |
| Commit yazmak | 👤 Sen |
| Partner'a açıklamak | 👤 Sen |
