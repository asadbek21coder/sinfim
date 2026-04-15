package classroom

import (
	"go-enterprise-blueprint/internal/modules/catalog/infra/postgres"
	"go-enterprise-blueprint/internal/modules/classroom/ctrl/http"
	"go-enterprise-blueprint/internal/modules/classroom/domain"
	classroompg "go-enterprise-blueprint/internal/modules/classroom/infra/postgres"
	"go-enterprise-blueprint/internal/modules/classroom/usecase"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/addstudent"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/assignmentor"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/createclass"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/getclassdetail"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/listclasses"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/updateaccess"
	organizationpg "go-enterprise-blueprint/internal/modules/organization/infra/postgres"
	portalauth "go-enterprise-blueprint/internal/portal/auth"

	"github.com/rise-and-shine/pkg/http/server"
	"github.com/uptrace/bun"
)

type Module struct{ httpCTRL *http.Controller }

type Config struct {
	HashingCost int `yaml:"hashing_cost" default:"10" validate:"oneof=4 10 31"`
}

func New(cfg Config, dbConn *bun.DB, authPortal portalauth.Portal, httpServer *server.HTTPServer) (*Module, error) {
	classRepo := classroompg.NewClassRepo(dbConn)
	mentorRepo := classroompg.NewMentorRepo(dbConn)
	enrollmentRepo := classroompg.NewEnrollmentRepo(dbConn)
	accessRepo := classroompg.NewAccessRepo(dbConn)
	courseRepo := postgres.NewCourseRepo(dbConn)
	membershipRepo := organizationpg.NewMembershipRepo(dbConn)
	dc := domain.NewContainer(classRepo, mentorRepo, enrollmentRepo, accessRepo, courseRepo, membershipRepo)
	uc := usecase.NewContainer(
		createclass.New(dc),
		listclasses.New(dc),
		getclassdetail.New(dc),
		addstudent.New(dbConn, dc, cfg.HashingCost),
		updateaccess.New(dc),
		assignmentor.New(dc),
	)
	return &Module{httpCTRL: http.NewController(uc, authPortal, httpServer)}, nil
}
