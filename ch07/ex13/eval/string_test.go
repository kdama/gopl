package eval

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
	}{
		{"sqrt(A / pi)"},
		{"pow(x, 3) + pow(y, 3)"},
		{"5 / 9 * (F - 32)"},
		{"-1 + -x"},
		{"-1 - x"},
	}
	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		got, err := Parse(expr.String())
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		if !expr.Equals(got) {
			t.Errorf("Parse(expr) != Parse(Parse(expr)): expr = %s", test.expr)
		}
	}
}
