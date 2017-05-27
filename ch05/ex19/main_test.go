package main

import (
	"testing"
)

func TestGet(t *testing.T) {
	var tests = []struct {
		in, want interface{}
	}{
		{nil, nil},
		{42, 42},
		{3.14159, 3.14159},
		{true, true},
		{"foo", "foo"},
	}

	for _, test := range tests {
		if got := get(test.in); got != test.want {
			t.Errorf("get(%v) = %v, want %v", test.in, got, test.want)
		}
	}
}
