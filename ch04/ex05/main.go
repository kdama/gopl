// ch04/ex05 は、隣接している重複をスライス内から除去します。
package main

import (
	"fmt"
)

func main() {
	s := []string{"A", "A", "B", "B", "A", "C"}
	fmt.Println(removeDup(s)) // "[A B A C]"
}

// removeDup は、隣接している重複をスライス内から除去します。
func removeDup(s []string) []string {
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}
