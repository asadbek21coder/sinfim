# Local Development

Bu dokuman Sinfim.uz local calisma standardini sabitler.

## Tek Komut Local Stack

Root klasorden:

```bash
make local-up
```

Bu komut asagidaki servisleri Docker ile kaldirir:

- Frontend: `http://localhost:5173`
- Backend: `http://localhost:9876`
- Backend health: `http://localhost:9876/health`
- Postgres: `localhost:5432`
- Redis: `localhost:6379`
- Kafka: `localhost:9092`
- AKHQ: `http://localhost:8081`
- Jaeger: `http://localhost:16686`
- MinIO API: `http://localhost:9100`
- MinIO Console: `http://localhost:9101`

Stack'i durdurmak:

```bash
make local-down
```

Volume'lari da silerek temiz kapatmak:

```bash
make local-down-volumes
```

## Frontend Local Proxy

Frontend Vite dev server backend'e proxy ile ulasir.

```text
/api/*   -> VITE_API_PROXY_TARGET
/health  -> VITE_API_PROXY_TARGET
```

Default deger:

```text
VITE_API_BASE_URL=/api/v1
VITE_API_PROXY_TARGET=http://localhost:9876
```

Not: macOS/Docker Desktop uzerinde `localhost` IPv6'ya gittiginde curl connection reset verebilir. Boyle bir durumda manuel health check icin `http://127.0.0.1:9876/health` kullan.

Docker compose icinde frontend `VITE_API_PROXY_TARGET=http://backend:9876` kullanir.

Backend container `ENVIRONMENT=local` ile calisir. Docker image build sirasinda `backend-go/config/docker-local.yaml`, container icinde `config/local.yaml` olarak kopyalanir; boylece config loader'in kabul ettigi environment isimleri korunur.

## Step 0 Route/Layout Haritasi

Public website layout:

- `/`
- `/enter`
- `/apply-school`
- `/demo`
- `/:schoolSlug`
- `/:schoolSlug/courses/:courseSlug`

Auth layout:

- `/auth/login`
- `/login` -> `/auth/login`

School app layout:

- `/admin/school-requests`
- `/admin/organizations/new`
- `/app/dashboard`
- `/app/courses`
- `/app/courses/:courseId`
- `/app/classes`
- `/app/classes/:classId`
- `/app/mentors`
- `/app/students`
- `/app/leads`
- `/app/lessons/:lessonId/edit`
- `/app/homework/review`
- `/app/settings/organization`

Student learning layout:

- `/learn/dashboard`
- `/learn/lessons/:lessonId`

Auth guard Step 2'de gercek hale getirilecek. Step 0/1'de app/student/admin preview route'lari acik tutulur.

## Step 1 School Request Manual Test

1. `make local-up` calistir.
2. `http://localhost:5173/apply-school` ac.
3. Ism, telefon, maktab/brend nomi ve istege bagli izoh alanlarini doldur.
4. Formu gonderince basari mesaji gorulur.
5. `http://localhost:5173/admin/school-requests` ac.
6. Yeni ariza listede gorulur.
7. Status select ile `Contacted`, `Approved` veya `Rejected` yapilir.

Direkt API smoke test:

```bash
curl -sS -X POST http://127.0.0.1:9876/api/v1/organization/create-school-request \
  -H 'Content-Type: application/json' \
  -d '{"full_name":"Ali Valiyev","phone_number":"+998 90 123 45 67","school_name":"Sinfim Demo School","category":"language","student_count":120,"note":"Telegram orqali kurs sotamiz"}'

curl -sS 'http://127.0.0.1:9876/api/v1/organization/list-school-requests?limit=5'
```

## Step 2 Auth Manual Test

Local seed platform admin:

```text
phone: +998900000001
password: admin12345
```

Browser testi:

1. `http://localhost:5173/admin/school-requests` ac. Login sayfasina yonlenir.
2. `+998900000001` / `admin12345` ile giris yap.
3. Admin request listesi acilir.
4. Logout UI henuz yoksa browser local storage temizlenerek cikis simule edilebilir.

Direkt API smoke test:

