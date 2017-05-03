package main

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
		got := isFinite(test.f)
		if got != test.want {
			t.Errorf("isFinite(%f) = %t, want %t", test.f, got, test.want)
		}
	}
}

func TestGetColor(t *testing.T) {
	var tests = []struct {
		height    float64
		maxHeight float64
		minHeight float64
		want      string
	}{
		{0, 1, 0, "#0000ff"},
		{0.5, 1, 0, "#7f0080"},
		{1, 1, 0, "#ff0000"},
	}

	for _, test := range tests {
		got := getColor(test.height, test.maxHeight, test.minHeight)
		if got != test.want {
			t.Errorf("getColor(%f, %f, %f) = %s, want %s", test.height, test.maxHeight, test.minHeight, got, test.want)
		}
	}
}
