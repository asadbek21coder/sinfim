package usecase

import (
	"go-enterprise-blueprint/internal/modules/catalog/usecase/createcourse"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/getcoursedetail"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/getpubliccoursepage"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/listcourses"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/updatecourse"
)

type Container struct {
	createCourse        createcourse.UseCase
	updateCourse        updatecourse.UseCase
	listCourses         listcourses.UseCase
	getCourseDetail     getcoursedetail.UseCase
	getPublicCoursePage getpubliccoursepage.UseCase
}

func NewContainer(
	createCourse createcourse.UseCase,
	updateCourse updatecourse.UseCase,
	listCourses listcourses.UseCase,
	getCourseDetail getcoursedetail.UseCase,
	getPublicCoursePage getpubliccoursepage.UseCase,
) *Container {
	return &Container{
		createCourse:        createCourse,
		updateCourse:        updateCourse,
		listCourses:         listCourses,
		getCourseDetail:     getCourseDetail,
		getPublicCoursePage: getPublicCoursePage,
	}
}

func (c *Container) CreateCourse() createcourse.UseCase { return c.createCourse }

func (c *Container) UpdateCourse() updatecourse.UseCase { return c.updateCourse }

func (c *Container) ListCourses() listcourses.UseCase { return c.listCourses }

func (c *Container) GetCourseDetail() getcoursedetail.UseCase { return c.getCourseDetail }

func (c *Container) GetPublicCoursePage() getpubliccoursepage.UseCase { return c.getPublicCoursePage }
