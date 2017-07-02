package crawl

import (
	"net/url"
	"testing"
)

func TestRelativeURL(t *testing.T) {
	tests := []struct {
		base, link string
		want       string
	}{
		{
			"http://example.com",
			"http://example.com",
			"./index.html",
		},
		{
			"http://example.com/",
			"http://example.com/",
			"./index.html",
		},
		{
			"http://example.com/foo.html",
			"http://example.com/bar.html",
			"./bar.html",
		},
		{
			"http://example.com/foo",
			"http://example.com/bar",
			"./bar/index.html",
		},
		{
			"http://example.com/foo",
			"http://example.com/foo/1/2/3.html",
			"./foo/1/2/3.html",
		},
		{
			"http://example.com/foo/1/2/3.html",
			"http://example.com/foo/4/5/6.pdf",
			"./../../../foo/4/5/6.pdf",
		},
		{
			"http://example.com/foo?x=1&y=2",
			"http://example.com/foo/1/2/3.html",
			"./foo/1/2/3.html",
		},
		{
			"http://example.com/foo/1/2/3.html?x=1&y=2",
			"http://example.com/foo/4/5/6.pdf?z=3",
			"./../../../foo/4/5/6.pdf?z=3",
		},
		{
			"http://example.com/foo",
			"https://example.com/foo",
			"https://example.com/foo",
		},
		{
			"http://example.COM/foo",
			"http://example.ORG/foo",
			"http://example.ORG/foo",
		},
	}
	for _, test := range tests {
		baseURL, err := url.Parse(test.base)
		if err != nil {
			t.Errorf("invalid base: %v", err)
		}
		linkURL, err := url.Parse(test.link)
		if err != nil {
			t.Errorf("invalid link: %v", err)
		}
		if got := relativeURL(baseURL, linkURL); got != test.want {
			t.Errorf("relativeURL(%s, %s) = %s, want %s", test.base, test.link, got, test.want)
		}
	}
}
