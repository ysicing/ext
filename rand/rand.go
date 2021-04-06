// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package rand

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var digits = []rune("0123456789")

const size = 62

// Letters 随机字母
func Letters(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(size)]
	}

	return fmt.Sprintf("%s", string(b))
}

// Digits 随机数
func Digits(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = digits[rand.Intn(10)]
	}

	return fmt.Sprintf("%s", string(b))
}

// Rand 随机数
func Rand() int {
	rand.Seed(time.Now().Unix())
	return rand.Int()
}

// NumRand 随机数
func NumRand(num int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(num)
}

// Rand64 随机数
func Rand64() int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63()
}

// NumRand64 随机数
func NumRand64(num int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(num)
}
