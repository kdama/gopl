// ch03/ex03 は、高さに基づいて個々のポリゴンに色を付けながら、3-D 面の関数の SVG レンダリングを計算します。
// 頂点が赤となり、谷が青となるように、色付けします。
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	maxHeight, minHeight := getMaxMinHeight()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			color := getColor(getHeight(i, j), maxHeight, minHeight)

			// 出力する前に、全ての値が有限かどうかを調べます。
			if isFinite(ax) && isFinite(ay) &&
				isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) &&
				isFinite(dx) && isFinite(dy) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// isFinite は、f が有限の値かどうかを返します。
func isFinite(f float64) bool {
	if math.IsInf(f, 0) {
		return false
	}
	if math.IsNaN(f) {
		return false
	}
	return true
}

// getHeight は、ポリゴンの高さを計算します。
func getHeight(i, j int) float64 {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	return f(x, y)
}

// getMaxMinHeight は、全てのポリゴンの高さを求めて、高さの最大値と最小値を返します。
func getMaxMinHeight() (float64, float64) {
	maxHeight := math.NaN()
	minHeight := math.NaN()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z := getHeight(i, j)

			if isFinite(z) {
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
func getColor(height, maxHeight, minHeight float64) string {
	if !isFinite(height) || !isFinite(maxHeight) || !isFinite(minHeight) {
		return "#0000FF"
	}

	n := int((height - minHeight) / (maxHeight - minHeight) * 255)
	rr := fmt.Sprintf("%02x", n)
	gg := "00"
	bb := fmt.Sprintf("%02x", 255-n)

	return fmt.Sprintf("#%s%s%s", rr, gg, bb)
}
