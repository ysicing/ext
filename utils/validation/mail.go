// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package validation

import "regexp"

const mail = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

//VerifyEmailFormat email verify
func VerifyEmailFormat(code string) bool {
	reg := regexp.MustCompile(mail)
	return reg.MatchString(code)
}
