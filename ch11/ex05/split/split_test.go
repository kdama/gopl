package word

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s, sep    string
		wantLen   int
		wantWords []string
	}{
		{"", "", 0, []string{}},
		{"", ":", 1, []string{""}},
		{"a", "", 1, []string{"a"}},
		{"a", ":", 1, []string{"a"}},
		{"a:b:c", "", 5, []string{"a", ":", "b", ":", "c"}},
		{"a:b:c", ":", 3, []string{"a", "b", "c"}},
		{"http://日本語/", "/", 4, []string{"http:", "", "日本語", ""}},
		{"http://日本語/", ":/", 2, []string{"http", "/日本語/"}},
	}

	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.wantLen {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, test.wantLen)
		} else if !reflect.DeepEqual(words, test.wantWords) {
			t.Errorf("Split(%q, %q) returned %v, want %v", test.s, test.sep, words, test.wantWords)
		}
	}
}
