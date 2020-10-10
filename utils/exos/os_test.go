// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exos

import "testing"

func TestOS(t *testing.T) {
	t.Logf("macos %t", IsMacOS())
	t.Logf("linux %t", IsLinux())
	t.Logf("unix %t", IsUnix())
	t.Logf("container %t", IsContainer())
	t.Logf("user %v", GetUser())
	t.Logf("username %v", GetUserName())
}
