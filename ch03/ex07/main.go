// ch03/ex07 は、
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			if equalsFloat64(real(z), 1) && equalsFloat64(imag(z), 0) {
				return color.RGBA{255 - contrast*i, 0, 0, 0xff}
			} else if equalsFloat64(real(z), -1) && equalsFloat64(imag(z), 0) {
				return color.RGBA{0, 255 - contrast*i, 0, 0xff}
			} else if equalsFloat64(real(z), 0) && equalsFloat64(imag(z), 1) {
				return color.RGBA{0, 0, 255 - contrast*i, 0xff}
			}
			return color.RGBA{255 - contrast*i, 255 - contrast*i, 0, 0xff}
		}
	}
	return color.Black
}

func equalsFloat64(x, y float64) bool {
	return math.Abs(x-y) < 1e-6
}
