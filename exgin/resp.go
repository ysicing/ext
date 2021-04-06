// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package exgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/ztime"
)

// done done
func respdone(data interface{}) gin.H {
	return gin.H{
		"data":      data,
		"message":   "请求成功",
		"timestamp": ztime.NowUnix(),
		"code":      200,
	}
}

// error error
func resperror(code int64, data interface{}) gin.H {
	return gin.H{
		"data":      nil,
		"message":   data,
		"timestamp": ztime.NowUnix(),
		"code":      code,
	}
}

func renderMessage(c *gin.Context, v interface{}) {
	if v == nil {
		c.JSON(200, respdone(nil))
		return
	}

	switch t := v.(type) {
	case string:
		c.JSON(200, resperror(10400, t))
	case error:
		c.JSON(200, resperror(10400, t.Error()))
	}
}

// GinsData gins resp data
func GinsData(c *gin.Context, data interface{}, err error) {
	if err == nil {
		c.JSON(200, respdone(data))
		return
	}

	renderMessage(c, err.Error())
}

// GinsAbort gins abort
func GinsAbort(c *gin.Context, msg string, args ...interface{}) {
	c.AbortWithStatusJSON(200, resperror(10400, fmt.Sprintf(msg, args...)))
}
