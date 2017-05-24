// ch05/ex13 は、対象のページから到達可能なページの複製を保存します。
// ただし、ドメインの異なるページはクロールせず、保存もしません。
package main

import (
	"io"
	"log"
	"net/http"
	neturl "net/url"
	"os"
	"path/filepath"
	"strings"

	"gopl.io/ch5/links"
)

// out は、ページの複製の保存先となるディレクトリです。
const out = "./out"

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	err := save(url)
	if err != nil {
		log.Print(err)
	}

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	var sameDomains []string
	for _, target := range list {
		shouldAppend, err := sameDomain(url, target)
		if err != nil {
			log.Print(err)
			continue
		}
		if shouldAppend {
			sameDomains = append(sameDomains, target)
		}
	}
	return sameDomains
}

func sameDomain(x, y string) (bool, error) {
	xParsed, err := neturl.Parse(x)
	if err != nil {
		return false, err
	}
	yParsed, err := neturl.Parse(y)
	if err != nil {
		return false, err
	}
	return xParsed.Host == yParsed.Host, nil
}

func save(url string) error {
	parsed, err := neturl.Parse(url)
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	localPath := out + parsed.Path
	if !strings.Contains(filepath.Base(localPath), ".") {
		localPath += "/index.html"
	}

	err = os.MkdirAll(filepath.Dir(localPath), os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(localPath)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return err
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
