package services

import (
	"testing"

	"appml/src/domain/entities"
)

type fakeSatelliteDB9 struct {
}

// GetByNumber comment
func (db *fakeSatelliteDB9) GetByNumber(num int) entities.Satellite {
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
func (db *fakeSatelliteDB9) Save(satellite entities.Satellite) {}
func (db *fakeSatelliteDB9) FindByName(name string) *entities.Satellite {
	if name == "sat4" {
		return nil
	}
	return &entities.Satellite{}
}

type fakeSatelliteMeasureDB9 struct {
}

// GetByNumber comment
func (db *fakeSatelliteMeasureDB9) GetByNumber(num int) entities.SatelliteMeasure {
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

func (db *fakeSatelliteMeasureDB9) Save(satelliteMeasure entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB9) ClearAll()                                       {}
func (db *fakeSatelliteMeasureDB9) FindByName(name string) *entities.SatelliteMeasure {
	return &entities.SatelliteMeasure{}
}
func (db *fakeSatelliteMeasureDB9) Update(satellite entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB9) Count() int {
	return 3
}

func TestSaveSatellitesMeasuresNotExistSatellite(t *testing.T) {

	fakeBodyRequest := entities.SatellitesRequest{Satellites: []entities.SatelliteRequest{
		entities.SatelliteRequest{Name: "sat1", Distance: 1.0, Message: []string{"", "", "", "D"}},
		entities.SatelliteRequest{Name: "sat2", Distance: 1.0, Message: []string{"A", "", "C", ""}},
		entities.SatelliteRequest{Name: "sat4", Distance: 1.0, Message: []string{"", "B", "C", ""}}}}

	fakeDbSatellite := &fakeSatelliteDB9{}
	fakeDbSatelliteMeasure := &fakeSatelliteMeasureDB9{}

	message, err := SaveSatellitesMeasures(fakeDbSatellite, fakeDbSatelliteMeasure, fakeBodyRequest)

	if message != "Not valid Satellites Measures - sat4" && err {
		t.Errorf("SaveSatellites Error - Not valid Satellites Measures %s.", message)
	}

}
