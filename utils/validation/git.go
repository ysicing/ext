// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package validation

import (
	"regexp"
	"strings"
)

const git string = "^(http(s)?|git|ssh).*(\\.git)(/)?$"

//VerifyGitFormat git addr verify
func VerifyGitFormat(gitaddr string) bool {
	reg := regexp.MustCompile(git)
	return reg.MatchString(gitaddr)
}

// ParseGitAddr 解析git repo
func ParseGitAddr(gitAddr string) (string, string) {
	gitAddr = strings.Replace(gitAddr, ".git", "", -1)
	addr := strings.Split(gitAddr, ":")
	names := strings.Split(addr[len(addr)-1], "/")
	owner := names[len(names)-2]
	repo := names[len(names)-1]
	return owner, repo
}

// ParseGitAddrv2 获取项目仓库组织和项目仓库名
func ParseGitAddrv2(gitaddr string) (string, string) {
	if strings.HasPrefix(gitaddr, "git") {
		s0 := strings.Split(gitaddr, ":")[1]
		s1 := strings.Split(s0, ".")[0]
		s2 := strings.Split(s1, "/")
		return s2[0], s2[1]
	}
	s0 := strings.Split(gitaddr, "/")
	s1 := strings.Split(s0[len(s0)-1], ".")[0]
	s2 := s0[len(s0)-2]
	return s2, s1
}
