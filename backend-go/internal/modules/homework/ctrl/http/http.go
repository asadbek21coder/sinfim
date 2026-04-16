package http

import (
	"go-enterprise-blueprint/internal/modules/homework/usecase"
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
	v1 := r.Group("/api/v1/homework", auth.NewAuthMiddleware(c.authPortal))
	v1.Post("/save-definition", forward.ToUserAction(c.usecaseContainer.SaveDefinition()))
	v1.Get("/get-lesson-homework", forward.ToUserAction(c.usecaseContainer.GetLessonHomework()))
	v1.Get("/get-student-homework", forward.ToUserAction(c.usecaseContainer.GetStudentHomework()))
	v1.Post("/submit-homework", forward.ToUserAction(c.usecaseContainer.SubmitHomework()))
	v1.Get("/list-review-submissions", forward.ToUserAction(c.usecaseContainer.ListReviewSubmissions()))
	v1.Get("/get-review-submission", forward.ToUserAction(c.usecaseContainer.GetReviewSubmission()))
	v1.Post("/review-submission", forward.ToUserAction(c.usecaseContainer.ReviewSubmission()))
}