```bash
LOGIN=$(curl -sS -X POST http://127.0.0.1:9876/api/v1/auth/admin-login \
  -H 'Content-Type: application/json' \
  -d '{"phone_number":"+998900000001","password":"admin12345"}')

TOKEN=$(printf '%s' "$LOGIN" | node -e "let s='';process.stdin.on('data',d=>s+=d);process.stdin.on('end',()=>console.log(JSON.parse(s).accessToken))")

curl -sS http://127.0.0.1:9876/api/v1/auth/me \
  -H "Authorization: Bearer $TOKEN"

curl -sS -o /tmp/noauth -w '%{http_code}' \
  'http://127.0.0.1:9876/api/v1/organization/list-school-requests?limit=1'

curl -sS -o /tmp/authorg -w '%{http_code}' \
  'http://127.0.0.1:9876/api/v1/organization/list-school-requests?limit=1' \
  -H "Authorization: Bearer $TOKEN"
```

## Step 3 Organization Create Manual Test

Browser testi:

1. Platform admin ile login ol: `+998900000001` / `admin12345`.
2. `http://localhost:5173/admin/organizations/new` ac.
3. School workspace ve owner access alanlarini doldur.
4. Create basildiginda organization ve owner bilgisi success kutusunda gorulur.
5. Cikis simule etmek icin local storage temizlenebilir.
6. Owner telefonu ve temporary password ile login ol.
7. Owner `mustChangePassword` nedeniyle `/auth/change-password` ekranina yonlenir.
8. Password degistirdikten sonra `/app/settings/organization` acilir.
9. Organization settings alanlari guncellenir ve success mesaji gorulur.

Direkt API smoke test:

```bash
ADMIN_LOGIN=$(curl -sS -X POST http://127.0.0.1:9876/api/v1/auth/admin-login \
  -H 'Content-Type: application/json' \
  -d '{"phone_number":"+998900000001","password":"admin12345"}')

ADMIN_TOKEN=$(printf '%s' "$ADMIN_LOGIN" | node -e "let s='';process.stdin.on('data',d=>s+=d);process.stdin.on('end',()=>console.log(JSON.parse(s).accessToken))")

curl -sS -X POST http://127.0.0.1:9876/api/v1/organization/create-organization \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"name":"Demo School","slug":"demo-school-local","owner":{"full_name":"Ali Owner","phone_number":"+998901234001","temporary_password":"TempPass123"}}'

curl -sS -X POST http://127.0.0.1:9876/api/v1/auth/admin-login \
  -H 'Content-Type: application/json' \
  -d '{"phone_number":"+998901234001","password":"TempPass123"}'

OWNER_LOGIN=$(curl -sS -X POST http://127.0.0.1:9876/api/v1/auth/admin-login \
  -H 'Content-Type: application/json' \
  -d '{"phone_number":"+998901234001","password":"TempPass123"}')

OWNER_TOKEN=$(printf '%s' "$OWNER_LOGIN" | node -e "let s='';process.stdin.on('data',d=>s+=d);process.stdin.on('end',()=>console.log(JSON.parse(s).accessToken))")

curl -sS http://127.0.0.1:9876/api/v1/organization/list-my-workspaces \
  -H "Authorization: Bearer $OWNER_TOKEN"
```

## Step 4 Public School Page and Lead Manual Test

Browser testi:

1. Platform admin ile login ol: `+998900000001` / `admin12345`.
2. `http://localhost:5173/admin/organizations/new` uzerinden bir organization ve owner olustur.
3. Owner veya platform admin ile `http://localhost:5173/app/settings/organization` ekraninda `Public status` alanini `public` yap.
4. `http://localhost:5173/{schoolSlug}` public school page acilir.
5. Public formdan isim, telefon ve opsiyonel not gonderilir.
6. Owner ile login ol ve `http://localhost:5173/app/leads` ekraninda lead'i gor.
7. Lead status alanini `contacted`, `converted` veya `archived` olarak degistir.

API smoke akisi:

