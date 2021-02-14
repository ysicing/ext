// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package exstr

import (
	"strconv"
)

// Str2Int string to int
func Str2Int(s string) int {
	si, _ := strconv.Atoi(s)
	return si
}

// Str2Int32 string to int32
func Str2Int32(s string) int32 {
	si32, _ := strconv.ParseInt(s, 10, 32)
	return int32(si32)
}

// Str2Int64 string to int64
func Str2Int64(s string) int64 {
	si64, _ := strconv.ParseInt(s, 10, 64)
	return si64
}

// Str2Float64 string to float64
func Str2Float64(s string) float64 {
	sf, _ := strconv.ParseFloat(s, 64)
	return sf
}

// Str2Byte string to byte
func Str2Byte(s string) []byte {
	return []byte(s)
}

// Int642Str int64 to string
func Int642Str(i int64) string {
	return strconv.FormatInt(i, 10)
}
