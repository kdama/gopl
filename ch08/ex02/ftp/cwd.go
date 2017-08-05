package ftp

import (
	"log"
	"os"
	"path/filepath"
)

func (c *Conn) cwd(args []string) {
	if len(args) != 1 {
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	workDir := filepath.Join(c.workDir, args[0])
	target := filepath.Join(c.rootDir, workDir)
	_, err := os.Stat(target)
	if err != nil {
		log.Print(err)
		c.respond("550 Requested action not taken. File unavailable.")
		return
	}
	c.workDir = workDir
	c.respond("200 Command okay.")
}
