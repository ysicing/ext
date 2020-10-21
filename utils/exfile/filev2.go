// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exfile

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	FilePath = "."
)

func init() {
	FilePath, _ = filepath.Abs(".")
}

// RealPath get an absolute path
func RealPath(path string, addSlash ...bool) (realPath string) {
	if !filepath.IsAbs(path) {
		path = FilePath + "/" + path
	}
	realPath, _ = filepath.Abs(path)
	realPath = pathAddSlash(filepath.ToSlash(realPath), addSlash...)

	return
}

// WorkDirPath program directory path
func WorkDirPath(addSlash ...bool) (path string) {
	ePath, err := os.Executable()
	if err != nil {
		ePath = FilePath
	}
	path = pathAddSlash(filepath.Dir(ePath), addSlash...)
	return
}

func pathAddSlash(path string, addSlash ...bool) string {
	if len(addSlash) > 0 && addSlash[0] && !strings.HasSuffix(path, "/") {
		path += "/"
	}
	return path
}
