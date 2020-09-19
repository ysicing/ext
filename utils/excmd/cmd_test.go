// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package excmd

import "testing"

func TestDownloadFile(t *testing.T) {
	DownloadFile("https://github.com/ysicing/crtools/releases/download/0.0.2/crtools_darwin_amd64.sha256sum", "/tmp/crtools_darwin_amd64.sha256sum")
}
