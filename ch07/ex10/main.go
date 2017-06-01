// ch07/ex10 は、sort.Interface を用いた IsPalindrome の実装です。
package main

import (
	"fmt"
	"sort"
)

// IsPalindrome は、列が "回文" であるかを返します。
func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i <= j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	x := sortInt([]int{1, 2, 3, 2, 1})
	fmt.Printf("IsPalindrome(%v) = %t\n", x, IsPalindrome(x))
	y := sortInt([]int{1, 2, 3, 3, 1})
	fmt.Printf("IsPalindrome(%v) = %t\n", y, IsPalindrome(y))
}

type sortInt []int

func (x sortInt) Len() int           { return len(x) }
func (x sortInt) Less(i, j int) bool { return x[i] < x[j] }
func (x sortInt) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
