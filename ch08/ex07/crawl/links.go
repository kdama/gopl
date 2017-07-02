package crawl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"strings"

	"path/filepath"

	"golang.org/x/net/html"
)

// Extract は、対象のページ内のリンクを抽出します。
// また、可能ならば、そのページ内のリンクを相対パスに変換します。
func Extract(url string) ([]string, []byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	fmt.Println(resp.Header.Get("Content-Type"))

	if strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") {
		doc, err := html.Parse(resp.Body)
		if err != nil {
			return nil, nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
		}

		var links []string
		visitNode := func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for i, a := range n.Attr {
					if a.Key != "href" {
						continue
					}
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue // ignore bad URLs
					}
					links = append(links, link.String())
					n.Attr[i].Val = relativeURL(link, resp.Request.URL)
				}
			}
		}
		forEachNode(doc, visitNode, nil)

		var b bytes.Buffer
		html.Render(&b, doc)
		return links, b.Bytes(), nil
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return nil, b, nil
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

func relativeURL(link, base *url.URL) (result string) {
	if link.Scheme != base.Scheme {
		return link.String()
	}
	if link.Host != base.Host {
		return link.String()
	}
	depth := strings.Count(strings.Split(base.Path, "?")[0], "/")
	result = "./" + strings.Repeat("../", depth) + link.Path
	if !strings.Contains(filepath.Base(link.Path), ".") {
		result += "/index.html"
	}
	if link.RawQuery != "" {
		result += "?" + link.RawQuery
	}
	return
}
