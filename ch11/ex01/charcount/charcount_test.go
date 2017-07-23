// Package charcount は、Unicode 文字の数を計算します。
package charcount

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCount(t *testing.T) {
	tests := []struct {
		in      string
		counts  map[rune]int         // counts of Unicode characters
		utflen  [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
		invalid int                  // count of invalid UTF-8 characters
		err     bool
	}{
		{
			"",
			map[rune]int{},
			[utf8.UTFMax + 1]int{},
			0,
			false,
		},
		{
			"abc",
			map[rune]int{
				'a': 1,
				'b': 1,
				'c': 1,
			},
			[utf8.UTFMax + 1]int{
				1: 3,
			},
			0,
			false,
		},
		{
			"Hello, 世界! 🍣",
			map[rune]int{
				'H': 1,
				'e': 1,
				'l': 2,
				'o': 1,
				',': 1,
				' ': 2,
				'世': 1,
				'界': 1,
				'!': 1,
				'🍣': 1,
			},
			[utf8.UTFMax + 1]int{
				1: 9,
				3: 2,
				4: 1,
			},
			0,
			false,
		},
		{
			string([]byte{128, 128, 128}),
			map[rune]int{},
			[utf8.UTFMax + 1]int{},
			3,
			false,
		},
	}

	for _, test := range tests {
		counts, utflen, invalid, err := Count(strings.NewReader(test.in))

		if err == nil && test.err {
			t.Errorf("Expects error for Count(%q), but no error", test.in)
		} else if err != nil && !test.err {
			t.Errorf("Expects no error for Count(%q), but error: %v", test.in, err)
		} else if !reflect.DeepEqual(counts, test.counts) {
			t.Errorf("Count(%q).counts = %v, want %v", test.in, counts, test.counts)
		} else if !reflect.DeepEqual(utflen, test.utflen) {
			t.Errorf("Count(%q).utflen = %v, want %v", test.in, utflen, test.utflen)
		} else if invalid != test.invalid {
			t.Errorf("Count(%q).invalid = %v, want %v", test.in, invalid, test.invalid)
		}
	}
}
