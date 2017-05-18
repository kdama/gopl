// ch04/ex02 は、標準入力の SHA256, SHA384 または SHA512 ハッシュを表示します。
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var modeFlag = flag.String("mode", "sha256", "hash: 'sha256' | 'sha384' | 'sha512'")

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ch04/ex02: %v", err)
		os.Exit(1)
	}

	switch *modeFlag {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256(bytes))
		return
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384(bytes))
		return
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512(bytes))
		return
	default:
		fmt.Fprintf(os.Stderr, "ch04/ex02: Invalid mode: %s\n", *modeFlag)
		os.Exit(1)
	}
}
