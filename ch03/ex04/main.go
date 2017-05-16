// ch03/ex04 は、URL からパラメータ値を読み込んで、対応する 3-D 面の SVG を生成するサーバです。
// URL クエリを通して、以下のパラメータに対応します。
// - width       : キャンバスの幅
// - height      : キャンバスの高さ
// - cells       : 格子のます目の数
// - xyrange     : 軸の範囲 (-xyrange .. xyrange)
// - xyscale     : x 単位および y 単位当たりの画素数
// - zscale      : z 単位当たりの画素数
// - angle       : x, y 軸の角度
// - topColor    : 頂点の色 (例: ff0000)
// - bottomColor : 谷の色 (例: 0000ff)
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"image/color"

	"github.com/kdama/gopl/ch03/ex04/colors"
	"github.com/kdama/gopl/ch03/ex04/surface"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		width := parseFirstIntOrDefault(r.Form["width"], 600)
		height := parseFirstIntOrDefault(r.Form["height"], 320)
		cells := parseFirstIntOrDefault(r.Form["size"], 100)
		xyrange := parseFirstFloat64OrDefault(r.Form["xyrange"], 30)
		xyscale := parseFirstFloat64OrDefault(r.Form["xyscale"], float64(width/2)/xyrange)
		zscale := parseFirstFloat64OrDefault(r.Form["zscale"], float64(height)*0.4)
		angle := parseFirstFloat64OrDefault(r.Form["angle"], math.Pi/6)
		topColor := parseFirstColorOrDefault(r.Form["topColor"], color.RGBA{0xff, 0x00, 0x00, 0xff})
		bottomColor := parseFirstColorOrDefault(r.Form["bottomColor"], color.RGBA{0x00, 0x00, 0xff, 0xff})
		w.Header().Set("Content-Type", "image/svg+xml")
		surface.Render(w, width, height, cells, xyrange, xyscale, zscale, angle, topColor, bottomColor)
	}
	http.HandleFunc("/", handler)

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

// parseFirstFloat64OrDefault は、与えられた文字列の配列のうち最初の要素を、整数にパースして返します。
// パース可能な要素が 1 個もない場合は、与えられたデフォルト値を返します。
func parseFirstIntOrDefault(array []string, defaultValue int) int {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := strconv.Atoi(array[0])
	if err != nil {
		return defaultValue
	}
	return value
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

// parseFirstColorOrDefault は、与えられた文字列の配列のうち最初の要素を、color.Color にパースして返します。
// パース可能な要素が 1 個もない場合は、与えられたデフォルト値を返します。
func parseFirstColorOrDefault(array []string, defaultValue color.Color) color.Color {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := colors.ColorFromString(array[0])
	if err != nil {
		return defaultValue
	}
	return value
}
