// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package exstr

import "strings"

// Blacklist
func Blacklist(s string) bool {
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

	if strings.Contains(s, "%") {
		return true
	}

	if strings.Contains(s, "=") {
		return true
	}

	if strings.Contains(s, "--") {
		return true
	}

	return false
}
