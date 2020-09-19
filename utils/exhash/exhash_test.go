// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exhash

import "testing"

func TestGenMd5(t *testing.T) {
	code := "md5"
	md5code := GenMd5(code)
	if md5code != "1bc29b36f623ba82aaf6724fd3b16718" {
		t.Errorf("md5 gen error: %v", md5code)
	}
}

func TestGenSha256(t *testing.T) {
	code := "sha250"
	encode := GenSha256(code)
	if encode != "d057148fe3c86aea5f179b08029051d7979ce384bf22661086413eae06124738" {
		t.Errorf("sha256 gen error: %v", encode)
	}
}

func TestB64EnCode(t *testing.T) {
	code := "hahaha"
	encode := B64EnCode(code)
	decode, err := B64Decode(encode)
	if err != nil {
		t.Errorf("decode err: %v", err)
	}
	if code != decode {
		t.Errorf("encode err: %v --> %v", code, encode)
	}
}

func TestB32EnCode(t *testing.T) {
	code := "hahaha"
	encode := B32EnCode(code)
	decode, err := B32Decode(encode)
	if err != nil {
		t.Errorf("decode err: %v", err)
	}
	if code != decode {
		t.Errorf("encode err: %v --> %v", code, encode)
	}
}
