package main

import (
	"testing"
)

func TestCommaSigned(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"1234567890", "1,234,567,890"},
		{"1.0", "1.0"},
		{"1234.0", "1,234.0"},
		{"1234567890.0987654321", "1,234,567,890.0987654321"},
		{"-1.0", "-1.0"},
		{"-1234.0", "-1,234.0"},
		{"-1234567890.0987654321", "-1,234,567,890.0987654321"},
	}

	for _, test := range tests {
		if got := commaSigned(test.s); got != test.want {
			t.Errorf("commaSigned(%s) = %s, want %s", test.s, got, test.want)
		}
	}
}
