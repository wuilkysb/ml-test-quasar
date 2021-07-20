package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"ml-test-quasar/dtos"
	"ml-test-quasar/interfaces/controller"
	"ml-test-quasar/mocks"
	"ml-test-quasar/utils"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	someError        = fmt.Errorf("error")
	satelliteRequest = []dtos.Satellites{
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

type SatelliteControllerTestSuite struct {
	suite.Suite
	intelligenceService *mocks.IntelligenceServiceInterface
	underTest           controller.SatelliteController
}

func TestSatelliteControllerSuite(t *testing.T) {
	suite.Run(t, new(SatelliteControllerTestSuite))
}

func (suite *SatelliteControllerTestSuite) SetupTest() {
	suite.intelligenceService = &mocks.IntelligenceServiceInterface{}
	suite.underTest = NewSatelliteController(suite.intelligenceService)
}

func (suite *SatelliteControllerTestSuite) TestTopSecret_WhenBindFail() {
	body, _ := json.Marshal("")
	c := SetupControllerCase(http.MethodPost, "/", bytes.NewBuffer(body))
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	err := suite.underTest.TopSecret(c.context)
	suite.Error(err)
}

func (suite *SatelliteControllerTestSuite) TestTopSecret_WhenServiceFail() {
	body, _ := json.Marshal(satelliteRequest)
	c := SetupControllerCase(http.MethodPost, "/", bytes.NewBuffer(body))
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.intelligenceService.Mock.On("GetLocation", satelliteRequest).Return(utils.Point{}, someError)

	err := suite.underTest.TopSecret(c.context)
	suite.Error(err)
}

func (suite *SatelliteControllerTestSuite) TestTopSecret_WhenSuccess() {
	body, _ := json.Marshal(satelliteRequest)
	c := SetupControllerCase(http.MethodPost, "/", bytes.NewBuffer(body))
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.intelligenceService.Mock.On("GetLocation", satelliteRequest).Return(utils.Point{}, nil)
	suite.intelligenceService.Mock.On("GetMessage", satelliteRequest).Return("")

	err := suite.underTest.TopSecret(c.context)
	suite.NoError(err)
}
