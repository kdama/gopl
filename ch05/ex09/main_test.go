package main

import (
	"testing"
)

func TestExpand(t *testing.T) {
	var tests = []struct {
		s    string
		f    func(string) string
		want string
	}{
		{"", double, ""},
		{"foo", double, "foo"},
		{"$foo", double, "foofoo"},
		{"$$foo", double, "$foo$foo"},
		{"$foo $foo", double, "foofoo foofoo"},
		{"$日本語", double, "日本語日本語"},
		{"＄日本語", double, "＄日本語"},
	}

	for _, test := range tests {
		if got := expand(test.s, test.f); got != test.want {
			t.Errorf("expand(%q, double) = %q, want %q", test.s, got, test.want)
		}
	}
}
