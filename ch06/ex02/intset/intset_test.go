package intset

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	var tests = []struct {
		in  []int
		add []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{2}},
		{[]int{1, 2, 42}, []int{3, 4, 42}},
	}

	for _, test := range tests {
		intset := &IntSet{}
		for _, num := range test.in {
			intset.Add(num)
		}
		intset.AddAll(test.add...)
		for _, i := range test.in {
			if !intset.Has(i) {
				t.Errorf("(%v).AddAll(%v) does not have %d", test.in, test.add, i)
			}
		}
		for _, i := range test.add {
			if !intset.Has(i) {
				t.Errorf("(%v).AddAll(%v) does not add %d", test.in, test.add, i)
			}
		}
	}
}
