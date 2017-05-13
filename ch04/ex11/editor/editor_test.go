package editor

import (
	"testing"
)

func TestRemoveUTF8BOM(t *testing.T) {
	var tests = []struct {
		b    []byte
		want []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("A"), []byte("A")},
		{[]byte("あいうえお"), []byte("あいうえお")},
		{[]byte("\xef\xbb\xbf"), []byte("")},
		{[]byte("\xef\xbb\xbfA"), []byte("A")},
		{[]byte("\xef\xbb\xbfあいうえお"), []byte("あいうえお")},
	}

	for _, test := range tests {
		got := removeUTF8BOM(test.b)
		if string(got) != string(test.want) {
			t.Errorf("removeUTF8BOM(%q) = %q, want %q", string(test.b), string(got), string(test.want))
		}
	}
}
