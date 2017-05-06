package main

import (
	"testing"
)

func TestReverseUTF8(t *testing.T) {
	var tests = []struct {
		b    []byte
		want []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("A"), []byte("A")},
		{[]byte("ABC"), []byte("CBA")},
	}

	for _, test := range tests {
		got := reverseUTF8(test.b)
		if string(got) != string(test.want) {
			t.Errorf("reverseUTF8(%q) = %q, want %q", string(test.b), string(got), string(test.want))
		}
	}
}
