package services

import (
	"testing"

	"ml-test-quasar/dtos"
	"ml-test-quasar/interfaces/services"
	"ml-test-quasar/utils"

	"github.com/stretchr/testify/suite"
)

var (
	satellites = []dtos.Satellites{
		{
			Name:     "kenobi",
			Distance: 1,
			Message:  []string{"es", "", ""},
		},
		{
			Name:     "skywalker",
			Distance: 1,
			Message:  []string{"", "", "mensaje"},
		},
		{
			Name:     "sato",
			Distance: 1,
			Message:  []string{"", "un", ""},
		},
	}
)

type IntelligenceServiceTestSuite struct {
	suite.Suite
	underTest services.IntelligenceServiceInterface
}

func TestIntelligenceServiceSuite(t *testing.T) {
	suite.Run(t, new(IntelligenceServiceTestSuite))
}

func (suite *IntelligenceServiceTestSuite) SetupTest() {
	suite.underTest = NewIntelligenceService()
}

func (suite *IntelligenceServiceTestSuite) TestGetLocation() {
	location, _ := suite.underTest.GetLocation(satellites)
	suite.IsType(utils.Point{}, location)
}

func (suite *IntelligenceServiceTestSuite) TestGetMessage() {
	message := suite.underTest.GetMessage(satellites)
	suite.IsType("es un mensaje", message)
}
