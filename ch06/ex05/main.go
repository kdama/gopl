// ch06/ex05 は、uint 型を使った IntSet の Len, Remove, Clear, Copy 関数を実行します。
package main

import (
	"fmt"

	"github.com/kdama/gopl/ch06/ex05/intset"
)

func main() {
	is := &intset.IntSet{}

	is.Add(1)
	is.Add(2)
	is.Add(3)
	is.Add(42)

	fmt.Println(is)       // "{1 2 3 42}"
	fmt.Println(is.Len()) // "4"

	fmt.Println(is) // "{1 2 3 42}"
	is.Remove(2)
	fmt.Println(is) // "{1 3 42}"

	is2 := is.Copy()
	is.Clear()
	fmt.Println(is)  // "{}"
	fmt.Println(is2) // "{1 3 42}"
}
