// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/zos"
	"net/http"
)

// ExGin gin engine
func ExGin(debug bool) *gin.Engine {
	if debug || zos.IsMacOS() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()
	return gin.New()
}

// ReadCookie 读cookie
func ReadCookie(c *gin.Context, key string) (string, error) {
	ck, err := c.Cookie(key)
	if err != nil {
		return "", err
	}
	return ck, nil
}

// WriteDefaultCookie write default cookie
func WriteDefaultCookie(c *gin.Context, key, value string, args ...string) {
	if len(args) > 0 {
		c.SetCookie(key, value, 3600*24, "/", args[0], false, true)

	} else {
		c.SetCookie(key, value, 3600*24, "/", "", false, true)
	}
}

// FoundRedirect redirect with the StatusFound
func FoundRedirect(c *gin.Context, location string) {
	c.Redirect(http.StatusFound, location)
	c.Abort()
}

// MovedRedirect redirect with the StatusMovedPermanently
func MovedRedirect(c *gin.Context, location string) {
	c.Redirect(http.StatusMovedPermanently, location)
	c.Abort()
}

// TemporaryRedirect redirect with the StatusTemporaryRedirect
func TemporaryRedirect(c *gin.Context, location string) {
	c.Redirect(http.StatusTemporaryRedirect, location)
	c.Abort()
}
