// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exos

import (
	"net"
)

//CheckIP Verify IP address validity
func CheckIP(ip string) bool {
	i := net.ParseIP(ip)
	return i != nil && !i.IsLoopback() && !i.IsUnspecified()
}

//LocalIP 获取本机非loopback ip,默认第一个
func LocalIP() string {
	tables, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, t := range tables {
		addrs, err := t.Addrs()
		if err != nil {
			return ""
		}
		for _, a := range addrs {
			ipnet, ok := a.(*net.IPNet)
			if !ok || ipnet.IP.IsLoopback() {
				continue
			}
			if v4 := ipnet.IP.To4(); v4 != nil {
				return v4.String()
			}
		}
	}
	return ""
}
