package main

import (
	"testing"

	"strings"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	var tests = []struct {
		html string
		want map[string]int
	}{
		{"<html><head></head><body></body></html>", map[string]int{
			"html": 1,
			"head": 1,
			"body": 1,
		}},
		{"<html><head></head><body><p>foo</p><P>bar</P></body></html>", map[string]int{
			"html": 1,
			"head": 1,
			"body": 1,
			"p":    2,
		}},
		{"<html><head></head><body><web-component/></body></html>", map[string]int{
			"html":          1,
			"head":          1,
			"body":          1,
			"web-component": 1,
		}},
	}

	for _, test := range tests {
		doc, err := html.Parse(strings.NewReader(test.html))
		if err != nil {
			t.Errorf("failed to parse HTML: %s", test.html)
		}

		got := visit(make(map[string]int), doc)
		if len(got) != len(test.want) {
			t.Errorf("len of visit of %q = %d, want %d", test.html, len(got), len(test.want))
		}
		for key := range got {
			if got[key] != test.want[key] {
				t.Errorf("%q of %q = %d, want %d", key, test.html, got[key], test.want[key])
			}
		}
	}
}
