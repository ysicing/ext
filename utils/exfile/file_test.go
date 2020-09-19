// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exfile

import "testing"

func TestWriteFile(t *testing.T) {
	fp := "/tmp/localtest/localtest"
	WriteFile(fp, "test")
	if !CheckFileExistsv2(fp) {
		t.Error("error")
	}
}
