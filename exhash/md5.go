// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exhash

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 md5
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
