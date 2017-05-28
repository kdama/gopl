package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestWordCounterWrite(t *testing.T) {
	var tests = []struct {
		str  string
		n    int64
		want string
	}{
		{"", 6, ""},
		{"a", 6, "a"},
		{"Hello, world!", 6, "Hello,"},
		{"こんにちは、世界。", 6, "こん"},
	}

	for _, test := range tests {
		b, err := ioutil.ReadAll(LimitReader(strings.NewReader(test.str), test.n))
		if err != nil {
			t.Errorf("%v", err)
		} else if got := string(b); got != test.want {
			t.Errorf("LimitReader(%q, %d) reads %q, want %q", test.str, test.n, got, test.want)
		}
	}
}
