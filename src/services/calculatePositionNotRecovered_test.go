package services

import (
	"testing"

	"appml/src/domain/entities"
)

type fakeSatelliteDB3 struct {
}

// GetByNumber comment
func (db *fakeSatelliteDB3) GetByNumber(num int) entities.Satellite {
	var satellite entities.Satellite
	switch num {
	case 0:
		satellite = entities.Satellite{Name: "sat1", Position: entities.Position{X: 1.0, Y: 0.0}}
	case 1:
		satellite = entities.Satellite{Name: "sat2", Position: entities.Position{X: 2.0, Y: 0.0}}
	case 2:
		satellite = entities.Satellite{Name: "sat3", Position: entities.Position{X: 3.0, Y: 0.0}}
	}
	return satellite
}
func (db *fakeSatelliteDB3) Save(satellite entities.Satellite) {}
func (db *fakeSatelliteDB3) FindByName(name string) *entities.Satellite {
	return &entities.Satellite{}
}

type fakeSatelliteMeasureDB3 struct {
}

// Distances cant return a valid position.
// GetByNumber comment
func (db *fakeSatelliteMeasureDB3) GetByNumber(num int) entities.SatelliteMeasure {
	var satelliteMeasure entities.SatelliteMeasure
	switch num {
	case 0:
		satelliteMeasure = entities.SatelliteMeasure{Distance: 1.0}
	case 1:
		satelliteMeasure = entities.SatelliteMeasure{Distance: 1.0}
	case 2:
		satelliteMeasure = entities.SatelliteMeasure{Distance: 1.0}
	}
	return satelliteMeasure
}

func (db *fakeSatelliteMeasureDB3) Save(satelliteMeasure entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB3) ClearAll()                                       {}
func (db *fakeSatelliteMeasureDB3) FindByName(name string) *entities.SatelliteMeasure {
	return &entities.SatelliteMeasure{}
}
func (db *fakeSatelliteMeasureDB3) Update(satellite entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB3) Count() int {
	return 3
}

func TestCalculatePositionNotRecovered(t *testing.T) {

	fakeDbSatellite := &fakeSatelliteDB3{}
	fakeDbSatelliteMeasure := &fakeSatelliteMeasureDB3{}

	_, err := CalculatePosition(fakeDbSatellite, fakeDbSatelliteMeasure)

	if !err {
		t.Errorf("Position Error - Distances cant return a valid position.")
	}

}
