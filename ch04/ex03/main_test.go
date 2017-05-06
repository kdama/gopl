package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		a, acopy *[6]int
		want     [6]int
	}{
		{&[...]int{0, 1, 2, 3, 4, 5}, &[...]int{0, 1, 2, 3, 4, 5}, [...]int{5, 4, 3, 2, 1, 0}},
	}

	for _, test := range tests {
		reverse(test.a)
		if *test.a != test.want {
			t.Errorf("reverse(%v) = %v, want %v", *test.acopy, *test.a, test.want)
		}
	}
}
