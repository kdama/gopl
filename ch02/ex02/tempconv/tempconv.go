// Package tempconv は、摂氏と華氏の温度計算を行います。
package tempconv

import "fmt"

// Celsius は、摂氏の温度を表します。
type Celsius float64

// Fahrenheit は、華氏の温度を表します。
type Fahrenheit float64

const (
	// AbsoluteZeroC は、絶対零度です。
	AbsoluteZeroC Celsius = -273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
