package tempconv

import (
	"testing"
)

func TestCelsiusString(t *testing.T) {
	var tests = []struct {
		celsius Celsius
		want  string
	}{
		{0, "0°C"},
		{1, "1°C"},
		{2.34, "2.34°C"},
		{-5, "-5°C"},
		{-6.78, "-6.78°C"},
	}

	for _, test := range tests {
		if got := test.celsius.String(); got != test.want {
			t.Errorf("Celsius(%f).String() = %s, want %s", test.celsius, got, test.want)
		}
	}
}

func TestFahrenheitString(t *testing.T) {
	var tests = []struct {
		fahrenheit Fahrenheit
		want     string
	}{
		{0, "0°F"},
		{1, "1°F"},
		{2.34, "2.34°F"},
		{-5, "-5°F"},
		{-6.78, "-6.78°F"},
	}

	for _, test := range tests {
		if got := test.fahrenheit.String(); got != test.want {
			t.Errorf("Fahrenheit(%f).String() = %s, want %s", test.fahrenheit, got, test.want)
		}
	}
}
