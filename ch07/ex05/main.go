// ch07/ex05 は、io.LimitReader の再実装です。
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// LimitedByteCounter は、読み出し可能なバイト数が限定されたバイトカウンタの実装です。
type LimitedByteCounter struct {
	reader io.Reader
	// 読み出し可能な残りのバイト数です。
	rest int64
}

func (c *LimitedByteCounter) Read(p []byte) (int, error) {
	length := int64(len(p))
	if c.rest < length {
		length = c.rest
	}
	c.rest -= length

	n, err := c.reader.Read(p[:length])
	if c.rest == 0 {
		return n, io.EOF
	}
	return n, err
}

// LimitReader は、n バイトまでのみを読み出す io.Reader の実装を返します。
func LimitReader(r io.Reader, n int64) io.Reader {
	lbc := LimitedByteCounter{
		reader: r,
		rest:   n,
	}
	return &lbc
}

func main() {
	r := strings.NewReader("Hello, world!")
	s, err := ioutil.ReadAll(LimitReader(r, 5))
	if err != nil {
		log.Fatalf("ch07/ex05: %v", err)
	}
	fmt.Println(string(s)) // "Hello"
}
