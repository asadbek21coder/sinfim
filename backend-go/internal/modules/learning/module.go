package learning

import (
	"go-enterprise-blueprint/internal/modules/learning/ctrl/http"
	"go-enterprise-blueprint/internal/modules/learning/usecase"
	"go-enterprise-blueprint/internal/modules/learning/usecase/getlessondetail"
	"go-enterprise-blueprint/internal/modules/learning/usecase/getstudentdashboard"
	"go-enterprise-blueprint/internal/modules/learning/usecase/marklessoncompleted"
	portalauth "go-enterprise-blueprint/internal/portal/auth"

	"github.com/rise-and-shine/pkg/http/server"
	"github.com/uptrace/bun"
)

type Module struct {
	httpCTRL *http.Controller
}

type Config struct{}

func New(_ Config, dbConn *bun.DB, authPortal portalauth.Portal, httpServer *server.HTTPServer) (*Module, error) {
	usecaseContainer := usecase.NewContainer(
		getstudentdashboard.New(dbConn),
		getlessondetail.New(dbConn),
		marklessoncompleted.New(dbConn),
	)
	return &Module{httpCTRL: http.NewController(usecaseContainer, authPortal, httpServer)}, nil
}
