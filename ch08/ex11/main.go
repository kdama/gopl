// ch08/ex11 は、最初のレスポンスが到着すると他のリクエストをキャンセルする fetch です。
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kdama/gopl/ch08/ex07/crawl"
)

// Response は、URL の内容を取得する処理の結果を表します。
type Response struct {
	url  string
	body []byte
	err  error
}

func fetch(url string, cancel <-chan struct{}) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// 全ての実行中の fetch をキャンセルするためのチャンネルです。
var cancel = make(chan struct{})

// 保存先のディレクトリ名です。
var out = "./out"

func main() {
	urls := os.Args[1:]

	responses := make(chan Response, len(urls))
	for _, url := range urls {
		url := url
		go func() {
			body, err := fetch(url, cancel)
			responses <- Response{url, body, err}
		}()
	}
	first := <-responses
	close(cancel)
	if first.err != nil {
		log.Fatalf("fetch %s: %v", first.url, first.err)
	}
	crawl.Save(first.url, out, first.body)
	fmt.Fprintf(os.Stderr, "Fetched %s (%d bytes).\n", first.url, len(first.body))
}
