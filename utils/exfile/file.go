// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exfile

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//CheckFileExists check file exist
func CheckFileExists(filename string) (bool, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

//CheckFileExistsv2 check file exist
func CheckFileExistsv2(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//CheckAndCreateDir check and create dir
func CheckAndCreateDir(path string) error {
	if subPathExists, err := CheckFileExists(path); err != nil {
		return fmt.Errorf("Could not determine if subPath %s exists; will not attempt to change its permissions", path)
	} else if !subPathExists {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to mkdir:%s", path)
		}

		if err := os.Chmod(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

//CheckAndCreateDirByMode check and create dir
func CheckAndCreateDirByMode(path string, mode os.FileMode) error {
	if subPathExists, err := CheckFileExists(path); err != nil {
		return fmt.Errorf("Could not determine if subPath %s exists; will not attempt to change its permissions", path)
	} else if !subPathExists {
		if err := os.MkdirAll(path, mode); err != nil {
			return fmt.Errorf("failed to mkdir:%s", path)
		}

		if err := os.Chmod(path, mode); err != nil {
			return err
		}
	}
	return nil
}

//DirIsEmpty 验证目录是否为空
func DirIsEmpty(dir string) bool {
	infos, err := ioutil.ReadDir(dir)
	if len(infos) == 0 || err != nil {
		return true
	}
	return false
}

// WriteFile 写文件
func WriteFile(filePath string, content string) error {
	filedir := filepath.Dir(filePath)
	if err := CheckAndCreateDir(filedir); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm); err != nil {
		return err
	}
	return nil
}

//SearchFileBody 搜索文件中是否含有指定字符串
func SearchFileBody(filename, searchStr string) bool {
	body, _ := ioutil.ReadFile(filename)
	return strings.Contains(string(body), searchStr)
}

//IsHaveFile 指定目录是否含有文件
//.开头文件除外
func IsHaveFile(path string) bool {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			return true
		}
	}
	return false
}

//SearchFile 搜索指定目录是否有指定文件，指定搜索目录层数，-1为全目录搜索
func SearchFile(pathDir, name string, level int) bool {
	if level == 0 {
		return false
	}
	files, _ := ioutil.ReadDir(pathDir)
	var dirs []os.FileInfo
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file)
			continue
		}
		if file.Name() == name {
			return true
		}
	}
	if level == 1 {
		return false
	}
	for _, dir := range dirs {
		ok := SearchFile(path.Join(pathDir, dir.Name()), name, level-1)
		if ok {
			return ok
		}
	}
	return false
}

//CheckFileExistsWithSuffix 指定目录是否含有指定后缀的文件
func CheckFileExistsWithSuffix(pathDir, suffix string) bool {
	files, _ := ioutil.ReadDir(pathDir)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), suffix) {
			return true
		}
	}
	return false
}

// RemoveFiles 删除文件
func RemoveFiles(path string) bool {
	if err := os.RemoveAll(path); err != nil {
		return false
	}
	return true
}

// ReadFileOneLine 读取文件一行
func ReadFileOneLine(fileName string) string {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return ""
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	line, err := buf.ReadString('\n')
	if err != nil {
		return ""
	}
	return line
}
