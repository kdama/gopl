// ch08/ex02 は、簡易な FTP サーバです。
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/kdama/gopl/ch08/ex02/ftp"
)

var port int
var rootDir string

func init() {
	flag.IntVar(&port, "port", 8000, "port number")
	flag.StringVar(&rootDir, "rootDir", ".", "root directory")
	flag.Parse()
}

func main() {
	server := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	ftp.Serve(ftp.NewConn(c, rootDir))
}
