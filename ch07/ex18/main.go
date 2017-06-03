// ch07/ex18 は、XML ドキュメントを読み込んで、ツリーを構築します。
package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/kdama/gopl/ch07/ex18/xmlnode"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	node, err := xmlnode.Parse(dec)
	if err != nil {
		log.Fatalf("ch07/ex18: %v", err)
	}
	fmt.Printf("%s\n", node)
}
