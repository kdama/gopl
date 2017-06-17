// ch08/ex03 は、標準入力が閉じられた後でも reverb1 サーバの最後のエコーを表示可能な netcat3 です。
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	tcpconn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("ch08/ex03: conn is not TCP network connection")
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)

	tcpconn.CloseWrite()
	<-done // wait for background goroutine to finish
	tcpconn.CloseRead()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
