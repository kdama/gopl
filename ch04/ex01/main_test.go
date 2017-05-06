package main

import (
	"testing"
)

func TestSha256PopCount(t *testing.T) {
	var tests = []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"a", "a", 0},
		{"あいうえお", "あいうえお", 0},
		{"a", "b", 126},
		{"あいうえお", "かきくけこ", 128},
	}

	for _, test := range tests {
		got := sha256PopCount(test.a, test.b)
		if got != test.want {
			t.Errorf("sha256PopCount(%s, %s) = %d, want %d", test.a, test.b, got, test.want)
		}
	}
}
