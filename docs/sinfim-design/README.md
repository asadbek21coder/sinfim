# Sinfim.uz Design Output Notes

Bu klasor Stitch AI tarafindan uretilen raw UI ciktisini icerir. Frontend implementation icin kaynak olarak kullanilabilir, fakat birebir production UI kabul edilmez.

## Canonical Brand

- Platform adi: `Sinfim.uz`
- Domain: `sinfim.uz`
- Organization URL: `sinfim.uz/{school-slug}`
- Course URL: `sinfim.uz/{school-slug}/courses/{course-slug}`

Detayli marka sabitleri: `../projects/lms/docs/brand-constants.md`.

## Important Warning

Raw Stitch ciktisinda su placeholder metinler kalmistir ve frontend'e tasinmayacak:

- `Architectural Academic`
- `Scholarly Canvas`
- `Sterling Institute of Design`
- `platform.uz`
- `sterling-edu.lms.com`
- SMS ile temporary code metinleri
- Physics/design/architecture demo ders icerikleri

Canonical design review: `../projects/lms/docs/ui-design-review.md`.

## Implementation Rule

`code.html` dosyalari birebir kopyalanmayacak. Screenshot ve HTML yapilari sadece layout, spacing, component hiyerarsisi ve interaction fikri icin okunacak. Vue 3 + TypeScript tarafinda componentler proje standardina gore yeniden kurulacak.
