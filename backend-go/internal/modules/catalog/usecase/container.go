package usecase

import (
	"go-enterprise-blueprint/internal/modules/catalog/usecase/createcourse"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/createlesson"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/getcoursedetail"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/getlessondetail"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/getpubliccoursepage"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/listcourses"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/listlessons"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/updatecourse"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/updatelesson"
)

type Container struct {
	createCourse        createcourse.UseCase
	updateCourse        updatecourse.UseCase
	listCourses         listcourses.UseCase
	getCourseDetail     getcoursedetail.UseCase
	getPublicCoursePage getpubliccoursepage.UseCase
	createLesson        createlesson.UseCase
	updateLesson        updatelesson.UseCase
	listLessons         listlessons.UseCase
	getLessonDetail     getlessondetail.UseCase
}

func NewContainer(
	createCourse createcourse.UseCase,
	updateCourse updatecourse.UseCase,
	listCourses listcourses.UseCase,
	getCourseDetail getcoursedetail.UseCase,
	getPublicCoursePage getpubliccoursepage.UseCase,
	createLesson createlesson.UseCase,
	updateLesson updatelesson.UseCase,
	listLessons listlessons.UseCase,
	getLessonDetail getlessondetail.UseCase,
) *Container {
	return &Container{
		createCourse:        createCourse,
		updateCourse:        updateCourse,
		listCourses:         listCourses,
		getCourseDetail:     getCourseDetail,
		getPublicCoursePage: getPublicCoursePage,
		createLesson:        createLesson,
		updateLesson:        updateLesson,
		listLessons:         listLessons,
		getLessonDetail:     getLessonDetail,
	}
}

func (c *Container) CreateCourse() createcourse.UseCase { return c.createCourse }

func (c *Container) UpdateCourse() updatecourse.UseCase { return c.updateCourse }

func (c *Container) ListCourses() listcourses.UseCase { return c.listCourses }

func (c *Container) GetCourseDetail() getcoursedetail.UseCase { return c.getCourseDetail }

func (c *Container) GetPublicCoursePage() getpubliccoursepage.UseCase { return c.getPublicCoursePage }

func (c *Container) CreateLesson() createlesson.UseCase { return c.createLesson }

func (c *Container) UpdateLesson() updatelesson.UseCase { return c.updateLesson }

func (c *Container) ListLessons() listlessons.UseCase { return c.listLessons }

func (c *Container) GetLessonDetail() getlessondetail.UseCase { return c.getLessonDetail }
