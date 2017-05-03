// Package colors は、色に関する処理を行います。
package colors

import (
	"fmt"
	"image/color"

	"strconv"

	"github.com/kdama/gopl/ch03/ex04/floats"
)

// GetIntermediateColor は、n (0 <= n <= 1) の値に対応する、color0 と color1 の間の色を返します。
// n が 0 以下のときは、color0 を返します。n が 1 以上のときは、color1 を返します。
func GetIntermediateColor(n float64, color0, color1 color.Color) color.Color {
	if n <= 0 || !floats.IsFinite(n) {
		return color0
	}
	if n >= 1 {
		return color1
	}

	r0, g0, b0, a0 := color0.RGBA()
	r1, g1, b1, a1 := color1.RGBA()

	r := uint8(n*(float64(r1%256)-float64(r0%256)) + float64(r0%256))
	g := uint8(n*(float64(g1%256)-float64(g0%256)) + float64(g0%256))
	b := uint8(n*(float64(b1%256)-float64(b0%256)) + float64(b0%256))
	a := uint8(n*(float64(a1%256)-float64(a0%256)) + float64(a0%256))

	return color.RGBA{r, g, b, a}
}

// ColorToString は、color.Color に対応する、#RRGGBB 形式の文字列を返します。
func ColorToString(c color.Color) string {
	r, g, b, _ := c.RGBA()

	rr := fmt.Sprintf("%02x", r%256)
	gg := fmt.Sprintf("%02x", g%256)
	bb := fmt.Sprintf("%02x", b%256)

	return fmt.Sprintf("#%s%s%s", rr, gg, bb)
}

// ColorFromString は、RRGGBB または #RRGGBB 形式の文字列に対応する、color.Color を返します。
func ColorFromString(colorCode string) (color.Color, error) {
	if len(colorCode) < 6 {
		return nil, fmt.Errorf("Invalid color code: %s", colorCode)
	}

	var s string

	if colorCode[0] == '#' {
		s = colorCode[1:]
	} else {
		s = colorCode
	}

	rr := s[0:2]
	gg := s[2:4]
	bb := s[4:6]

	r, err := strconv.ParseUint(rr, 16, 16)

	if err != nil {
		return nil, err
	}

	g, err := strconv.ParseUint(gg, 16, 16)

	if err != nil {
		return nil, err
	}

	b, err := strconv.ParseUint(bb, 16, 16)

	if err != nil {
		return nil, err
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}, nil
}
