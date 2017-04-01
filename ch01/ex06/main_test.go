package main

import (
	"image/color"
	"testing"
)

func TestGetPalette(t *testing.T) {
	var tests = []struct {
		colorRes int
		want     []color.Color
	}{
		{1, []color.Color{color.Gray{0x00}}},
		{2, []color.Color{color.Gray{0x00}, color.Gray{0xff}}},
		{3, []color.Color{color.Gray{0x00}, color.Gray{0x7f}, color.Gray{0xff}}},
		{0, []color.Color{}},
		{-1, []color.Color{}},
	}

	for _, test := range tests {
		got := getPalette(test.colorRes)
		if len(got) != len(test.want) {
			t.Errorf("len(getPalette(%d)) = %d, want %d", test.colorRes, len(got), len(test.want))
		}
		for idx, gotValue := range got {
			if gotValue != test.want[idx] {
				t.Errorf("getPalette(%d)[%d] = %v, want %v", test.colorRes, idx, gotValue, test.want[idx])
			}
		}
	}
}

func TestGetGrayIndex(t *testing.T) {
	var tests = []struct {
		brightness float64
		palette    []color.Color
		want       uint8
	}{
		{0, []color.Color{color.Gray{0x00}}, 0},
		{1, []color.Color{color.Gray{0x00}}, 0},
		{0, []color.Color{color.Gray{0x00}, color.Gray{0xff}}, 0},
		{1, []color.Color{color.Gray{0x00}, color.Gray{0xff}}, 1},
		{0, []color.Color{color.Gray{0x00}, color.Gray{0x7f}, color.Gray{0xff}}, 0},
		{0.5, []color.Color{color.Gray{0x00}, color.Gray{0x7f}, color.Gray{0xff}}, 1},
		{1, []color.Color{color.Gray{0x00}, color.Gray{0x7f}, color.Gray{0xff}}, 2},
		{0, []color.Color{}, 0},
		{-1, []color.Color{}, 0},
	}

	for _, test := range tests {
		got := getGrayIndex(test.brightness, test.palette)
		if got != test.want {
			t.Errorf("getPalette(%f, %v) = %d, want %d", test.brightness, test.palette, got, test.want)
		}
	}
}
