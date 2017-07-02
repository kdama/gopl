// ch08/ex06 は、指定された深さまでをクロールする crawl2 です。URL と根からの深さを表示します。
package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

var maxdepth int
var args []string

type Work struct {
	url   string
	depth int
}

func init() {
	flag.IntVar(&maxdepth, "depth", 3, "max depth to crawl")
	flag.Parse()
	args = flag.Args()
}

var tokens = make(chan struct{}, 20)

func crawl(work Work) []Work {
	fmt.Printf("%d\t%s\n", work.depth, work.url)

	// 対象の深さが最大深さ以上のとき、それ以上クロールする必要はないので nil を返します。
	if work.depth >= maxdepth {
		return nil
	}

	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(work.url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}

	works := []Work{}
	for _, link := range list {
		works = append(works, Work{link, work.depth + 1})
	}
	return works
}

func main() {
	worklist := make(chan []Work)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() {
		works := []Work{}
		for _, url := range args {
			works = append(works, Work{url, 1})
		}
		worklist <- works
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				n++
				go func(link Work) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
