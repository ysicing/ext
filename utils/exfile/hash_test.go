// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exfile

import "testing"

func TestMd5file(t *testing.T) {
	local := "/Users/ysicing/go/src/github.com/ysicing/ext/go.sum"
	md5key, err := Md5file(local)
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("path %v, md5 %v", local, md5key)
}

func TestMd5FromLocal(t *testing.T) {
	local := "/Users/ysicing/go/src/github.com/ysicing/ext/go.sum"
	md5key, err := Md5FromLocal(local)
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("path %v, md5 %v", local, md5key)
}
