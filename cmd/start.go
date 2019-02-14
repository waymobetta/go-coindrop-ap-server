//go:generate goagen bootstrap -d github.com/waymobetta/go-coindrop-ap-server/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/waymobetta/go-coindrop-ap-server/app"
	"github.com/waymobetta/go-coindrop-ap-server/controllers"
)

func main() {
	// Create service
	service := goa.New("coindrop-ap")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "healthcheck" controller
	c := controllers.NewHealthcheckController(service)
	app.MountHealthcheckController(service, c)
	// Mount "paycheck" controller
	c2 := controllers.NewPaycheckController(service)
	app.MountPaycheckController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8888"); err != nil {
		service.LogError("startup", "err", err)
	}

}
