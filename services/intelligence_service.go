package services

import (
	"ml-test-quasar/config"
	"ml-test-quasar/dtos"
	"ml-test-quasar/interfaces/services"
	"ml-test-quasar/utils"
)

type IntelligenceService struct {
}

func NewIntelligenceService() services.IntelligenceServiceInterface {
	return &IntelligenceService{
	}
}

func (s IntelligenceService) GetLocation(satellites []dtos.Satellites) (utils.Point, error) {
	satellitesAlliance := map[string]*utils.Point{
		"kenobi":    &config.Kenobi,
		"skywalker": &config.Skywalker,
		"sato":      &config.Sato,
	}

	for _, satellite := range satellites {
		satellitesAlliance[satellite.Name].R = satellite.Distance
	}

	return utils.Solve(config.Kenobi, config.Skywalker, config.Sato)
}

func (s IntelligenceService) GetMessage(satellites []dtos.Satellites) string {
	return utils.GetCompleteMessage(satellites[0].Message, satellites[1].Message, satellites[2].Message)
}
