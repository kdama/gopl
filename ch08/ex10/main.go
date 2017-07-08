// ch08/ex10 は、HTTP リクエストのキャンセルをサポートする crawl2 です。
// リターンキーなどが入力されると、クロールを中止します。
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kdama/gopl/ch08/ex10/links"
)

var cancel = make(chan struct{})

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url, cancel)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
	}()

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
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
