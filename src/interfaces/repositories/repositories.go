package repositories

import (
	"appml/src/domain/entities"
)

// SatelliteRepository comment
type SatelliteRepository interface {
	Save(satellite entities.Satellite)
	GetByNumber(num int) entities.Satellite
	FindByName(name string) *entities.Satellite
}

// SatelliteMeasureRepository comment
type SatelliteMeasureRepository interface {
	Save(satelliteMeasure entities.SatelliteMeasure)
	ClearAll()
	GetByNumber(num int) entities.SatelliteMeasure
	FindByName(name string) *entities.SatelliteMeasure
	Update(satellite entities.SatelliteMeasure)
	Count() int
}
