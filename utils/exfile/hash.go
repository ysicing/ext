// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exfile

import (
	"bufio"
	"crypto"
	"encoding/hex"
	"io"
	"os"
)

// Md5file 计算文件md5
func Md5file(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	hash := crypto.MD5.New()
	_, err = io.Copy(hash, r)
	if err != nil {
		return "", err
	}

	out := hex.EncodeToString(hash.Sum(nil))
	return out, nil
}
