package main

import (
	"testing"
)

func TestJoin(t *testing.T) {
	var tests = []struct {
		sep  string
		vals []string
		want string
	}{
		{"", []string{}, ""},
		{"", []string{"foo"}, "foo"},
		{" ", []string{"foo", "bar", "baz"}, "foo bar baz"},

		{",", []string{}, ""},
		{",", []string{"foo"}, "foo"},
		{",", []string{"foo", "bar", "baz"}, "foo,bar,baz"},
	}

	for _, test := range tests {
		if got := join(test.sep, test.vals...); got != test.want {
			t.Errorf("join(%q, %q) = %q, want %q", test.sep, test.vals, got, test.want)
		}
	}
}
