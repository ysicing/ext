// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>
package validation

import "regexp"

const sha string = "[0-9a-z]{0,31}"

var shaRegexp = regexp.MustCompile("^" + sha + "$")
