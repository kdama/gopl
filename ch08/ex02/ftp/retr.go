package ftp

import (
	"bufio"
	"io"
	"log"
	"os"
)

func (c *Conn) retr(args []string) {
	if len(args) != 1 {
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	target := c.rootDir + "/" + args[0]
	file, err := os.Open(target)
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

	switch c.datatype {
	case ascii:
		r, w := bufio.NewReader(file), bufio.NewWriter(conn)
		for {
			line, isPrefix, err := r.ReadLine()
			if err == io.EOF {
				err := w.Flush()
				if err != nil {
					log.Print(err)
				}
				c.respond("226 Closing data connection. Requested file action successful.")
				return
			}
			if err != nil {
				log.Print(err)
				c.respond("450 Requested file action not taken. File unavailable.")
				return
			}
			_, err = w.Write(line)
			if err != nil {
				log.Print(err)
			}
			if !isPrefix {
				w.Write([]byte(c.eol()))
			}
		}
	case image:
		_, err := io.Copy(conn, file)
		if err != nil {
			log.Print(err)
			c.respond("450 Requested file action not taken. File unavailable.")
			return
		}
		c.respond("226 Closing data connection. Requested file action successful.")
		return
	}
}
