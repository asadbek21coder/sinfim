package domain

import (
	"go-enterprise-blueprint/internal/modules/catalog/domain/course"
	"go-enterprise-blueprint/internal/modules/catalog/domain/lesson"
	"go-enterprise-blueprint/internal/modules/catalog/domain/lessonmaterial"
	"go-enterprise-blueprint/internal/modules/catalog/domain/lessonvideo"
	"go-enterprise-blueprint/internal/modules/organization/domain/membership"
	"go-enterprise-blueprint/internal/modules/organization/domain/org"
)

type Container struct {
	courseRepo       course.Repo
	lessonRepo       lesson.Repo
	lessonVideoRepo  lessonvideo.Repo
	lessonMatRepo    lessonmaterial.Repo
	membershipRepo   membership.Repo
	organizationRepo org.Repo
}

func NewContainer(courseRepo course.Repo, lessonRepo lesson.Repo, lessonVideoRepo lessonvideo.Repo, lessonMatRepo lessonmaterial.Repo, membershipRepo membership.Repo, organizationRepo org.Repo) *Container {
	return &Container{courseRepo: courseRepo, lessonRepo: lessonRepo, lessonVideoRepo: lessonVideoRepo, lessonMatRepo: lessonMatRepo, membershipRepo: membershipRepo, organizationRepo: organizationRepo}
}

func (c *Container) CourseRepo() course.Repo { return c.courseRepo }

func (c *Container) LessonRepo() lesson.Repo { return c.lessonRepo }

func (c *Container) LessonVideoRepo() lessonvideo.Repo { return c.lessonVideoRepo }

func (c *Container) LessonMaterialRepo() lessonmaterial.Repo { return c.lessonMatRepo }

func (c *Container) MembershipRepo() membership.Repo { return c.membershipRepo }

func (c *Container) OrganizationRepo() org.Repo { return c.organizationRepo }
