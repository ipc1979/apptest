package services

import (
	"testing"

	"appml/src/domain/entities"
)

type fakeSatelliteDB8 struct {
}

// GetByNumber comment
func (db *fakeSatelliteDB8) GetByNumber(num int) entities.Satellite {
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
func (db *fakeSatelliteDB8) Save(satellite entities.Satellite) {}
func (db *fakeSatelliteDB8) FindByName(name string) *entities.Satellite {
	return &entities.Satellite{}
}

type fakeSatelliteMeasureDB8 struct {
}

// GetByNumber comment
func (db *fakeSatelliteMeasureDB8) GetByNumber(num int) entities.SatelliteMeasure {
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

func (db *fakeSatelliteMeasureDB8) Save(satelliteMeasure entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB8) ClearAll()                                       {}
func (db *fakeSatelliteMeasureDB8) FindByName(name string) *entities.SatelliteMeasure {
	return &entities.SatelliteMeasure{}
}
func (db *fakeSatelliteMeasureDB8) Update(satellite entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB8) Count() int {
	return 3
}

func TestSaveSatellitesMeasuresNotSavedInvalidDistance(t *testing.T) {

	// Not valid Satellites Measures - sat3
	fakeBodyRequest := entities.SatellitesRequest{Satellites: []entities.SatelliteRequest{
		entities.SatelliteRequest{Name: "sat1", Distance: 1.0, Message: []string{"", "", "", "D"}},
		entities.SatelliteRequest{Name: "sat2", Distance: 1.0, Message: []string{"A", "", "C", ""}},
		entities.SatelliteRequest{Name: "sat3", Distance: 0.0, Message: []string{"", "B", "C", ""}}}}

	fakeDbSatellite := &fakeSatelliteDB8{}
	fakeDbSatelliteMeasure := &fakeSatelliteMeasureDB8{}

	message, err := SaveSatellitesMeasures(fakeDbSatellite, fakeDbSatelliteMeasure, fakeBodyRequest)

	if message != "Not valid Satellites Measures - sat3" && !err {
		t.Errorf("SaveSatellites Error - Message %s.", message)
	}

}
