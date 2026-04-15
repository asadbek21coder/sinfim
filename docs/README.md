# Sinfim.uz Docs

Bu klasor Sinfim.uz urun, tasarim, teknik karar ve AI-agent calisma hafizasinin ana merkezidir.

## Klasor Haritasi

```text
docs/
  product/      Urun fikri, UX, teknik kararlar, implementation plan
  specs/        Modul/use-case bazli detay sozlesmeleri
  design/       Stitch AI ciktisi ve gorsel referanslar
  prompts/      UI/AI promptlari
  ai-context/   Codex/Claude ortak session, handoff ve worklog dosyalari
  templates/    Blueprint/project-builder'dan kalan yeniden kullanilabilir sablonlar
```

## En Once Okunacak Dosyalar

- `docs/product/brand-constants.md`
- `docs/product/startup-idea.md`
- `docs/product/ux-doc.md`
- `docs/product/tech-doc.md`
- `docs/product/implementation-plan.md`
- `docs/product/local-development.md`
- `docs/ai-context/SESSION.md`
- `docs/ai-context/HANDOFF.md`

## Tasarim Referansi

Stitch AI ciktisi `docs/design/sinfim-design/` altindadir.

Bu ciktidaki HTML dosyalari production koda birebir kopyalanmayacak. Screenshot, layout, component hiyerarsisi, spacing ve genel tasarim dili icin referans alinacak.

Canonical marka karari:

- Platform adi: `Sinfim.uz`
- Domain: `sinfim.uz`
- Organization URL: `sinfim.uz/{school-slug}`

## Implementation Plani

Ana uygulama sirasi `docs/product/implementation-plan.md` icindedir. Her step frontend ve backend'i birlikte ele alir ve manuel test edilebilir bir urun parcasi olarak tamamlanir.
