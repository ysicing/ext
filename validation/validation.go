// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package validation

import (
	"net"
	"regexp"
)

const (
	wsName string = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"
)

// IsValidIP tests that the argument is a valid IP address.
func IsValidIP(value string) []string {
	if net.ParseIP(value) == nil {
		return []string{"must be a valid IP address, (e.g. 10.9.8.7)"}
	}
	return nil
}

var wsNameRegexp = regexp.MustCompile("^" + wsName + "$")

// IsValidWsName 是否是合法项目名
func IsValidWsName(value string) []string {
	var errs []string
	if !wsNameRegexp.MatchString(value) {
		errs = append(errs, "不符合要求,示例 a-b, aaa, a0b")
	}
	return errs
}
