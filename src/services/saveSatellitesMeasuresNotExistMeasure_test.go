package services

import (
	"testing"

	"appml/src/domain/entities"
)

type fakeSatelliteDB10 struct {
}

// GetByNumber comment
func (db *fakeSatelliteDB10) GetByNumber(num int) entities.Satellite {
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
func (db *fakeSatelliteDB10) Save(satellite entities.Satellite) {}
func (db *fakeSatelliteDB10) FindByName(name string) *entities.Satellite {
	return &entities.Satellite{}
}

type fakeSatelliteMeasureDB10 struct {
}

// GetByNumber comment
func (db *fakeSatelliteMeasureDB10) GetByNumber(num int) entities.SatelliteMeasure {
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

func (db *fakeSatelliteMeasureDB10) Save(satelliteMeasure entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB10) ClearAll()                                       {}

// Measure not exist
func (db *fakeSatelliteMeasureDB10) FindByName(name string) *entities.SatelliteMeasure {
	if name == "sat3" {
		return nil
	}
	return &entities.SatelliteMeasure{}
}
func (db *fakeSatelliteMeasureDB10) Update(satellite entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB10) Count() int {
	return 3
}

func TestSaveSatellitesMeasuresNotExistMeasure(t *testing.T) {

	fakeBodyRequest := entities.SatellitesRequest{Satellites: []entities.SatelliteRequest{
		entities.SatelliteRequest{Name: "sat1", Distance: 1.0, Message: []string{"", "", "", "D"}},
		entities.SatelliteRequest{Name: "sat2", Distance: 1.0, Message: []string{"A", "", "C", ""}},
		entities.SatelliteRequest{Name: "sat3", Distance: 1.0, Message: []string{"", "B", "C", ""}}}}

	fakeDbSatellite := &fakeSatelliteDB10{}
	fakeDbSatelliteMeasure := &fakeSatelliteMeasureDB10{}

	message, err := SaveSatellitesMeasures(fakeDbSatellite, fakeDbSatelliteMeasure, fakeBodyRequest)

	if message != "Satellites loaded" && err {
		t.Errorf("SaveSatellites Error - Satellites loaded %s.", message)
	}

}
