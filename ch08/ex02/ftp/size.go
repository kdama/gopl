package ftp

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (c *Conn) size(args []string) {
	if len(args) != 1 {
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	target := filepath.Join(c.rootDir, c.workDir, args[0])
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
	c.respond(fmt.Sprintf("213 %d", stat.Size()))
}
