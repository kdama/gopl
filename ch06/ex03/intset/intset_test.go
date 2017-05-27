package intset

import (
	"testing"
)

func TestIntersectWith(t *testing.T) {
	var tests = []struct {
		s, t, want []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{1}, []int{2}, []int{}},
		{[]int{1, 2, 42}, []int{3, 4, 42}, []int{42}},
	}

	for _, test := range tests {
		intsetS := &IntSet{}
		for _, num := range test.s {
			intsetS.Add(num)
		}
		intsetT := &IntSet{}
		for _, num := range test.t {
			intsetT.Add(num)
		}
		intsetS.IntersectWith(intsetT)
		if got := intsetS.Len(); got != len(test.want) {
			t.Errorf("(%v).IntersectWith(%v).Len() = %d, want %d", test.s, test.t, got, len(test.want))
		}
		for _, want := range test.want {
			if !intsetS.Has(want) {
				t.Errorf("(%v).IntersectWith(%v) does not have %d", test.s, test.t, want)
			}
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	var tests = []struct {
		s, t, want []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{0}, []int{0}, []int{}},
		{[]int{1}, []int{2}, []int{1}},
		{[]int{1, 2, 42}, []int{3, 4, 42}, []int{1, 2}},
	}

	for _, test := range tests {
		intsetS := &IntSet{}
		for _, num := range test.s {
			intsetS.Add(num)
		}
		intsetT := &IntSet{}
		for _, num := range test.t {
			intsetT.Add(num)
		}
		intsetS.DifferenceWith(intsetT)
		if got := intsetS.Len(); got != len(test.want) {
			t.Errorf("(%v).DifferenceWith(%v).Len() = %d, want %d", test.s, test.t, got, len(test.want))
		}
		for _, want := range test.want {
			if !intsetS.Has(want) {
				t.Errorf("(%v).DifferenceWith(%v) does not have %d", test.s, test.t, want)
			}
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	var tests = []struct {
		s, t, want []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{0}, []int{0}, []int{}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{1, 2, 42}, []int{3, 4, 42}, []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		intsetS := &IntSet{}
		for _, num := range test.s {
			intsetS.Add(num)
		}
		intsetT := &IntSet{}
		for _, num := range test.t {
			intsetT.Add(num)
		}
		intsetS.SymmetricDifference(intsetT)
		if got := intsetS.Len(); got != len(test.want) {
			t.Errorf("(%v).SymmetricDifference(%v).Len() = %d, want %d", test.s, test.t, got, len(test.want))
		}
		for _, want := range test.want {
			if !intsetS.Has(want) {
				t.Errorf("(%v).SymmetricDifference(%v) does not have %d", test.s, test.t, want)
			}
		}
	}
}
