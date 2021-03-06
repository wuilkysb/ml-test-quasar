// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// SatelliteController is an autogenerated mock type for the SatelliteController type
type SatelliteController struct {
	mock.Mock
}

// TopSecret provides a mock function with given fields: c
func (_m *SatelliteController) TopSecret(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
