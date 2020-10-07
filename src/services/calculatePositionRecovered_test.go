package services

import (
	"testing"

	"appml/src/domain/entities"
)

type fakeSatelliteDB1 struct {
}

// GetByNumber comment
func (db *fakeSatelliteDB1) GetByNumber(num int) entities.Satellite {
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
func (db *fakeSatelliteDB1) Save(satellite entities.Satellite) {}
func (db *fakeSatelliteDB1) FindByName(name string) *entities.Satellite {
	return &entities.Satellite{}
}

type fakeSatelliteMeasureDB1 struct {
}

// GetByNumber comment
func (db *fakeSatelliteMeasureDB1) GetByNumber(num int) entities.SatelliteMeasure {
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

func (db *fakeSatelliteMeasureDB1) Save(satelliteMeasure entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB1) ClearAll()                                       {}
func (db *fakeSatelliteMeasureDB1) FindByName(name string) *entities.SatelliteMeasure {
	return &entities.SatelliteMeasure{}
}
func (db *fakeSatelliteMeasureDB1) Update(satellite entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB1) Count() int {
	return 3
}

func TestCalculatePositionRecovered(t *testing.T) {

	fakeDbSatellite := &fakeSatelliteDB1{}
	fakeDbSatelliteMeasure := &fakeSatelliteMeasureDB1{}

	position, err := CalculatePosition(fakeDbSatellite, fakeDbSatelliteMeasure)

	if err {
		t.Errorf("Position Error - Position Expected : 2.00007 , 0.99992985 , Received %f, want: %f.", position.X, position.Y)
	}

}
