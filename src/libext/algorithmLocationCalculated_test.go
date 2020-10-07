package libext

import (
	"testing"
)

func TestLocation(t *testing.T) {
	location := Location{}
	x, y, err := location.Calculate(1.0, 0.0, 2.0, 0.0, 3.0, 0.0, 1.4142135623, 1.0, 1.4142135623)
	if err {
		t.Errorf("Position Error - Distances can recover a position. x=%f , y=%f.", x, y)
	}
}
