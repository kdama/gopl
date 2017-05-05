package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"1234567890", "1,234,567,890"},
		{"あいうえお", "あい,うえお"},
	}

	for _, test := range tests {
		if got := comma(test.s); got != test.want {
			t.Errorf("comma(%s) = %s, want %s", test.s, got, test.want)
		}
	}
}