```bash
BASE=http://127.0.0.1:9876/api/v1

ADMIN_LOGIN=$(curl -sS -X POST "$BASE/auth/admin-login" \
  -H 'Content-Type: application/json' \
  -d '{"phone_number":"+998900000001","password":"admin12345"}')

ADMIN_TOKEN=$(printf '%s' "$ADMIN_LOGIN" | node -e "let s='';process.stdin.on('data',d=>s+=d);process.stdin.on('end',()=>console.log(JSON.parse(s).accessToken))")

curl -sS "$BASE/organization/get-public-school-page?slug={schoolSlug}"

curl -sS -X POST "$BASE/organization/create-lead" \
  -H 'Content-Type: application/json' \
  -d '{"organization_id":"{organizationId}","full_name":"Lead Name","phone_number":"+998901112233","note":"Public page lead"}'

curl -sS "$BASE/organization/list-leads?organization_id={organizationId}" \
  -H "Authorization: Bearer $ADMIN_TOKEN"

curl -sS -X POST "$BASE/organization/update-lead-status" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"id":"{leadId}","status":"contacted"}'
```

## Step 5 Course Management Manual Test

Browser testi:

1. Platform admin ile organization olustur ve owner temporary password al.
2. Organization settings icinde school public status alanini `public` yap.
3. Owner ile login ol ve `http://localhost:5173/app/courses` ac.
4. Title, slug, description, category ve level girerek course olustur.
5. Course detail ekraninda status alanini `active`, public status alanini `public` yap.
6. Public link `http://localhost:5173/{schoolSlug}/courses/{courseSlug}` acilir.
7. Public course page uzerinden lead formu gonderilir.
8. `http://localhost:5173/app/leads` ekraninda course notu ile yeni lead gorulur.

API endpointleri:

```bash
POST /api/v1/catalog/create-course
GET  /api/v1/catalog/list-courses?organization_id={organizationId}
GET  /api/v1/catalog/get-course-detail?id={courseId}
POST /api/v1/catalog/update-course
GET  /api/v1/catalog/get-public-course-page?school_slug={schoolSlug}&course_slug={courseSlug}
```

## Step 6 Class/Group and Access Manual Test

Browser testi:

1. Platform admin ile organization olustur ve owner temporary password al.
2. Owner ile login ol ve `http://localhost:5173/app/courses` uzerinden bir course olustur.
3. `http://localhost:5173/app/courses/{courseId}` ac.
4. Class adi ve start date girerek yeni sinif/grup olustur.
5. `http://localhost:5173/app/classes` ekraninda sinifin listelendigini kontrol et.
6. `http://localhost:5173/app/classes/{classId}` ac.
7. Student name, phone ve temporary password ile ogrenci ekle.
8. Ogrencinin access/payment durumunu `active` / `confirmed` yap.
9. Sayfayi yenileyince class detail icinde ogrenci ve statuslar korunur.

API endpointleri:

```bash
POST /api/v1/classroom/create-class
GET  /api/v1/classroom/list-classes?organization_id={organizationId}
GET  /api/v1/classroom/list-classes?course_id={courseId}
GET  /api/v1/classroom/get-class-detail?id={classId}
POST /api/v1/classroom/add-student
POST /api/v1/classroom/update-access
POST /api/v1/classroom/assign-mentor
```

## Step 7 Lesson, Video Reference and Materials Manual Test

Browser testi:

1. Backend ve frontend calisir durumda olsun: `http://localhost:9876/health`, `http://localhost:5173`.
2. Owner ile login ol ve `http://localhost:5173/app/courses` uzerinden bir course ac.
3. `http://localhost:5173/app/courses/{courseId}` ekraninda `New lesson` formundan ders olustur.
4. Otomatik olarak `http://localhost:5173/app/lessons/{lessonId}/edit` editorune gec veya lesson listesinden `Edit` bas.
5. Title, order, publish day ve status alanlarini duzenle.
6. Video reference bolumunu ac, provider `Telegram` sec ve stream/channel/message/embed metadata gir.
7. Bir veya daha fazla material ekle: title, type `PDF`, URL.
8. Save basinca success mesaji gorulur.
9. Course detail'e donunce lesson listesinde video/material sayisi gorulur.

API endpointleri:

