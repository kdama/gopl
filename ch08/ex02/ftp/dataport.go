package ftp

import "fmt"

type dataport struct {
	h1, h2, h3, h4 int // host
	p1, p2         int // port
}

func dataportFromAddress(hostport string) (*dataport, error) {
	var h1, h2, h3, h4, port int
	_, err := fmt.Sscanf(hostport, "%d.%d.%d.%d:%d", &h1, &h2, &h3, &h4, &port)
	if err != nil {
		return nil, err
	}

	p1 := port >> 8
	p2 := port & (1<<8 - 1)
	return &dataport{h1, h2, h3, h4, p1, p2}, nil
}

func dataportFromHostport(address string) (*dataport, error) {
	var h1, h2, h3, h4, p1, p2 int
	_, err := fmt.Sscanf(address, "%d,%d,%d,%d,%d,%d", &h1, &h2, &h3, &h4, &p1, &p2)
	if err != nil {
		return nil, err
	}
	return &dataport{h1, h2, h3, h4, p1, p2}, nil
}

func (d *dataport) toAddress() string {
	port := d.p1<<8 + d.p2
	return fmt.Sprintf("%d.%d.%d.%d:%d", d.h1, d.h2, d.h3, d.h4, port)
}

func (d *dataport) toHostport() string {
	return fmt.Sprintf("%d,%d,%d,%d,%d,%d", d.h1, d.h2, d.h3, d.h4, d.p1, d.p2)
}

func (d *dataport) String() string {
	return fmt.Sprintf("%v", *d)
}
