// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exid

import (
	guuid "github.com/google/uuid"
)

// GenUUID 生成新的uuid
func GenUUID() string {
	u, _ := guuid.NewRandom()
	return u.String()
}

// CheckUUID 检查uuid是否合法
func CheckUUID(uid string) bool {
	_, err := guuid.Parse(uid)
	if err != nil {
		return false
	}
	return true
}
