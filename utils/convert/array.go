// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package convert

// StringArrayContains 字符串数组是否包含某字符串
func StringArrayContains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

// Int64ArrayContains 数组是否包含某字符串
func Int64ArrayContains(addrs []int64, i int64) bool {
	for _, s := range addrs {
		if s == i {
			return true
		}
	}
	return false
}

// RemoveDuplicateElementInt64 去重
func RemoveDuplicateElementInt64(addrs []int64) []int64 {
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

// RemoveDuplicateElement 去重
func RemoveDuplicateElement(addrs []string) []string {
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
