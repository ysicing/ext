// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exos

import (
	"github.com/mitchellh/go-homedir"
	"path"
	"path/filepath"
)

// DefaultPrivateKey 默认ssh key
func DefaultPrivateKey() string {
	home, err := homedir.Dir()
	if err != nil {
		return ""
	}
	key := filepath.ToSlash(path.Join(home, ".ssh/id_rsa"))
	return key
}
