package services

import (
	"appml/src/domain/entities"
	"appml/src/interfaces/repositories"
	"strings"
)

// SaveSatellitesMeasures comment
func SaveSatellitesMeasures(satelliteRepository repositories.SatelliteRepository,
	satelliteMeasureRepository repositories.SatelliteMeasureRepository,
	satellitesRequest entities.SatellitesRequest) (string, bool) {

	// Check if exists invalid satellites on the request
	// Distances != 0 and Empty messages
	notValidSatellites := ""
	for _, satellite := range satellitesRequest.Satellites {
		if satelliteRepository.FindByName(satellite.Name) == nil {
			notValidSatellites += " - " + satellite.Name
		} else if satellite.Distance == 0 || len(strings.Join(satellite.Message, "")) == 0 {
			notValidSatellites += " - " + satellite.Name
		}
	}
	if len(notValidSatellites) > 0 {
		return "Not valid Satellites Measures" + notValidSatellites, true
	}

	// Make Create or Update of the satellites
	for _, satellite := range satellitesRequest.Satellites {
		satelliteMeasureTemp := entities.SatelliteMeasure{Name: satellite.Name,
			Distance: satellite.Distance,
			Message:  satellite.Message}
		if satelliteMeasureRepository.FindByName(satellite.Name) != nil {
			satelliteMeasureRepository.Update(satelliteMeasureTemp)
		} else {
			satelliteMeasureRepository.Save(satelliteMeasureTemp)
		}
	}

	return "Satellites loaded", false
}
