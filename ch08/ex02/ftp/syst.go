package ftp

func (c *Conn) syst(args []string) {
	c.respond("215 UNIX system type.")
}
