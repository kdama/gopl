// Package tempconv は、摂氏、華氏、および絶対温度の温度計算を行います。
// ただし、Celsius や Fahrenheit が文字列として表示される際に ° を表示しません。
package tempconv

import (
	"flag"
	"fmt"
)

// Celsius は、摂氏の温度を表します。
type Celsius float64

// Fahrenheit は、華氏の温度を表します。
type Fahrenheit float64

const (
	// AbsoluteZeroC は、絶対零度です。
	AbsoluteZeroC Celsius = -273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%gC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }

// CToF は、摂氏の温度を華氏の温度に変換します。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC は、華氏の温度を摂氏の温度に変換します。
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

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
