// ch12/ex01 は、キーが構造体から配列であるマップを表示できるように拡張された Display です。
package main

import (
	"github.com/kdama/gopl/ch12/ex01/display"
)

func main() {
	display.Display("map with struct key", map[struct{ x, y int }]int{
		struct{ x, y int }{}:      0,
		struct{ x, y int }{1, 10}: 100,
	})

	display.Display("map with array key", map[[3]int]int{
		[3]int{}:        100,
		[3]int{0, 0, 1}: 100,
		[3]int{1, 4, 7}: 100,
	})
}
