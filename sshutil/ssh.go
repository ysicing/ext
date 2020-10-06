// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package sshutil

import (
	"bufio"
	"fmt"
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/exmisc"
	"io"
	"strings"
)

//Cmd is in host exec cmd
func (ss *SSH) Cmd(host string, cmd string) []byte {
	logger.Slog.Infof("[ssh][%s] %s", host, cmd)
	session, err := ss.Connect(host)
	defer func() {
		if r := recover(); r != nil {
			logger.Slog.Errorf("[ssh][%s]Error create ssh session failed,%s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	defer session.Close()
	b, err := session.CombinedOutput(cmd)
	logger.Slog.Debugf("[ssh][%s]command result is: %s", host, string(b))
	defer func() {
		if r := recover(); r != nil {
			logger.Slog.Errorf("[ssh][%s]Error exec command failed: %s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	return b
}

func readPipe(host string, pipe io.Reader, isErr bool, debugmsg ...bool) {
	r := bufio.NewReader(pipe)
	for {
		line, _, err := r.ReadLine()
		if line == nil {
			return
		} else if err != nil {
			logger.Slog.Infof("[%s] %s", exmisc.SGreen(host), line)
			logger.Slog.Errorf("[ssh] [%s] %s", exmisc.SRed(host), err)
			return
		} else {
			if isErr {
				if len(debugmsg) > 0 && debugmsg[0] {
					logger.Slog.Errorf("[%s] %s", exmisc.SRed(host), line)
				} else {
					msg, _ := fmt.Printf("%s", line)
					fmt.Println(msg)
				}
			} else {
				if len(debugmsg) > 0 && debugmsg[0] {
					logger.Slog.Infof("[%s] %s", exmisc.SGreen(host), line)
				} else {
					msg, _ := fmt.Printf("%s", line)
					fmt.Println(msg)
				}
			}
		}
	}
}

func (ss *SSH) CmdAsync(host string, cmd string, debugmsg ...bool) error {
	fmt.Printf("[ssh][%s] ➜   %s\n", exmisc.SGreen(host), cmd)
	session, err := ss.Connect(host)
	if err != nil {
		logger.Slog.Errorf("[ssh][%s]Error create ssh session failed,%s", host, err)
		return err
	}
	defer session.Close()
	stdout, err := session.StdoutPipe()
	if err != nil {
		logger.Slog.Errorf("[ssh][%s]Unable to request StdoutPipe(): %s", host, err)
		return err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		logger.Slog.Errorf("[ssh][%s]Unable to request StderrPipe(): %s", host, err)
		return err
	}
	if err := session.Start(cmd); err != nil {
		logger.Slog.Errorf("[ssh][%s]Unable to execute command: %s", host, err)
		return err
	}
	doneout := make(chan bool, 1)
	doneerr := make(chan bool, 1)
	go func() {
		readPipe(host, stderr, true, debugmsg...)
		doneerr <- true
	}()
	go func() {
		readPipe(host, stdout, false, debugmsg...)
		doneout <- true
	}()
	<-doneerr
	<-doneout
	return session.Wait()
}

//CmdToString is in host exec cmd and replace to spilt str
func (ss *SSH) CmdToString(host, cmd, spilt string) string {
	data := ss.Cmd(host, cmd)
	if data != nil {
		str := string(data)
		str = strings.ReplaceAll(str, "\r\n", spilt)
		return str
	}
	return ""
}
