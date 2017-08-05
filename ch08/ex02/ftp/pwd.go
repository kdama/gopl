package ftp

import "fmt"

func (c *Conn) pwd(args []string) {
	if len(args) > 0 {
		c.respond("501 Syntax error in parameters or arguments.")
	}
	c.respond(fmt.Sprintf("257 %q is current directory.", c.workDir))
}
