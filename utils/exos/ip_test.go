// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exos

import (
	"testing"
)

func TestCheckIP(t *testing.T) {
	t.Logf("ip %s %t", "1829.123", CheckIP("1829.123"))
	t.Logf("ip %s %t", "1829.123.1", CheckIP("1829.123.1"))
	t.Logf("ip %s %t", "1829.123.2.1", CheckIP("1829.123.2.1"))
	t.Logf("ip %s %t", "y.123", CheckIP("y.123"))
	t.Logf("ip %s %t", "0.0.0.0", CheckIP("0.0.0.0"))
	t.Logf("ip %s %t", "127.0.0.1", CheckIP("127.0.0.1"))
	t.Logf("ip %s %t", "localhost", CheckIP("localhost"))
	t.Logf("ip %s %t", "192.168.0.1", CheckIP("192.168.0.1"))
}
