// ch08/ex07 は、指定されたアドレスから到達可能な同じドメイン内のページをローカルディスクに保存します。
package main

import (
	"os"

	"github.com/kdama/gopl/ch08/ex07/crawl"
)

// out は、ページの複製の保存先となるローカルディスク上のディレクトリです。
const out = "./out"

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl.Crawl(link, out)
				}(link)
			}
		}
	}
}
