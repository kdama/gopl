// ch13/ex04 は、Go による純粋な bzip 圧縮ツールです。
package main

import (
	"io"
	"log"
	"os"

	"github.com/kdama/gopl/ch13/ex04/bzip"
)

func main() {
	w, err := bzip.NewWriter(os.Stdout)
	if err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: write: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}
