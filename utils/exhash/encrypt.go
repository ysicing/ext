// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exhash

//import (
//	"encoding/base64"
//	"github.com/forgoer/openssl"
//)
//
//func AesCBCEncrypt(src, key, iv []byte) (string, error) {
//	var dst, err = openssl.AesCBCEncrypt(src, key, iv, openssl.PKCS7_PADDING)
//	if err != nil {
//		return "", err
//	}
//	return base64.StdEncoding.EncodeToString(dst), nil
//}
//
//func AesCBCDecrypt(dst, key, iv []byte) (string, error) {
//	result, err := openssl.AesCBCDecrypt(dst, key, iv, openssl.PKCS7_PADDING)
//	if err != nil {
//		return "", err
//	}
//	return string(result), nil
//}
