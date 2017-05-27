// ch05/ex16 は、strings.Join の可変個引数関数としての実装です。
package main

import (
	"fmt"
	"strings"
)

func join(sep string, vals ...string) string {
	return strings.Join(vals, sep)
}

func main() {
	fmt.Println(join(","))               // ""
	fmt.Println(join(",", "foo"))        // "foo"
	fmt.Println(join(",", "foo", "bar")) // "foo,bar"
	fmt.Println(join(" ", "foo", "bar")) // "foo bar"
}
