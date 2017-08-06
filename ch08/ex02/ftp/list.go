package ftp

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (c *Conn) list(args []string) {
	target := filepath.Join(c.rootDir, c.workDir)
	if len(args) > 0 {
		target = filepath.Join(target, args[0])
	}
	f, err := os.Open(target)
	if err != nil {
		log.Print(err)
		c.respond("550 Requested action not taken. File unavailable.")
		return
	}
	c.respond("150 File status okay; about to open data connection.")

	w, err := c.dataconnect()
	if err != nil {
		log.Print(err)
		c.respond("425 Can't open data connection.")
		return
	}
	defer w.Close()
	stat, err := f.Stat()
	if err != nil {
		log.Print(err)
		c.respond("450 Requested file action not taken. File unavailable.")
		return
	}
	if stat.IsDir() {
		filenames, err := f.Readdirnames(0)
		if err != nil {
			log.Print(err)
			c.respond("550 Requested action not taken. File unavailable.")
			return
		}
		for _, filename := range filenames {
			_, err = fmt.Fprint(w, filename, c.eol())
			if err != nil {
				log.Print(err)
				c.respond("426 Connection closed; transfer aborted.")
				return
			}
		}
		c.respond("226 Closing data connection. Requested file action successful.")
		return
	}
	rel, err := filepath.Rel(c.rootDir, target)
	if err != nil {
		log.Print(err)
		c.respond("550 Requested action not taken. File unavailable.")
		return
	}
	_, err = fmt.Fprint(w, rel, c.eol())
	if err != nil {
		log.Print(err)
		c.respond("426 Connection closed; transfer aborted.")
		return
	}
	c.respond("226 Closing data connection. Requested file action successful.")
	return
}
