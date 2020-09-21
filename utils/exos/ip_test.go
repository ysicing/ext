// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exos

import (
	"testing"
)

func TestIsValidIP(t *testing.T) {
	t.Logf("ip %s %t", "1829.123", IsValidIP("1829.123"))
	t.Logf("ip %s %t", "1829.123.1", IsValidIP("1829.123.1"))
	t.Logf("ip %s %t", "1829.123.2.1", IsValidIP("1829.123.2.1"))
	t.Logf("ip %s %t", "y.123", IsValidIP("y.123"))
	t.Logf("ip %s %t", "0.0.0.0", IsValidIP("0.0.0.0"))
	t.Logf("ip %s %t", "127.0.0.1", IsValidIP("127.0.0.1"))
	t.Logf("ip %s %t", "localhost", IsValidIP("localhost"))
	t.Logf("ip %s %t", "192.168.0.1", IsValidIP("192.168.0.1"))
	t.Logf("ipv2 %s %t", "1829.123", IsValidIPV2("1829.123"))
	t.Logf("ipv2 %s %t", "1829.123.1", IsValidIPV2("1829.123.1"))
	t.Logf("ipv2 %s %t", "1829.123.2.1", IsValidIPV2("1829.123.2.1"))
	t.Logf("ipv2 %s %t", "y.123", IsValidIPV2("y.123"))
	t.Logf("ipv2 %s %t", "0.0.0.0", IsValidIPV2("0.0.0.0"))
	t.Logf("ipv2 %s %t", "127.0.0.1", IsValidIPV2("127.0.0.1"))
	t.Logf("ipv2 %s %t", "localhost", IsValidIPV2("localhost"))
	t.Logf("ipv2 %s %t", "192.168.0.1", IsValidIPV2("192.168.0.1"))
}

func TestIsValidPortNum(t *testing.T) {
	t.Logf("port %v %t", 80, IsValidPortNum(80))
	t.Logf("port %v %t", -1, IsValidPortNum(-1))
	t.Logf("port %v %t", 64400, IsValidPortNum(64400))
	t.Logf("port %v %t", 1235000, IsValidPortNum(1235000))
}

func TestIsValidSocketAddr(t *testing.T) {
	t.Logf("addr %v %t", "127.0.0.1", IsValidSocketAddr("127.0.0.1"))
	t.Logf("addr %v %t", ":80", IsValidSocketAddr(":80"))
	t.Logf("addr %v %t", "127.0.0.0.0:890", IsValidSocketAddr("127.0.0.0.0:890"))
	t.Logf("addr %v %t", "11.11.11.11:80", IsValidSocketAddr("11.11.11.11:80"))
}

func TestLocalIP(t *testing.T) {
	t.Logf("localip %v, %v", LocalIP(), len(LocalIP()))
}
