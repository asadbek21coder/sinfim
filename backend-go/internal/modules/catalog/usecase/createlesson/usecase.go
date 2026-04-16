package createlesson

import (
	"context"
	"strings"

	"go-enterprise-blueprint/internal/modules/catalog/domain"
	"go-enterprise-blueprint/internal/modules/catalog/domain/course"
	"go-enterprise-blueprint/internal/modules/catalog/domain/lesson"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/shared"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	CourseID    string  `json:"course_id" validate:"required,uuid"`
	Title       string  `json:"title" validate:"required,min=2,max=180"`
	Description *string `json:"description" validate:"omitempty,max=2000"`
	OrderNumber int     `json:"order_number" validate:"omitempty,min=1,max=1000"`
	PublishDay  int     `json:"publish_day" validate:"omitempty,min=1,max=1000"`
	Status      string  `json:"status" validate:"omitempty,oneof=draft published archived"`
}

type Response struct {
	Item lesson.Lesson `json:"item"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(dc *domain.Container) UseCase { return &usecase{dc: dc} }

type usecase struct{ dc *domain.Container }

func (uc *usecase) OperationID() string { return "create-lesson" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	courseItem, err := uc.dc.CourseRepo().Get(ctx, course.Filter{ID: &in.CourseID})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if accessErr := shared.EnsureCourseWriteAccess(ctx, uc.dc, courseItem.OrganizationID); accessErr != nil {
		return nil, errx.Wrap(accessErr)
	}
	status := in.Status
	if status == "" {
		status = lesson.StatusDraft
	}
	orderNumber := in.OrderNumber
	if orderNumber == 0 {
		existing, listErr := uc.dc.LessonRepo().List(ctx, lesson.Filter{CourseID: &in.CourseID, Limit: 1000})
		if listErr != nil {
			return nil, errx.Wrap(listErr)
		}
		orderNumber = len(existing) + 1
	}
	publishDay := in.PublishDay
	if publishDay == 0 {
		publishDay = orderNumber
	}
	item, createErr := uc.dc.LessonRepo().Create(ctx, &lesson.Lesson{OrganizationID: courseItem.OrganizationID, CourseID: courseItem.ID, Title: strings.TrimSpace(in.Title), Description: shared.TrimPtr(in.Description), OrderNumber: orderNumber, PublishDay: publishDay, Status: status})
	if createErr != nil {
		return nil, errx.Wrap(createErr)
	}
	return &Response{Item: *item}, nil
}
