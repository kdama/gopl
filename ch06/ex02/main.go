// ch06/ex01 は、IntSet のいくつかの関数を実行します。
package main

import (
	"fmt"

	"github.com/kdama/gopl/ch06/ex02/intset"
)

func main() {
	is := &intset.IntSet{}

	is.Add(1)

	fmt.Println(is.Has(1)) // "true"
	fmt.Println(is.Has(2)) // "false"
	fmt.Println(is.Has(3)) // "false"
	fmt.Println(is.Has(4)) // "false"
	fmt.Println(is.Has(5)) // "false"

	is.AddAll(2, 3, 4)

	fmt.Println(is.Has(1)) // "true"
	fmt.Println(is.Has(2)) // "true"
	fmt.Println(is.Has(3)) // "true"
	fmt.Println(is.Has(4)) // "true"
	fmt.Println(is.Has(5)) // "false"
}
