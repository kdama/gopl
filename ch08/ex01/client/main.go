// ch08/ex01/client は、ビジネスオフィスで見かける壁にかかった複数の時計に似せた表を表示します。
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/kdama/gopl/ch08/ex01/client/console"
)

// Server は、時刻を報告するサーバーです。
// サーバの名前、アドレス、最新の出力を持ちます。
type Server struct {
	name, address, output string
}

func main() {
	servers, err := parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		conn, err := net.Dial("tcp", server.address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(conn, server)
	}

	// 1 秒ごとに表を更新します。
	for {
		var data [][]string
		for _, server := range servers {
			data = append(data, []string{server.name, server.output})
		}

		table := console.SprintTable(data)
		console.Clear()
		fmt.Fprintf(os.Stdout, table)

		time.Sleep(time.Second)
	}
}

// mustCopy は、Reader から値を読み込む度に、その値をサーバーの最新の出力として更新します。
func mustCopy(src io.Reader, server *Server) {
	sc := bufio.NewScanner(src)
	for sc.Scan() {
		server.output = sc.Text()
		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

// parse は、"name=address" の形式の文字列を受け取って、時刻を報告するサーバーを返します。
func parse(args []string) (servers []*Server, err error) {
	for _, arg := range args {
		s := strings.SplitN(arg, "=", 2)

		if len(s) != 2 {
			return nil, fmt.Errorf("failed to parse 'name=address': %s", arg)
		}

		name, address := s[0], s[1]
		servers = append(servers, &Server{name, address, ""})
	}
	return
}
