// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package validation

import (
	"errors"
	"regexp"
	"strconv"
)

const domainrule = "[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?"

// ValidatePort validates port numbers
func ValidatePort(port int32) error {
	if _, err := ParsePort(strconv.Itoa(int(port))); err != nil {
		return errors.New("port number is not valid")
	}
	return nil
}

// ParsePort parses a string representing a TCP port.
// If the string is not a valid representation of a TCP port, ParsePort returns an error.
func ParsePort(port string) (int, error) {
	portInt, err := strconv.Atoi(port)
	if err == nil && (1 <= portInt && portInt <= 65535) {
		return portInt, nil
	}

	return 0, errors.New("port must be a valid number between 1 and 65535, inclusive")
}

//VerifyDomainFormat domain verify
func VerifyDomainFormat(domain string) bool {
	reg := regexp.MustCompile(domainrule)
	return reg.MatchString(domain)
}
