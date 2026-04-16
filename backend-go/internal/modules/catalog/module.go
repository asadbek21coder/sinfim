package catalog

import (
	"go-enterprise-blueprint/internal/modules/catalog/ctrl/http"
	"go-enterprise-blueprint/internal/modules/catalog/domain"
	"go-enterprise-blueprint/internal/modules/catalog/infra/postgres"
	"go-enterprise-blueprint/internal/modules/catalog/usecase"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/createcourse"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/createlesson"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/getcoursedetail"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/getlessondetail"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/getpubliccoursepage"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/listcourses"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/listlessons"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/updatecourse"
	"go-enterprise-blueprint/internal/modules/catalog/usecase/updatelesson"
	organizationpg "go-enterprise-blueprint/internal/modules/organization/infra/postgres"
	portalauth "go-enterprise-blueprint/internal/portal/auth"

	"github.com/rise-and-shine/pkg/http/server"
	"github.com/uptrace/bun"
)

type Module struct {
	httpCTRL *http.Controller
}

type Config struct{}

func New(_ Config, dbConn *bun.DB, authPortal portalauth.Portal, httpServer *server.HTTPServer) (*Module, error) {
	courseRepo := postgres.NewCourseRepo(dbConn)
	lessonRepo := postgres.NewLessonRepo(dbConn)
	lessonVideoRepo := postgres.NewLessonVideoRepo(dbConn)
	lessonMaterialRepo := postgres.NewLessonMaterialRepo(dbConn)
	membershipRepo := organizationpg.NewMembershipRepo(dbConn)
	organizationRepo := organizationpg.NewOrganizationRepo(dbConn)
	domainContainer := domain.NewContainer(courseRepo, lessonRepo, lessonVideoRepo, lessonMaterialRepo, membershipRepo, organizationRepo)
	usecaseContainer := usecase.NewContainer(
		createcourse.New(domainContainer),
		updatecourse.New(domainContainer),
		listcourses.New(domainContainer),
		getcoursedetail.New(domainContainer),
		getpubliccoursepage.New(domainContainer),
		createlesson.New(domainContainer),
		updatelesson.New(domainContainer),
		listlessons.New(domainContainer),
		getlessondetail.New(domainContainer),
	)

	return &Module{httpCTRL: http.NewController(usecaseContainer, authPortal, httpServer)}, nil
}
