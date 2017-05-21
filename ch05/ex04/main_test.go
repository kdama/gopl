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
		{"<p>foo</p>", []string{}},
		{"<img src='http://example.com/src.png'>", []string{"http://example.com/src.png"}},
		{"<link rel='stylesheet' href='/src.css'>", []string{"/src.css"}},
		{"<script src='../../../src.js'>", []string{"../../../src.js"}},
		{"<video src='src.mp4' poster='poster.jpg'>", []string{"src.mp4", "poster.jpg"}},
		{"<a href='./href.html' onClick='function(){ window.location=\"http://example.com/\" }'>", []string{"./href.html"}},
		{"<p>foo</p><img src='src.png'><p>bar</p>", []string{"src.png"}},
	}

	for _, test := range tests {
		doc, err := html.Parse(strings.NewReader(test.html))
		if err != nil {
			t.Errorf("failed to parse HTML: %s", test.html)
		}

		got := visit([]string{}, doc)

		if len(got) != len(test.want) {
			t.Errorf("visit of %q = %q, want %q", test.html, got, test.want)
		} else {
			for i := range got {
				if got[i] != test.want[i] {
					t.Errorf("visit of %q [%d] = %q, want %q", test.html, i, got[i], test.want[i])
				}
			}
		}
	}
}
