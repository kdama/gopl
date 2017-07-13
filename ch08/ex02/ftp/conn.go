package ftp

import (
	"log"
	"net"
)

// Conn は、FTP 接続のためのコネクションをラップして、状態を保持します。
// 状態は、データポート、データタイプ、ルートディレクトリを含みます。
type Conn struct {
	conn     net.Conn
	dataconn dataconn
	dataport *dataport
	datatype datatype
	passive  net.Listener
	rootDir  string
}

// NewConn は、FTP 接続のためのラップされたコネクションを返します。
func NewConn(conn net.Conn, rootDir string) Conn {
	return Conn{
		conn:    conn,
		rootDir: rootDir,
	}
}

// Close は、コネクションを閉じます。
func (c *Conn) Close() error {
	err := c.conn.Close()
	if err != nil {
		log.Print(err)
	}
	return err
}
