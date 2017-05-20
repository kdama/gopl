package main

import (
	"bytes"
	"testing"

	"strings"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	var tests = []struct {
		html string
		want string
	}{
		{"", ""},
		{"<p>foo</p>", "foo"},
		{"<div>foo<p>bar</p>baz</div>", "foobarbaz"},
		{"<p>foo</p><script>alert();</script>", "foo"},
		{"<head><style>body {}</style></head><body>foo</body>", "foo"},
	}

	for _, test := range tests {
		doc, err := html.Parse(strings.NewReader(test.html))
		if err != nil {
			t.Errorf("failed to parse HTML: %s", test.html)
		}

		var b bytes.Buffer
		visit(&b, doc)

		got := b.String()
		if b.String() != test.want {
			t.Errorf("visit of %q = %q, want %q", test.html, got, test.want)
		}
	}
}
