package http

import (
	"go-enterprise-blueprint/internal/modules/organization/usecase"
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
	v1 := r.Group("/api/v1/organization")

	v1.Post("/create-school-request", forward.ToUserAction(c.usecaseContainer.CreateSchoolRequest()))
	v1.Get("/get-demo-access", forward.ToUserAction(c.usecaseContainer.GetDemoAccess()))
	v1.Get("/get-public-school-page", forward.ToUserAction(c.usecaseContainer.GetPublicSchoolPage()))
	v1.Post("/create-lead", forward.ToUserAction(c.usecaseContainer.CreateLead()))

	v1Auth := v1.Group("", auth.NewAuthMiddleware(c.authPortal))
	platformAdmin := auth.RequirePermission(auth.PermissionUserRead)
	platformManage := auth.RequirePermission(auth.PermissionUserManage)
	v1Auth.Post("/create-organization", platformManage, forward.ToUserAction(c.usecaseContainer.CreateOrganization()))
	v1Auth.Get("/list-my-workspaces", forward.ToUserAction(c.usecaseContainer.ListMyWorkspaces()))
	v1Auth.Get("/get-owner-dashboard", forward.ToUserAction(c.usecaseContainer.GetOwnerDashboard()))
	v1Auth.Post("/update-organization", forward.ToUserAction(c.usecaseContainer.UpdateOrganization()))
	v1Auth.Get("/list-leads", forward.ToUserAction(c.usecaseContainer.ListLeads()))
	v1Auth.Post("/update-lead-status", forward.ToUserAction(c.usecaseContainer.UpdateLeadStatus()))
	v1Auth.Get("/list-school-requests", platformAdmin, forward.ToUserAction(c.usecaseContainer.ListSchoolRequests()))
	v1Auth.Post("/update-school-request-status", platformAdmin, forward.ToUserAction(c.usecaseContainer.UpdateSchoolRequestStatus()))
}
