// ch07/ex01 は、io.Writer の契約を満足するワードカウンタや行カウンタの実装です。
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter は、ワードカウンタを表します。
type WordCounter int

// LineCounter は、行カウンタを表します。
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	i := 0
	for sc.Scan() {
		i++
	}
	*c += WordCounter(i)
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	i := 0
	for sc.Scan() {
		i++
	}
	*c += LineCounter(i)
	return len(p), nil
}

func main() {
	var wc WordCounter
	fmt.Fprint(&wc, "Hello, world!\nHello, 世界。")
	fmt.Println(wc) // "4"

	var lc LineCounter
	fmt.Fprint(&lc, "Hello, world!\nHello, 世界。")
	fmt.Println(lc) // "2"
}
