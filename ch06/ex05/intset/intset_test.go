package intset

import (
	"testing"
)

// uint 型を使った場合でも、ch06/ex01 のテストが同様に成功することを確認します。

func TestLen(t *testing.T) {
	var tests = []struct {
		in   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{0}, 1},
		{[]int{1}, 1},
		{[]int{1, 2, 42}, 3},
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
		{[]int{1, 2, 42}, 1, 42, true},
		{[]int{1, 2, 42}, 42, 42, false},
	}

	for _, test := range tests {
		intset := &IntSet{}
		for _, num := range test.in {
			intset.Add(num)
		}
		intset.Remove(test.remove)
		if got := intset.Has(test.check); got != test.has {
			t.Errorf("(%v).Remove(%d) removes? = %t, want %t", intset, test.remove, got, test.has)
		}
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		in []int
	}{
		{[]int{}},
		{[]int{0}},
		{[]int{1}},
		{[]int{1, 2, 42}},
	}

	for _, test := range tests {
		intset := &IntSet{}
		for _, num := range test.in {
			intset.Add(num)
		}
		intset.Clear()
		for _, num := range test.in {
			if intset.Has(num) {
				t.Errorf("(%v).Clear() does not clear %d", intset, num)
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
		{[]int{1, 2, 42}},
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
		for _, num := range test.in {
			if !intset.Has(num) {
				t.Errorf("(%v).Copy() does not copy %d", intset, num)
			}
		}
	}
}
