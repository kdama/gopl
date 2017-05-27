// ch05/ex19 は、return 文を含まないのに与えられた値を返す関数 get を実行します。
package main

import (
	"fmt"
)

func get(in interface{}) (out interface{}) {
	out = in
	defer func() { recover() }()
	panic(in)
}

func main() {
	fmt.Println(get("Hello, world!"))
}
