package tempconv

import (
	"testing"
)

func TestCelsiusString(t *testing.T) {
	var tests = []struct {
		celsius Celsius
		want    string
	}{
		{0, "0C"},
		{1, "1C"},
		{2.34, "2.34C"},
		{-5, "-5C"},
		{-6.78, "-6.78C"},
	}

	for _, test := range tests {
		if got := test.celsius.String(); got != test.want {
			t.Errorf("Celsius(%f).String() = %s, want %s", test.celsius, got, test.want)
		}
	}
}
