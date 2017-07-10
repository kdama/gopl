package ftp

// ASCII モードにおいて、end-of-line シーケンスは CRLF であるべきです。
// https://tools.ietf.org/html/rfc959#page-19
func (c *Conn) eol() string {
	switch c.datatype {
	case ascii:
		return "\r\n"
	case image:
		return "\n"
	default:
		return "\n"
	}
}
