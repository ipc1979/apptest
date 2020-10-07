package infra

import (
	"appml/src/domain/entities"
	"sync"
)

// SatelliteMeasureDB comment
type SatelliteMeasureDB struct {
	satellitesMeasures []entities.SatelliteMeasure
}

var satelliteRepositoryDB *SatelliteMeasureDB
var satelliteRepositoryOnce sync.Once

// GetSatelliteMeasureDB comment
func GetSatelliteMeasureDB() *SatelliteMeasureDB {
	satelliteRepositoryOnce.Do(func() {
		satelliteRepositoryDB = &SatelliteMeasureDB{satellitesMeasures: []entities.SatelliteMeasure{}}
	})
	return satelliteRepositoryDB
}

// Save comment
func (db *SatelliteMeasureDB) Save(satelliteMeasure entities.SatelliteMeasure) {
	db.satellitesMeasures = append(db.satellitesMeasures, satelliteMeasure)
}

// GetByNumber comment
func (db *SatelliteMeasureDB) GetByNumber(num int) entities.SatelliteMeasure {
	return db.satellitesMeasures[num]
}

// FindByName comment
func (db *SatelliteMeasureDB) FindByName(name string) *entities.SatelliteMeasure {
	for _, satelliteMeasure := range db.satellitesMeasures {
		if satelliteMeasure.Name == name {
			return &satelliteMeasure
		}
	}
	return nil
}

// Update comment
func (db *SatelliteMeasureDB) Update(satelliteMeasure entities.SatelliteMeasure) {
	for n, dbSatelliteMeasure := range db.satellitesMeasures {
		if dbSatelliteMeasure.Name == satelliteMeasure.Name {
			db.satellitesMeasures[n].Distance = satelliteMeasure.Distance
			db.satellitesMeasures[n].Message = satelliteMeasure.Message
		}
	}
}

// Count comment
func (db *SatelliteMeasureDB) Count() int {
	return len(db.satellitesMeasures)
}

// ClearAll comment
func (db *SatelliteMeasureDB) ClearAll() {
	db.satellitesMeasures = nil
}
