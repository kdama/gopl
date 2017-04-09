// ex01 は、それを起動したコマンド名とコマンドライン引数を表示します。
package main

import (
	"fmt"
	"os"
	"strings"
)

func joinWithSpace(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	fmt.Println(joinWithSpace(os.Args))
}
