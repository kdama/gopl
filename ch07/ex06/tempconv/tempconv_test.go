package tempconv

import (
	"math"
	"testing"
)

const Epsilon = 1e-10

func equalsFloat64(x, y float64) bool {
	return math.Abs(x-y) < Epsilon
}

func TestKelvinString(t *testing.T) {
	var tests = []struct {
		kelvin Kelvin
		want   string
	}{
		{0, "0K"},
		{1, "1K"},
		{2.34, "2.34K"},
		{-5, "-5K"},
		{-6.78, "-6.78K"},
	}

	for _, test := range tests {
		if got := test.kelvin.String(); got != test.want {
			t.Errorf("Kelvin(%f).String() = %s, want %s", test.kelvin, got, test.want)
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct {
		celsius Celsius
		want    Kelvin
	}{
		{0, 273.15},
		{1, 274.15},
		{-1, 272.15},
		{273.15, 546.3},
		{-273.15, 0},
	}

	for _, test := range tests {
		if got := CToK(test.celsius); !equalsFloat64(float64(got), float64(test.want)) {
			t.Errorf("CToK(%f) = %f, want %f", test.celsius, got, test.want)
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct {
		kelvin Kelvin
		want   Celsius
	}{
		{273.15, 0},
		{274.15, 1},
		{272.15, -1},
		{546.3, 273.15},
		{0, -273.15},
	}

	for _, test := range tests {
		if got := KToC(test.kelvin); !equalsFloat64(float64(got), float64(test.want)) {
			t.Errorf("KToC(%f) = %f, want %f", test.kelvin, got, test.want)
		}
	}
}

func TestFToK(t *testing.T) {
	var tests = []struct {
		fahrenheit Fahrenheit
		want       Kelvin
	}{
		{32, 273.15},
		{212, 373.15},
		{-40, 233.15},
		{-459.67, 0},
		{-461.47, -1},
	}

	for _, test := range tests {
		if got := FToK(test.fahrenheit); !equalsFloat64(float64(got), float64(test.want)) {
			t.Errorf("FToK(%f) = %f, want %f", test.fahrenheit, got, test.want)
		}
	}
}

func TestKToF(t *testing.T) {
	var tests = []struct {
		kelvin Kelvin
		want   Fahrenheit
	}{
		{273.15, 32},
		{373.15, 212},
		{233.15, -40},
		{0, -459.67},
		{-1, -461.47},
	}

	for _, test := range tests {
		if got := KToF(test.kelvin); !equalsFloat64(float64(got), float64(test.want)) {
			t.Errorf("KToF(%f) = %f, want %f", test.kelvin, got, test.want)
		}
	}
}
