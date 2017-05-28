// ch07/ex06 は、-temp フラグで指定された温度を、絶対温度で表示します。
package main

import (
	"flag"
	"fmt"

	"github.com/kdama/gopl/ch07/ex06/tempconv"
)

var temp = tempconv.KelvinFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
