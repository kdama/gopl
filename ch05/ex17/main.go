// ch05/ex17 は、HTML を取得して、指定されたタグ名の要素を出力します。
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kdama/gopl/ch05/ex07/prettyhtml"
	"github.com/kdama/gopl/ch05/ex17/element"
	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: ex17 url tag_names...")
	}
	url, tagNames := os.Args[1], os.Args[2:]

	doc, err := fetch(url)
	if err != nil {
		log.Fatalf("ch05/ex17: %v", err)
	}

	nodes := element.ElementsByTagName(doc, tagNames...)

	for _, node := range nodes {
		prettyhtml.WriteHTML(os.Stdout, node)
		fmt.Println()
	}
}

func fetch(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return html.Parse(resp.Body)
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
