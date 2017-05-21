package main

import (
	"testing"

	"strings"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	var tests = []struct {
		html string
		want []string
	}{
		{"", []string{}},
		{"<html></html>", []string{}},
		{"<html><a href='foo'>f</a></html>", []string{"foo"}},
		{"<html><a href='foo'>f</a><a href='日本語'>日</a></html>", []string{"foo", "日本語"}},
	}

	for _, test := range tests {
		doc, err := html.Parse(strings.NewReader(test.html))
		if err != nil {
			t.Errorf("failed to parse HTML: %s", test.html)
		}

		got := visit(nil, doc)
		if len(got) != len(test.want) {
			t.Errorf("len of visit of %q = %d, want %d", test.html, len(got), len(test.want))
		}
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("visit of %q = %q, want %q", test.html, got[i], test.want[i])
			}
		}
	}
}
