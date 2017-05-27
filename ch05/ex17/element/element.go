// Package element は、HTML の要素に関する処理を行います。
package element

import (
	"golang.org/x/net/html"
)

// ElementsByTagName は、指定されたタグ名のいずれかに一致する要素を全て返します。
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var result []*html.Node

	scan := func(node *html.Node) {
		if node.Type == html.ElementNode {
			for _, n := range name {
				if node.Data == n {
					result = append(result, node)
				}
			}
		}
	}

	// Do nothing
	noop := func(node *html.Node) {}

	forEachNode(doc, scan, noop)

	return result
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
