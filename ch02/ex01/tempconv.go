// Package tempconv は、摂氏、華氏、および絶対温度の温度計算を行います。
package tempconv

import "fmt"

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
