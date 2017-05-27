// ch06/ex03 は、IntSet の IntersectWith, DifferenceWith, SymmetricDifference 関数を実行します。
package main

import (
	"fmt"

	"github.com/kdama/gopl/ch06/ex03/intset"
)

func main() {
	s := &intset.IntSet{}
	t := &intset.IntSet{}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	t.Add(1)
	t.Add(4)

	fmt.Println(s) // "{1 2 3}"

	s.IntersectWith(t)
	fmt.Println(s) // "{1}"

	s.Add(1)
	s.Add(2)
	s.Add(3)

	s.DifferenceWith(t)
	fmt.Println(s) // "{2 3}"

	s.Add(1)
	s.Add(2)
	s.Add(3)

	s.SymmetricDifference(t)
	fmt.Println(s) // "{2 3 4}"
}
