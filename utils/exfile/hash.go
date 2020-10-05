// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exfile

import (
	"bufio"
	"crypto"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
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

func Md5FromLocal(localPath string) (string, error) {
	cmd := fmt.Sprintf("md5sum %s | cut -d\" \" -f1", localPath)
	c := exec.Command("sh", "-c", cmd)
	out, err := c.CombinedOutput()
	if err != nil {
		return "", err
	}
	md5 := string(out)
	md5 = strings.ReplaceAll(md5, "\n", "")
	md5 = strings.ReplaceAll(md5, "\r", "")

	return md5, nil
}
