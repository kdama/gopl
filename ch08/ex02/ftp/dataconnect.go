package ftp

import (
	"fmt"
	"io"
	"net"
)

func (c *Conn) dataconnect() (io.ReadWriteCloser, error) {
	switch c.dataconn {
	case pasv:
		conn, err := c.passive.Accept()
		if err != nil {
			return nil, err
		}
		return conn, nil
	case port:
		conn, err := net.Dial("tcp", c.dataport.toAddress())
		if err != nil {
			return nil, err
		}
		return conn, nil
	default:
		return nil, fmt.Errorf("No data connection")
	}
}

type dataconn int

const (
	none dataconn = iota
	pasv
	port
)
