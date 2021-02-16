// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package gerr

import (
	"fmt"
)

type GinsError struct {
	Message string
}

func (ge *GinsError) Error() string {
	return ge.Message
}

func (ge *GinsError) String() string {
	return ge.Message
}

func Bomb(format string, args ...interface{}) {
	panic(GinsError{Message: fmt.Sprintf(format, args...)})
}

func Boka(value string, v interface{}) {
	if v == nil {
		return
	}
	Bomb(value)
}

func Dangerous(v interface{}) {
	if v == nil {
		return
	}

	switch t := v.(type) {
	case string:
		if t != "" {
			panic(GinsError{Message: t})
		}
	case error:
		panic(GinsError{Message: t.Error()})
	}
}
