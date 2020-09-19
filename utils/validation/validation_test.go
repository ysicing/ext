// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package validation

import "testing"

func TestParseGitAddr(t *testing.T) {
	gitssh := "git@github.com:ysicing/ergo.git"
	githttp := "https://github.com/ysicing/ergo.git"
	gitother := "https://github.com/ysicing/ergo"
	ower, repo := ParseGitAddr(gitssh)
	t.Logf("git %s ower %s repo %s", gitssh, ower, repo)
	ower, repo = ParseGitAddr(githttp)
	t.Logf("git %s ower %s repo %s", githttp, ower, repo)
	ower, repo = ParseGitAddr(gitother)
	t.Logf("git %s ower %s repo %s", gitother, ower, repo)
}

func TestParseGitAddrv2(t *testing.T) {
	gitssh := "git@github.com:ysicing/ergo.git"
	githttp := "https://github.com/ysicing/ergo.git"
	gitother := "https://github.com/ysicing/ergo"
	ower, repo := ParseGitAddrv2(gitssh)
	t.Logf("git %s ower %s repo %s", gitssh, ower, repo)
	ower, repo = ParseGitAddr(githttp)
	t.Logf("git %s ower %s repo %s", githttp, ower, repo)
	ower, repo = ParseGitAddr(gitother)
	t.Logf("git %s ower %s repo %s", gitother, ower, repo)
}
