package domain

import (
	"go-enterprise-blueprint/internal/modules/catalog/domain/course"
	"go-enterprise-blueprint/internal/modules/classroom/domain/accessgrant"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classgroup"
	"go-enterprise-blueprint/internal/modules/classroom/domain/classmentor"
	"go-enterprise-blueprint/internal/modules/classroom/domain/enrollment"
	"go-enterprise-blueprint/internal/modules/organization/domain/membership"
)

type Container struct {
	classRepo      classgroup.Repo
	mentorRepo     classmentor.Repo
	enrollmentRepo enrollment.Repo
	accessRepo     accessgrant.Repo
	courseRepo     course.Repo
	membershipRepo membership.Repo
}

func NewContainer(classRepo classgroup.Repo, mentorRepo classmentor.Repo, enrollmentRepo enrollment.Repo, accessRepo accessgrant.Repo, courseRepo course.Repo, membershipRepo membership.Repo) *Container {
	return &Container{classRepo: classRepo, mentorRepo: mentorRepo, enrollmentRepo: enrollmentRepo, accessRepo: accessRepo, courseRepo: courseRepo, membershipRepo: membershipRepo}
}

func (c *Container) ClassRepo() classgroup.Repo { return c.classRepo }
func (c *Container) MentorRepo() classmentor.Repo { return c.mentorRepo }
func (c *Container) EnrollmentRepo() enrollment.Repo { return c.enrollmentRepo }
func (c *Container) AccessRepo() accessgrant.Repo { return c.accessRepo }
func (c *Container) CourseRepo() course.Repo { return c.courseRepo }
func (c *Container) MembershipRepo() membership.Repo { return c.membershipRepo }
