package main

import (
	"testing"
)

func TestAppendAsSet(t *testing.T) {
	var tests = []struct {
		set  []string
		str  string
		want []string
	}{
		{[]string{"a", "b", "c"}, "a", []string{"a", "b", "c"}},
		{[]string{"a", "b", "c"}, "b", []string{"a", "b", "c"}},
		{[]string{"a", "b", "c"}, "c", []string{"a", "b", "c"}},
		{[]string{"a", "b", "c"}, "d", []string{"a", "b", "c", "d"}},
		{[]string{}, "e", []string{"e"}},
	}

	for _, test := range tests {
		got := appendAsSet(test.set, test.str)
		for idx, gotValue := range got {
			if gotValue != test.want[idx] {
				t.Errorf("appendAsSet(%q, %q)[%d] = %q, want %q", test.set, test.str, idx, gotValue, test.want[idx])
			}
		}
	}
}

func TestIncludes(t *testing.T) {
	var tests = []struct {
		array []string
		str   string
		want  bool
	}{
		{[]string{"a", "b", "c"}, "a", true},
		{[]string{"a", "b", "c"}, "b", true},
		{[]string{"a", "b", "c"}, "c", true},
		{[]string{"a", "b", "c"}, "d", false},
		{[]string{}, "e", false},
	}

	for _, test := range tests {
		if got := includes(test.array, test.str); got != test.want {
			t.Errorf("includes(%q, %q) = %t, want %t", test.array, test.str, got, test.want)
		}
	}
}
