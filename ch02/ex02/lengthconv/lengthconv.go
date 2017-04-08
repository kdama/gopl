// Package lengthconv は、フィートとメートルの長さの計算を行います。
package lengthconv

import "fmt"

// Foot は、フィートの長さを表します。
type Foot float64

// Meter は、メートルの長さを表します。
type Meter float64

const (
	// FootPerMeter は、メートルあたりフィートの値です。
	FootPerMeter float64 = 0.3048
)

func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
