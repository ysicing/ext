// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exhash

import (
	"crypto/sha256"
	"encoding/hex"
)

// GenSha256 生成sha256
func GenSha256(code string) string {
	s := sha256.New()
	s.Write([]byte(code))
	return hex.EncodeToString(s.Sum(nil))
}
