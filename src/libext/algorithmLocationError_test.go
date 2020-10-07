package libext

import (
	"testing"
)

func TestLocationError(t *testing.T) {
	location := Location{}

	x, y, err := location.Calculate(1.0, 0.0, 2.0, 0.0, 3.0, 0.0, 1.0, 1.0, 1.0)
	if !err {
		t.Errorf("Position Error - Distances cant recover a position. x=%f , y=%f", x, y)
	}
}
