// Package tempconv は、摂氏と華氏の温度計算を行います。
package tempconv

// CToF は、摂氏の温度を華氏の温度に変換します。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC は、華氏の温度を摂氏の温度に変換します。
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
