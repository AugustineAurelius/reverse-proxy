package main

import (
	"flag"
	"io"
	"os"

	"github.com/AugustineAurelius/reverse-proxy/internal/listener"
	"github.com/AugustineAurelius/reverse-proxy/pkg/cli"
)

var (
	protocolFlag string
	portFlag     int
	clientsFlag  cli.StringList
)

func main() {
	flag.StringVar(&protocolFlag, "protocol", "tcp", `flag for choosing protocol connect: tcp, tcp4, tcp6, udp, udp4, udp6, ip, ip4, ip6, unix, unixgram, unixpacket`)
	flag.IntVar(&portFlag, "port", 8080, "port on which reverse-proxy would run")
	flag.Var(&clientsFlag, "clients", "client addresses. Write separated by commas")

	flag.Parse()

	if !cli.ValidateProtocols(protocolFlag) {
		io.WriteString(os.Stdout, "bad protocol chose one of: tcp, tcp4, tcp6, udp, udp4, udp6, ip, ip4, ip6, unix, unixgram, unixpacket")
		return
	}

	listener.New().Do(protocolFlag)
}
