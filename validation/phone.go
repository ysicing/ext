// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package validation

import "regexp"

const phone = "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

// VerifyPhoneFormat mobile verify
func VerifyPhoneFormat(mobileNum string) bool {
	reg := regexp.MustCompile(phone)
	return reg.MatchString(mobileNum)
}
