// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package exhash

import "github.com/google/uuid"

// GenUUID 生成新的uuid
func GenUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

// CheckUUID 检查uuid是否合法
func CheckUUID(uid string) bool {
	_, err := uuid.Parse(uid)
	if err != nil {
		return false
	}
	return true
}
