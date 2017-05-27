// ch06/ex01 は、IntSet のいくつかの関数を実行します。
package main

import (
	"fmt"

	"github.com/kdama/gopl/ch06/ex01/intset"
)

func main() {
	is := &intset.IntSet{}

	is.Add(1)
	is.Add(2)
	is.Add(3)
	is.Add(42)

	fmt.Println(is.Len()) // "4"

	fmt.Println(is.Has(2)) // "true"
	is.Remove(2)
	fmt.Println(is.Len())  // "3"
	fmt.Println(is.Has(2)) // "false"

	is2 := is.Copy()
	is.Clear()
	fmt.Println(is.Len())  // "0"
	fmt.Println(is2.Len()) // "3"
}
