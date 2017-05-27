package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		vals []int
		want int
		err  bool
	}{
		{[]int{}, 0, true},
		{[]int{1}, 1, false},
		{[]int{1, 2, 3}, 3, false},
		{[]int{-1, -2, -3}, -1, false},
	}

	for _, test := range tests {
		got, err := max(test.vals...)
		if err != nil && !test.err {
			t.Errorf("Expect no error for max(%v), but there is error: %v", test.vals, err)
		} else if err == nil && test.err {
			t.Errorf("Expect error for max(%v), but there is no error", test.vals)
		} else if err == nil && got != test.want {
			t.Errorf("max(%v) = %d, want %d", test.vals, got, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		vals []int
		want int
		err  bool
	}{
		{[]int{}, 0, true},
		{[]int{1}, 1, false},
		{[]int{1, 2, 3}, 1, false},
		{[]int{-1, -2, -3}, -3, false},
	}

	for _, test := range tests {
		got, err := min(test.vals...)
		if err != nil && !test.err {
			t.Errorf("Expect no error for min(%v), but there is error: %v", test.vals, err)
		} else if err == nil && test.err {
			t.Errorf("Expect error for min(%v), but there is no error", test.vals)
		} else if err == nil && got != test.want {
			t.Errorf("min(%v) = %d, want %d", test.vals, got, test.want)
		}
	}
}

func TestAlternativeMax(t *testing.T) {
	var tests = []struct {
		val    int
		others []int
		want   int
	}{
		{1, []int{}, 1},
		{1, []int{2, 3}, 3},
		{-1, []int{-2, -3}, -1},
	}

	for _, test := range tests {
		if got := alternativeMax(test.val, test.others...); got != test.want {
			t.Errorf("alternativeMax(%d, %v) = %d, want %d", test.val, test.others, got, test.want)
		}
	}
}

func TestAlternativeMin(t *testing.T) {
	var tests = []struct {
		val    int
		others []int
		want   int
	}{
		{1, []int{}, 1},
		{1, []int{2, 3}, 1},
		{-1, []int{-2, -3}, -3},
	}

	for _, test := range tests {
		if got := alternativeMin(test.val, test.others...); got != test.want {
			t.Errorf("alternativeMin(%d, %v) = %d, want %d", test.val, test.others, got, test.want)
		}
	}
}
