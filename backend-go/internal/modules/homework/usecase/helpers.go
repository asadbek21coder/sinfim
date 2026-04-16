package usecase

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"go-enterprise-blueprint/internal/modules/organization/domain/membership"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

func trimPtr(value *string) *string {
	if value == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}

func loadLessonContext(ctx context.Context, db *bun.DB, lessonID string) (lessonContext, error) {
	var item lessonContext
	err := db.NewSelect().TableExpr("catalog.lessons AS l").
		ColumnExpr("l.id AS lesson_id, l.organization_id, l.course_id").
		Where("l.id = ?", lessonID).
		Limit(1).
		Scan(ctx, &item)
	if errors.Is(err, sql.ErrNoRows) {
		return lessonContext{}, errx.New("lesson not found", errx.WithType(errx.T_NotFound), errx.WithCode(CodeLessonNotFound))
	}
	return item, err
}

func ensureTeacherAccess(ctx context.Context, db *bun.DB, organizationID string) error {
	userCtx := auth.MustUserContext(ctx)
	if auth.HasPermission(userCtx, auth.PermissionUserManage) {
		return nil
	}
	var ok bool
	err := db.NewSelect().TableExpr("organization.memberships AS m").
		ColumnExpr("COUNT(*) > 0").
		Where("m.organization_id = ?", organizationID).
		Where("m.user_id = ?", userCtx.UserID).
		Where("m.role IN (?)", bun.In([]string{membership.RoleOwner, membership.RoleTeacher})).
		Scan(ctx, &ok)
	if err != nil {
		return errx.Wrap(err)
	}
	if ok {
		return nil
	}
	return errx.New("owner or teacher membership required", errx.WithType(errx.T_Forbidden), errx.WithCode(auth.CodeForbidden))
}

func ensureReviewAccess(ctx context.Context, db *bun.DB, organizationID string, classID string) error {
	userCtx := auth.MustUserContext(ctx)
	if auth.HasPermission(userCtx, auth.PermissionUserManage) {
		return nil
	}
	var memberOK bool
	err := db.NewSelect().TableExpr("organization.memberships AS m").
		ColumnExpr("COUNT(*) > 0").
		Where("m.organization_id = ?", organizationID).
		Where("m.user_id = ?", userCtx.UserID).
		Where("m.role IN (?)", bun.In([]string{membership.RoleOwner, membership.RoleTeacher})).
		Scan(ctx, &memberOK)
	if err != nil {
		return errx.Wrap(err)
	}
	if memberOK {
		return nil
	}
	var mentorOK bool
	err = db.NewSelect().TableExpr("classroom.class_mentors AS cm").
		ColumnExpr("COUNT(*) > 0").
		Where("cm.class_id = ?", classID).
		Where("cm.mentor_user_id = ?", userCtx.UserID).
		Scan(ctx, &mentorOK)
	if err != nil {
		return errx.Wrap(err)
	}
	if mentorOK {
		return nil
	}
	return errx.New("homework review access required", errx.WithType(errx.T_Forbidden), errx.WithCode(auth.CodeForbidden))
}

func applyReviewAccessFilter(ctx context.Context, query *bun.SelectQuery, db *bun.DB) *bun.SelectQuery {
	userCtx := auth.MustUserContext(ctx)
	if auth.HasPermission(userCtx, auth.PermissionUserManage) {
		return query
	}
	return query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("EXISTS (SELECT 1 FROM organization.memberships m WHERE m.organization_id = hs.organization_id AND m.user_id = ? AND m.role IN (?))", userCtx.UserID, bun.In([]string{membership.RoleOwner, membership.RoleTeacher})).
			WhereOr("EXISTS (SELECT 1 FROM classroom.class_mentors cm WHERE cm.class_id = hs.class_id AND cm.mentor_user_id = ?)", userCtx.UserID)
	})
}

func ensureStudentLessonAccess(ctx context.Context, db *bun.DB, classID string, lessonID string) (lessonContext, error) {
	userID := auth.MustUserContext(ctx).UserID
	var item lessonContext
	err := db.NewSelect().TableExpr("classroom.enrollments AS e").
		ColumnExpr("l.id AS lesson_id, l.organization_id, l.course_id").
		Join("JOIN classroom.classes AS cl ON cl.id = e.class_id").
		Join("JOIN catalog.lessons AS l ON l.course_id = cl.course_id").
		Join("LEFT JOIN classroom.access_grants AS ag ON ag.class_id = e.class_id AND ag.student_user_id = e.student_user_id").
		Where("e.student_user_id = ?", userID).
		Where("e.class_id = ?", classID).
		Where("e.status = 'active'").
		Where("cl.status = 'active'").
		Where("l.id = ?", lessonID).
		Where("l.status = 'published'").
		Where("COALESCE(ag.access_status, 'pending') = 'active'").
		Limit(1).
		Scan(ctx, &item)
	if errors.Is(err, sql.ErrNoRows) {
		return lessonContext{}, errx.New("student lesson access required", errx.WithType(errx.T_Forbidden), errx.WithCode(CodeAccessDenied))
	}
	return item, err
}

func getDefinition(ctx context.Context, db *bun.DB, definitionID string) (*DefinitionDTO, error) {
	item := new(DefinitionDTO)
	err := db.NewSelect().TableExpr("homework.definitions AS hd").
		ColumnExpr("hd.id, hd.organization_id, hd.course_id, hd.lesson_id, hd.title, hd.instructions, hd.submission_type, hd.status, hd.max_score, hd.due_days_after_publish, hd.allow_resubmission, hd.created_at, hd.updated_at").
		Where("hd.id = ?", definitionID).
		Limit(1).
		Scan(ctx, item)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errx.New("homework definition not found", errx.WithType(errx.T_NotFound), errx.WithCode(CodeDefinitionNotFound))
	}
	return item, err
}

func loadQuiz(ctx context.Context, db *bun.DB, definitionID string, includeCorrect bool) ([]QuizQuestionDTO, error) {
	questions := make([]QuizQuestionDTO, 0)
	err := db.NewSelect().TableExpr("homework.quiz_questions AS qq").
		ColumnExpr("qq.id, qq.prompt, qq.order_number, qq.points").
		Where("qq.definition_id = ?", definitionID).
		Order("qq.order_number ASC").
		Scan(ctx, &questions)
	if err != nil {
		return nil, err
	}
	for index := range questions {
		options := make([]QuizOptionDTO, 0)
		query := db.NewSelect().TableExpr("homework.quiz_options AS qo").
			ColumnExpr("qo.id, qo.label, qo.order_number").
			Where("qo.question_id = ?", questions[index].ID).
			Order("qo.order_number ASC")
		if includeCorrect {
			query = query.ColumnExpr("qo.is_correct")
		} else {
			query = query.ColumnExpr("false AS is_correct")
		}
		if err := query.Scan(ctx, &options); err != nil {
			return nil, err
		}
		questions[index].Options = options
	}
	return questions, nil
}
