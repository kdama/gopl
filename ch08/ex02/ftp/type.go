package ftp

func (c *Conn) typeCommand(args []string) {
	if len(args) < 1 || len(args) > 2 {
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	if args[0] == "A" && (len(args) == 1 || args[1] == "N") {
		c.datatype = ascii
		c.respond("200 Command okay.")
	} else if args[0] == "I" && (len(args) == 1 || args[1] == "N") {
		c.datatype = image
		c.respond("200 Command okay.")
	} else {
		c.respond("504 Command not implemented for that parameter.")
	}
	return
}
