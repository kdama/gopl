// ch05/ex07 は、汎用の HTML プリティプリンタです。
// コメントノード、テキストノード、個々の要素の属性を表示します。
// また、要素が子を持たない場合には、self-closing 形式を使用します。
package main

import (
	"net/http"
	"os"

	"github.com/kdama/gopl/ch05/ex07/prettyhtml"
	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	prettyhtml.WriteHTML(os.Stdout, doc)
	return nil
}
