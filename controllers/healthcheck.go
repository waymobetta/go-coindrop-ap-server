package controllers

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-ap-server/app"
)

// HealthcheckController implements the healthcheck resource.
type HealthcheckController struct {
	*goa.Controller
}

// NewHealthcheckController creates a healthcheck controller.
func NewHealthcheckController(service *goa.Service) *HealthcheckController {
	return &HealthcheckController{Controller: service.NewController("HealthcheckController")}
}

// Show runs the show action.
func (c *HealthcheckController) Show(ctx *app.ShowHealthcheckContext) error {
	// HealthcheckController_Show: start_implement

	// Put your logic here

	res := &app.Healthcheck{
		Status: "OK",
	}
	return ctx.OK(res)
	// HealthcheckController_Show: end_implement
}
