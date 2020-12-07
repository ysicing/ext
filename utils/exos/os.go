// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exos

import (
	"github.com/ysicing/ext/utils/exfile"
	"net"
	"os"
	ou "os/user"
	"runtime"
)

// IsMacOS is Mac OS
func IsMacOS() bool {
	if runtime.GOOS == "darwin" {
		return true
	}
	return false
}

// IsLinux is linux
func IsLinux() bool {
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}

// IsUnix macos or linux
func IsUnix() bool {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		return true
	}
	return false
}

// IsContainer 是否是容器
func IsContainer() bool {
	return exfile.CheckFileExistsv2("/.dockerenv")
}

// GetUserName 获取当前系统登录用户
func GetUserName() string {
	user, err := ou.Current()
	if err != nil {
		return ""
	}
	return user.Username
}

// GetUser 获取当前系统登录用户
func GetUser() *ou.User {
	user, err := ou.Current()
	if err != nil {
		return nil
	}
	return user
}

func GetHostname() []string {
	host, err := os.Hostname()
	if err != nil {
		return nil
	}
	return []string{host}
}

// GetFreePort 获取空闲端口
func GetFreePort() int {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0
	}

	port := listener.Addr().(*net.TCPAddr).Port
	err = listener.Close()
	if err != nil {
		return 0
	}

	return port
}