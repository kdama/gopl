package main

import (
	"testing"
)

func TestJoinWithSpace(t *testing.T) {
	input := []string{"a", "b", "c"}
	actual := joinWithSpace(input)
	expected := "a b c"
	if actual != expected {
		t.Errorf(`joinWithSpace(%q) = %q, want %q`, input, actual, expected)
	}
}
