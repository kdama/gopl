package colors

import (
	"image/color"
	"testing"
)

func TestGetIntermediateColor(t *testing.T) {
	var tests = []struct {
		n      float64
		color0 color.Color
		color1 color.Color
		want   color.Color
	}{
		{0.0, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}, color.RGBA{0xff, 0x00, 0x00, 0xff}},
		{0.5, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}, color.RGBA{0x7f, 0x00, 0x7f, 0xff}},
		{1.0, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}},
	}

	for _, test := range tests {
		got := GetIntermediateColor(test.n, test.color0, test.color1)
		if got != test.want {
			t.Errorf("GetIntermediateColor(%f, %v, %v) = %v, want %v", test.n, test.color0, test.color1, got, test.want)
		}
	}
}

func TestColorToString(t *testing.T) {
	var tests = []struct {
		c    color.Color
		want string
	}{
		{color.RGBA{0x00, 0x00, 0x00, 0xff}, "#000000"},
		{color.RGBA{0x00, 0x00, 0xff, 0xff}, "#0000ff"},
		{color.RGBA{0x7f, 0x00, 0x7f, 0xff}, "#7f007f"},
		{color.RGBA{0xff, 0x00, 0x00, 0xff}, "#ff0000"},
	}

	for _, test := range tests {
		got := ColorToString(test.c)
		if got != test.want {
			t.Errorf("ColorToString(%v) = %s, want %s", test.c, got, test.want)
		}
	}
}

func TestColorFromString(t *testing.T) {
	var tests = []struct {
		colorCode string
		wantColor color.Color
		wantError bool
	}{
		{"#000000", color.RGBA{0x00, 0x00, 0x00, 0xff}, false},
		{"#ff0000", color.RGBA{0xff, 0x00, 0x00, 0xff}, false},
		{"#00ff00", color.RGBA{0x00, 0xff, 0x00, 0xff}, false},
		{"0000ff", color.RGBA{0x00, 0x00, 0xff, 0xff}, false},
		{"#zzzzzz", nil, true}, // Expects error.
	}

	for _, test := range tests {
		got, err := ColorFromString(test.colorCode)
		if test.wantError && err == nil {
			t.Errorf("ColorFromString(%s) = %v, expects error", test.colorCode, got)
		} else if got != test.wantColor {
			t.Errorf("ColorFromString(%s) = %v, want %v", test.colorCode, got, test.wantColor)
		}
	}
}
