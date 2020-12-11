// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exrand

import (
	"fmt"
	"math/rand"
	"time"
)

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

// Rand 随机数
func Rand64() int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63()
}

// NumRand 随机数
func NumRand64(num int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(num)
}

// StringRand 生成随机字符串
func StringRand(len int) string {
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var digits = []rune("0123456789")

const size = 62

func RandLetters(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(size)]
	}

	return fmt.Sprintf("%s", string(b))
}

func RandDigits(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = digits[rand.Intn(10)]
	}

	return fmt.Sprintf("%s", string(b))
}