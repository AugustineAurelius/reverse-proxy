package cli

import "strings"

type StringList []string

func (s StringList) String() string {
	return strings.Join(s, ",")
}

func (s *StringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

func TrimAddresses(addresses string) []string {
	addresses = strings.Trim(addresses, "[]{}")
	return strings.Split(addresses, ",")
}

func ValidateProtocols(protocol string) bool {
	switch protocol {
	case "tcp":
		return true
	case "tcp4":
		return true
	case "tcp6":
		return true
	case "udp":
		return true
	case "udp4":
		return true
	case "udp6":
		return true
	case "ip":
		return true
	case "ip4":
		return true
	case "ip6":
		return true
	case "unix":
		return true
	case "unixgram":
		return true
	case "unixpacket":
		return true
	default:
		return false
	}
}
