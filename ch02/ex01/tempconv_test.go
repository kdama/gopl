package tempconv

import (
	"testing"
)

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
