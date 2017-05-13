// ch03/ex09 は、URL からパラメータ値を読み込んで、対応するマンデルブロフラクタルの画像を生成するサーバです。
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		x := parseFirstFloat64OrDefault(r.Form["x"], 0)
		y := parseFirstFloat64OrDefault(r.Form["y"], 0)
		zoom := parseFirstFloat64OrDefault(r.Form["zoom"], 0)
		renderPNG(w, x, y, zoom)
	}
	http.HandleFunc("/", handler)

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func renderPNG(out io.Writer, x, y, zoom float64) {
	const (
		width, height = 1024, 1024
	)

	m := math.Exp2(1 - zoom)
	xmin, xmax := x-m, x+m
	ymin, ymax := y-m, y+m

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
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

// parseFirstFloat64OrDefault は、与えられた文字列の配列のうち最初の要素を、浮動小数点数にパースして返します。
// パース可能な要素が 1 個もない場合は、与えられたデフォルト値を返します。
func parseFirstFloat64OrDefault(array []string, defaultValue float64) float64 {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := strconv.ParseFloat(array[0], 64)
	if err != nil {
		return defaultValue
	}
	return value
}
