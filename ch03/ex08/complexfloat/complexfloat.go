// Package complexfloat は、big.Float を用いて複素数を表現します。
package complexfloat

import "math/big"

// ComplexFloat は、big.Float を用いて複素数を表現します。
type ComplexFloat struct {
	Re   *big.Float
	Im   *big.Float
	Prec uint
}

// AbsCompare は、複素数の絶対値を n と比較します。
// 複素数の絶対値のほうが大きいとき、正の値を返します。
func (c *ComplexFloat) AbsCompare(n *big.Float) int {
	re2 := big.NewFloat(0).SetPrec(c.Prec).Mul(c.Re, c.Re)
	im2 := big.NewFloat(0).SetPrec(c.Prec).Mul(c.Im, c.Im)
	n2 := big.NewFloat(0).SetPrec(c.Prec).Mul(n, n)
	return big.NewFloat(0).SetPrec(c.Prec).Add(re2, im2).Cmp(n2)
}

// Square は、複素数の自乗を返します。
func (c *ComplexFloat) Square() *ComplexFloat {
	re2 := big.NewFloat(0).SetPrec(c.Prec).Mul(c.Re, c.Re)
	im2 := big.NewFloat(0).SetPrec(c.Prec).Mul(c.Im, c.Im)
	reim := big.NewFloat(0).SetPrec(c.Prec).Mul(c.Re, c.Im)

	c.Re.Sub(re2, im2)
	c.Im.Add(reim, reim)
	return c
}

// Add は、複素数 z との和を返します。
func (c *ComplexFloat) Add(z *ComplexFloat) *ComplexFloat {
	c.Re.Add(c.Re, z.Re)
	c.Im.Add(c.Im, z.Im)
	return c
}
