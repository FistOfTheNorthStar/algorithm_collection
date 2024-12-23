package geometry

import (
	"math"
	"testing"
)

func TestPointDistance(t *testing.T) {
	point1 := NewPoint(2, 3)
	point2 := NewPoint(5, 7)

	expected := 5.0
	result := point1.Distance(point2)

	if math.Abs(result-expected) > 1e-10 {
		t.Errorf("Distance calculation incorrect. Got %v, want %v", result, expected)
	}
}

func TestPointDistanceNilCheck(t *testing.T) {
	point1 := NewPoint(2, 3)
	var point2 *Point = nil

	defer func() {
		if r := recover(); r == nil {
			t.Error("nissing points no panic?")
		}
	}()

	point1.Distance(point2)
}
