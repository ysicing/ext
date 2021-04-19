// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exhash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

// MD5 md5
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// CryptoPass crypto password use salt
func CryptoPass(salt, raw string) string {
	return MD5(salt + "<-*Uk30^96eY*->" + raw)
}

// GenUUIDByStr genuuid str
func GenUUIDByStr(str string) string {
	return MD5(str + fmt.Sprint(time.Now().UnixNano()))
}
