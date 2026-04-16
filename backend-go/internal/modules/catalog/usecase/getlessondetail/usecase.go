package getlessondetail

import (
	"context"

	"go-enterprise-blueprint/internal/modules/catalog/domain"
	"go-enterprise-blueprint/internal/modules/catalog/domain/lesson"
	"go-enterprise-blueprint/internal/modules/catalog/domain/lessonmaterial"
	"go-enterprise-blueprint/internal/modules/catalog/domain/lessonvideo"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/shared"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	ID string `query:"id" validate:"required,uuid"`
}

type Response struct {
	Item      lesson.Lesson                   `json:"item"`
	Video     *lessonvideo.LessonVideo        `json:"video"`
	Materials []lessonmaterial.LessonMaterial `json:"materials"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(dc *domain.Container) UseCase { return &usecase{dc: dc} }

type usecase struct{ dc *domain.Container }

func (uc *usecase) OperationID() string { return "get-lesson-detail" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	item, err := uc.dc.LessonRepo().Get(ctx, lesson.Filter{ID: &in.ID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if accessErr := shared.EnsureCourseReadAccess(ctx, uc.dc, item.OrganizationID); accessErr != nil {
		return nil, errx.Wrap(accessErr)
	}
	video, videoErr := uc.dc.LessonVideoRepo().GetByLesson(ctx, item.ID)
	if videoErr != nil {
		return nil, errx.Wrap(videoErr)
	}
	materials, matErr := uc.dc.LessonMaterialRepo().ListByLesson(ctx, item.ID)
	if matErr != nil {
		return nil, errx.Wrap(matErr)
	}
	return &Response{Item: *item, Video: video, Materials: materials}, nil
}
