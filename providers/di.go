package providers

import (
	"ml-test-quasar/controllers"
	"ml-test-quasar/router"
	"ml-test-quasar/server"
	"ml-test-quasar/services"

	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(server.NewServer)
	_ = Container.Provide(router.NewRouter)

	_ = Container.Provide(controllers.NewSatelliteController)
	_ = Container.Provide(services.NewIntelligenceService)

	return Container
}
