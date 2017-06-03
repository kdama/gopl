package main

import (
	"testing"
)

func TestPushAsSet(t *testing.T) {
	var tests = []struct {
		slice []string
		s     string
		want  []string
	}{
		{[]string{}, "", []string{""}},
		{[]string{}, "a", []string{"a"}},
		{[]string{"a"}, "a", []string{"a"}},
		{[]string{"a"}, "b", []string{"a", "b"}},
		{[]string{"a", "b", "c"}, "a", []string{"b", "c", "a"}},
		{[]string{"a", "b", "c"}, "b", []string{"a", "c", "b"}},
		{[]string{"a", "b", "c"}, "c", []string{"a", "b", "c"}},
		{[]string{"a", "b", "c"}, "d", []string{"a", "b", "c", "d"}},
	}

	for _, test := range tests {
		if got := pushAsSet(test.slice, test.s); !equals(got, test.want) {
			t.Errorf("pushAsSet(%v, %q) = %v, want %v", test.slice, test.s, got, test.want)
		}
	}
}

func equals(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
