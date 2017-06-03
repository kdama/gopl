package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"x ? y : z", Env{"x": 1, "y": 2, "z": 3}, "2"},
		{"x ? y : z", Env{"x": 0, "y": 2, "z": 3}, "3"},
		{"x ? y : z", Env{"x": math.NaN(), "y": 2, "z": 3}, "3"},
		{"x ? y : z", Env{"x": math.Inf(0), "y": 2, "z": 3}, "3"},
		{"x ? y : z", Env{"x": -math.Inf(-1), "y": 2, "z": 3}, "3"},

		{"a ? b ? c : d : e", Env{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, "3"},
		{"a ? b : c ? d : e", Env{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, "2"},
		{"a + b * c ? d * e ? f - g : h * i : j", Env{
			"a": 1, "b": 2, "c": 3, "d": 4, "e": 0, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
		}, "72"},
	}
	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.want {
			t.Errorf("(%s).Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

func TestErrors(t *testing.T) {
	for _, test := range []struct{ expr, wantErr string }{
		{"x ? y", "unexpected end of file"},
		{"x : y", "unexpected ':'"},
		{"?true", "unexpected '?'"},
		{`::hello`, "unexpected ':'"},
		{"log(?)", "unexpected '?'"},
		{"sqrt(?, :)", "unexpected '?'"},
	} {
		expr, err := Parse(test.expr)
		if err == nil {
			vars := make(map[Var]bool)
			err = expr.Check(vars)
			if err == nil {
				t.Errorf("unexpected success: %s", test.expr)
				continue
			}
		}
		if err.Error() != test.wantErr {
			t.Errorf("got error %s, want %s", err, test.wantErr)
		}
	}
}
