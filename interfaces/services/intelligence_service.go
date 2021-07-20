package services

import (
	"ml-mutant-test/dtos"
	"ml-mutant-test/utils"
)

type IntelligenceServiceInterface interface {
	GetLocation(satellites []dtos.Satellites) (utils.Point, error)
	GetMessage(satellites []dtos.Satellites) string
}
