// ch08/ex13 は、5 分間、何も発言しないクライアントを切断する chat です。
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

type client struct {
	name string
	ch   chan<- string // an outgoing message channel
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
				cli.ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true

			// 新しいクライアントに、現在のクライアントの集まりを知らせます。
			var onlines []string
			for c := range clients {
				onlines = append(onlines, c.name)
			}
			cli.ch <- fmt.Sprintf("%d clients: %s", len(clients), strings.Join(onlines, ", "))

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	// クライアントが発言したことを通知します。
	talk := make(chan struct{})

	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{who, ch}

	input := bufio.NewScanner(conn)
	go func() {
		for {
			if input.Scan() {
				messages <- who + ": " + input.Text()
				talk <- struct{}{}
			} else {
				leaving <- client{who, ch}
				messages <- who + " has left"
				conn.Close()
				return
			}
		}
	}()
	// NOTE: ignoring potential errors from input.Err()
	for {
		select {
		case _, ok := <-talk:
			if !ok {
				conn.Close()
				return
			}
		case <-time.After(timeout):
			conn.Close()
			return
		}
	}
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
