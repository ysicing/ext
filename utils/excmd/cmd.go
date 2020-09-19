// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package excmd

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
)

//RunCmd is exec on os ,no return
func RunCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

//RunCmdRes is exec on os , return result
func RunCmdRes(name string, arg ...string) string {
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdin = os.Stdin
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	err := cmd.Run()
	if err != nil {
		log.Println("[os]os call error.", err)
		return ""
	}
	return b.String()
}

func downloadCmd(url string) string {
	u, ishttp := isURL(url)
	var c = ""
	if ishttp {
		param := ""
		if u.Scheme == "https" {
			param = "--no-check-certificate"
		}
		c = fmt.Sprintf(" wget -c %s %s", param, url)
	}
	return c
}

func isURL(u string) (url.URL, bool) {
	if uu, err := url.Parse(u); err == nil && uu != nil && uu.Host != "" {
		return *uu, true
	}
	return url.URL{}, false
}

// DownloadFile 下载文件
func DownloadFile(url string, location string) {
	dwncmd := downloadCmd(url)
	RunCmd("/bin/sh", "-c", "mkdir -p /tmp/ysicing && cd /tmp/ysicing && "+dwncmd)
	RunCmd("/bin/sh", "-c", "cp -a /tmp/ysicing/* "+location)
}
