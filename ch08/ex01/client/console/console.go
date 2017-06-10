// Package console は、コンソールに対する操作を提供します。
package console

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode/utf8"
)

// Clear は、コンソールをクリアします。
func Clear() error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

// SprintTable は、表の行のスライスを受け取って、文字列による表現を返します。
func SprintTable(rows [][]string) string {
	if len(rows) == 0 || len(rows[0]) == 0 {
		return ""
	}

	var b bytes.Buffer
	lens := columnLens(rows)

	for i := range rows[0] {
		fmt.Fprintf(&b, "+-%s-", strings.Repeat("-", lens[i]))
	}
	fmt.Fprintf(&b, "+\n")

	for _, row := range rows {
		for i := range rows[0] {
			var val string
			if len(row) > i {
				val = row[i]
			}
			fmt.Fprintf(&b, "| % -*s ", lens[i], val)
		}
		fmt.Fprintf(&b, "|\n")
		for i := range rows[0] {
			fmt.Fprintf(&b, "+-%s-", strings.Repeat("-", lens[i]))
		}
		fmt.Fprintf(&b, "+\n")
	}

	return b.String()
}

// columnLens は、各列の、最大の文字列長を返します。
func columnLens(rows [][]string) []int {
	result := []int{}

	if len(rows) == 0 || len(rows[0]) == 0 {
		return result
	}

	for i := range rows[0] {
		max := 0
		for _, row := range rows {
			if len(row) > i {
				val := utf8.RuneCountInString(row[i])
				if max < val {
					max = val
				}
			}
		}
		result = append(result, max)
	}
	return result
}
