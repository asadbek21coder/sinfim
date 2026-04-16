package http

import (
	"go-enterprise-blueprint/internal/modules/learning/usecase"
	"go-enterprise-blueprint/internal/portal/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/rise-and-shine/pkg/http/server"
	"github.com/rise-and-shine/pkg/http/server/forward"
)

type Controller struct {
	usecaseContainer *usecase.Container
	authPortal       auth.Portal
}

func NewController(usecaseContainer *usecase.Container, authPortal auth.Portal, httpServer *server.HTTPServer) *Controller {
	ctrl := &Controller{usecaseContainer: usecaseContainer, authPortal: authPortal}
	httpServer.RegisterRouter(ctrl.initRoutes)
	return ctrl
}

func (c *Controller) initRoutes(r fiber.Router) {
	v1 := r.Group("/api/v1/learning", auth.NewAuthMiddleware(c.authPortal))
	v1.Get("/get-student-dashboard", forward.ToUserAction(c.usecaseContainer.GetStudentDashboard()))
	v1.Get("/get-lesson-detail", forward.ToUserAction(c.usecaseContainer.GetLessonDetail()))
	v1.Post("/mark-lesson-completed", forward.ToUserAction(c.usecaseContainer.MarkLessonCompleted()))
}
