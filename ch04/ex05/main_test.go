package main

import (
	"testing"
)

func equals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRemoveDup(t *testing.T) {
	var tests = []struct {
		s, scopy []string
		want     []string
	}{
		{[]string{}, []string{}, []string{}},
		{[]string{"A"}, []string{"A"}, []string{"A"}},
		{[]string{"A", "A"}, []string{"A", "A"}, []string{"A"}},
		{[]string{"A", "A", "B", "B", "A", "C"}, []string{"A", "A", "B", "B", "A", "C"}, []string{"A", "B", "A", "C"}},
	}

	for _, test := range tests {
		got := removeDup(test.s)
		if !equals(got, test.want) {
			t.Errorf("removeDup(%v) = %v, want %v", test.scopy, got, test.want)
		}
	}
}
