package services

import (
	"testing"

	"appml/src/domain/entities"
)

type fakeSatelliteDB2 struct {
}

// GetByNumber comment
func (db *fakeSatelliteDB2) GetByNumber(num int) entities.Satellite {
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
func (db *fakeSatelliteDB2) Save(satellite entities.Satellite) {}
func (db *fakeSatelliteDB2) FindByName(name string) *entities.Satellite {
	return &entities.Satellite{}
}

type fakeSatelliteMeasureDB2 struct {
}

// GetByNumber comment
func (db *fakeSatelliteMeasureDB2) GetByNumber(num int) entities.SatelliteMeasure {
	var satelliteMeasure entities.SatelliteMeasure
	switch num {
	case 0:
		satelliteMeasure = entities.SatelliteMeasure{Distance: 1.4142135623}
	case 1:
		satelliteMeasure = entities.SatelliteMeasure{Distance: 1.0}
	case 2:
		satelliteMeasure = entities.SatelliteMeasure{Distance: 1.4142135623}
	}
	return satelliteMeasure
}

func (db *fakeSatelliteMeasureDB2) Save(satelliteMeasure entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB2) ClearAll()                                       {}
func (db *fakeSatelliteMeasureDB2) FindByName(name string) *entities.SatelliteMeasure {
	return &entities.SatelliteMeasure{}
}
func (db *fakeSatelliteMeasureDB2) Update(satellite entities.SatelliteMeasure) {}

// Not enough Satellites Measures
func (db *fakeSatelliteMeasureDB2) Count() int {
	return 2
}

func TestCalculatePositionNotEnoughSatellitesMeasures(t *testing.T) {

	fakeDbSatellite := &fakeSatelliteDB2{}
	fakeDbSatelliteMeasure := &fakeSatelliteMeasureDB2{}

	_, err := CalculatePosition(fakeDbSatellite, fakeDbSatelliteMeasure)

	if !err {
		t.Errorf("Position Error - Not enough Satellites Measures")
	}

}
