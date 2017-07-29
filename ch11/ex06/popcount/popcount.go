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

// BitShiftPopCount は、ビットシフトを用いて x のポピュレーションカウントを返します。
func BitShiftPopCount(x uint64) int {
	count := 0
	for i := uint64(0); i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

// LSBPopCount は、LSB へのビット演算を用いて x のポピュレーションカウントを返します。
func LSBPopCount(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}
