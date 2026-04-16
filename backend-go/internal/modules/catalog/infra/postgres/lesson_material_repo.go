package postgres

import (
	"context"
	"time"

	"go-enterprise-blueprint/internal/modules/catalog/domain/lessonmaterial"

	"github.com/code19m/errx"
	"github.com/uptrace/bun"
)

type lessonMaterialRepo struct{ db bun.IDB }

func NewLessonMaterialRepo(db bun.IDB) lessonmaterial.Repo { return &lessonMaterialRepo{db: db} }

func (r *lessonMaterialRepo) ReplaceByLesson(ctx context.Context, organizationID string, lessonID string, items []lessonmaterial.ReplaceItem) ([]lessonmaterial.LessonMaterial, error) {
	_, err := r.db.NewDelete().Model((*lessonmaterial.LessonMaterial)(nil)).Where("lesson_id = ?", lessonID).Exec(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	result := make([]lessonmaterial.LessonMaterial, 0, len(items))
	if len(items) == 0 {
		return result, nil
	}
	now := time.Now()
	for idx, input := range items {
		order := input.OrderNumber
		if order == 0 {
			order = idx + 1
		}
		materialType := input.MaterialType
		if materialType == "" {
			materialType = lessonmaterial.TypePDF
		}
		sourceType := input.SourceType
		if sourceType == "" {
			sourceType = lessonmaterial.SourceURL
		}
		result = append(result, lessonmaterial.LessonMaterial{OrganizationID: organizationID, LessonID: lessonID, Title: input.Title, MaterialType: materialType, SourceType: sourceType, URL: input.URL, FileID: input.FileID, OrderNumber: order, CreatedAt: now, UpdatedAt: now})
	}
	_, err = r.db.NewInsert().Model(&result).Returning("*").Exec(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return result, nil
}

func (r *lessonMaterialRepo) ListByLesson(ctx context.Context, lessonID string) ([]lessonmaterial.LessonMaterial, error) {
	items := make([]lessonmaterial.LessonMaterial, 0)
	err := r.db.NewSelect().Model(&items).Where("lesson_id = ?", lessonID).Order("order_number ASC").Scan(ctx)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	return items, nil
}
