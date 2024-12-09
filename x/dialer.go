package x

import "net"

type Dialer struct{}

var DefaultDial Dialer

func (Dialer) Dial(network string, addr string) (net.Conn, error) {
	return net.Dial(network, addr)
}
