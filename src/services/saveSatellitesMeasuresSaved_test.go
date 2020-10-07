package services

import (
	"testing"

	"appml/src/domain/entities"
)

type fakeSatelliteDB7 struct {
}

// GetByNumber comment
func (db *fakeSatelliteDB7) GetByNumber(num int) entities.Satellite {
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
func (db *fakeSatelliteDB7) Save(satellite entities.Satellite) {}
func (db *fakeSatelliteDB7) FindByName(name string) *entities.Satellite {
	return &entities.Satellite{}
}

type fakeSatelliteMeasureDB7 struct {
}

// GetByNumber comment
func (db *fakeSatelliteMeasureDB7) GetByNumber(num int) entities.SatelliteMeasure {
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

func (db *fakeSatelliteMeasureDB7) Save(satelliteMeasure entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB7) ClearAll()                                       {}
func (db *fakeSatelliteMeasureDB7) FindByName(name string) *entities.SatelliteMeasure {
	return &entities.SatelliteMeasure{}
}
func (db *fakeSatelliteMeasureDB7) Update(satellite entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB7) Count() int {
	return 3
}

func TestSaveSatellitesMeasuresSaved(t *testing.T) {

	fakeBodyRequest := entities.SatellitesRequest{Satellites: []entities.SatelliteRequest{
		entities.SatelliteRequest{Name: "sat1", Distance: 1.0, Message: []string{"", "", "", "D"}},
		entities.SatelliteRequest{Name: "sat2", Distance: 1.0, Message: []string{"A", "", "C", ""}},
		entities.SatelliteRequest{Name: "sat3", Distance: 1.0, Message: []string{"", "B", "C", ""}}}}

	fakeDbSatellite := &fakeSatelliteDB7{}
	fakeDbSatelliteMeasure := &fakeSatelliteMeasureDB7{}

	message, err := SaveSatellitesMeasures(fakeDbSatellite, fakeDbSatelliteMeasure, fakeBodyRequest)

	if message != "Satellites loaded" && err {
		t.Errorf("SaveSatellites Error - Satellites loaded %s.", message)
	}

}
