package main

import (
	"testing"
)

func TestRemoveDupSpace(t *testing.T) {
	var tests = []struct {
		b    []byte
		want []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("A"), []byte("A")},
		{[]byte("A B C"), []byte("A B C")},
		{[]byte("A  B  C"), []byte("A B C")},
		{[]byte("AB C  D   "), []byte("AB C D ")},
		{[]byte("   A  B CD"), []byte(" A B CD")},
		{[]byte("AB　C　　D　　　"), []byte("AB C D ")},
		{[]byte("　　　A　　B　CD"), []byte(" A B CD")},
	}

	for _, test := range tests {
		got := removeDupSpace(test.b)
		if string(got) != string(test.want) {
			t.Errorf("removeDupSpace(%q) = %q, want %q", string(test.b), string(got), string(test.want))
		}
	}
}
