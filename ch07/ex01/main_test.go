package main

import (
	"fmt"
	"testing"
)

func TestWordCounterWrite(t *testing.T) {
	var tests = []struct {
		str  string
		want int
	}{
		{"", 0},
		{" ", 0},
		{"\n", 0},
		{" \n", 0},
		{"\n ", 0},
		{" \n ", 0},
		{"a", 1},
		{"a b c d e", 5},
		{"a b\nc d\ne", 5},
		{"日本語 にほんご　ニホンゴ\nﾆﾎﾝｺﾞ", 4},
	}

	for _, test := range tests {
		var wc WordCounter
		fmt.Fprint(&wc, test.str)
		if got := int(wc); got != test.want {
			t.Errorf("Word count of %q = %d, want %d", test.str, got, test.want)
		}
	}
}

func TestLineCounterWrite(t *testing.T) {
	var tests = []struct {
		str  string
		want int
	}{
		{"", 0},
		{" ", 1},
		{"\n", 1},
		{" \n", 1},
		{"\n ", 2},
		{" \n ", 2},
		{"a", 1},
		{"a b c d e", 1},
		{"a b\nc d\ne", 3},
		{"日本語 にほんご　ニホンゴ\nﾆﾎﾝｺﾞ", 2},
	}

	for _, test := range tests {
		var wc LineCounter
		fmt.Fprint(&wc, test.str)
		if got := int(wc); got != test.want {
			t.Errorf("Line count of %q = %d, want %d", test.str, got, test.want)
		}
	}
}
