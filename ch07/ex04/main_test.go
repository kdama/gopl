package main

import (
	"io/ioutil"
	"testing"
)

func TestWordCounterWrite(t *testing.T) {
	var tests = []struct {
		str string
	}{
		{""},
		{" "},
		{"\n"},
		{" \n"},
		{"\n "},
		{" \n "},
		{"a"},
		{"a b c d e"},
		{"a b\nc d\ne"},
		{"日本語 にほんご　ニホンゴ\nﾆﾎﾝｺﾞ"},
	}

	for _, test := range tests {
		b, err := ioutil.ReadAll(NewReader(test.str))
		if err != nil {
			t.Errorf("%v", err)
		} else if got := string(b); got != test.str {
			t.Errorf("NewReader(%q) reads %q, want %q", test.str, got, test.str)
		}
	}
}
