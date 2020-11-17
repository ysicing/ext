// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exfile

import (
	"fmt"
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

// Rmdir rmdir,support to keep the current directory
func Rmdir(path string, notIncludeSelf ...bool) (ok bool) {
	realPath := RealPath(path)
	err := os.RemoveAll(realPath)
	ok = err == nil
	if ok && len(notIncludeSelf) > 0 && notIncludeSelf[0] {
		_ = os.Mkdir(path, os.ModePerm)
	}
	return
}

func FileSize(path string) float64 {
	fi, err := os.Stat(path)
	if err == nil {
		bs := float64(fi.Size())
		return bs
	}
	return 0
}

func FileSize2Str(path string) string {
	fi, err := os.Stat(path)
	if err == nil {
		bs := float64(fi.Size())
		kbs := bs / 1024.0
		mbs := kbs / 1024.0
		if mbs < 1024.0 {
			return fmt.Sprintf("%v M", mbs)
		}
		gbs := mbs / 1024.0
		if gbs < 1024.0 {
			return fmt.Sprintf("%v G", gbs)
		}
		tbs := gbs / 1024.0
		return fmt.Sprintf("%v T", tbs)
	}
	return ""
}
