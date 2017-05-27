package prettyhtml

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func equals(a, b *html.Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		fmt.Fprintf(os.Stderr, "nil and not nil: %v != %v\n", a, b)
		return false
	}
	if a.Type != b.Type {
		fmt.Fprintf(os.Stderr, "different type: %v != %v\n", a.Type, b.Type)
		return false
	}
	if a.Data != b.Data {
		fmt.Fprintf(os.Stderr, "different data: %q != %q\n", a.Data, b.Data)
		return false
	}
	if len(a.Attr) != len(b.Attr) {
		fmt.Fprintf(os.Stderr, "different len of attr: %v != %v\n", len(a.Attr), len(b.Attr))
		return false
	}
	for i := range a.Attr {
		if a.Attr[i].Key != b.Attr[i].Key || a.Attr[i].Val != b.Attr[i].Val {
			fmt.Fprintf(os.Stderr, "different attr: (%s, %s) != (%s, %s)\n", a.Attr[i].Key, b.Attr[i].Key, a.Attr[i].Val, b.Attr[i].Val)
			return false
		}
	}
	if !equals(a.FirstChild, b.FirstChild) {
		fmt.Fprintf(os.Stderr, "different child: %v != %v\n", a, b)
		return false
	}
	if !equals(a.NextSibling, b.NextSibling) {
		fmt.Fprintf(os.Stderr, "different sibling: %v != %v\n", a, b)
		return false
	}
	return true
}

func TestForEachNode(t *testing.T) {
	var tests = []struct {
		html string
	}{
		{`
<html>
  <head>
		<link rel='stylesheet' href='style.css'/>
  </head>
  <body>
    <p>
      Hello, world.
      <img src='src.png'></img>
      こんにちは世界。
    </p>
		<script src='script.js'></script>
  </body>
</html>
`},
	}

	for _, test := range tests {
		doc, err := html.Parse(strings.NewReader(test.html))
		if err != nil {
			t.Errorf("failed to parse HTML: %s", test.html)
		}

		var parsedBuffer, reparsedBuffer bytes.Buffer

		forEachNode(&parsedBuffer, doc, startElement, endElement)
		parsed, err := html.Parse(&parsedBuffer)

		forEachNode(&reparsedBuffer, parsed, startElement, endElement)
		reparsed, err := html.Parse(&reparsedBuffer)

		// パースの結果と、パースの結果をパースした結果とが、等しいことを調べます。
		if !equals(parsed, reparsed) {
			t.Errorf("Parse(%q) != Parse(Parse(%q))", test.html, test.html)
		}
	}
}
