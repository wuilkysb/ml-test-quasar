package router

import (
	_ "ml-test-quasar/docs"
	"ml-test-quasar/enums"
	"ml-test-quasar/interfaces/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	server              *echo.Echo
	satelliteController controller.SatelliteController
}

func NewRouter(
	server *echo.Echo,
	satelliteController controller.SatelliteController) *Router {
	return &Router{
		server,
		satelliteController,
	}
}

func (r *Router) Init() {
	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))
	apiGroup := r.server.Group(enums.BasePath)

	apiGroup.POST("/topsecret/", r.satelliteController.TopSecret)

	apiGroup.GET("/swagger/*", echoSwagger.WrapHandler)
}
