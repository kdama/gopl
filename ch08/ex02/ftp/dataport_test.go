package ftp

import (
	"testing"
)

func TestDataportFromAddress(t *testing.T) {
	tests := []struct {
		address string
		want    dataport
		err     bool
	}{
		{
			"192.168.1.1:10",
			dataport{192, 168, 1, 1, 0, 10},
			false,
		},
		{
			"10.0.2.2:50000",
			dataport{10, 0, 2, 2, 195, 80},
			false,
		},
		{
			"192.168.1.1:foo",
			dataport{},
			true,
		},
		{
			"www.example.com:10",
			dataport{},
			true,
		},
	}
	for _, test := range tests {
		got, err := dataportFromAddress(test.address)
		if err == nil && test.err {
			t.Fatalf("Expects error but no error when dataportFromAddress(%q)", test.address)
		} else if err != nil && !test.err {
			t.Fatalf("Expects no error but error when dataportFromAddress(%q): %v", test.address, err)
		} else if !test.err && got.String() != test.want.String() {
			t.Fatalf("dataportFromAddress(%q) = %v, want %v", test.address, *got, test.want)
		}
	}
}

func TestDataportFromHostport(t *testing.T) {
	tests := []struct {
		hostport string
		want     dataport
		err      bool
	}{
		{
			"192,168,1,1,0,10",
			dataport{192, 168, 1, 1, 0, 10},
			false,
		},
		{
			"10,0,2,2,195,80",
			dataport{10, 0, 2, 2, 195, 80},
			false,
		},
		{
			"192,168,1,1,0,foo",
			dataport{},
			true,
		},
	}
	for _, test := range tests {
		got, err := dataportFromHostport(test.hostport)
		if err == nil && test.err {
			t.Fatalf("Expects error but no error when dataportFromHostport(%q)", test.hostport)
		} else if err != nil && !test.err {
			t.Fatalf("Expects no error but error when dataportFromHostport(%q): %v", test.hostport, err)
		} else if !test.err && got.String() != test.want.String() {
			t.Fatalf("dataportFromHostport(%q) = %v, want %v", test.hostport, *got, test.want)
		}
	}
}

func TestToAddress(t *testing.T) {
	tests := []struct {
		dataport dataport
		want     string
	}{
		{
			dataport{192, 168, 1, 1, 0, 10},
			"192.168.1.1:10",
		},
		{
			dataport{10, 0, 2, 2, 195, 80},
			"10.0.2.2:50000",
		},
	}
	for _, test := range tests {
		if got := test.dataport.toAddress(); got != test.want {
			t.Fatalf("(%v).toAddress() = %q, want %q", test.dataport, got, test.want)
		}
	}
}

func TestToHostport(t *testing.T) {
	tests := []struct {
		dataport dataport
		want     string
	}{
		{
			dataport{192, 168, 1, 1, 0, 10},
			"192,168,1,1,0,10",
		},
		{
			dataport{10, 0, 2, 2, 195, 80},
			"10,0,2,2,195,80",
		},
	}
	for _, test := range tests {
		if got := test.dataport.toHostport(); got != test.want {
			t.Fatalf("(%v).toHostport() = %q, want %q", test.dataport, got, test.want)
		}
	}
}
