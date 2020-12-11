// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package validation

import "strings"

//Dangerous 是否危险
func Dangerous(s string) bool {
	if strings.Contains(s, "<") {
		return true
	}

	if strings.Contains(s, ">") {
		return true
	}

	if strings.Contains(s, "&") {
		return true
	}

	if strings.Contains(s, "'") {
		return true
	}

	if strings.Contains(s, "\"") {
		return true
	}

	if strings.Contains(s, "file://") {
		return true
	}

	if strings.Contains(s, "../") {
		return true
	}

	return false
}