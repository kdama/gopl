package main

import (
	"testing"
)

func TestAnagram(t *testing.T) {
	var tests = []struct {
		a    string
		b    string
		want bool
	}{
		{"a", "a", true},
		{"a", "b", false},
		{"abc", "cba", true},
		{"aabb", "abab", true},
		{"aAbBcC", "abcABC", true},
		{"aAbBcC", "abcABCD", false},
	}

	for _, test := range tests {
		if got := anagram(test.a, test.b); got != test.want {
			t.Errorf("anagram(%s, %s) = %t, want %t", test.a, test.b, got, test.want)
		}
	}
}
