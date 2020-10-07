package services

import (
	"testing"

	"appml/src/domain/entities"
)

type fakeSatelliteMeasureDB6 struct {
}

// GetByNumber comment
func (db *fakeSatelliteMeasureDB6) GetByNumber(num int) entities.SatelliteMeasure {
	var satelliteMeasure entities.SatelliteMeasure
	switch num {
	case 0:
		satelliteMeasure = entities.SatelliteMeasure{Message: []string{"", "", "C", ""}}
	case 1:
		satelliteMeasure = entities.SatelliteMeasure{Message: []string{"", "A", "", "D"}}
	case 2:
		satelliteMeasure = entities.SatelliteMeasure{Message: []string{"", "B", "C", ""}}
	}
	return satelliteMeasure
}

func (db *fakeSatelliteMeasureDB6) Save(satelliteMeasure entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB6) ClearAll()                                       {}
func (db *fakeSatelliteMeasureDB6) FindByName(name string) *entities.SatelliteMeasure {
	return &entities.SatelliteMeasure{}
}
func (db *fakeSatelliteMeasureDB6) Update(satellite entities.SatelliteMeasure) {}

// Not enough Satellites Measures
func (db *fakeSatelliteMeasureDB6) Count() int {
	return 2
}

func TestRecoverMessageNotEnoughSatellitesMeasures(t *testing.T) {

	fakeDbSatelliteMeasure := &fakeSatelliteMeasureDB6{}

	_, err := RecoverMessage(fakeDbSatelliteMeasure)

	if !err {
		t.Errorf("Message Error - Not enough Satellites Measures.")
	}

}
