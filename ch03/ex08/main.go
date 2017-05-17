// ch03/ex05 はマンデルブロフラクタルのフルカラーの PNG 画像を生成します。
// 点 z が半径 2 の円を「出る」までに掛かった回数 mod 6 について、0 から順に、赤、白、緑、白、青、白、で塗ります。
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"

	"github.com/kdama/gopl/ch03/ex08/complexfloat"
	"github.com/kdama/gopl/ch03/ex08/complexrat"
)

const (
	width, height = 1024, 1024
	contrast      = 15
)

// complex64, complex128, Float, Rat
var typeFlag = flag.String("type", "complex128", "type: 'complex64' | 'complex128' | 'bigfloat' | 'bigrat'")

var iterationsFlag = flag.Int("iterations", 200, "iterations: int")
var verboseFlag = flag.Bool("verbose", false, "verbose: bool")
var xFlag = flag.Int("x", 0, "center x: int")
var yFlag = flag.Int("y", 0, "center y: int")
var precFlag = flag.Uint("prec", 1024, "bigfloat prec: uint")

// 2^(1-zoom) = -xmin = xmax = -ymin = ymax
var zoomFlag = flag.Uint("zoom", 0, "zoom bit: uint")

func main() {
	flag.Parse()

	if *typeFlag == "complex64" {
		mainComplex64()
		return
	} else if *typeFlag == "complex128" {
		mainComplex128()
		return
	} else if *typeFlag == "bigfloat" {
		mainBigFloat()
		return
	} else if *typeFlag == "bigrat" {
		mainBigRat()
		return
	}

	fmt.Fprintf(os.Stderr, "Invalid type: %s\n", *typeFlag)
	os.Exit(1)
}

func mainComplex64() {
	xcenter := float32(*xFlag)
	ycenter := float32(*yFlag)
	invzoom := 2 / float32(int64(1)<<*zoomFlag)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*invzoom*2 - invzoom + ycenter
		for px := 0; px < width; px++ {
			x := float32(px)/height*invzoom*2 - invzoom + xcenter
			z := complex(x, y)
			img.Set(px, py, mandelbrotComplex64(z))
		}
		if *verboseFlag {
			fmt.Fprintf(os.Stderr, "c064 (%d/%d)\n", py, height)
		}
	}
	png.Encode(os.Stdout, img)
}

func mainComplex128() {
	xcenter := float64(*xFlag)
	ycenter := float64(*yFlag)
	invzoom := 2 / float64(int64(1)<<*zoomFlag)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*invzoom*2 - invzoom + ycenter
		for px := 0; px < width; px++ {
			x := float64(px)/height*invzoom*2 - invzoom + xcenter
			z := complex(x, y)
			img.Set(px, py, mandelbrotComplex128(z))
		}
		if *verboseFlag {
			fmt.Fprintf(os.Stderr, "c128 (%d/%d)\n", py, height)
		}
	}
	png.Encode(os.Stdout, img)
}

func mainBigFloat() {
	xcenter := float64(*xFlag)
	ycenter := float64(*yFlag)
	invzoom := 2 / float64(int64(1)<<*zoomFlag)
	prec := *precFlag

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	dy := big.NewFloat(invzoom * 2 / height)
	dx := big.NewFloat(invzoom * 2 / width)
	y := big.NewFloat(-invzoom + ycenter).SetPrec(prec)
	for py := 0; py < height; py++ {
		x := big.NewFloat(-invzoom + xcenter).SetPrec(prec)
		for px := 0; px < width; px++ {
			z := &complexfloat.ComplexFloat{
				Re:   x,
				Im:   y,
				Prec: prec,
			}
			img.Set(px, py, mandelbrotBigFloat(z))
			x.Add(x, dx)
		}
		y.Add(y, dy)
		if *verboseFlag {
			fmt.Fprintf(os.Stderr, "bflo (%d/%d)\n", py, height)
		}
	}
	png.Encode(os.Stdout, img)
}

func mainBigRat() {
	xcenter := int64(*xFlag)
	ycenter := int64(*yFlag)
	zoom := int64(1) << *zoomFlag

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	dy := big.NewRat(4, height*zoom)
	dx := big.NewRat(4, width*zoom)
	y := big.NewRat(ycenter*zoom-2, zoom)
	for py := 0; py < height; py++ {
		x := big.NewRat(xcenter*zoom-2, zoom)
		for px := 0; px < width; px++ {
			z := &complexrat.ComplexRat{
				Re: x,
				Im: y,
			}
			img.Set(px, py, mandelbrotBigRat(z))
			x.Add(x, dx)
		}
		if *verboseFlag {
			fmt.Fprintf(os.Stderr, "brat (%d/%d)\n", py, height)
		}
		y.Add(y, dy)
	}
	png.Encode(os.Stdout, img)
}

func mandelbrotComplex64(z complex64) color.Color {
	iterations := uint8(*iterationsFlag)
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotComplex128(z complex128) color.Color {
	iterations := uint8(*iterationsFlag)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z *complexfloat.ComplexFloat) color.Color {
	prec := *precFlag
	iterations := uint8(*iterationsFlag)
	two := big.NewFloat(2).SetPrec(prec)
	v := &complexfloat.ComplexFloat{
		Re:   big.NewFloat(0).SetPrec(prec),
		Im:   big.NewFloat(0).SetPrec(prec),
		Prec: prec,
	}
	for n := uint8(0); n < iterations; n++ {
		v.Square().Add(z)
		if v.AbsCompare(two) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotBigRat(z *complexrat.ComplexRat) color.Color {
	iterations := uint8(*iterationsFlag)
	two := big.NewRat(2, 1)
	v := &complexrat.ComplexRat{
		Re: big.NewRat(0, 1),
		Im: big.NewRat(0, 1),
	}
	for n := uint8(0); n < iterations; n++ {
		v.Square().Add(z)
		if v.AbsCompare(two) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
