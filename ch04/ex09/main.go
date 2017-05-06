// ch04/ex09 は、入力テキストファイル内のそれぞれの単語の出現頻度を報告します。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "ch04/ex09: must have 1 argument.")
		os.Exit(1)
	}

	counts := make(map[string]int) // counts of words

	in, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch04/ex09: %v", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(in)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		counts[word]++
	}

	fmt.Printf("word\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
