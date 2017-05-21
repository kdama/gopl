package main

import (
	"testing"

	"strings"

	"golang.org/x/net/html"
)

func TestCountWordsAndImages(t *testing.T) {
	var tests = []struct {
		html string
		want struct {
			words  int
			images int
		}
	}{
		{"", struct {
			words  int
			images int
		}{words: 0, images: 0}},
		{"<p>foo</p>", struct {
			words  int
			images int
		}{words: 1, images: 0}},
		{"<img src='http://example.com/src.png'/>", struct {
			words  int
			images int
		}{words: 0, images: 1}},
		{"<script>foo</script>", struct {
			words  int
			images int
		}{words: 0, images: 0}},
		{"<p>Hello, world.<img src='src.png'/>こんにちは<br>世界</p>", struct {
			words  int
			images int
		}{words: 4, images: 1}},
		{"<p>Hello,world.<img src='src.png'/>こんにちは世界</p>", struct {
			words  int
			images int
		}{words: 2, images: 1}},
	}

	for _, test := range tests {
		doc, err := html.Parse(strings.NewReader(test.html))
		if err != nil {
			t.Errorf("failed to parse HTML: %s", test.html)
		}

		words, images := countWordsAndImages(doc)

		if words != test.want.words {
			t.Errorf("words of countWordsAndImages of %q = %d, want %d", test.html, words, test.want.words)
		}
		if images != test.want.images {
			t.Errorf("images of countWordsAndImages of %q = %d, want %d", test.html, images, test.want.images)
		}
	}
}