```bash
POST /api/v1/catalog/create-lesson
GET  /api/v1/catalog/list-lessons?course_id={courseId}
GET  /api/v1/catalog/get-lesson-detail?id={lessonId}
POST /api/v1/catalog/update-lesson
```

Not: MVP'de material dosyasini direkt yuklemek yerine URL/file metadata saklanir. Filevault/MinIO upload baglantisi sonraki dosya polish adiminda eklenecek.

## Step 8 Student Learning Manual Test

Browser testi:

1. Backend ve frontend calisir durumda olsun: `http://localhost:9876/health`, `http://localhost:5173`.
2. Owner ile bir course, class, published lesson ve active access'li student olustur.
3. Student phone/password ile login ol.
4. `http://localhost:5173/learn/dashboard` ac.
5. Aktif sinif, progress ve lesson listesi gorulur; kilitli derslerde locked badge gorulur.
6. Acik bir derste `Open lesson` bas ve `http://localhost:5173/learn/lessons/{lessonId}` sayfasini kontrol et.
7. Video reference ve material linkleri gorulur.
8. `Mark completed` basinca dashboard progress degeri artar.

API endpointleri:

```bash
GET  /api/v1/learning/get-student-dashboard?class_id={classId}
GET  /api/v1/learning/get-lesson-detail?class_id={classId}&lesson_id={lessonId}
POST /api/v1/learning/mark-lesson-completed
```

Not: Lesson availability class `start_date`, class `cadence`, lesson `publish_day` ve lesson status alanlarindan hesaplanir. Access status active degilse student dashboard dersi gosterir ama icerik kilitli kalir.

## Step 9 Homework Definition and Submission Manual Test

Browser testi:

1. Backend ve frontend calisir durumda olsun: `http://localhost:9876/health`, `http://localhost:5173`.
2. Owner veya platform admin ile login ol ve `http://localhost:5173/app/lessons/{lessonId}/edit` ac.
3. `Homework` bolumunu enabled yap.
4. Text/file/audio/quiz type sec, title ve instructions gir.
5. Quiz type icin en az bir question ve correct option gir.
6. Status `Published` sec ve `Save homework` bas.
7. Active access'li student ile login ol.
8. `http://localhost:5173/learn/lessons/{lessonId}?class_id={classId}` ac.
9. Homework formunu doldur ve submit et.
10. Quiz ise score hemen gorulur; text/file/audio ise status `submitted` kalir ve Step 10 review inbox'a gider.

API endpointleri:

```bash
POST /api/v1/homework/save-definition
GET  /api/v1/homework/get-lesson-homework?lesson_id={lessonId}
GET  /api/v1/homework/get-student-homework?class_id={classId}&lesson_id={lessonId}
POST /api/v1/homework/submit-homework
```

Not: MVP'de file/photo/audio teslimlerinde gercek upload yerine URL metadata saklanir. Filevault upload polish sonraki dosya adiminda eklenecek.

## Step 10 Mentor Homework Review Manual Test

Browser testi:

1. Backend ve frontend calisir durumda olsun: `http://localhost:9876/health`, `http://localhost:5173`.
2. Student tarafindan text/file/audio homework submission olustur.
3. Owner, teacher veya class'a atanmis mentor ile login ol.
4. `http://localhost:5173/app/homework/review` ac.
5. Pending review listesinde submission gorulur.
6. Submission secilince ogrenci cevabi, ders/sinif bilgisi ve review formu acilir.
7. Score ve feedback gir, status `Reviewed` veya `Needs revision` sec, `Save review` bas.
8. Student tekrar lesson detail sayfasini actiginda score, status ve feedback gorur.

API endpointleri:

```bash
GET  /api/v1/homework/list-review-submissions?status=submitted&limit=100
GET  /api/v1/homework/get-review-submission?id={submissionId}
POST /api/v1/homework/review-submission
GET  /api/v1/homework/get-student-homework?class_id={classId}&lesson_id={lessonId}
```

Permission notu: platform admin hepsini gorebilir; owner/teacher kendi organization'indaki submissions'i gorebilir; mentor sadece `classroom.class_mentors` ile atandigi class submissions'ini gorebilir.

## Step 11 Owner Operational Dashboard Manual Test

