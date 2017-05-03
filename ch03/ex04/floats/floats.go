// Package floats は、浮動小数点数に関する計算を行います。
package floats

import "math"

// IsFinite は、f が有限の値かどうかを返します。
func IsFinite(f float64) bool {
	if math.IsInf(f, 0) {
		return false
	}
	if math.IsNaN(f) {
		return false
	}
	return true
}
