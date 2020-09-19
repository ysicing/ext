// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exjson

import "encoding/json"

// Encode encode
func Encode(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Encodev1 encode
func Encodev1(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// Decode decode
func Decode(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}
