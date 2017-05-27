package intset

import (
	"testing"
)

func equals(x, y []int) bool {
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
func TestElems(t *testing.T) {
	var tests = []struct {
		s []int
	}{
		{[]int{}},
		{[]int{0}},
		{[]int{1}},
		{[]int{1, 2, 42}},
	}

	for _, test := range tests {
		is := &IntSet{}
		for _, num := range test.s {
			is.Add(num)
		}
		if got := is.Elems(); !equals(got, test.s) {
			t.Errorf("(%v).Elems() = %v, want %v", test.s, got, test.s)
		}
	}
}
