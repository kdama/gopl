package ftp

import (
	"io"
	"net"
)

func (c *Conn) dataconnect() (io.ReadWriteCloser, error) {
	conn, err := net.Dial("tcp", c.dataport.toAddress())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
