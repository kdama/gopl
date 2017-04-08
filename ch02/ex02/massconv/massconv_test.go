package massconv

import (
	"testing"
)

func TestPoundString(t *testing.T) {
	var tests = []struct {
		pound Pound
		want  string
	}{
		{0, "0lb"},
		{1, "1lb"},
		{2.34, "2.34lb"},
		{-5, "-5lb"},
		{-6.78, "-6.78lb"},
	}

	for _, test := range tests {
		if got := test.pound.String(); got != test.want {
			t.Errorf("pound(%f).String() = %s, want %s", test.pound, got, test.want)
		}
	}
}

func TestKilogramString(t *testing.T) {
	var tests = []struct {
		kilogram Kilogram
		want     string
	}{
		{0, "0kg"},
		{1, "1kg"},
		{2.34, "2.34kg"},
		{-5, "-5kg"},
		{-6.78, "-6.78kg"},
	}

	for _, test := range tests {
		if got := test.kilogram.String(); got != test.want {
			t.Errorf("kilogram(%f).String() = %s, want %s", test.kilogram, got, test.want)
		}
	}
}
