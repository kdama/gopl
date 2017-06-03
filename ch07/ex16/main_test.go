package main

import (
	"testing"

	"github.com/kdama/gopl/ch07/ex15/eval"
)

func TestParseEnv(t *testing.T) {
	tests := []struct {
		s    string
		want map[eval.Var]float64
		err  bool
	}{
		{"", map[eval.Var]float64{}, false},
		{"a = 1\nb = 2", map[eval.Var]float64{"a": 1, "b": 2}, false},
		{"{ a: 3, b: 4 }", map[eval.Var]float64{"a": 3, "b": 4}, false},

		{"a = a", nil, true},
		{"{ a: 1+2 }", nil, true},
	}
	for _, test := range tests {
		got, err := parseEnv(test.s)
		if err != nil && !test.err {
			t.Errorf("expects no error for parseEnv(%q), but error", test.s)
		} else if err == nil && test.err {
			t.Errorf("expects error for parseEnv(%q), but no error", test.s)
		} else if !equals(got, test.want) {
			t.Errorf("parseEnv(%q) = %v, want %v", test.s, got, test.want)
		}
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		expr string
		env  map[eval.Var]float64
		err  bool
	}{
		{"1", map[eval.Var]float64{}, false},
		{"a + b", map[eval.Var]float64{"a": 1, "b": 2}, false},
		{"a ? 42 : b", map[eval.Var]float64{"a": 3, "b": 4}, false},

		{"a + b", map[eval.Var]float64{}, true},
		{"a ? 42 : b", map[eval.Var]float64{"a": 1}, true},
	}
	for _, test := range tests {
		expr, err := eval.Parse(test.expr)
		if err != nil {
			t.Errorf("%s: %v", test.expr, err)
		} else {
			err = validate(expr, test.env)
			if err != nil && !test.err {
				t.Errorf("expects no error for validate(%q, %v), but error", test.expr, test.env)
			} else if err == nil && test.err {
				t.Errorf("expects error for validate(%q, %v), but no error", test.expr, test.env)
			}
		}
	}
}

func equals(a, b map[eval.Var]float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if _, ok := b[i]; !ok {
			return false
		} else if a[i] != b[i] {
			return false
		}
	}
	return true
}
