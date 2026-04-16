package usecase

import (
	"go-enterprise-blueprint/internal/modules/learning/usecase/getlessondetail"
	"go-enterprise-blueprint/internal/modules/learning/usecase/getstudentdashboard"
	"go-enterprise-blueprint/internal/modules/learning/usecase/marklessoncompleted"
)

type Container struct {
	getStudentDashboard getstudentdashboard.UseCase
	getLessonDetail     getlessondetail.UseCase
	markLessonCompleted marklessoncompleted.UseCase
}

func NewContainer(getStudentDashboard getstudentdashboard.UseCase, getLessonDetail getlessondetail.UseCase, markLessonCompleted marklessoncompleted.UseCase) *Container {
	return &Container{getStudentDashboard: getStudentDashboard, getLessonDetail: getLessonDetail, markLessonCompleted: markLessonCompleted}
}

func (c *Container) GetStudentDashboard() getstudentdashboard.UseCase { return c.getStudentDashboard }

func (c *Container) GetLessonDetail() getlessondetail.UseCase { return c.getLessonDetail }

func (c *Container) MarkLessonCompleted() marklessoncompleted.UseCase { return c.markLessonCompleted }
