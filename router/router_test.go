package router

import (
	"testing"

	"ml-mutant-test/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

type RouterTestSuite struct {
	suite.Suite

	server              *echo.Echo
	satelliteController *mocks.SatelliteController
	underTest           *Router
}

func (suite *RouterTestSuite) SetupTest() {
	suite.server = echo.New()
	suite.satelliteController = &mocks.SatelliteController{}
	suite.underTest = NewRouter(suite.server, suite.satelliteController)
}

func (suite *RouterTestSuite) TestInit() {
	suite.underTest.Init()
}
