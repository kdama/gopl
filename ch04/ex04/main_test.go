package main

import (
	"testing"
)

func equals(a, b []int) bool {
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

func TestRotate(t *testing.T) {
	var tests = []struct {
		a, acopy []int
		n        int
		want     []int
	}{
		{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}, 0, []int{0, 1, 2, 3}},
		{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}, 1, []int{1, 2, 3, 0}},
		{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}, 2, []int{2, 3, 0, 1}},
		{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}, 3, []int{3, 0, 1, 2}},
		{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}, 4, []int{0, 1, 2, 3}},
		{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}, 400, []int{0, 1, 2, 3}},
	}

	for _, test := range tests {
		rotate(test.a, test.n)
		if !equals(test.a, test.want) {
			t.Errorf("reverse(%v, %d) = %v, want %v", test.acopy, test.n, test.a, test.want)
		}
	}
}
