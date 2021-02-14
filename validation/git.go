// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package validation

import "regexp"

const git string = "^(http(s)?|git|ssh).*(\\.git)(/)?$"

//VerifyGitFormat git addr verify
func VerifyGitFormat(gitaddr string) bool {
	reg := regexp.MustCompile(git)
	return reg.MatchString(gitaddr)
}
