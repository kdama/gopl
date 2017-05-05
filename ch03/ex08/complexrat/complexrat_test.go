package complexrat

import (
	"math/big"
	"testing"
)

func equalsBigRat(a, b *big.Rat) bool {
	sub := big.NewRat(0, 1).Sub(a, b)
	abssub := big.NewRat(0, 1).Abs(sub)
	return abssub.Cmp(big.NewRat(1, 1e6)) < 0
}

func equals(a, b *ComplexRat) bool {
	return equalsBigRat((*a).Re, (*b).Re) && equalsBigRat((*a).Im, (*b).Im)
}

func TestSquare(t *testing.T) {
	var tests = []struct {
		c, want *ComplexRat
	}{
		{
			&ComplexRat{big.NewRat(0, 1), big.NewRat(0, 1)},
			&ComplexRat{big.NewRat(0, 1), big.NewRat(0, 1)},
		},
		{
			&ComplexRat{big.NewRat(1, 1), big.NewRat(1, 1)},
			&ComplexRat{big.NewRat(0, 1), big.NewRat(2, 1)},
		},
		{
			&ComplexRat{big.NewRat(0, 1), big.NewRat(-2, 1)},
			&ComplexRat{big.NewRat(-4, 1), big.NewRat(0, 1)},
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
		a, b, want *ComplexRat
	}{
		{
			&ComplexRat{big.NewRat(0, 1), big.NewRat(0, 1)},
			&ComplexRat{big.NewRat(0, 1), big.NewRat(0, 1)},
			&ComplexRat{big.NewRat(0, 1), big.NewRat(0, 1)},
		},
		{
			&ComplexRat{big.NewRat(1, 1), big.NewRat(2, 1)},
			&ComplexRat{big.NewRat(3, 1), big.NewRat(4, 1)},
			&ComplexRat{big.NewRat(4, 1), big.NewRat(6, 1)},
		},
		{
			&ComplexRat{big.NewRat(-4, 1), big.NewRat(0, 1)},
			&ComplexRat{big.NewRat(0, 1), big.NewRat(-2, 1)},
			&ComplexRat{big.NewRat(-4, 1), big.NewRat(-2, 1)},
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
		c    *ComplexRat
		n    *big.Rat
		want int
	}{
		{
			&ComplexRat{big.NewRat(0, 1), big.NewRat(0, 1)},
			big.NewRat(0, 1),
			0,
		},
		{
			&ComplexRat{big.NewRat(10, 1), big.NewRat(20, 1)},
			big.NewRat(10, 1),
			1,
		},
		{
			&ComplexRat{big.NewRat(-4, 1), big.NewRat(-2, 1)},
			big.NewRat(2, 1),
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
