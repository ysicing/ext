// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package convert

import "testing"

func TestStringArrayContains(t *testing.T) {
	code := "hahaha"
	code1 := []string{"hahahaha", "hahaha"}
	if !StringArrayContains(code1, code) {
		t.Error("error")
	}
	code2 := []string{"xxxx", "1233"}
	if StringArrayContains(code2, code) {
		t.Error("error")
	}
}
