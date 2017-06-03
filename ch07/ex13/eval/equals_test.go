package eval

import (
	"testing"
)

func TestEquals(t *testing.T) {
	tests := []struct {
		e1, e2 string
		want   bool
	}{
		{"sqrt(A / pi)", "sqrt(A / pi)", true},
		{"pow(x, 3) + pow(y, 3)", "pow(x, 3) + pow(y, 3)", true},
		{"5 / 9 * (F - 32)", "5 / 9 * (F - 32)", true},
		{"-1 + -x", "-1 + -x", true},
		{"-1 - x", "-1 - x", true},

		{"sqrt(A / pi)", "sqrt(A/pi)", true},
		{"pow(x, 3) + pow(y, 3)", "pow(x,3)+pow(y,3)", true},
		{"5 / 9 * (F - 32)", "5 / 9 * ( F - 32 )", true},
		{"-1 + -x", "(-1) + (-(x))", true},
		{"-1 - x", "((-1) - (x))", true},

		{"sqrt(A / pi)", "sqrt(A * pi)", false},
		{"sqrt(A / pi)", "SQRT(A / pi)", false},
		{"pow(x, 3) + pow(y, 3)", "pow(x, 3) - pow(y, 3)", false},
		{"pow(x, 3) + pow(y, 3)", "pow(y, 3) + pow(y, 3)", false},
		{"pow(x, 3) + pow(y, 3)", "pow(x, 2) + pow(y, 3)", false},
		{"5 / 9 * (F - 32)", "5 / 9 * (F + 32)", false},
		{"5 / 9 * (F - 32)", "5 / 9 *  F - 32 ", false},
		{"-1 + -x", "-1 - x", false},
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
