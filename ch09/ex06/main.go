// ch08/ex05 は、マンデルブロフラクタルのレンダリングを並列に実行します。
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

func main() {
	// 8 個のゴルーチンを使ってマンデルブロをレンダリングします。
	img := render(8)
	png.Encode(os.Stdout, img)
}

func render(worker int) *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	for w := 0; w < worker; w++ {
		wg.Add(1)
		minHeight := height * w / worker
		maxHeight := height * (w + 1) / worker
		if w == worker-1 {
			// 最後のゴルーチンは、残りの全ての領域を担当します。
			maxHeight = height
		}
		go func() {
			defer wg.Done()
			for py := minHeight; py < maxHeight; py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
		}()
	}
	wg.Wait()
	return img
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
