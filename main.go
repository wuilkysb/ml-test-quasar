package main

import (
	"fmt"

	"ml-test-quasar/config"
	"ml-test-quasar/providers"
	"ml-test-quasar/router"

	"github.com/labstack/echo/v4"
	"github.com/signalfx/signalfx-go-tracing/tracing"
)

func main() {
	tracing.Start()
	defer tracing.Stop()
	container := providers.BuildContainer()

	err := container.Invoke(func(server *echo.Echo, route *router.Router) {
		address := fmt.Sprintf(":%s", config.Environments().Port)

		route.Init()
		server.Logger.Fatal(server.Start(address))
	})

	if err != nil {
		panic(err)
	}
}
