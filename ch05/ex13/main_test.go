package main

import (
	"testing"
)

func TestSameDomain(t *testing.T) {
	var tests = []struct {
		x, y      string
		want, err bool
	}{
		{"", "", true, false},
		{"a", "a", true, false},
		{"あ", "あ", true, false},
		{"http://example.com/foo", "http://example.com/foo", true, false},
		{"http://example.com/foo", "https://example.com/foo", true, false},
		{"http://example.com/foo", "//example.com/foo", true, false},
		{"http://example.com/foo", "http://example.com?foo=bar", true, false},
		{"http://example.com/foo", "http://example.com/bar", true, false},
		{"http://example.com/foo", "http://sub.example.com/foo", false, false},
		{"http://example.com/foo", "http://example.co.jp/foo", false, false},
	}

	for _, test := range tests {
		got, err := sameDomain(test.x, test.y)
		if err != nil && !test.err {
			t.Errorf("Expect no error for getDomainName(%q, %q), but there is error: %v", test.x, test.y, err)
		} else if err == nil && test.err {
			t.Errorf("Expect error for getDomainName(%q, %q), but there is no error", test.x, test.y)
		} else if err == nil && got != test.want {
			t.Errorf("getDomainName(%q, %q) = %t, want %t", test.x, test.y, got, test.want)
		}
	}
}
