package ftp

import (
	"fmt"
	"log"
	"net"
)

func (c *Conn) pasv(args []string) {
	if len(args) > 0 {
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Print(err)
		c.respond("451 Requested action aborted. Local error in processing. (Failed to listen)")
		return
	}
	c.passive = listener
	_, port, err := net.SplitHostPort(c.passive.Addr().String())
	if err != nil {
		c.passive.Close()
		c.passive = nil
		log.Print(err)
		c.respond("451 Requested action aborted. Local error in processing. (Failed to get host)")
		return
	}
	host, _, err := net.SplitHostPort(c.conn.LocalAddr().String())
	if err != nil {
		c.passive.Close()
		c.passive = nil
		log.Print(err)
		c.respond("451 Requested action aborted. Local error in processing. (Failed to get port)")
		return
	}
	addr, err := dataportFromAddress(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		c.passive.Close()
		c.passive = nil
		log.Print(err)
		c.respond("451 Requested action aborted. Local error in processing. (Failed to get dataport)")
		return
	}
	c.dataconn = pasv
	c.respond(fmt.Sprintf("227 Entering Passive Mode (%s).", addr.toHostport()))
}
