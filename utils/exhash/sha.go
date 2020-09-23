// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exhash

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// GenSha256 生成sha256
func GenSha256(code string) string {
	s := sha256.New()
	s.Write([]byte(code))
	return hex.EncodeToString(s.Sum(nil))
}

// GenSha1 生成sha1
func GenSha1(code string) string {
	s := sha1.New()
	s.Write([]byte(code))
	return hex.EncodeToString(s.Sum(nil))
}
