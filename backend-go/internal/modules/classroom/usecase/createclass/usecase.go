package createclass

import (
	"context"
	"time"

	"go-enterprise-blueprint/internal/modules/catalog/domain/course"
	"go-enterprise-blueprint/internal/modules/classroom/domain"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classgroup"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classmentor"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/shared"
	"go-enterprise-blueprint/internal/modules/organization/domain/membership"

	"github.com/code19m/errx"
	"github.com/rise-and-shine/pkg/ucdef"
)

type Request struct {
	OrganizationID string   `json:"organization_id" validate:"required,uuid"`
	CourseID       string   `json:"course_id" validate:"required,uuid"`
	Name           string   `json:"name" validate:"required,min=2,max=160"`
	StartDate      *string  `json:"start_date" validate:"omitempty"`
	LessonCadence  string   `json:"lesson_cadence" validate:"omitempty,oneof=daily every_other_day weekly_3 manual"`
	MentorUserIDs  []string `json:"mentor_user_ids" validate:"omitempty,dive,required"`
}

type Response struct {
	Item classgroup.Summary `json:"item"`
}

type UseCase = ucdef.UserAction[*Request, *Response]

func New(dc *domain.Container) UseCase { return &usecase{dc: dc} }

type usecase struct{ dc *domain.Container }

func (uc *usecase) OperationID() string { return "create-class" }

func (uc *usecase) Execute(ctx context.Context, in *Request) (*Response, error) {
	if err := shared.EnsureClassWrite(ctx, uc.dc, in.OrganizationID); err != nil {
		return nil, errx.Wrap(err)
	}
	courseItem, err := uc.dc.CourseRepo().Get(ctx, course.Filter{ID: &in.CourseID, OrganizationID: &in.OrganizationID})
	if err != nil {
		return nil, errx.Wrap(err)
	}

	startDate, err := parseDate(in.StartDate)
	if err != nil {
		return nil, errx.Wrap(err)
	}
	cadence := in.LessonCadence
	if cadence == "" {
		cadence = classgroup.CadenceEveryOtherDay
	}
	created, err := uc.dc.ClassRepo().Create(ctx, &classgroup.Class{OrganizationID: in.OrganizationID, CourseID: courseItem.ID, Name: in.Name, StartDate: startDate, LessonCadence: cadence, Status: classgroup.StatusActive})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	for _, mentorID := range in.MentorUserIDs {
		ok, existsErr := uc.dc.MembershipRepo().Exists(ctx, mentorID, in.OrganizationID, membership.RoleMentor)
		if existsErr != nil {
			return nil, errx.Wrap(existsErr)
		}
		if !ok {
			return nil, errx.New("mentor role required", errx.WithType(errx.T_Validation), errx.WithCode("MENTOR_ROLE_REQUIRED"))
		}
		_, assignErr := uc.dc.MentorRepo().Assign(ctx, &classmentor.ClassMentor{OrganizationID: in.OrganizationID, ClassID: created.ID, MentorUserID: mentorID})
		if assignErr != nil {
			return nil, errx.Wrap(assignErr)
		}
	}
	items, err := uc.dc.ClassRepo().List(ctx, classgroup.Filter{ID: &created.ID, Limit: 1})
	if err != nil {
		return nil, errx.Wrap(err)
	}
	if len(items) == 0 {
		return nil, errx.New("class not found", errx.WithType(errx.T_NotFound), errx.WithCode(classgroup.CodeClassNotFound))
	}
	return &Response{Item: items[0]}, nil
}

func parseDate(value *string) (*time.Time, error) {
	if value == nil || *value == "" {
		return nil, nil
	}
	parsed, err := time.Parse("2006-01-02", *value)
	if err != nil {
		return nil, errx.New("start_date must be YYYY-MM-DD", errx.WithType(errx.T_Validation), errx.WithCode("VALIDATION_ERROR"))
	}
	return &parsed, nil
}
