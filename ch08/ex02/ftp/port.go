package ftp

import "log"

func (c *Conn) port(args []string) {
	if len(args) != 1 {
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	dataport, err := dataportFromHostport(args[0])
	if err != nil {
		log.Print(err)
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	c.dataport = *dataport
	c.respond("200 Command okay.")
}
