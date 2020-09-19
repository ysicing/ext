// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package extls

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// TLSCheck check tls Certificate
func TLSCheck(pemData string) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, errors.New("Certificate decoding error")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.New("Certificate Parsing error")
	}
	return cert, nil
}
