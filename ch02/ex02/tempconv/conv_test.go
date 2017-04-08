package tempconv

import (
	"math"
	"testing"
)

const EPSILON = 1e-10

func equalsFloat64(x, y float64) bool {
	return math.Abs(x-y) < EPSILON
}

func TestCelsiusToFahrenheit(t *testing.T) {
	var tests = []struct {
		celsius Celsius
		want    Fahrenheit
	}{
		{32, 89.6},
		{212, 413.6},
		{-40, -40},
	}

	for _, test := range tests {
		if got := CToF(test.celsius); !equalsFloat64(float64(got), float64(test.want)) {
			t.Errorf("CToF(%f) = %f, want %f", test.celsius, got, test.want)
		}
	}
}

func TestFahrenheitTocelsius(t *testing.T) {
	var tests = []struct {
		fahrenheit Fahrenheit
		want       Celsius
	}{
		{89.6, 32},
		{413.6, 212},
		{-40, -40},
	}

	for _, test := range tests {
		if got := FToC(test.fahrenheit); !equalsFloat64(float64(got), float64(test.want)) {
			t.Errorf("FToC(%f) = %f, want %f", test.fahrenheit, got, test.want)
		}
	}
}
