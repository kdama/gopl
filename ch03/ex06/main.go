// ch03/ex06 は、スーパーサンプリングによる高画質なマンデルブロフラクタルの PNG 画像を生成します。
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"

	"github.com/kdama/gopl/ch03/ex06/colors"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y0 := float64(py)/height*(ymax-ymin) + ymin
		y1 := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x0 := float64(px)/width*(xmax-xmin) + xmin
			x1 := (float64(px)+0.5)/width*(xmax-xmin) + xmin
			z0 := complex(x0, y0)
			z1 := complex(x1, y0)
			z2 := complex(x0, y1)
			z3 := complex(x1, y1)

			color := colors.GetAverageColor([]color.Color{
				mandelbrot(z0),
				mandelbrot(z1),
				mandelbrot(z2),
				mandelbrot(z3),
			})
			img.Set(px, py, color)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
