// ex04 は、標準入力または指定されたファイルに 2 回以上現れた行の数とそのテキストを表示します。
// ファイルが指定された場合は、重複した行のそれぞれが含まれていた全てのファイル名を表示します。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	occurrences := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, occurrences)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, occurrences)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, occurrences[line])
		}
	}
}

// countLines は、入力を 1 行ずつ読み込み、出現した行の回数と入力の名前を記録します。
func countLines(f *os.File, counts map[string]int, occurrences map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		occurrences[input.Text()] = appendAsSet(occurrences[input.Text()], f.Name())
	}
}

// appendAsSet は、与えられた文字列の配列 set に、与えられた文字列 str を追加します。
// ただし、set が str を既に含んでいる場合は、str を追加しません。
func appendAsSet(set []string, str string) []string {
	if !includes(set, str) {
		return append(set, str)
	}
	return set
}

// includes は、与えられた文字列の配列 array が、与えられた文字列 str を含んでいるかどうかを返します。
func includes(array []string, str string) bool {
	for _, value := range array {
		if value == str {
			return true
		}
	}
	return false
}
