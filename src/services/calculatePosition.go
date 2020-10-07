package services

import (
	"appml/src/domain/entities"
	"appml/src/interfaces/repositories"
	"appml/src/libext"
)

// CalculatePosition comment
func CalculatePosition(satelliteRepository repositories.SatelliteRepository,
	satelliteMeasureRepository repositories.SatelliteMeasureRepository) (entities.Position, bool) {

	if satelliteMeasureRepository.Count() != 3 {
		return entities.Position{}, true
	}

	alliance := &entities.Alliance{}
	alliance.AddSatellite(satelliteRepository.GetByNumber(0))
	alliance.AddSatellite(satelliteRepository.GetByNumber(1))
	alliance.AddSatellite(satelliteRepository.GetByNumber(2))

	loc := &libext.Location{}
	alliance.SetLocationAlgorithm(loc)
	return alliance.GetLocation(satelliteMeasureRepository.GetByNumber(0).Distance,
		satelliteMeasureRepository.GetByNumber(1).Distance,
		satelliteMeasureRepository.GetByNumber(2).Distance)
}
