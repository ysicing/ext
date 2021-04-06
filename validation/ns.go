// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package validation

import "fmt"

// CheckName 校验是否符合DNS规范
func CheckName(name string) error {
	res := IsDNS1123Label(name)
	if len(res) != 0 {
		return fmt.Errorf(res[0])
	}
	return nil
}
