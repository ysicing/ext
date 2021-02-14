// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package validation

import "fmt"

// ValidationName 校验团队名或者项目名
func ValidationName(name string) error {
	res := IsDNS1123Label(name)
	if len(res) != 0 {
		return fmt.Errorf(res[0])
	}
	return nil
}
