package complexfloat

import (
	"math/big"
	"testing"
)

func equalsBigFloat(a, b *big.Float) bool {
	sub := big.NewFloat(0).Sub(a, b)
	abssub := big.NewFloat(0).Abs(sub)
	return abssub.Cmp(big.NewFloat(1e-6)) < 0
}

func equals(a, b *ComplexFloat) bool {
	return equalsBigFloat((*a).Re, (*b).Re) && equalsBigFloat((*a).Im, (*b).Im)
}

func TestSquare(t *testing.T) {
	var tests = []struct {
		c, want *ComplexFloat
	}{
		{
			&ComplexFloat{big.NewFloat(0), big.NewFloat(0), 1024},
			&ComplexFloat{big.NewFloat(0), big.NewFloat(0), 1024},
		},
		{
			&ComplexFloat{big.NewFloat(1), big.NewFloat(1), 1024},
			&ComplexFloat{big.NewFloat(0), big.NewFloat(2), 1024},
		},
		{
			&ComplexFloat{big.NewFloat(0), big.NewFloat(-2), 1024},
			&ComplexFloat{big.NewFloat(-4), big.NewFloat(0), 1024},
		},
	}

	for _, test := range tests {
		got := test.c.Square()
		if !equals(got, test.want) {
			t.Errorf("(%v).Square() = %v, want %v", test.c, *got, *test.want)
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b, want *ComplexFloat
	}{
		{
			&ComplexFloat{big.NewFloat(0), big.NewFloat(0), 1024},
			&ComplexFloat{big.NewFloat(0), big.NewFloat(0), 1024},
			&ComplexFloat{big.NewFloat(0), big.NewFloat(0), 1024},
		},
		{
			&ComplexFloat{big.NewFloat(1), big.NewFloat(2), 1024},
			&ComplexFloat{big.NewFloat(3), big.NewFloat(4), 1024},
			&ComplexFloat{big.NewFloat(4), big.NewFloat(6), 1024},
		},
		{
			&ComplexFloat{big.NewFloat(-4), big.NewFloat(0), 1024},
			&ComplexFloat{big.NewFloat(0), big.NewFloat(-2), 1024},
			&ComplexFloat{big.NewFloat(-4), big.NewFloat(-2), 1024},
		},
	}

	for _, test := range tests {
		got := test.a.Add(test.b)
		if !equals(got, test.want) {
			t.Errorf("(%v).Add(%v) = %v, want %v", test.a, test.b, got, *test.want)
		}
	}
}

func TestAbsCompare(t *testing.T) {
	var tests = []struct {
		c    *ComplexFloat
		n    *big.Float
		want int
	}{
		{
			&ComplexFloat{big.NewFloat(0), big.NewFloat(0), 1024},
			big.NewFloat(0),
			0,
		},
		{
			&ComplexFloat{big.NewFloat(10), big.NewFloat(20), 1024},
			big.NewFloat(10),
			1,
		},
		{
			&ComplexFloat{big.NewFloat(-4), big.NewFloat(-2), 1024},
			big.NewFloat(2),
			1,
		},
	}

	for _, test := range tests {
		got := test.c.AbsCompare(test.n)
		if got != test.want {
			t.Errorf("(%v).AbsCompare(%v) = %v, want %v", test.c, test.n, got, test.want)
		}
	}
}
