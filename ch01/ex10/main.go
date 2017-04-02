package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fetchToFile(os.Args[1:], "out/1.txt")
	fetchToFile(os.Args[1:], "out/2.txt")
}

func fetchToFile(urls []string, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err) // send to channel ch
		return
	}
	defer file.Close()

	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, file, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

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
