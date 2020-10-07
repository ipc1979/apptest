package infra

import (
	"appml/src/domain/entities"
	"sync"
)

// SatelliteDB comment
type SatelliteDB struct {
	satellites []entities.Satellite
}

var satelliteDB *SatelliteDB
var satelliteOnce sync.Once

// GetSatelliteDB comment
func GetSatelliteDB() *SatelliteDB {
	satelliteOnce.Do(func() {
		satelliteDB = &SatelliteDB{satellites: []entities.Satellite{}}
	})
	return satelliteDB
}

// Save comment
func (db *SatelliteDB) Save(satellite entities.Satellite) {
	db.satellites = append(db.satellites, satellite)
}

// GetByNumber comment
func (db *SatelliteDB) GetByNumber(num int) entities.Satellite {
	return db.satellites[num]
}

// FindByName comment
func (db *SatelliteDB) FindByName(name string) *entities.Satellite {
	for _, satellite := range db.satellites {
		if satellite.Name == name {
			return &satellite
		}
	}
	return nil
}
