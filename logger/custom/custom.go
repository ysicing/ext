// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package custom

import (
	"bufio"
	"fmt"
	"github.com/ysicing/ext/zos"
	"os"
)

const logprefix = "/var/exlog"

func getLogFilePath(svcname, logname string) (logpath string) {
	if zos.IsLinux() {
		logpath = fmt.Sprintf("%s/%s", logprefix, svcname)
	} else {
		logpath = fmt.Sprintf("%s/%s", "/tmp/custom/logs", svcname)
	}
	os.MkdirAll(logpath, os.ModePerm)
	return fmt.Sprintf("%s/%s", logpath, logname)
}

// ReadLogs read logs
func ReadLogs(svcname, logname string) (logs []map[string]interface{}, err error) {
	logpath := getLogFilePath(svcname, logname)
	file, err := os.Open(logpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		logs = append(logs, map[string]interface{}{
			"id":      i,
			"svcname": svcname,
			"logname": logname,
			"message": scanner.Text(),
		})
		i++
	}
	return logs, nil
}

// WriteLogs write logs
func WriteLogs(svcname, logname, msg string) (err error) {
	logpath := getLogFilePath(svcname, logname)
	file, err := os.OpenFile(logpath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(fmt.Sprintf("%v\n", msg))
	write.Flush()
	return nil
}
