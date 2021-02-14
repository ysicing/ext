// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package zos

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// TLSCheck check tls Certificate
func TLSCheck(pemData string) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, fmt.Errorf("Certificate decoding error")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Certificate Parsing error")
	}
	return cert, nil
}
