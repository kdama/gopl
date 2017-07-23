// Package charcount は、Unicode 文字の数を計算します。
package charcount

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

// Count は、Unicode 文字の数を計算します。
func Count(r io.Reader) (
	counts map[rune]int, // counts of Unicode characters
	utflen [utf8.UTFMax + 1]int, // count of lengths of UTF-8 encodings
	invalid int, // count of invalid UTF-8 characters
	err error,
) {
	counts = make(map[rune]int) // counts of Unicode characters

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			return counts, utflen, invalid, err
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return
}
