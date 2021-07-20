package services

import (
	"ml-test-quasar/dtos"
	"ml-test-quasar/utils"
)

type IntelligenceServiceInterface interface {
	GetLocation(satellites []dtos.Satellites) (utils.Point, error)
	GetMessage(satellites []dtos.Satellites) string
}
