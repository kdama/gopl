// Package tempconv は、摂氏、華氏、および絶対温度の温度計算を行います。
package tempconv

import (
	"flag"
	"fmt"
)

// Celsius は、摂氏の温度を表します。
type Celsius float64

// Fahrenheit は、華氏の温度を表します。
type Fahrenheit float64

// Kelvin は、絶対温度を表します。
type Kelvin float64

const (
	// AbsoluteZeroC は、絶対零度です。
	AbsoluteZeroC Celsius = -273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

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

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag は、指定された名前、デフォルト値、使い方を持つ Celsius フラグ
// を定義しており、そのフラグ変数のアドレスを返します。
// フラグ引数は度数と単位です。例えば、"100C" です。
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

type kelvinFlag struct{ Kelvin }

func (f *kelvinFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Kelvin = CToK(Celsius(value))
		return nil
	case "F", "°F":
		f.Kelvin = FToK(Fahrenheit(value))
		return nil
	case "K":
		f.Kelvin = Kelvin(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// KelvinFlag は、指定された名前、デフォルト値、使い方を持つ Kelvin フラグ
// を定義しており、そのフラグ変数のアドレスを返します。
// フラグ引数は度数と単位です。例えば、"100C" です。
func KelvinFlag(name string, value Kelvin, usage string) *Kelvin {
	f := kelvinFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Kelvin
}
