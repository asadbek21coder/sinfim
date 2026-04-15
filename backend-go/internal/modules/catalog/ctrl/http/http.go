package http

import (
	"go-enterprise-blueprint/internal/modules/catalog/usecase"
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
	v1 := r.Group("/api/v1/catalog")
	v1.Get("/get-public-course-page", forward.ToUserAction(c.usecaseContainer.GetPublicCoursePage()))

	v1Auth := v1.Group("", auth.NewAuthMiddleware(c.authPortal))
	v1Auth.Post("/create-course", forward.ToUserAction(c.usecaseContainer.CreateCourse()))
	v1Auth.Post("/update-course", forward.ToUserAction(c.usecaseContainer.UpdateCourse()))
	v1Auth.Get("/list-courses", forward.ToUserAction(c.usecaseContainer.ListCourses()))
	v1Auth.Get("/get-course-detail", forward.ToUserAction(c.usecaseContainer.GetCourseDetail()))
}
