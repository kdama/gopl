// Package massconv は、ポンドとキログラムの重さの計算を行います。
package massconv

// PoundToKilogram は、ポンドの重さをキログラムの重さに変換します。
func PoundToKilogram(p Pound) Kilogram { return Kilogram(float64(p) / PoundPerKilogram) }

// KilogramToPound は、キログラムの重さをポンドの重さに変換します。
func KilogramToPound(k Kilogram) Pound { return Pound(float64(k) * PoundPerKilogram) }
