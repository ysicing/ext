// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package zos

import (
	"github.com/ysicing/ext/file"
	"os"
	"os/user"
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
	return file.CheckFileExists("/.dockerenv")
}

// GetUserName 获取当前系统登录用户
func GetUserName() string {
	user, err := user.Current()
	if err != nil {
		return ""
	}
	return user.Username
}

// GetUser 获取当前系统登录用户
func GetUser() *user.User {
	user, err := user.Current()
	if err != nil {
		return nil
	}
	return user
}

func GetHostnames() []string {
	host, err := os.Hostname()
	if err != nil {
		return nil
	}
	return []string{host}
}

func GetHostname() string {
	hosts := GetHostnames()
	if len(hosts) == 0 {
		return "unknow"
	}
	return hosts[0]
}

func GetOS() string {
	return runtime.GOOS
}
