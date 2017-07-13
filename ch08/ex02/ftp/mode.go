package ftp

// MODE は "S" (ストリーム) のみに対応します。
func (c *Conn) mode(args []string) {
	if len(args) != 1 {
		c.respond("501 Syntax error in parameters or arguments.")
		return
	}
	if args[0] == "S" {
		c.respond("200 Command okay.")
	} else {
		c.respond("504 Command not implemented for that parameter.")
	}
	return
}
