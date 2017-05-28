// ch07/ex04 は、文字列から読み込む io.Reader の実装によって HTML をパースします。
package main

import (
	"io"
	"log"
	"os"

	"github.com/kdama/gopl/ch05/ex07/prettyhtml"
	"golang.org/x/net/html"
)

// StringReader は、文字列を読み込む io.Reader の実装です。
type StringReader string

func (s *StringReader) Read(p []byte) (int, error) {
	copy(p, *s)
	return len(*s), io.EOF
}

// NewReader は、StringReader を返します。
func NewReader(s string) io.Reader {
	sr := StringReader(s)
	return &sr
}

func main() {
	doc, err := html.Parse(NewReader("<p>Hello, world!</p>"))
	if err != nil {
		log.Fatalf("ch07/ex04: %v", err)
	}
	prettyhtml.WriteHTML(os.Stdout, doc)
}
