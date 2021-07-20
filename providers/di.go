package providers

import (
	"ml-mutant-test/controllers"
	"ml-mutant-test/router"
	"ml-mutant-test/server"
	"ml-mutant-test/services"

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
