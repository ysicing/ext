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

func Error(data interface{}) gin.H {
	return gin.H{
		"data":      data,
		"timestamp": extime.NowUnix(),
		"code":      "10400",
	}
}

func Errorv2(code int, message string, data interface{}) gin.H {
	return gin.H{
		"data":      data,
		"timestamp": extime.NowUnix(),
		"code":      code,
		"message":   message,
	}
}
