// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exid

import "testing"

var tests = []string{"f47ac10b-58cc-0372-8567-0e02b2c3d479", "f47ac10b58cc037285670e02b2c3d47"}

func TestGenUUID(t *testing.T) {
	m := make(map[string]bool)
	for i := 1; i < 10; i++ {
		uuid := GenUUID()
		if !CheckUUID(uuid) {
			t.Error("Random UUID returned which does not decode")
		}
		if m[uuid] {
			t.Errorf("New returned duplicated UUID %s", uuid)
		}
		m[uuid] = true
	}
}
