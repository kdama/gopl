package ftp

import (
	"fmt"
	"log"
	"os"
)

func (c *Conn) size(args []string) {
	if len(args) != 1 {
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	target := c.rootDir + "/" + c.workDir + "/" + args[0]
	file, err := os.Open(target)
	if err != nil {
		log.Print(err)
		c.respond("550 Requested action not taken. File unavailable.")
		return
	}
	stat, err := file.Stat()
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
	c.respond(fmt.Sprintf("213 %d", stat.Size()))
}