Browser testi:

1. Backend ve frontend calisir durumda olsun: `http://localhost:9876/health`, `http://localhost:5173`.
2. Owner veya teacher ile login ol.
3. `http://localhost:5173/app/dashboard` ac.
4. Active courses/classes/students, pending homework ve pending access metriclerini kontrol et.
5. Quick action linkleri `Courses`, `Classes`, `Homework`, `Leads` ekranlarina gider.
6. Bir homework submission review edilince pending homework sayisinin azalmasini kontrol et.
7. Pending access/payment update edilince access confirmation kartlarinin degismesini kontrol et.

API endpointi:

```bash
GET /api/v1/organization/get-owner-dashboard?organization_id={organizationId}
```

Not: `organization_id` platform admin smoke/test icin kullanilabilir. Owner/teacher/mentor kendi ilk workspace'i icin parametresiz dashboard alabilir. Mentor dashboard'u organization geneli yerine sadece access yetkisi olan operasyonlara gore ileride daraltilabilir; MVP'de owner dashboard ana odaktir.

Docker Desktop kapaliysa local smoke icin Homebrew servisleriyle de calisilabilir:

```bash
brew services start postgresql@16
createuser -s postgres 2>/dev/null || true
psql -d postgres -c "ALTER USER postgres WITH PASSWORD 'postgres';"
createdb -O postgres blueprint 2>/dev/null || true

MINIO_ROOT_USER=minioadmin MINIO_ROOT_PASSWORD=minioadmin \
  minio server --address :9100 --console-address :9101 /opt/homebrew/var/minio-sinfim
```

Bu durumda backend yine `ENVIRONMENT=local go run ./cmd run` ile `backend-go/config/local.yaml` degerlerini kullanir.

## Migration ve Seed Standardi

Backend migration dosyalari `backend-go/migrations/` altinda kalacak.

Kurallar:

- Her yeni tablo/schema migration ile eklenecek.
- Migration dosyalari blueprint'in mevcut goose standardini takip edecek.
- Uygulama baslarken `app.init()` icinde migration up calismaya devam edecek.
- Seed data MVP boyunca CLI komutu veya `backend-go/scripts/seeds/` altindaki SQL dosyalari ile yonetilecek.
- Demo school seed'i `GET /api/v1/organization/get-demo-access` endpoint'i ile resetlenebilir/idempotent sekilde hazirlanir.

## Step 12 Demo School Manual Test

Browser testi:

1. Backend ve frontend calisir durumda olsun: `http://localhost:9876/health`, `http://localhost:5173`.
2. `http://localhost:5173/demo` ac.
3. Demo kartinda owner, mentor ve student hesaplari gorulur.
4. `Open public school` ile `http://localhost:5173/demo-school` acilir.
5. Owner account ile login ol; login formu phone/password alanlarini doldurur ve `/app/dashboard` redirect eder.
6. Student account ile login ol; lesson detail sayfasina gider ve video/material/homework gorulur.
7. Mentor account ile login ol; `/app/homework/review` acilir.

Demo hesaplari:

```text
Owner:   +998900000777 / DemoPass123
Student: +998900000778 / DemoPass123
Mentor:  +998900000779 / DemoPass123
```

API smoke:

```bash
curl -sS http://127.0.0.1:9876/api/v1/organization/get-demo-access | jq .
curl -sS "http://127.0.0.1:9876/api/v1/organization/get-public-school-page?slug=demo-school" | jq .organization
```

Not: MVP demo modu strict read-only degildir. Karar: demo school resetlenebilir seed olarak calisir; gercek kullanici verisi buraya girilmemeli. Global demo read-only guard MVP hardening/polish adiminda ele alinabilir.

## Step 0 Manual Test

1. `make local-up` calistir.
2. `http://localhost:5173` acilir.
3. Landing ekranindaki Backend health karti online olur.
4. `http://127.0.0.1:9876/health` direkt acildiginda `{ "status": "ok" }` doner.
5. `/enter`, `/apply-school`, `/demo`, `/app/dashboard`, `/learn/dashboard` route'lari acilir.
6. `make local-down` ile servisler kapanir.
