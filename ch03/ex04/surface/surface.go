// Package surface は、高さに基づいて個々のポリゴンに色を付けながら、3-D 面の関数の SVG レンダリングを計算します。
package surface

import (
	"fmt"
	"image/color"
	"io"
	"math"

	"github.com/kdama/gopl/ch03/ex04/colors"
	"github.com/kdama/gopl/ch03/ex04/floats"
)

// Render は 3-D 面の関数の SVG レンダリングの結果を返します。
func Render(w io.Writer, width, height, cells int, xyrange, xyscale, zscale, angle float64, topColor, bottomColor color.Color) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	maxHeight, minHeight := getMaxMinHeight(cells, xyrange)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height, cells, xyrange, xyscale, zscale, angle)
			bx, by := corner(i, j, width, height, cells, xyrange, xyscale, zscale, angle)
			cx, cy := corner(i, j+1, width, height, cells, xyrange, xyscale, zscale, angle)
			dx, dy := corner(i+1, j+1, width, height, cells, xyrange, xyscale, zscale, angle)
			color := getColor(getHeight(i, j, cells, xyrange), maxHeight, minHeight, topColor, bottomColor)

			// 出力する前に、全ての値が有限かどうかを調べます。
			if floats.IsFinite(ax) && floats.IsFinite(ay) &&
				floats.IsFinite(bx) && floats.IsFinite(by) &&
				floats.IsFinite(cx) && floats.IsFinite(cy) &&
				floats.IsFinite(dx) && floats.IsFinite(dy) {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j, width, height, cells int, xyrange, xyscale, zscale, angle float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*math.Cos(angle)*xyscale
	sy := float64(height)/2 + (x+y)*math.Sin(angle)*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// getHeight は、ポリゴンの高さを計算します。
func getHeight(i, j, cells int, xyrange float64) float64 {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	return f(x, y)
}

// getMaxMinHeight は、全てのポリゴンの高さを求めて、高さの最大値と最小値を返します。
func getMaxMinHeight(cells int, xyrange float64) (float64, float64) {
	maxHeight := math.NaN()
	minHeight := math.NaN()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z := getHeight(i, j, cells, xyrange)

			if floats.IsFinite(z) {
				if math.IsNaN(maxHeight) || maxHeight < z {
					maxHeight = z
				}
				if math.IsNaN(minHeight) || minHeight > z {
					minHeight = z
				}
			}
		}
	}

	return maxHeight, minHeight
}

// getColor は、対象のポリゴンの高さから、ポリゴンの色を計算して、#RRGGBB 形式の文字列を返します。
// 計算には、対象のポリゴンの高さと、全てのポリゴンの高さの最大値と最小値を利用します。
func getColor(height, maxHeight, minHeight float64, topColor, bottomColor color.Color) string {
	if !floats.IsFinite(height) || !floats.IsFinite(maxHeight) || !floats.IsFinite(minHeight) {
		return colors.ColorToString(bottomColor)
	}

	n := (height - minHeight) / (maxHeight - minHeight)
	intermediate := colors.GetIntermediateColor(n, bottomColor, topColor)

	return colors.ColorToString(intermediate)
}
