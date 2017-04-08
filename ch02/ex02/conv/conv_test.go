package conv

import (
	"math"
	"testing"
)

const EPSILON = 1e-10

func equalsFloat64(x, y float64) bool {
	return math.Abs(x-y) < EPSILON
}

func TestParse(t *testing.T) {
	var tests = []struct {
		str       string
		wantValue float64
		wantUnit  string
		wantError error
	}{
		{"0a", 0, "a", nil},
		{"-1B", -1, "B", nil},
		{"12.345e10cde", 12.345e10, "cde", nil},
		{"-12.345e-10cde", -12.345e-10, "cde", nil},
	}

	for _, test := range tests {
		gotValue, gotUnit, gotError := Parse(test.str)
		if !equalsFloat64(float64(gotValue), float64(test.wantValue)) {
			t.Errorf("Parse(%s).value = %f, want %f", test.str, gotValue, test.wantValue)
		}
		if gotUnit != test.wantUnit {
			t.Errorf("Parse(%s).unit = %s, want %s", test.str, gotUnit, test.wantUnit)
		}
		if gotError != test.wantError {
			t.Errorf("Parse(%s).error = %s, want %s", test.str, gotError, test.wantError)
		}
	}
}

func TestConvert(t *testing.T) {
	var tests = []struct {
		value     float64
		unit      string
		wantFrom  string
		wantTo    string
		wantError error
	}{
		{0, "ft", "0ft", "0m", nil},
		{1, "m", "1m", "0.3048ft", nil},
		{32, "℃", "32°C", "89.6°F", nil},
		{212, "°F", "212°F", "100°C", nil},
	}

	for _, test := range tests {
		gotFrom, gotTo, gotError := Convert(test.value, test.unit)
		if gotFrom != test.wantFrom {
			t.Errorf("Convert(%f, %s).from = %s, want %s", test.value, test.unit, gotFrom, test.wantFrom)
		}
		if gotTo != test.wantTo {
			t.Errorf("Convert(%f, %s).to = %s, want %s", test.value, test.unit, gotTo, test.wantTo)
		}
		if gotError != test.wantError {
			t.Errorf("Convert(%f, %s).error = %s, want %s", test.value, test.unit, gotError, test.wantError)
		}
	}
}
