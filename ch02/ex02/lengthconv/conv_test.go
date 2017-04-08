package lengthconv

import (
	"math"
	"testing"
)

const EPSILON = 1e-10

func equalsFloat64(x, y float64) bool {
	return math.Abs(x-y) < EPSILON
}

func TestFootToMeter(t *testing.T) {
	var tests = []struct {
		foot Foot
		want Meter
	}{
		{0, 0},
		{1, 1 / 0.3048},
		{-1, -1 / 0.3048},
		{0.3048, 1},
		{-0.3048, -1},
	}

	for _, test := range tests {
		if got := FootToMeter(test.foot); !equalsFloat64(float64(got), float64(test.want)) {
			t.Errorf("FootToMeter(%f) = %f, want %f", test.foot, got, test.want)
		}
	}
}

func TestMeterToFoot(t *testing.T) {
	var tests = []struct {
		meter Meter
		want  Foot
	}{
		{0, 0},
		{1 / 0.3048, 1},
		{-1 / 0.3048, -1},
		{1, 0.3048},
		{-1, -0.3048},
	}

	for _, test := range tests {
		if got := MeterToFoot(test.meter); !equalsFloat64(float64(got), float64(test.want)) {
			t.Errorf("MeterToFoot(%f) = %f, want %f", test.meter, got, test.want)
		}
	}
}
