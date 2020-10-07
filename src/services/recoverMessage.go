package services

import (
	"appml/src/domain/entities"
	"appml/src/interfaces/repositories"
	"appml/src/libext"
)

// RecoverMessage comment
func RecoverMessage(satelliteMeasureRepository repositories.SatelliteMeasureRepository) (string, bool) {

	if satelliteMeasureRepository.Count() != 3 {
		return "", true
	}

	alliance := &entities.Alliance{}

	message := &libext.Message{}
	alliance.SetMessageAlgorithm(message)
	return alliance.GetMessage(satelliteMeasureRepository.GetByNumber(0).Message,
		satelliteMeasureRepository.GetByNumber(1).Message,
		satelliteMeasureRepository.GetByNumber(2).Message)
}
