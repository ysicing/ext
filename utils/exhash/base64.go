// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exhash

import (
	b32 "encoding/base32"
	b64 "encoding/base64"
)

// B64EnCode base64加密
func B64EnCode(code string) string {
	return b64.StdEncoding.EncodeToString([]byte(code))
}

// B64Decode base64解密
func B64Decode(code string) (string, error) {
	ds, err := b64.StdEncoding.DecodeString(code)
	return string(ds), err
}

// B32EnCode base32加密
func B32EnCode(code string) string {
	return b32.StdEncoding.EncodeToString([]byte(code))
}

// B32Decode base32解密
func B32Decode(code string) (string, error) {
	ds, err := b32.StdEncoding.DecodeString(code)
	return string(ds), err
}
