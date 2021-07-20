package controllers

import (
	"net/http"

	"ml-mutant-test/dtos"
	"ml-mutant-test/interfaces/controller"
	"ml-mutant-test/interfaces/services"

	"github.com/labstack/echo/v4"
)

type SatelliteController struct {
	service services.IntelligenceServiceInterface
}

func NewSatelliteController(service services.IntelligenceServiceInterface) controller.SatelliteController {
	return &SatelliteController{
		service,
	}
}

// @Tags Satellite Manager
// @Summary satellite get top secret message
// @Description satellite top secret
// @Accept  json
// @Produce  json
// @Param request body []dtos.Satellites true "Request body"
// @Success 200 {object} dtos.Response
// @Router /satellite/topsecret/ [post]
func (controller SatelliteController) TopSecret(c echo.Context) error {
	var satellites []dtos.Satellites
	if err := c.Bind(&satellites); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	position, err := controller.service.GetLocation(satellites)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, dtos.Response{
		Message:  controller.service.GetMessage(satellites),
		Position: position,
	})
}
