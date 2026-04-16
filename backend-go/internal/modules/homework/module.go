package homework

import (
	"go-enterprise-blueprint/internal/modules/homework/ctrl/http"
	"go-enterprise-blueprint/internal/modules/homework/usecase"
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
		usecase.NewSaveDefinition(dbConn),
		usecase.NewGetLessonHomework(dbConn),
		usecase.NewGetStudentHomework(dbConn),
		usecase.NewSubmitHomework(dbConn),
		usecase.NewListReviewSubmissions(dbConn),
		usecase.NewGetReviewSubmission(dbConn),
		usecase.NewReviewSubmission(dbConn),
	)
	return &Module{httpCTRL: http.NewController(usecaseContainer, authPortal, httpServer)}, nil
}
