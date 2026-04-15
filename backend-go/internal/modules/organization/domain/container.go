package domain

import (
	"go-enterprise-blueprint/internal/modules/organization/domain/lead"
	"go-enterprise-blueprint/internal/modules/organization/domain/membership"
	"go-enterprise-blueprint/internal/modules/organization/domain/org"
	"go-enterprise-blueprint/internal/modules/organization/domain/schoolrequest"
)

type Container struct {
	organizationRepo  org.Repo
	membershipRepo    membership.Repo
	leadRepo          lead.Repo
	schoolRequestRepo schoolrequest.Repo
}

func NewContainer(organizationRepo org.Repo, membershipRepo membership.Repo, leadRepo lead.Repo, schoolRequestRepo schoolrequest.Repo) *Container {
	return &Container{organizationRepo: organizationRepo, membershipRepo: membershipRepo, leadRepo: leadRepo, schoolRequestRepo: schoolRequestRepo}
}

func (c *Container) OrganizationRepo() org.Repo {
	return c.organizationRepo
}

func (c *Container) MembershipRepo() membership.Repo {
	return c.membershipRepo
}

func (c *Container) LeadRepo() lead.Repo {
	return c.leadRepo
}

func (c *Container) SchoolRequestRepo() schoolrequest.Repo {
	return c.schoolRequestRepo
}
