// ex10 は、URL を並行に取り出して時間と大きさを表示することを、2 回行います。
// また、取り出しによって得られた内容を、ファイルに保存します。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fetchToDir(os.Args[1:], "out/1")
	fetchToDir(os.Args[1:], "out/2")
}

// fetchToDir は、URL を並行に取り出して、時間を表示します。
// また、取り出しによって得られた内容を、指定されたディレクトリに保存します。
func fetchToDir(urls []string, dirName string) {
	start := time.Now()
	ch := make(chan string)
	for idx, url := range urls {
		file, err := os.Create(fmt.Sprintf("%s/%d.txt", dirName, idx))
		if err != nil {
			fmt.Println(err) // send to channel ch
			return
		}
		defer file.Close()
		go fetch(url, file, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// fetch は、URL を取り出して、時間と大きさを表示します。
// また、取り出しによって得られた内容を、与えられた writer に書き込みます。
func fetch(url string, writer io.Writer, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(writer, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
