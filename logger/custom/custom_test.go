// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package custom

import (
	"testing"
)

func TestWriteLogs(t *testing.T) {
	err := WriteLogs("default", "demo", "hhhhhh")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestReadLogs(t *testing.T) {
	info, err := ReadLogs("default", "demo")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("data: %v", info)
}
