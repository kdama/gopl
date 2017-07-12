package ftp

import (
	"io"
	"log"
	"os"
)

func (c *Conn) stor(args []string) {
	if len(args) != 1 {
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	target := c.rootDir + "/" + args[0]
	file, err := os.Create(target)
	if err != nil {
		log.Print(err)
		c.respond("550 Requested action not taken. File unavailable.")
		return
	}
	c.respond("150 File status okay; about to open data connection.")

	conn, err := c.dataconnect()
	if err != nil {
		c.respond("425 Can't open data connection.")
		return
	}
	defer conn.Close()
	_, err = io.Copy(file, conn)
	if err != nil {
		log.Print(err)
		c.respond("450 Requested file action not taken. File unavailable.")
		return
	}
	c.respond("226 Closing data connection. Requested file action successful.")
}
