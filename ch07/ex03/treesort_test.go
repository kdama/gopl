package treesort

import (
	"testing"
)

func TestString(t *testing.T) {
	var tests = []struct {
		values []int
		want   string
	}{
		{[]int{}, ""},
		{[]int{0}, "0"},
		{[]int{1}, "1"},
		{[]int{0, 1}, "0, 1"},
		{[]int{-21, 0, 1, 2, 42}, "-21, 0, 1, 2, 42"},
		{[]int{0, 42, 2, 1, -21}, "-21, 0, 1, 2, 42"},
	}

	for _, test := range tests {
		var root *tree
		for _, v := range test.values {
			root = add(root, v)
		}
		if got := root.String(); got != test.want {
			t.Errorf("(tree of %v).String() = %q, want %q", test.values, got, test.want)
		}
	}
}
