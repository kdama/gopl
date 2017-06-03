package eval

import (
	"testing"
)

func TestVars(t *testing.T) {
	tests := []struct {
		expr string
		want []Var
	}{
		{"sqrt(A / pi)", []Var{"A", "pi"}},
		{"pow(x, 3) + pow(y, 3)", []Var{"x", "y"}},
		{"5 / 9 * (F - 32)", []Var{"F"}},
		{"-1 + -x", []Var{"x"}},
		{"-1 - x", []Var{"x"}},
	}
	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Errorf("%s: %v", test.expr, err) // parse error
			continue
		}

		if got := expr.Vars(); !equals(got, test.want) {
			t.Errorf("(%s).Vars() = %v, want %v", test.expr, got, test.want)
		}
	}
}

func equals(a, b []Var) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
