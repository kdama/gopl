// ch07/ex15 は、標準入力から読み込んだ式に対する評価結果を出力します。変数に対する値は、ユーザーに問い合わせます。
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"strconv"

	"github.com/kdama/gopl/ch07/ex15/eval"
)

func main() {
	fmt.Print("Expr: ")

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("ch07/ex14: %v", err)
	}

	expr, err := eval.Parse(string(b))
	if err != nil {
		log.Fatalf("ch07/ex14: %v", err)
	}

	env := inputEnv(expr)
	fmt.Printf("\n%s = %g\n", expr, expr.Eval(env))
}

func inputEnv(expr eval.Expr) map[eval.Var]float64 {
	env := make(map[eval.Var]float64)
	sc := bufio.NewScanner(os.Stdin)
	for _, v := range expr.Vars() {
		fmt.Printf("%s: ", v)
		if !sc.Scan() {
			log.Fatalf("ch07/ex14: failed to scan input")
		} else if err := sc.Err(); err != nil {
			log.Fatalf("ch07/ex14: %v", err)
		}

		val, err := strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			log.Fatalf("ch07/ex14: %v", err)
		}

		env[v] = val
	}
	return env
}
