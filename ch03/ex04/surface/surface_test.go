package surface

import (
	"image/color"
	"testing"
)

func TestGetColor(t *testing.T) {
	var tests = []struct {
		height      float64
		maxHeight   float64
		minHeight   float64
		topColor    color.Color
		bottomColor color.Color
		want        string
	}{
		{0.0, 1, 0, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}, "#0000ff"},
		{0.5, 1, 0, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}, "#7f007f"},
		{1.0, 1, 0, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}, "#ff0000"},
	}

	for _, test := range tests {
		got := getColor(test.height, test.maxHeight, test.minHeight, test.topColor, test.bottomColor)
		if got != test.want {
			t.Errorf("getColor(%f, %f, %f) = %s, want %s", test.height, test.maxHeight, test.minHeight, got, test.want)
		}
	}
}
