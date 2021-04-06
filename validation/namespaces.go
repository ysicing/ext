// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package validation

import (
	"fmt"
	"regexp"
)

const (
	maxLength = 76
	alpha     = `[A-Za-z]`
	alphanum  = `[A-Za-z0-9]+`
	label     = alpha + alphanum + `(:?[-]+` + alpha + alphanum + `)*`
)

var (
	// namespaceRe validates that a namespace matches valid identifiers.
	//
	// Rules for domains, defined in RFC 1035, section 2.3.1, are used for
	// namespaces.
	namespaceRe = regexp.MustCompile(reAnchor(label + reGroup("[.]"+reGroup(label)) + "*"))
)

// NamespaceValidate returns nil if the string s is a valid namespace.
func NamespaceValidate(s string) error {
	if len(s) > maxLength {
		return fmt.Errorf("namespace %q greater than maximum length (%d characters)", s, maxLength)
	}

	if !namespaceRe.MatchString(s) {
		return fmt.Errorf("namespace %q must match %v", s, namespaceRe)
	}
	return nil
}

func reGroup(s string) string {
	return `(?:` + s + `)`
}

func reAnchor(s string) string {
	return `^` + s + `$`
}
