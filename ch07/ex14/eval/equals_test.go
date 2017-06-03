package eval

import (
	"testing"
)

func TestEquals(t *testing.T) {
	tests := []struct {
		e1, e2 string
		want   bool
	}{
		{"x ? y : z", "x ? y : z", true},
		{"x ? y ? a : b : c", "x ? y ? a : b : c", true},
		{"x ? y : a ? b : c", "x ? y : a ? b : c", true},

		{"x ? y : z", "y ? z : x", false},
		{"x ? y : z", "z ? x : y", false},
		{"x ? y : z", "y ? x : z", false},
		{"x ? y : z", "z ? y : x", false},
		{"x ? y ? a : b : c", "x ? y ? a : c : b", false},
		{"x ? y ? a : b : c", "x ? y ? b : a : c", false},
		{"x ? y ? a : b : c", "x ? y ? b : c : a", false},
		{"x ? y ? a : b : c", "x ? y ? c : a : b", false},
		{"x ? y ? a : b : c", "x ? y ? c : b : a", false},
		{"x ? y ? a : b : c", "x ? y : a ? c : b", false},
		{"x ? y ? a : b : c", "x ? y : b ? a : c", false},
		{"x ? y ? a : b : c", "x ? y : b ? c : a", false},
		{"x ? y ? a : b : c", "x ? y : c ? a : b", false},
		{"x ? y ? a : b : c", "x ? y : c ? b : a", false},
	}
	for _, test := range tests {
		e1, err := Parse(test.e1)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		e2, err := Parse(test.e2)
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		if got := e1.Equals(e2); got != test.want {
			t.Errorf("Parse(%q).Equals(Parse(%q))) = %t, want %t", test.e1, test.e2, got, test.want)
		}
	}
}
