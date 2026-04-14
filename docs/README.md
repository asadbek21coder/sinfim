# Project Builder

AI agent'larla solo/iki kişilik geliştirme için eksiksiz sistem.
Startup fikrinden veya outsource projeden lansmana.

**Başlamak için:** `new-project-template/kickoff.md`'yi yeni projeye kopyala ve üstten aşağı işaretle.

---

## Dosya Haritası

```
project-builder/
│
├── roadmap.md                          ← Genel akış, neden bu sıra
│
├── templates/                          ← Doldurulacak şablonlar
│   ├── startup-idea.md                 ← Fikir → problem, persona, varsayım, scope
│   ├── ux-doc.md                       ← Ekranlar, akışlar, state'ler
│   ├── tech-doc.md                     ← Modüller, SQL, API contract, UC listesi
│   ├── ui-prompts.md                   ← AI'a verilecek UI prompt şablonları
│   ├── dev-kickoff.md                  ← Blueprint kurulumu, agent pipeline, sprint yapısı
│   └── outsource-discovery.md          ← Client'tan ne alınır, hangi sorular sorulur
│
└── new-project-template/               ← Her yeni projede kopyalanacak dosyalar
    ├── kickoff.md                      ← BURADAN BAŞLA: adım adım görev listesi (👤/🤖)
    ├── CLAUDE.md                       ← AI context — her session'da okunur
    ├── partner.md                      ← Ortak için: sprint durumu, bekleyen kararlar
```

---

## Yeni Proje Açarken

```bash
# new-project-template içindeki dosyaları projeye kopyala
cp blueprints/project-builder/new-project-template/kickoff.md  ./kickoff.md
cp blueprints/project-builder/new-project-template/CLAUDE.md   ./CLAUDE.md
cp blueprints/project-builder/new-project-template/partner.md  ./partner.md
```

Sonra `kickoff.md`'yi aç ve üstten işaretlemeye başla.
