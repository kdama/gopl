// ch05/ex05 は、<a> 以外も含めて、全ての種類のリンクをドキュメントから抽出します。
package main

import (
	"fmt"
	"net/http"
	"os"

	"strings"

	"golang.org/x/net/html"
)

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return countWordsAndImages(n.NextSibling)
	} else if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	cwords, cimages := countWordsAndImages(n.FirstChild)
	swords, simages := countWordsAndImages(n.NextSibling)
	words += cwords + swords
	images += cimages + simages
	return
}

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ch05/ex05: %v\n", err)
			continue
		}
		fmt.Printf("words:  %d\n", words)
		fmt.Printf("images: %d\n", images)
	}
}

// CountWordsAndImages は、HTML ドキュメントに対する HTTP GET リクエストを url へ行い、
// そのドキュメント内に含まれる単語と画像の数を返します。
func CountWordsAndImages(url string) (word, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("getting %s: %s", url, resp.Status)
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}

	word, images = countWordsAndImages(doc)
	return
}
