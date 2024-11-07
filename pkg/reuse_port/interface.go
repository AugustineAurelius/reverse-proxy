package reuseport

import (
	"context"
	"fmt"
	"net"
	"time"
)

func Available() bool {
	return true
}

var listenConfig = net.ListenConfig{
	Control: Control,
}

func Listen(network, address string) (net.Listener, error) {
	return listenConfig.Listen(context.Background(), network, address)
}

func ListenPacket(network, address string) (net.PacketConn, error) {
	return listenConfig.ListenPacket(context.Background(), network, address)
}

func Dial(network, laddr, raddr string) (net.Conn, error) {
	return DialTimeout(network, laddr, raddr, time.Duration(0))
}

func DialTimeout(network, laddr, raddr string, timeout time.Duration) (net.Conn, error) {
	nla, err := ResolveAddr(network, laddr)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve local addr: %w", err)
	}
	d := net.Dialer{
		Control:   Control,
		LocalAddr: nla,
		Timeout:   timeout,
	}
	return d.Dial(network, raddr)
}
