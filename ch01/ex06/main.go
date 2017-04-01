package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles   = 5     // number of complete x oscillator revolutions
		res      = 0.001 // angular resolution
		size     = 100   // image canvas covers [-size..+size]
		nframes  = 64    // number of animation frames
		delay    = 8     // delay between frames in 10ms units
		colorRes = 32    // color resolution
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	palette := getPalette(colorRes)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), getGrayIndex(t/(cycles*2*math.Pi), palette))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func getPalette(colorRes int) []color.Color {
	if colorRes < 1 {
		return []color.Color{}
	}

	palette := []color.Color{color.Gray{0x00}}
	for i := 1; i < colorRes; i++ {
		palette = append(palette, color.Gray{uint8(0xff * i / (colorRes - 1))})
	}
	return palette
}

func getGrayIndex(brightness float64, palette []color.Color) uint8 {
	if brightness < 0 || brightness > 1 {
		return 0
	}
	return uint8(brightness * float64(len(palette)-1))
}
