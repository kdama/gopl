package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var tests = []struct {
		str  string
		want int64
	}{
		{"", 0},
		{" ", 1},
		{"\n", 1},
		{" \n", 2},
		{"\n ", 2},
		{" \n ", 3},
		{"a", 1},
		{"a b c d e", 9},
		{"a b\nc d\ne", 9},
		{"日本語", 9},
	}

	for _, test := range tests {
		w, count := CountingWriter(ioutil.Discard)
		fmt.Fprint(w, test.str)
		if got := *count; got != test.want {
			t.Errorf("Byte count of %q = %d, want %d", test.str, got, test.want)
		}
	}
}
