// ch05/ex04 は、<a> 以外も含めて、全ての種類のリンクをドキュメントから抽出します。
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var linkAttrs = map[string][]string{
	"a":      []string{"href"},
	"link":   []string{"href"},
	"img":    []string{"src"},
	"script": []string{"src"},
	"iframe": []string{"src"},
	"form":   []string{"action"},
	"html":   []string{"manifest"},
	"video":  []string{"src", "poster"},
	// ...
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch05/ex04: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		for k, v := range linkAttrs {
			if n.Data == k {
				for _, attr := range v {
					for _, a := range n.Attr {
						if a.Key == attr {
							links = append(links, a.Val)
						}
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
