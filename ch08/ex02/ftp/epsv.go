package ftp

func (c *Conn) epsv(args []string) {
	c.respond("522 Network protocol not supported. Use IPv4.")
}
