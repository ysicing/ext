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

// Validate returns nil if the string s is a valid namespace.
//
// To allow such namespace identifiers to be used across various contexts
// safely, the character set has been restricted to that defined for domains in
// RFC 1035, section 2.3.1. This will make namespace identifiers safe for use
// across networks, filesystems and other media.
//
// The identifier specification departs from RFC 1035 in that it allows
// "labels" to start with number and only enforces a total length restriction
// of 76 characters.
//
// While the character set may be expanded in the future, namespace identifiers
// are guaranteed to be safely used as filesystem path components.
//
// For the most part, this doesn't need to be called directly when using the
// context-oriented functions.
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
