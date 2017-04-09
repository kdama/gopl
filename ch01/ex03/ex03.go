// Package ex03 は、コマンドライン引数の表示を行います。
package ex03

import (
	"os"
	"strings"
)

// Echo1 は、for ループを用いてコマンドライン引数を表示します。
func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// fmt.Println(s)
}

// Echo2 は、for ループと range を用いてコマンドライン引数を表示します。
func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

// Echo3 は、strings.Join を用いてコマンドライン引数を表示します。
func Echo3() {
	strings.Join(os.Args[1:], " ")
	// fmt.Println(strings.Join(os.Args[1:], " "))
}
