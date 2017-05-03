// Package colors は、色に関する処理を行います。
package colors

import (
	"image/color"
)

// GetAverageColor は、与えられた複数の色の、平均の色を返します。
func GetAverageColor(colors []color.Color) color.Color {
	if len(colors) < 1 {
		return nil
	}

	var r, g, b, a float64

	for _, color := range colors {
		dr, dg, db, da := color.RGBA()
		r += float64(dr>>8) / float64(len(colors))
		g += float64(dg>>8) / float64(len(colors))
		b += float64(db>>8) / float64(len(colors))
		a += float64(da>>8) / float64(len(colors))
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}
