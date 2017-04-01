package main

import (
	"testing"
)

func TestAppendPrefix(t *testing.T) {
	var tests = []struct {
		str  string
		want string
	}{
		{"gopl.io", "http://gopl.io"},
		{"http://gopl.io", "http://gopl.io"},
		{"", "http://"},
		{"http://", "http://"},
		{"https://gopl.io", "http://https://gopl.io"},
	}

	for _, test := range tests {
		if got := appendPrefix(test.str); got != test.want {
			t.Errorf("appendPrefix(%s) = %s, want %s", test.str, got, test.want)
		}
	}
}
