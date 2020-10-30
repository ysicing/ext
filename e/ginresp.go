// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package e

import (
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/utils/extime"
)

func Done(data interface{}) gin.H {
	return gin.H{
		"data":      data,
		"timestamp": extime.NowUnix(),
		"code":      "200",
	}
}

func Error(code string, data interface{}, msg ...string) gin.H {
	if len(msg) == 0 {
		return gin.H{
			"data":      data,
			"timestamp": extime.NowUnix(),
			"code":      code,
		}
	}
	return gin.H{
		"data":      data,
		"timestamp": extime.NowUnix(),
		"code":      code,
		"message":   msg[0],
	}
}
