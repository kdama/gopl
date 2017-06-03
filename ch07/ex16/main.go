// ch07/ex16 は、Web ベースの電卓です。
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/kdama/gopl/ch07/ex15/eval"
)

var calcTemplate = template.Must(template.New("calc").Parse(`
<!doctype html>
<html>
<body>
<form method="post" action="/">
<p>Expr<br>
<textarea name="expr" cols="30" rows="5">{{.Expr}}</textarea></p>
<p>Env<br>
<textarea name="env" cols="30" rows="5">{{.Env}}</textarea></p>
<p><input type="submit" value="Calc"></p>
</form>
<p>{{.Message}}</p>
</body>
</html>
`))

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type calcdata struct {
	Expr    string
	Env     string
	Message interface{}
}

var defaults = calcdata{
	"a + b",
	"a = 1\nb = 2",
	"",
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		calcTemplate.Execute(w, defaults)
	} else if req.Method == "POST" {
		exprStr := req.FormValue("expr")
		envStr := req.FormValue("env")
		if expr, err := eval.Parse(exprStr); err != nil {
			w.WriteHeader(http.StatusBadRequest) // 400
			calcTemplate.Execute(w, calcdata{exprStr, envStr, err})
		} else {
			if env, err := parseEnv(envStr); err != nil {
				w.WriteHeader(http.StatusBadRequest) // 400
				calcTemplate.Execute(w, calcdata{exprStr, envStr, err})
			} else {
				if err := validate(expr, env); err != nil {
					w.WriteHeader(http.StatusBadRequest) // 400
					calcTemplate.Execute(w, calcdata{exprStr, envStr, err})
				} else {
					result := fmt.Sprintf("%s = %g", expr, expr.Eval(env))
					calcTemplate.Execute(w, calcdata{exprStr, envStr, result})
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		fmt.Fprintf(w, "method not allowed\n")
	}
}

func parseEnv(s string) (map[eval.Var]float64, error) {
	env := make(map[eval.Var]float64)
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return strings.ContainsRune(":=,{}'\"", r) || unicode.IsSpace(r)
	})
	for i := 0; i+1 < len(fields); i += 2 {
		k := strings.TrimSpace(fields[i])
		v := strings.TrimSpace(fields[i+1])

		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		env[eval.Var(k)] = val
	}
	return env, nil
}

func validate(expr eval.Expr, env map[eval.Var]float64) error {
	for _, v := range expr.Vars() {
		if _, ok := env[v]; !ok {
			return fmt.Errorf("variable %q is not defined", v)
		}
	}
	return nil
}
