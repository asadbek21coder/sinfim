package domain

import (
	"go-enterprise-blueprint/internal/modules/catalog/domain/course"
	"go-enterprise-blueprint/internal/modules/organization/domain/membership"
	"go-enterprise-blueprint/internal/modules/organization/domain/org"
)

type Container struct {
	courseRepo       course.Repo
	membershipRepo   membership.Repo
	organizationRepo org.Repo
}

func NewContainer(courseRepo course.Repo, membershipRepo membership.Repo, organizationRepo org.Repo) *Container {
	return &Container{courseRepo: courseRepo, membershipRepo: membershipRepo, organizationRepo: organizationRepo}
}

func (c *Container) CourseRepo() course.Repo { return c.courseRepo }

func (c *Container) MembershipRepo() membership.Repo { return c.membershipRepo }

func (c *Container) OrganizationRepo() org.Repo { return c.organizationRepo }
