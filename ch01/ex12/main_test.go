package main

import (
	"testing"
)

func TestParseFirstIntOrDefault(t *testing.T) {
	var tests = []struct {
		array        []string
		defaultValue int
		want         int
	}{
		{[]string{"1"}, 0, 1},
		{[]string{"2", "3", "4"}, 0, 2},
		{[]string{"-1"}, 5, -1},
		{[]string{"1.1"}, 6, 6},
		{[]string{"1e10"}, 7, 7},
		{[]string{"A"}, 8, 8},
		{[]string{}, 9, 9},
	}

	for _, test := range tests {
		if got := parseFirstIntOrDefault(test.array, test.defaultValue); got != test.want {
			t.Errorf("parseFirstIntOrDefault(%s, %d) = %d, want %d", test.array, test.defaultValue, got, test.want)
		}
	}
}

func TestParseFirstFloat64OrDefault(t *testing.T) {
	var tests = []struct {
		array        []string
		defaultValue float64
		want         float64
	}{
		{[]string{"1"}, 0, 1},
		{[]string{"2", "3", "4"}, 0, 2},
		{[]string{"-1"}, 5, -1},
		{[]string{"1.1"}, 6, 1.1},
		{[]string{"1e10"}, 7, 1e10},
		{[]string{"A"}, 8, 8},
		{[]string{}, 9, 9},
	}

	for _, test := range tests {
		if got := parseFirstFloat64OrDefault(test.array, test.defaultValue); got != test.want {
			t.Errorf("parseFirstFloat64OrDefault(%s, %f) = %f, want %f", test.array, test.defaultValue, got, test.want)
		}
	}
}
