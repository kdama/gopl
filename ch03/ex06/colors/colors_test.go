package colors

import (
	"image/color"
	"testing"
)

func TestGetAverageColor(t *testing.T) {
	var tests = []struct {
		colors []color.Color
		want   color.Color
	}{
		{[]color.Color{color.RGBA{0x11, 0x22, 0x33, 0xff}}, color.RGBA{0x11, 0x22, 0x33, 0xff}},
		{[]color.Color{color.RGBA{0x00, 0x00, 0x00, 0x00}, color.RGBA{0x22, 0x22, 0x22, 0x22}}, color.RGBA{0x11, 0x11, 0x11, 0x11}},
		{[]color.Color{color.RGBA{0x00, 0x00, 0x00, 0x00}, color.RGBA{0x22, 0x22, 0x22, 0x22}, color.RGBA{0x44, 0x44, 0x44, 0x44}}, color.RGBA{0x22, 0x22, 0x22, 0x22}},
	}

	for _, test := range tests {
		got := GetAverageColor(test.colors)
		if got != test.want {
			t.Errorf("GetAverageColor(%v) = %v, want %v", test.colors, got, test.want)
		}
	}
}
