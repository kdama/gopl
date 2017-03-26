package main

import (
	"testing"
)

func TestAddIndex(t *testing.T) {
	input := []string{"a", "b", "c"}
	actual := addIndex(input)
	expected := []string{
		"0 a",
		"1 b",
		"2 c",
	}
	for i := range input {
		if actual[i] != expected[i] {
			t.Errorf(`addIndex(%q)[%d] = %q, want %q`, input, i, actual[i], expected[i])
		}
	}
}
