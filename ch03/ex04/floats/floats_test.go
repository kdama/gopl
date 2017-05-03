package floats

import (
	"math"
	"testing"
)

func TestIsFinite(t *testing.T) {
	var tests = []struct {
		f    float64
		want bool
	}{
		{0, true},
		{1, true},
		{-1, true},
		{1e100, true},
		{1e-100, true},
		{math.Inf(0), false},
		{math.Inf(-1), false},
		{math.NaN(), false},
	}

	for _, test := range tests {
		got := IsFinite(test.f)
		if got != test.want {
			t.Errorf("IsFinite(%f) = %t, want %t", test.f, got, test.want)
		}
	}
}
