// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package file

import (
	"bufio"
	"fmt"
	"github.com/ysicing/ext/ztime"
	"os"
)

//CheckFileExists check file exist
func CheckFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// Writefile 写文件
func Writefile(logpath, msg string) (err error) {
	file, err := os.OpenFile(logpath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(fmt.Sprintf("%v - %v\n", ztime.NowFormat(), msg))
	write.Flush()
	return nil
}
