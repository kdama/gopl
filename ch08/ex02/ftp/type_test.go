package ftp

import (
	"net"
	"testing"
)

func TestTypeCommand(t *testing.T) {
	tests := []struct {
		args []string
		from datatype
		want datatype
	}{
		{[]string{"A"}, ascii, ascii},
		{[]string{"I"}, ascii, image},
		{[]string{"A"}, image, ascii},
		{[]string{"I"}, image, image},
		{[]string{"A", "N"}, ascii, ascii},
		{[]string{"I", "N"}, ascii, image},
		{[]string{"A", "N"}, image, ascii},
		{[]string{"I", "N"}, image, image},

		{[]string{"X"}, ascii, ascii},
		{[]string{"ASCII"}, image, image},
		{[]string{"IMAGE"}, ascii, ascii},
		{[]string{"A", "X"}, image, image},
		{[]string{"I", "X"}, ascii, ascii},
	}
	for _, test := range tests {
		server, client := net.Pipe()
		defer server.Close()
		client.Close()
		conn := &Conn{
			conn:     server,
			datatype: test.from,
		}
		conn.typeCommand(test.args)
		if got := conn.datatype; got != test.want {
			t.Fatalf("typeCommand(%q) sets type to %v, want %v", test.args, got, test.want)
		}
	}
}
