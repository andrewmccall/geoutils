package geoutils

import (
	"testing"
)

func TestGetDistance(t *testing.T) {

	lat1 := 53.75734328397976
	lon1 := -2.0180395172300716

	lat2 := 53.7492235165634
	lon2 := -2.070567899356706

	distance := CalculateDistance(lat1, lon1, lat2, lon2)

	if distance != 3568 {
		t.Errorf("Distance should be 3569m got %d", distance)
	}

	// should get the same backwards.
	distance2 := CalculateDistance(lat2, lon2, lat1, lon1)
	if distance != distance2 {
		t.Errorf("Distance should be internally consistent! 1->2 = %d 2->1 = %d", distance, distance2)
	}
}

func TestFloatEquals(t *testing.T) {
	if !floatEquals(1, 1, 0) {
		t.Errorf("Floats should be equal")
	}
}
