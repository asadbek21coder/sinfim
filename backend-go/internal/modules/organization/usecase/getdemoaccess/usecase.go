package getdemoaccess

import (
	"context"
	"time"

	"go-enterprise-blueprint/internal/modules/organization/domain/membership"

	"github.com/code19m/errx"
	"github.com/google/uuid"
	"github.com/rise-and-shine/pkg/hasher"
	"github.com/rise-and-shine/pkg/ucdef"
	"github.com/uptrace/bun"
)

const (
	demoOwnerPhone   = "+998900000777"
	demoStudentPhone = "+998900000778"
	demoMentorPhone  = "+998900000779"
	demoPassword     = "DemoPass123"
	demoSlug         = "demo-school"
	demoCourseSlug   = "russian-a1-demo"
	demoOwnerName    = "Demo Owner"
	demoStudentName  = "Demo Student"
	demoMentorName   = "Demo Mentor"
	demoSchoolName   = "Sinfim Demo School"
)

type Request struct{}

type DemoUserDTO struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}

type Response struct {
	OrganizationID string      `json:"organization_id"`
	SchoolSlug     string      `json:"school_slug"`
	CourseID       string      `json:"course_id"`
	ClassID        string      `json:"class_id"`
	LessonID       string      `json:"lesson_id"`
	PublicURL      string      `json:"public_url"`
	OwnerURL       string      `json:"owner_url"`
	StudentURL     string      `json:"student_url"`
	Owner          DemoUserDTO `json:"owner"`
	Mentor         DemoUserDTO `json:"mentor"`
	Student        DemoUserDTO `json:"student"`
	Message        string      `json:"message"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(db *bun.DB, hashingCost int) UseCase { return &usecase{db: db, hashingCost: hashingCost} }

type usecase struct {
	db          *bun.DB
	hashingCost int
}

func (uc *usecase) OperationID() string { return "get-demo-access" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	passwordHash, err := hasher.Hash(demoPassword, hasher.WithCost(uc.hashingCost))
	if err != nil {
		return nil, errx.Wrap(err)
	}
	now := time.Now()
	var orgID, ownerID, studentID, mentorID, courseID, lessonID, classID string
	err = uc.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		orgID = uuid.NewString()
		if err := tx.NewRaw(`
INSERT INTO organization.organizations (id, name, slug, description, category, contact_phone, telegram_url, public_status, is_demo, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, 'public', true, ?, ?)
ON CONFLICT (slug) DO UPDATE SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  category = EXCLUDED.category,
  contact_phone = EXCLUDED.contact_phone,
  telegram_url = EXCLUDED.telegram_url,
  public_status = 'public',
  is_demo = true,
  updated_at = EXCLUDED.updated_at
RETURNING id`, orgID, demoSchoolName, demoSlug, "A safe demo workspace for exploring Sinfim.uz owner, mentor and student flows.", "Language school", "+998 90 000 07 77", "https://t.me/sinfim_demo", now, now).Scan(ctx, &orgID); err != nil {
			return err
		}
		ownerID = uuid.NewString()
		if err := upsertUser(ctx, tx, ownerID, demoOwnerPhone, demoOwnerName, passwordHash, false, &ownerID); err != nil {
			return err
		}
		studentID = uuid.NewString()
		if err := upsertUser(ctx, tx, studentID, demoStudentPhone, demoStudentName, passwordHash, false, &studentID); err != nil {
			return err
		}
		mentorID = uuid.NewString()
		if err := upsertUser(ctx, tx, mentorID, demoMentorPhone, demoMentorName, passwordHash, false, &mentorID); err != nil {
			return err
		}
		if err := assignRole(ctx, tx, ownerID, membership.RoleOwner); err != nil {
			return err
		}
		if err := assignRole(ctx, tx, studentID, membership.RoleStudent); err != nil {
			return err
		}
		if err := assignRole(ctx, tx, mentorID, membership.RoleMentor); err != nil {
			return err
		}
		if err := assignMembership(ctx, tx, ownerID, orgID, membership.RoleOwner, now); err != nil {
			return err
		}
		if err := assignMembership(ctx, tx, studentID, orgID, membership.RoleStudent, now); err != nil {
			return err
		}
		if err := assignMembership(ctx, tx, mentorID, orgID, membership.RoleMentor, now); err != nil {
			return err
		}
		courseID = uuid.NewString()
		if err := tx.NewRaw(`
INSERT INTO catalog.courses (id, organization_id, title, slug, description, category, level, status, public_status, created_at, updated_at)
VALUES (?, ?, 'Russian A1 Demo', ?, 'A compact demo course with lesson, materials and homework.', 'Language', 'Beginner', 'active', 'public', ?, ?)
ON CONFLICT (organization_id, slug) DO UPDATE SET title = EXCLUDED.title, description = EXCLUDED.description, status = 'active', public_status = 'public', updated_at = EXCLUDED.updated_at
RETURNING id`, courseID, orgID, demoCourseSlug, now, now).Scan(ctx, &courseID); err != nil {
			return err
		}
		lessonID = uuid.NewString()
		if err := tx.NewRaw(`
INSERT INTO catalog.lessons (id, organization_id, course_id, title, description, order_number, publish_day, status, created_at, updated_at)
VALUES (?, ?, ?, 'Alphabet and first words', 'Watch the demo lesson and submit a short answer.', 1, 1, 'published', ?, ?)
ON CONFLICT (course_id, order_number) DO UPDATE SET title = EXCLUDED.title, description = EXCLUDED.description, publish_day = 1, status = 'published', updated_at = EXCLUDED.updated_at
RETURNING id`, lessonID, orgID, courseID, now, now).Scan(ctx, &lessonID); err != nil {
			return err
		}
		if _, err := tx.NewRaw(`
INSERT INTO catalog.lesson_videos (organization_id, lesson_id, provider, stream_ref, embed_url, duration_seconds, created_at, updated_at)
VALUES (?, ?, 'external', 'demo-russian-a1-video', 'https://www.youtube.com/watch?v=dQw4w9WgXcQ', 420, ?, ?)
ON CONFLICT (lesson_id) DO UPDATE SET provider = EXCLUDED.provider, stream_ref = EXCLUDED.stream_ref, embed_url = EXCLUDED.embed_url, duration_seconds = EXCLUDED.duration_seconds, updated_at = EXCLUDED.updated_at`, orgID, lessonID, now, now).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewRaw(`DELETE FROM catalog.lesson_materials WHERE lesson_id = ?`, lessonID).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewRaw(`
INSERT INTO catalog.lesson_materials (organization_id, lesson_id, title, material_type, source_type, url, order_number, created_at, updated_at)
VALUES (?, ?, 'Demo workbook PDF', 'pdf', 'url', 'https://example.com/demo-workbook.pdf', 1, ?, ?)
`, orgID, lessonID, now, now).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewRaw(`
INSERT INTO homework.definitions (organization_id, course_id, lesson_id, title, instructions, submission_type, status, max_score, allow_resubmission, created_at, updated_at)
VALUES (?, ?, ?, 'Write three Russian words', 'Type three words you remember from the lesson.', 'text', 'published', 10, true, ?, ?)
ON CONFLICT (lesson_id) DO UPDATE SET title = EXCLUDED.title, instructions = EXCLUDED.instructions, submission_type = 'text', status = 'published', max_score = 10, updated_at = EXCLUDED.updated_at`, orgID, courseID, lessonID, now, now).Exec(ctx); err != nil {
			return err
		}
		if err := tx.NewSelect().TableExpr("classroom.classes").ColumnExpr("id").Where("organization_id = ?", orgID).Where("course_id = ?", courseID).Where("name = ?", "Demo Group A").Limit(1).Scan(ctx, &classID); err != nil {
			classID = uuid.NewString()
			if err := tx.NewRaw(`
INSERT INTO classroom.classes (id, organization_id, course_id, name, start_date, lesson_cadence, status, created_at, updated_at)
VALUES (?, ?, ?, 'Demo Group A', CURRENT_DATE, 'daily', 'active', ?, ?)
RETURNING id`, classID, orgID, courseID, now, now).Scan(ctx, &classID); err != nil {
				return err
			}
		}
		if _, err := tx.NewRaw(`
INSERT INTO classroom.class_mentors (organization_id, class_id, mentor_user_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
ON CONFLICT (class_id, mentor_user_id) DO UPDATE SET updated_at = EXCLUDED.updated_at`, orgID, classID, mentorID, now, now).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewRaw(`
INSERT INTO classroom.enrollments (organization_id, class_id, student_user_id, status, enrolled_at, created_at, updated_at)
VALUES (?, ?, ?, 'active', ?, ?, ?)
ON CONFLICT (class_id, student_user_id) DO UPDATE SET status = 'active', updated_at = EXCLUDED.updated_at`, orgID, classID, studentID, now, now, now).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewRaw(`
INSERT INTO classroom.access_grants (organization_id, class_id, student_user_id, access_status, payment_status, note, granted_by, granted_at, created_at, updated_at)
VALUES (?, ?, ?, 'active', 'confirmed', 'Demo access', ?, ?, ?, ?)
ON CONFLICT (class_id, student_user_id) DO UPDATE SET access_status = 'active', payment_status = 'confirmed', note = 'Demo access', granted_by = EXCLUDED.granted_by, granted_at = EXCLUDED.granted_at, updated_at = EXCLUDED.updated_at`, orgID, classID, studentID, ownerID, now, now, now).Exec(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return &Response{OrganizationID: orgID, SchoolSlug: demoSlug, CourseID: courseID, ClassID: classID, LessonID: lessonID, PublicURL: "/" + demoSlug, OwnerURL: "/app/dashboard", StudentURL: "/learn/lessons/" + lessonID + "?class_id=" + classID, Owner: DemoUserDTO{PhoneNumber: demoOwnerPhone, Password: demoPassword, Role: membership.RoleOwner}, Mentor: DemoUserDTO{PhoneNumber: demoMentorPhone, Password: demoPassword, Role: membership.RoleMentor}, Student: DemoUserDTO{PhoneNumber: demoStudentPhone, Password: demoPassword, Role: membership.RoleStudent}, Message: "Demo school is ready."}, nil
}

func upsertUser(ctx context.Context, tx bun.Tx, id string, phone string, fullName string, passwordHash string, mustChange bool, outID *string) error {
	return tx.NewRaw(`
INSERT INTO auth.users (id, username, phone_number, full_name, password_hash, is_active, must_change_password, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, true, ?, now(), now())
ON CONFLICT (username) DO UPDATE SET phone_number = EXCLUDED.phone_number, full_name = EXCLUDED.full_name, password_hash = EXCLUDED.password_hash, is_active = true, must_change_password = EXCLUDED.must_change_password, updated_at = now()
RETURNING id`, id, phone, phone, fullName, passwordHash, mustChange).Scan(ctx, outID)
}

func assignRole(ctx context.Context, tx bun.Tx, userID string, roleName string) error {
	_, err := tx.NewRaw(`INSERT INTO auth.user_roles (user_id, role_id, created_at, updated_at)
SELECT ?, r.id, now(), now() FROM auth.roles r
WHERE r.name = ?
  AND NOT EXISTS (
    SELECT 1 FROM auth.user_roles ur WHERE ur.user_id = ? AND ur.role_id = r.id
  )`, userID, roleName, userID).Exec(ctx)
	return err
}

func assignMembership(ctx context.Context, tx bun.Tx, userID string, organizationID string, role string, now time.Time) error {
	_, err := tx.NewRaw(`INSERT INTO auth.user_memberships (user_id, organization_id, role, is_active, created_at, updated_at)
VALUES (?, ?, ?, true, ?, ?)
ON CONFLICT (user_id, organization_id, role) DO UPDATE SET is_active = true, updated_at = EXCLUDED.updated_at`, userID, organizationID, role, now, now).Exec(ctx)
	return err
}
