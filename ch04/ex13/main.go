// ch04/ex13 は、Open Movie Database から映画のポスター画像を取得します。
package main

import (
	"fmt"
	"os"

	"github.com/kdama/gopl/ch04/ex13/omdb"
)

func main() {
	err := omdb.GetPoster(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch04/ex12: %v\n", err)
		os.Exit(1)
	}
}
