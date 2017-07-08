// ch08/ex14 は、接続したクライアントに名前を尋ねる chat です。
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const timeout = 5 * time.Minute
const outbuffer = 64 // 送信用メッセージチャンネルのバッファサイズ

type client struct {
	name  string
	inch  chan<- string // 受信用メッセージチャンネル
	outch chan<- string // 送信用メッセージチャンネル
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.outch <- msg
			}

		case cli := <-entering:
			clients[cli] = true

			// 新しいクライアントに、現在のクライアントの集まりを知らせます。
			var onlines []string
			for c := range clients {
				onlines = append(onlines, c.name)
			}
			cli.outch <- fmt.Sprintf("%d clients: %s", len(clients), strings.Join(onlines, ", "))

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.outch)
		}
	}
}

func handleConn(conn net.Conn) {
	inch := make(chan string)
	outch := make(chan string, outbuffer)

	go clientReader(conn, inch)
	go clientWriter(conn, outch)

	// クライアントに名前を尋ねます。
	var who string

	outch <- "Input your name:"

	// タイムアウト時間内に名前を答えないクライアントは切断します。
	select {
	case in, ok := <-inch:
		if !ok {
			conn.Close()
			return
		}
		who = in
	case <-time.After(timeout):
		conn.Close()
		return
	}

	messages <- who + " has arrived"
	entering <- client{who, inch, outch}

	for {
		select {
		case in, ok := <-inch:
			if ok {
				messages <- who + ": " + in
			} else {
				leaving <- client{who, inch, outch}
				messages <- who + " has left"
				conn.Close()
				return
			}
		case <-time.After(timeout):
			leaving <- client{who, inch, outch}
			messages <- who + " has left"
			conn.Close()
			return
		}
	}
}

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
	close(ch)
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
