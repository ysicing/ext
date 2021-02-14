// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package exstr

// DuplicateInt64ElementInt64 数组去重
func DuplicateInt64ElementInt64(addrs []int64) []int64 {
	result := make([]int64, 0, len(addrs))
	temp := map[int64]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// DuplicateStrElement 字符串去重
func DuplicateStrElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
