package lengthconv

import (
	"testing"
)

func TestFootString(t *testing.T) {
	var tests = []struct {
		foot Foot
		want string
	}{
		{0, "0ft"},
		{1, "1ft"},
		{2.34, "2.34ft"},
		{-5, "-5ft"},
		{-6.78, "-6.78ft"},
	}

	for _, test := range tests {
		if got := test.foot.String(); got != test.want {
			t.Errorf("foot(%f).String() = %s, want %s", test.foot, got, test.want)
		}
	}
}

func TestMeterString(t *testing.T) {
	var tests = []struct {
		meter Meter
		want  string
	}{
		{0, "0m"},
		{1, "1m"},
		{2.34, "2.34m"},
		{-5, "-5m"},
		{-6.78, "-6.78m"},
	}

	for _, test := range tests {
		if got := test.meter.String(); got != test.want {
			t.Errorf("meter(%f).String() = %s, want %s", test.meter, got, test.want)
		}
	}
}
