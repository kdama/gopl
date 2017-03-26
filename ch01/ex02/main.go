package main

import (
	"fmt"
	"os"
)

func addIndex(args []string) []string {
	var r []string
	for idx, arg := range args {
		r = append(r, fmt.Sprintf("%d %s", idx, arg))
	}
	return r
}

func main() {
	for _, arg := range addIndex(os.Args[1:]) {
		fmt.Println(arg)
	}
}
