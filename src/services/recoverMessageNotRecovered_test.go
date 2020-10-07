package services

import (
	"testing"

	"appml/src/domain/entities"
)

type fakeSatelliteMeasureDB5 struct {
}

// Messages cant return a valid position.
// GetByNumber comment
func (db *fakeSatelliteMeasureDB5) GetByNumber(num int) entities.SatelliteMeasure {
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

func (db *fakeSatelliteMeasureDB5) Save(satelliteMeasure entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB5) ClearAll()                                       {}
func (db *fakeSatelliteMeasureDB5) FindByName(name string) *entities.SatelliteMeasure {
	return &entities.SatelliteMeasure{}
}
func (db *fakeSatelliteMeasureDB5) Update(satellite entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB5) Count() int {
	return 3
}

func TestRecoverMessageNotRecovered(t *testing.T) {

	fakeDbSatelliteMeasure := &fakeSatelliteMeasureDB5{}

	_, err := RecoverMessage(fakeDbSatelliteMeasure)

	if !err {
		t.Errorf("Message Error - Distances cant return a valid position.")
	}

}
