// ch12/ex01 は、キーが構造体から配列であるマップを表示できるように拡張された Display です。
package main

import (
	"github.com/kdama/gopl/ch12/ex02/display"
)

type cycle struct {
	value int
	tail  *cycle
}

func main() {
	var c cycle
	c = cycle{42, &c}

	display.Display("cycle", c)
}
