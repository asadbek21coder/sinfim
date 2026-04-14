# Geliştirme Başlangıç Rehberi

**Ne zaman kullanılır:** UX + tech-doc hazır, CLAUDE.md dolduruldu, şimdi kod yazılıyor.
**Stack:** `go-enterprise-blueprint-main` + `vue-blueprint-web`

---

## Gün 0: Proje Kurulumu (Tek Seferlik)

### 1. Blueprint'leri Kopyala

```bash
# Ana proje klasörü
mkdir my-project && cd my-project

# Backend — go-enterprise-blueprint'i kopyala
cp -r ../blueprints/go-enterprise-blueprint-main/ backend/
cd backend && rm -rf .git && cd ..

# Frontend — vue-blueprint-web'i kopyala
cp -r ../blueprints/vue-blueprint-web/ frontend/
cd frontend && rm -rf .git && cd ..

# CLAUDE.md'yi projeye kopyala
cp ../blueprints/project-builder/new-project-template/CLAUDE.md ./CLAUDE.md

# Git başlat
git init
git add .
git commit -m "chore: init project from blueprints"
```

### 2. CLAUDE.md'yi Doldur

Kopyalanan `CLAUDE.md`'yi aç:
- Proje adı, amaç, hedef kullanıcı
- Aktif modüller
- Sprint hedefi

### 3. Backend Kurulumu

```bash
cd backend
cp .env.example .env          # .env'i doldur
make infra-up                 # Docker ile PostgreSQL + diğer servisler
make migrate-up               # migration'ları çalıştır
make run                      # sunucuyu başlat → http://localhost:8080
```

### 4. Frontend Kurulumu

```bash
cd frontend
cp .env.example .env          # VITE_API_BASE_URL=/api/v1
npm install
npm run dev                   # → http://localhost:5173
```

### 5. Vue Blueprint Customisation (Proje Başında Bir Kez)

```
src/types/auth.ts     → Role union type'ı güncelle
src/main.ts           → roleHomeMap güncelle
src/router/index.ts   → placeholder route'ları değiştir
src/layouts/AppLayout.vue  → SidebarItem girdileri ekle
src/layouts/AuthLayout.vue → appName + appSubtitle
index.html            → <title>
tailwind.config.js    → brand renkleri
package.json          → "name" alanı
```

---

## Feature Geliştirme Protokolü

### Belge Önce, Kod Sonra

Her feature için UC dokümanı yazılmadan implementation başlamaz.

```
docs/specs/modules/{module}/
├── overview.md              ← modül yoksa önce bunu yaz
├── ERD.md                   ← entity ilişkilerini çiz
└── usecases/{domain}/{operation}.md  ← her UC için
```

### Implementation Sırası (Bottom-Up)

```
1. Migration          migrations/{timestamp}_{module}_{aciklama}.sql
2. Domain             internal/modules/{module}/domain/
3. Infra              internal/modules/{module}/infra/postgres/
4. PBLC               internal/modules/{module}/pblc/ (gerekiyorsa)
5. Use Case           internal/modules/{module}/usecase/{domain}/{op}/usecase.go
6. Controller         internal/modules/{module}/ctrl/http/
7. DI Wiring          internal/modules/{module}/module.go
8. System Test        tests/system/modules/{module}/{domain}/{op}_test.go
```

### Agent Kullanımı (Go Blueprint Pipeline)

```bash
# Yeni UC için standart pipeline:

# 1. Architect — plan
Agent: architect
"Plan the implementation of [operation-id] UC in [module] module."

# 2. System Analyst — UC doc yaz
Agent: system-analyst
"Write the UC doc for [operation-id] based on architect's plan."

# 3. Go Coder + Go Tester — paralel
Agent: go-coder    → "Implement [operation-id] use case"
Agent: go-tester   → "Write system tests for [operation-id] use case"

# 4. Reviewer
Agent: reviewer
"Review [operation-id] implementation against backend-guidelines."

# 5. Fix + verify
make fmt && make lint && make test && make test-system
```

### Yeni Modül Oluşturma

```bash
# Go Blueprint slash command
/generate-module
```

Bu komut modülün iskeletini (`domain/`, `usecase/`, `infra/`, `ctrl/`, `embassy/`) oluşturur.

---

## Verification — Her Zaman

```bash
# Backend — hepsi geçmeden kod teslim edilmez
make fmt              # önce format
make lint             # golangci-lint
make test             # unit tests (pkg/)
make test-system      # system tests (100% UC coverage)

# Frontend
npm run type-check    # TypeScript strict
npm run lint          # ESLint
npm run build         # production build başarılı mı
```

---

## Sprint Yapısı (Haftalık)

### Pazartesi — Planlama (30–45 dk)

```
Bu haftanın hedefi: [hangi UC'ler bitecek]

Backend:
  UC 1: [module/domain/operation] — doc hazır mı? → evet/hayır
  UC 2: ...

Frontend:
  Ekran 1: [view adı] — API hazır mı? → evet/hayır

CLAUDE.md güncelle: "Şu An Ne Üzerinde Çalışıyoruz" bölümü
```

### Salı–Perşembe — Build

Her UC için sıra:
1. UC doc yaz
2. architect planla
3. go-coder + go-tester paralel
4. reviewer
5. fix → verify
6. Vue view/api yaz
7. Commit

### Cuma — Review + Deploy

```bash
# Tüm testler
make test && make test-system
npm run type-check && npm run lint

# CLAUDE.md güncelle
# "Şu An Ne Üzerinde Çalışıyoruz" bölümünü güncelle
# Biten UC'leri tabloya işle

# Deploy
[deploy komutu]
```

---

## Commit Kuralları

```bash
# Her UC veya ekran bitince — anlamlı commit
git commit -m "feat(auth): implement admin-login use case"
git commit -m "feat(auth): add login view with SSO flow"
git commit -m "feat(users): implement list-users with pagination"
git commit -m "test(auth): add system tests for admin-login"
git commit -m "docs(users): write UC doc for create-user"

# Branch stratejisi
main          → production
dev           → aktif geliştirme
feat/{konu}   → büyük feature
```

---

## Hata Ayıklama — Claude'a Ver

```
## Hata
[tam hata mesajı]

## Katman
[domain / usecase / infra / ctrl / vue-view / vue-api]

## Dosya
[tam dosya yolu]

## Ne yapmaya çalışıyorum
[açıklama]

## İlgili kod
[hata veren kısım]
```

---

## Deploy Checklist

**Backend:**
- [ ] `make fmt && make lint` temiz
- [ ] `make test && make test-system` geçiyor
- [ ] Production config doğru
- [ ] Migration'lar production'da çalıştırıldı

**Frontend:**
- [ ] `npm run type-check` temiz
- [ ] `npm run lint` temiz
- [ ] `npm run build` başarılı
- [ ] `VITE_API_BASE_URL` production değeri set

**Ürün:**
- [ ] Happy path manuel test edildi
- [ ] Auth flow çalışıyor
- [ ] Boş state'ler görünüyor
- [ ] Error state'ler görünüyor
