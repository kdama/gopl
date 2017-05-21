// ch05/ex08 は、指定された id 属性を持つ最初の HTML 要素を見つけます。
package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/kdama/gopl/ch05/ex07/prettyhtml"
	"golang.org/x/net/html"
)

var idFlag = flag.String("id", "", "element ID to search")

func main() {
	flag.Parse()
	for _, url := range flag.Args() {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("ch05/ex08: %v", err)
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Fatalf("ch05/ex08: %v", err)
		}

		node := ElementByID(doc, *idFlag)
		if node == nil {
			log.Fatalf("ch05/ex08: Node not found")
		} else {
			prettyhtml.WriteHTML(os.Stdout, node)
		}
	}
}

// ElementByID は、指定された ID の要素を取得します。
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, findElement, findElement)
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if !pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, id, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil {
		if !post(n, id) {
			return n
		}
	}

	return nil
}

var depth int

func findElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}
	}
	return true
}

// isSelfClosableTag は、tagName が self-closing 可能なタグ名かどうかを返します。
func isSelfClosableTag(tagName string) bool {
	return tagName == "area" ||
		tagName == "base" ||
		tagName == "br" ||
		tagName == "col" ||
		tagName == "command" ||
		tagName == "embed" ||
		tagName == "hr" ||
		tagName == "img" ||
		tagName == "input" ||
		tagName == "keygen" ||
		tagName == "link" ||
		tagName == "meta" ||
		tagName == "param" ||
		tagName == "source" ||
		tagName == "track" ||
		tagName == "wbr"
}
