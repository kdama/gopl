// ex08 は、URL にある内容を表示します。
// URL に接頭辞 http:// がない場合は、追加してから内容を取得します。
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		completeURL := appendPrefix(url)
		resp, err := http.Get(completeURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", completeURL, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// appendPrefix は、与えられた文字列が接頭辞 http:// を持たなければ、それを追加した文字列を返します。
// 既に接頭辞 http:// を持っていれば、与えられた文字列をそのまま返します。
func appendPrefix(str string) string {
	if !strings.HasPrefix(str, "http://") {
		return "http://" + str
	}
	return str
}
