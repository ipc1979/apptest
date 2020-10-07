package services

import (
	"testing"

	"appml/src/domain/entities"
)

type fakeSatelliteMeasureDB4 struct {
}

// GetByNumber comment
func (db *fakeSatelliteMeasureDB4) GetByNumber(num int) entities.SatelliteMeasure {
	var satelliteMeasure entities.SatelliteMeasure
	switch num {
	case 0:
		satelliteMeasure = entities.SatelliteMeasure{Message: []string{"", "", "C", ""}}
	case 1:
		satelliteMeasure = entities.SatelliteMeasure{Message: []string{"A", "", "", "D"}}
	case 2:
		satelliteMeasure = entities.SatelliteMeasure{Message: []string{"", "B", "C", ""}}
	}
	return satelliteMeasure
}

func (db *fakeSatelliteMeasureDB4) Save(satelliteMeasure entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB4) ClearAll()                                       {}
func (db *fakeSatelliteMeasureDB4) FindByName(name string) *entities.SatelliteMeasure {
	return &entities.SatelliteMeasure{}
}
func (db *fakeSatelliteMeasureDB4) Update(satellite entities.SatelliteMeasure) {}
func (db *fakeSatelliteMeasureDB4) Count() int {
	return 3
}

func TestRecoverMessageRecovered(t *testing.T) {

	fakeDbSatelliteMeasure := &fakeSatelliteMeasureDB4{}

	message, err := RecoverMessage(fakeDbSatelliteMeasure)

	if err {
		t.Errorf("Message Error - Message Expected : A,B,C,D , Received %s.", message)
	}

}
