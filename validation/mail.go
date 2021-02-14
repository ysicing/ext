// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package validation

import (
	"regexp"
	"strings"
)

const email string = "[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\\.){1,4}[a-z]{2,4}"

var emailRegexp = regexp.MustCompile("^" + email + "$")

//VerifyEmailFormat email verify
func VerifyEmailFormat(email string) bool {
	return emailRegexp.MatchString(email)
}

func MailDropHasSuffix(mail string) string {
	if strings.Contains(mail, "@") {
		return strings.Split(mail, "@")[0]
	}
	return mail
}
