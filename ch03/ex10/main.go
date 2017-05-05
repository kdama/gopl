// ch03/ex10 は、再帰呼び出しをせずに、文字列にカンマを挿入します。
package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

// comma は、再帰呼び出しをせずに、文字列にカンマを挿入します。
func comma(s string) string {
	var buf bytes.Buffer
	i := (3 - utf8.RuneCountInString(s)%3) % 3
	for _, r := range s {
		if i == 3 {
			buf.WriteByte(',')
			i = 0
		}
		buf.WriteRune(r)
		i++
	}
	return buf.String()
}
