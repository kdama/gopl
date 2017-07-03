// ch08/ex08 は、10 秒以内に何も叫ばないクライアントとの接続を切断する reverb2 です。
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	chat := make(chan struct{})
	disconnect := make(chan struct{})
	input := bufio.NewScanner(c)
	go func() {
		for {
			if input.Scan() {
				chat <- struct{}{}
			} else {
				disconnect <- struct{}{}
			}
		}
	}()
	for {
		select {
		case <-chat:
			// NOTE: ignoring potential errors from input.Err()
			go echo(c, input.Text(), 1*time.Second)
		case <-disconnect:
			c.Close()
			return
		case <-time.After(10 * time.Second):
			c.Close()
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
