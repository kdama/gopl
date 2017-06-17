// ch08/ex04 は、活動中の echo ゴルーチンを数え、0 になったら TCP 接続の書き込み側を閉じるように変更された reverb2 です。
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c *net.TCPConn) {
	var wg sync.WaitGroup

	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(c, input.Text(), 1*time.Second)
		}()
	}

	wg.Wait()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	tcpl, ok := l.(*net.TCPListener)
	if !ok {
		log.Fatal("ch08/ex04: l is not TCP network listener")
	}

	for {
		conn, err := tcpl.AcceptTCP()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
