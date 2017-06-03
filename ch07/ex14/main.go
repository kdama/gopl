// ch07/ex14 は、文字列を三項演算子を含む式としてパースした構文ツリーを、式のような文字列で表示します。
package main

import (
	"fmt"
	"log"

	"github.com/kdama/gopl/ch07/ex14/eval"
)

func main() {
	expr, err := eval.Parse("x ? pow(x, 3) : pow(y, 3)")
	if err != nil {
		log.Fatalf("ch07/ex14: %v", err)
	}
	fmt.Printf("%s\n", expr)
}
