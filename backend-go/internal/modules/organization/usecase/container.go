package usecase

import (
	"go-enterprise-blueprint/internal/modules/organization/usecase/createlead"
	"go-enterprise-blueprint/internal/modules/organization/usecase/createorganization"
	"go-enterprise-blueprint/internal/modules/organization/usecase/createschoolrequest"
	"go-enterprise-blueprint/internal/modules/organization/usecase/getpublicschoolpage"
	"go-enterprise-blueprint/internal/modules/organization/usecase/listleads"
	"go-enterprise-blueprint/internal/modules/organization/usecase/listmyworkspaces"
	"go-enterprise-blueprint/internal/modules/organization/usecase/listschoolrequests"
	"go-enterprise-blueprint/internal/modules/organization/usecase/updateleadstatus"
	"go-enterprise-blueprint/internal/modules/organization/usecase/updateorganization"
	"go-enterprise-blueprint/internal/modules/organization/usecase/updateschoolrequeststatus"
)

type Container struct {
	createOrganization        createorganization.UseCase
	createLead                createlead.UseCase
	createSchoolRequest       createschoolrequest.UseCase
	getPublicSchoolPage       getpublicschoolpage.UseCase
	listLeads                 listleads.UseCase
	listMyWorkspaces          listmyworkspaces.UseCase
	listSchoolRequests        listschoolrequests.UseCase
	updateOrganization        updateorganization.UseCase
	updateLeadStatus          updateleadstatus.UseCase
	updateSchoolRequestStatus updateschoolrequeststatus.UseCase
}

func NewContainer(
	createOrganization createorganization.UseCase,
	createLead createlead.UseCase,
	createSchoolRequest createschoolrequest.UseCase,
	getPublicSchoolPage getpublicschoolpage.UseCase,
	listLeads listleads.UseCase,
	listMyWorkspaces listmyworkspaces.UseCase,
	listSchoolRequests listschoolrequests.UseCase,
	updateOrganization updateorganization.UseCase,
	updateLeadStatus updateleadstatus.UseCase,
	updateSchoolRequestStatus updateschoolrequeststatus.UseCase,
) *Container {
	return &Container{
		createOrganization:        createOrganization,
		createLead:                createLead,
		createSchoolRequest:       createSchoolRequest,
		getPublicSchoolPage:       getPublicSchoolPage,
		listLeads:                 listLeads,
		listMyWorkspaces:          listMyWorkspaces,
		listSchoolRequests:        listSchoolRequests,
		updateOrganization:        updateOrganization,
		updateLeadStatus:          updateLeadStatus,
		updateSchoolRequestStatus: updateSchoolRequestStatus,
	}
}

func (c *Container) CreateOrganization() createorganization.UseCase {
	return c.createOrganization
}

func (c *Container) CreateLead() createlead.UseCase {
	return c.createLead
}

func (c *Container) CreateSchoolRequest() createschoolrequest.UseCase {
	return c.createSchoolRequest
}

func (c *Container) GetPublicSchoolPage() getpublicschoolpage.UseCase {
	return c.getPublicSchoolPage
}

func (c *Container) ListLeads() listleads.UseCase {
	return c.listLeads
}

func (c *Container) ListMyWorkspaces() listmyworkspaces.UseCase {
	return c.listMyWorkspaces
}

func (c *Container) ListSchoolRequests() listschoolrequests.UseCase {
	return c.listSchoolRequests
}

func (c *Container) UpdateOrganization() updateorganization.UseCase {
	return c.updateOrganization
}

func (c *Container) UpdateLeadStatus() updateleadstatus.UseCase {
	return c.updateLeadStatus
}

func (c *Container) UpdateSchoolRequestStatus() updateschoolrequeststatus.UseCase {
	return c.updateSchoolRequestStatus
}
