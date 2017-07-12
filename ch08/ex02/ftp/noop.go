package ftp

func (c *Conn) noop(args []string) {
	c.respond("200 Command okay.")
}
