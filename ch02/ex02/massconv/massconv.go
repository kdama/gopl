// Package massconv は、ポンドとキログラムの重さの計算を行います。
package massconv

import "fmt"

// Pound は、ポンドの重さを表します。
type Pound float64

// Kilogram は、キログラムの重さを表します。
type Kilogram float64

const (
	// PoundPerKilogram は、キログラムあたりポンドの値です。
	PoundPerKilogram float64 = 0.45359237
)

func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
