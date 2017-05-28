// Package treesort は、二分木による挿入ソートを提供します。
package treesort

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() (result string) {
	values := appendValues([]int{}, t)

	strs := []string{}
	for _, v := range values {
		strs = append(strs, fmt.Sprint(v))
	}
	return strings.Join(strs, ", ")
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
