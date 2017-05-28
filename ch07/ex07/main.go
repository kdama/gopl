// ch07/ex07 は、-temp フラグで指定された温度を、摂氏で表示します。
// ただし、ヘルプメッセージに ° を表示しません。
package main

import (
	"flag"
	"fmt"

	"github.com/kdama/gopl/ch07/ex07/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
