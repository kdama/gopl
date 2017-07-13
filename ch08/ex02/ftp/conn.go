package ftp

import (
	"log"
	"net"
)

// Conn は、FTP 接続のためのコネクションをラップして、状態を保持します。
// 状態は、データポート、データタイプ、ルートディレクトリ、ワーキングディレクトリを含みます。
type Conn struct {
	conn     net.Conn
	dataport *dataport
	datatype datatype
	rootDir  string
	workDir  string
}

// NewConn は、FTP 接続のためのラップされたコネクションを返します。
func NewConn(conn net.Conn, rootDir string) Conn {
	return Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: rootDir,
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
