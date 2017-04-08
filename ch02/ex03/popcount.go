// Package popcount は、ポピュレーションカウントを返します。
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// TablePopCount は、テーブル参照を用いて x のポピュレーションカウントを返します。
func TablePopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// LoopPopCount は、ループを用いて x のポピュレーションカウントを返します。
func LoopPopCount(x uint64) int {
	count := 0
	for i := uint64(0); i < 64; i++ {
		if x>>i%2 == 1 {
			count++
		}
	}
	return count
}
