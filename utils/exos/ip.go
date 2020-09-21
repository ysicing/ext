// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exos

import (
	"net"
	"strconv"
)

//IsValidIP Verify IP address validity
func IsValidIP(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	}
	return true
}

// IsValidIPV2 确定是合法ip且不是loop ip 或者 0.0.0.0
func IsValidIPV2(ip string) bool {
	i := net.ParseIP(ip)
	return i != nil && !i.IsLoopback() && !i.IsUnspecified()
}

// IsValidPortNum tests that the argument is a valid, non-zero port number.
func IsValidPortNum(port int) bool {
	if 1 <= port && port <= 65535 {
		return true
	}
	return false
}

// IsValidSocketAddr checks that string represents a valid socket address
// as defined in RFC 789. (e.g 0.0.0.0:10254 or [::]:10254))
func IsValidSocketAddr(value string) bool {
	ip, port, err := net.SplitHostPort(value)
	if err != nil {
		return false
	}
	portInt, _ := strconv.Atoi(port)
	return IsValidIP(ip) && IsValidPortNum(portInt)
}

//LocalIP 获取本机非loopback ip,默认第一个
func LocalIP() (addr []string) {
	tables, err := net.Interfaces()
	if err != nil {
		return nil
	}

	for _, t := range tables {
		addrs, err := t.Addrs()
		if err != nil {
			return nil
		}
		for _, a := range addrs {
			ipnet, ok := a.(*net.IPNet)
			if !ok || ipnet.IP.IsLoopback() {
				continue
			}
			if v4 := ipnet.IP.To4(); v4 != nil {
				addr = append(addr, v4.String())
			}
		}
	}
	return addr
}
