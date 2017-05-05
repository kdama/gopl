// Package complexrat は、big.Rat を用いて複素数を表現します。
package complexrat

import "math/big"

// ComplexRat は、big.Rat を用いて複素数を表現します。
type ComplexRat struct {
	Re *big.Rat
	Im *big.Rat
}

// AbsCompare は、複素数の絶対値を n と比較します。
// 複素数の絶対値のほうが大きいとき、正の値を返します。
func (c *ComplexRat) AbsCompare(n *big.Rat) int {
	re2 := big.NewRat(0, 1).Mul(c.Re, c.Re)
	im2 := big.NewRat(0, 1).Mul(c.Im, c.Im)
	n2 := big.NewRat(0, 1).Mul(n, n)
	return big.NewRat(0, 1).Add(re2, im2).Cmp(n2)
}

// Square は、複素数 c の自乗を返します。
func (c *ComplexRat) Square() *ComplexRat {
	re2 := big.NewRat(0, 1).Mul(c.Re, c.Re)
	im2 := big.NewRat(0, 1).Mul(c.Im, c.Im)
	reim := big.NewRat(0, 1).Mul(c.Re, c.Im)

	c.Re.Sub(re2, im2)
	c.Im.Add(reim, reim)
	return c
}

// Add は、複素数 a, b の和を返します。
func (c *ComplexRat) Add(z *ComplexRat) *ComplexRat {
	c.Re.Add(c.Re, z.Re)
	c.Im.Add(c.Im, z.Im)
	return c
}
