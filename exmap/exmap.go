// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package exmap

import (
	"encoding/json"
	"reflect"
	"strings"
)

// Struct2Map ...
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// Struct2Json2Map ...
func Struct2Json2Map(obj interface{}) (result map[string]interface{}, err error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonBytes, &result)
	return
}

// Map2String ...
func Map2String(data []string) (result string) {
	if len(data) <= 0 {
		return
	}
	for _, v := range data {
		if strings.Contains(v, "\"") {
			result += v
		} else {
			result += "\"" + v + "\""
		}
		result += ","
	}
	result = strings.Trim(result, ",")
	return
}
