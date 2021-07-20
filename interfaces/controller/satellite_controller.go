package controller

import "github.com/labstack/echo/v4"

type SatelliteController interface {
	TopSecret(c echo.Context) error
}
