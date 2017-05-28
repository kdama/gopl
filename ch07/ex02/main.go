// ch07/ex02 は、CountingWriter の実装です。
package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

// ByteCounter は、バイトカウンタを表します。
type ByteCounter struct {
	writer io.Writer
	count  int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	c.writer.Write(p)
	c.count += int64(len(p))
	return len(p), nil
}

// CountingWriter は、与えられた io.Writer を包む新たな Writer と、書き込まれたバイト数を常に保持する int64 変数へのポインタを返します。
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := ByteCounter{
		writer: w,
		count:  0,
	}
	return &bc, &bc.count
}

func main() {
	w, count := CountingWriter(ioutil.Discard)
	fmt.Fprint(w, "hello")
	fmt.Println(*count) // "5"
}
