// Package tempconv は、摂氏、華氏、および絶対温度の温度計算を行います。
package tempconv

// CToF は、摂氏の温度を華氏の温度に変換します。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC は、華氏の温度を摂氏の温度に変換します。
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK は、摂氏の温度を絶対温度に変換します。
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToC は、絶対温度を摂氏の温度に変換します。
func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }

// FToK は、華氏の温度を絶対温度に変換します。
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToF は、絶対温度を華氏の温度に変換します。
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
