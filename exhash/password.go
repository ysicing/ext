// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exhash

import (
	"fmt"
	"time"
)

// CryptoPass crypto password use salt
func CryptoPass(salt, raw string) string {
	return MD5(salt + "<-*Uk30^96eY*->" + raw)
}

func GenUUIDForUser(username string) string {
	return MD5(username + fmt.Sprint(time.Now().UnixNano()))
}
