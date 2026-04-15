package http

import (
	"go-enterprise-blueprint/internal/modules/classroom/usecase"
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
	v1 := r.Group("/api/v1/classroom", auth.NewAuthMiddleware(c.authPortal))
	v1.Post("/create-class", forward.ToUserAction(c.usecaseContainer.CreateClass()))
	v1.Get("/list-classes", forward.ToUserAction(c.usecaseContainer.ListClasses()))
	v1.Get("/get-class-detail", forward.ToUserAction(c.usecaseContainer.GetClassDetail()))
	v1.Post("/add-student", forward.ToUserAction(c.usecaseContainer.AddStudent()))
	v1.Post("/update-access", forward.ToUserAction(c.usecaseContainer.UpdateAccess()))
	v1.Post("/assign-mentor", forward.ToUserAction(c.usecaseContainer.AssignMentor()))
}
