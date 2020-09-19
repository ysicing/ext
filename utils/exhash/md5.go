// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exhash

import (
	"crypto/md5"
	"encoding/hex"
)

// GenMd5 生成Md5
func GenMd5(code string) string {
	s := md5.New()
	s.Write([]byte(code))
	return hex.EncodeToString(s.Sum(nil))
}
