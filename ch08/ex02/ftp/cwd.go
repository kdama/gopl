package ftp

func (c *Conn) cwd(args []string) {
	if len(args) != 1 {
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	c.workDir = args[0]
	c.respond("200 Command okay.")
}
