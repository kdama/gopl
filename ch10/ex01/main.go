// ch10/ex01 は、標準入力から読み込んだ画像を、-format フラグで指定された形式に変換します。
// 入力と出力ともに、GIF, PNG, JPEG に対応します。
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

var format string

func init() {
	flag.StringVar(&format, "format", "jpeg", "output format")
	flag.Parse()
}

func main() {
	if err := convert(os.Stdin, os.Stdout, format); err != nil {
		log.Fatalf("ch10/ex01: %v\n", err)
	}
}

func convert(in io.Reader, out io.Writer, format string) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	switch format {
	case "gif":
		return gif.Encode(out, img, nil)
	case "jpg", "jpeg":
		return jpeg.Encode(out, img, nil)
	case "png":
		return png.Encode(out, img)
	default:
		return fmt.Errorf("unknown format: %s", format)
	}
}
