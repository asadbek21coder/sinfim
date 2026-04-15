package organization

import (
	"go-enterprise-blueprint/internal/modules/organization/ctrl/http"
	"go-enterprise-blueprint/internal/modules/organization/domain"
	"go-enterprise-blueprint/internal/modules/organization/infra/postgres"
	"go-enterprise-blueprint/internal/modules/organization/usecase"
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
	portalauth "go-enterprise-blueprint/internal/portal/auth"

	"github.com/rise-and-shine/pkg/http/server"
	"github.com/uptrace/bun"
)

type Module struct {
	httpCTRL *http.Controller
}

type Config struct {
	HashingCost int `yaml:"hashing_cost" default:"10" validate:"oneof=4 10 31"`
}

func New(cfg Config, dbConn *bun.DB, authPortal portalauth.Portal, httpServer *server.HTTPServer) (*Module, error) {
	organizationRepo := postgres.NewOrganizationRepo(dbConn)
	membershipRepo := postgres.NewMembershipRepo(dbConn)
	leadRepo := postgres.NewLeadRepo(dbConn)
	schoolRequestRepo := postgres.NewSchoolRequestRepo(dbConn)
	domainContainer := domain.NewContainer(organizationRepo, membershipRepo, leadRepo, schoolRequestRepo)
	usecaseContainer := usecase.NewContainer(
		createorganization.New(dbConn, domainContainer, cfg.HashingCost),
		createlead.New(domainContainer),
		createschoolrequest.New(domainContainer),
		getpublicschoolpage.New(domainContainer),
		listleads.New(domainContainer),
		listmyworkspaces.New(domainContainer),
		listschoolrequests.New(domainContainer),
		updateorganization.New(domainContainer),
		updateleadstatus.New(domainContainer),
		updateschoolrequeststatus.New(domainContainer),
	)

	return &Module{httpCTRL: http.NewController(usecaseContainer, authPortal, httpServer)}, nil
}
