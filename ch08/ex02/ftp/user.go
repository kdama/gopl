package ftp

func (c *Conn) user(args []string) {
	c.respond("230 User logged in, proceed.")
}
