// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package validation

import "regexp"

const tag string = "^[a-z0-9]*(([a-z0-9-])*[a-z0-9]+)*$"

//VerifyTagFormat tag verify
func VerifyTagFormat(tagstr string) bool {
	reg := regexp.MustCompile(tag)
	// a-b-c
	return reg.MatchString(tagstr)
}
