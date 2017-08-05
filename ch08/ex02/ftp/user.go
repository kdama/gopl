package ftp

import "fmt"
import "strings"

func (c *Conn) user(args []string) {
	c.respond(fmt.Sprintf("230 User %s logged in, proceed.", strings.Join(args, " ")))
}
