package updatelesson

import (
	"context"
	"strings"

	"go-enterprise-blueprint/internal/modules/catalog/domain"
	"go-enterprise-blueprint/internal/modules/catalog/domain/lesson"
	"go-enterprise-blueprint/internal/modules/catalog/domain/lessonmaterial"
	"go-enterprise-blueprint/internal/modules/catalog/domain/lessonvideo"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/shared"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type VideoInput struct {
	Enabled           bool    `json:"enabled"`
	Provider          string  `json:"provider" validate:"omitempty,oneof=telegram external"`
	StreamRef         *string `json:"stream_ref" validate:"omitempty,max=2000"`
	TelegramChannelID *string `json:"telegram_channel_id" validate:"omitempty,max=255"`
	TelegramMessageID *string `json:"telegram_message_id" validate:"omitempty,max=255"`
	EmbedURL          *string `json:"embed_url" validate:"omitempty,max=2000"`
	DurationSeconds   *int    `json:"duration_seconds" validate:"omitempty,min=0,max=86400"`
}

type MaterialInput struct {
	Title        string  `json:"title" validate:"required,min=1,max=180"`
	MaterialType string  `json:"material_type" validate:"omitempty,oneof=pdf image doc link other"`
	SourceType   string  `json:"source_type" validate:"omitempty,oneof=url filevault"`
	URL          *string `json:"url" validate:"omitempty,max=2000"`
	FileID       *string `json:"file_id" validate:"omitempty,uuid"`
	OrderNumber  int     `json:"order_number" validate:"omitempty,min=1,max=1000"`
}

type Request struct {
	ID          string          `json:"id" validate:"required,uuid"`
	Title       string          `json:"title" validate:"required,min=2,max=180"`
	Description *string         `json:"description" validate:"omitempty,max=2000"`
	OrderNumber int             `json:"order_number" validate:"required,min=1,max=1000"`
	PublishDay  int             `json:"publish_day" validate:"required,min=1,max=1000"`
	Status      string          `json:"status" validate:"required,oneof=draft published archived"`
	Video       *VideoInput     `json:"video"`
	Materials   []MaterialInput `json:"materials" validate:"omitempty,max=50,dive"`
}

type Response struct {
	Item      lesson.Lesson                   `json:"item"`
	Video     *lessonvideo.LessonVideo        `json:"video"`
	Materials []lessonmaterial.LessonMaterial `json:"materials"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(dc *domain.Container) UseCase { return &usecase{dc: dc} }

type usecase struct{ dc *domain.Container }

func (uc *usecase) OperationID() string { return "update-lesson" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	item, err := uc.dc.LessonRepo().Get(ctx, lesson.Filter{ID: &in.ID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if accessErr := shared.EnsureCourseWriteAccess(ctx, uc.dc, item.OrganizationID); accessErr != nil {
		return nil, errx.Wrap(accessErr)
	}
	item.Title = strings.TrimSpace(in.Title)
	item.Description = shared.TrimPtr(in.Description)
	item.OrderNumber = in.OrderNumber
	item.PublishDay = in.PublishDay
	item.Status = in.Status
	updated, updateErr := uc.dc.LessonRepo().Update(ctx, item)
	if updateErr != nil {
		return nil, errx.Wrap(updateErr)
	}
	video, videoErr := uc.saveVideo(ctx, updated, in.Video)
	if videoErr != nil {
		return nil, errx.Wrap(videoErr)
	}
	materials, matErr := uc.saveMaterials(ctx, updated, in.Materials)
	if matErr != nil {
		return nil, errx.Wrap(matErr)
	}
	return &Response{Item: *updated, Video: video, Materials: materials}, nil
}

func (uc *usecase) saveVideo(ctx context.Context, item *lesson.Lesson, input *VideoInput) (*lessonvideo.LessonVideo, error) {
	if input == nil {
		return uc.dc.LessonVideoRepo().GetByLesson(ctx, item.ID)
	}
	if !input.Enabled {
		return nil, uc.dc.LessonVideoRepo().DeleteByLesson(ctx, item.ID)
	}
	provider := input.Provider
	if provider == "" {
		provider = lessonvideo.ProviderTelegram
	}
	return uc.dc.LessonVideoRepo().Upsert(ctx, &lessonvideo.LessonVideo{OrganizationID: item.OrganizationID, LessonID: item.ID, Provider: provider, StreamRef: shared.TrimPtr(input.StreamRef), TelegramChannelID: shared.TrimPtr(input.TelegramChannelID), TelegramMessageID: shared.TrimPtr(input.TelegramMessageID), EmbedURL: shared.TrimPtr(input.EmbedURL), DurationSeconds: input.DurationSeconds})
}

func (uc *usecase) saveMaterials(ctx context.Context, item *lesson.Lesson, inputs []MaterialInput) ([]lessonmaterial.LessonMaterial, error) {
	replaceItems := make([]lessonmaterial.ReplaceItem, 0, len(inputs))
	for idx, input := range inputs {
		order := input.OrderNumber
		if order == 0 {
			order = idx + 1
		}
		replaceItems = append(replaceItems, lessonmaterial.ReplaceItem{Title: strings.TrimSpace(input.Title), MaterialType: input.MaterialType, SourceType: input.SourceType, URL: shared.TrimPtr(input.URL), FileID: shared.TrimPtr(input.FileID), OrderNumber: order})
	}
	return uc.dc.LessonMaterialRepo().ReplaceByLesson(ctx, item.OrganizationID, item.ID, replaceItems)
}
