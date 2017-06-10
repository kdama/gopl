package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	var tests = []struct {
		in   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{0}, 1},
		{[]int{1}, 1},
		{[]int{1, 2, 42, 255, 256}, 5},
	}

	for _, test := range tests {
		intset := &IntSet{}
		for _, num := range test.in {
			intset.Add(num)
		}
		if got := intset.Len(); got != test.want {
			t.Errorf("Len of %v = %d, want %d", intset, got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		in     []int
		remove int
		check  int
		has    bool
	}{
		{[]int{}, 0, 0, false},
		{[]int{0}, 1, 0, true},
		{[]int{0}, 1, 1, false},
		{[]int{1}, 1, 0, false},
		{[]int{1}, 1, 1, false},
		{[]int{1, 2, 42, 255}, 1, 1, false},
		{[]int{1, 2, 42, 255}, 1, 2, true},
		{[]int{1, 2, 42, 255}, 1, 42, true},
		{[]int{1, 2, 42, 255}, 1, 255, true},
		{[]int{1, 2, 42, 255}, 1, 256, false},
		{[]int{1, 2, 42, 255}, 255, 1, true},
		{[]int{1, 2, 42, 255}, 255, 2, true},
		{[]int{1, 2, 42, 255}, 255, 42, true},
		{[]int{1, 2, 42, 255}, 255, 255, false},
		{[]int{1, 2, 42, 255}, 255, 256, false},
		{[]int{1, 2, 42, 255}, 256, 1, true},
		{[]int{1, 2, 42, 255}, 256, 2, true},
		{[]int{1, 2, 42, 255}, 256, 42, true},
		{[]int{1, 2, 42, 255}, 256, 255, true},
		{[]int{1, 2, 42, 255}, 256, 256, false},
	}

	for _, test := range tests {
		intset := &IntSet{}
		for _, num := range test.in {
			intset.Add(num)
		}
		intset.Remove(test.remove)
		if got := intset.Has(test.check); got != test.has {
			t.Errorf("(%v).Remove(%d) removes? = %t, want %t", test.in, test.remove, got, test.has)
		}
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		in  []int
		len int
	}{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{1}, 0},
		{[]int{1, 2, 42, 255, 256}, 0},
	}

	for _, test := range tests {
		intset := &IntSet{}
		for _, num := range test.in {
			intset.Add(num)
		}
		intset.Clear()
		if got := intset.Len(); got != test.len {
			t.Errorf("(%v).Len() is %d, want %d", test.in, got, test.len)
		}
		for _, num := range test.in {
			if intset.Has(num) {
				t.Errorf("(%v).Clear() does not clear %d", test.in, num)
			}
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		in []int
	}{
		{[]int{}},
		{[]int{0}},
		{[]int{1}},
		{[]int{1, 2, 42, 255, 256}},
	}

	for _, test := range tests {
		intset := &IntSet{}
		for _, num := range test.in {
			intset.Add(num)
		}
		got := intset.Copy()
		if &intset == &got {
			t.Errorf("(%v).Copy() returns itself", intset)
		}
		if &intset.words == &got.words {
			t.Errorf("(%v).Copy() returns 'words' of itself", intset)
		}
		if got.Len() != intset.Len() {
			t.Errorf("(%v).Len() is %d, want %d", test.in, got.Len(), intset.Len())
		}
		for _, num := range test.in {
			if !got.Has(num) {
				t.Errorf("(%v).Copy() does not copy %d", intset, num)
			}
		}
	}
}
