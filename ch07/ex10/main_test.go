package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		ints []int
		want bool
	}{
		{[]int{}, true},
		{[]int{0}, true},
		{[]int{1}, true},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{1, 2, 3, 3, 2, 1}, true},
		{[]int{0, 1}, false},
		{[]int{1, 2, 3, 3, 1}, false},
		{[]int{1, 2, 3, 3, 3, 1}, false},
	}

	for _, test := range tests {
		if got := IsPalindrome(sortInt(test.ints)); got != test.want {
			t.Errorf("IsPalindrome(%v) = %t, want %t", test.ints, got, test.want)
		}
	}
}
