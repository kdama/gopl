package massconv

import (
	"math"
	"testing"
)

const EPSILON = 1e-10

func equalsFloat64(x, y float64) bool {
	return math.Abs(x-y) < EPSILON
}

func TestPoundToKilogram(t *testing.T) {
	var tests = []struct {
		pound Pound
		want  Kilogram
	}{
		{0, 0},
		{1, 1 / 0.45359237},
		{-1, -1 / 0.45359237},
		{0.45359237, 1},
		{-0.45359237, -1},
	}

	for _, test := range tests {
		if got := PoundToKilogram(test.pound); !equalsFloat64(float64(got), float64(test.want)) {
			t.Errorf("PoundToKilogram(%f) = %f, want %f", test.pound, got, test.want)
		}
	}
}

func TestKilogramToPound(t *testing.T) {
	var tests = []struct {
		kilogram Kilogram
		want     Pound
	}{
		{0, 0},
		{1 / 0.45359237, 1},
		{-1 / 0.45359237, -1},
		{1, 0.45359237},
		{-1, -0.45359237},
	}

	for _, test := range tests {
		if got := KilogramToPound(test.kilogram); !equalsFloat64(float64(got), float64(test.want)) {
			t.Errorf("KilogramToPound(%f) = %f, want %f", test.kilogram, got, test.want)
		}
	}
}
