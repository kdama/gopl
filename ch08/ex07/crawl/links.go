package crawl

import (
	"bytes"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// extract は、対象の URL をローカルディスクに保存します。
// また、対象が HTML ページ内のリンクを抽出します。
// また、可能ならば、そのページ内のリンクを相対パスに変換します。
func extract(path string, data []byte) (links []string, converted []byte, err error) {
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, nil, fmt.Errorf("parsing %s as HTML: %v", path, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				pathURL, err := url.Parse(path)
				if err != nil {
					continue // ignore bad URLs
				}
				link, err := pathURL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
				n.Attr[i].Val = relativeURL(pathURL, link)
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	var b bytes.Buffer
	html.Render(&b, doc)
	return links, b.Bytes(), nil
}

// Copied from gopl.io/ch5/outline2.
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

// base URL を基点として、link URL を相対 URL に変換します。
func relativeURL(base, link *url.URL) (result string) {
	if link.Scheme != base.Scheme {
		return link.String()
	}
	if link.Host != base.Host {
		return link.String()
	}
	depth := strings.Count(base.Path, "/") - 1
	if depth < 0 {
		depth = 0
	}
	result = "./" + strings.Repeat("../", depth) + link.Path
	if link.Path == "" || !strings.Contains(filepath.Base(link.Path), ".") {
		result += "/index.html"
	}
	if link.RawQuery != "" {
		result += "?" + link.RawQuery
	}

	re := regexp.MustCompile(`/+`)
	result = re.ReplaceAllString(result, "/")
	return
}
