package ftp

func (c *Conn) stru(args []string) {
	if len(args) != 1 {
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	if args[0] == "F" {
		c.respond("200 Command okay.")
	} else {
		c.respond("504 Command not implemented for that parameter.")
	}
	return
}
