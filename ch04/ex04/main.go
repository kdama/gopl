// ch04/ex04 は、スライスの左方向への回転を、1 回のパスで行います。
package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	rotate(s, 2)
	fmt.Println(s) // "[2 3 4 5 0 1]"
}

func rotate(s []int, n int) {
	num := n % len(s)
	double := append(s, s[:num]...)
	copy(s, double[num:num+len(s)])
}
