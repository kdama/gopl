// Package conv は、いくつかの単位変換を行います。
package conv

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kdama/gopl/ch02/ex02/lengthconv"
	"github.com/kdama/gopl/ch02/ex02/massconv"
	"github.com/kdama/gopl/ch02/ex02/tempconv"
)

// Parse は、値とその単位からなる与えられた文字列を解釈して、その値と単位を返します。
func Parse(str string) (float64, string, error) {
	r := regexp.MustCompile("^([-+]?\\d*\\.?\\d+(?:[eE][-+]?\\d+)?)(.+)$")
	result := r.FindStringSubmatch(str)

	if len(result) != 3 {
		return 0, "", fmt.Errorf("invalid format: %s", str)
	}

	value, err := strconv.ParseFloat(result[1], 64)
	unit := result[2]

	if err != nil {
		return 0, "", err
	}

	return value, unit, nil
}

// Convert は、与えられた値と単位を結合した文字列と、関連する別の単位に変換した場合の文字列を返します。
func Convert(value float64, unit string) (string, string, error) {
	lowerUnit := strings.ToLower(unit)

	if lowerUnit == "ft" {
		f := lengthconv.Foot(value)
		return f.String(), lengthconv.FootToMeter(f).String(), nil
	} else if lowerUnit == "m" {
		m := lengthconv.Meter(value)
		return m.String(), lengthconv.MeterToFoot(m).String(), nil
	} else if lowerUnit == "lb" {
		p := massconv.Pound(value)
		return p.String(), massconv.PoundToKilogram(p).String(), nil
	} else if lowerUnit == "kg" {
		k := massconv.Kilogram(value)
		return k.String(), massconv.KilogramToPound(k).String(), nil
	} else if lowerUnit == "c" || lowerUnit == "°c" || lowerUnit == "℃" {
		c := tempconv.Celsius(value)
		return c.String(), tempconv.CToF(c).String(), nil
	} else if lowerUnit == "f" || lowerUnit == "°f" || lowerUnit == "℉" {
		f := tempconv.Fahrenheit(value)
		return f.String(), tempconv.FToC(f).String(), nil
	} else {
		return "", "", fmt.Errorf("unknown unit: %s", unit)
	}
}
