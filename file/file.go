// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package file

import (
	"bufio"
	"crypto"
	"encoding/hex"
	"fmt"
	"github.com/ysicing/ext/ztime"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var (
	FilePath = "."
)

func init() {
	FilePath, _ = filepath.Abs(".")
}

// RealPath get an absolute path
func RealPath(path string, addSlash ...bool) (realPath string) {
	if !filepath.IsAbs(path) {
		path = FilePath + "/" + path
	}
	realPath, _ = filepath.Abs(path)
	realPath = pathAddSlash(filepath.ToSlash(realPath), addSlash...)

	return
}

// WorkDirPath program directory path
func WorkDirPath(addSlash ...bool) (path string) {
	ePath, err := os.Executable()
	if err != nil {
		ePath = FilePath
	}
	path = pathAddSlash(filepath.Dir(ePath), addSlash...)
	return
}

func pathAddSlash(path string, addSlash ...bool) string {
	if len(addSlash) > 0 && addSlash[0] && !strings.HasSuffix(path, "/") {
		path += "/"
	}
	return path
}

// Rmdir rmdir,support to keep the current directory
func Rmdir(path string, notIncludeSelf ...bool) (ok bool) {
	realPath := RealPath(path)
	err := os.RemoveAll(realPath)
	ok = err == nil
	if ok && len(notIncludeSelf) > 0 && notIncludeSelf[0] {
		_ = os.Mkdir(path, os.ModePerm)
	}
	return
}

func FileSize(path string) float64 {
	fi, err := os.Stat(path)
	if err == nil {
		bs := float64(fi.Size())
		return bs
	}
	return 0
}

func FileSize2Str(path string) string {
	fi, err := os.Stat(path)
	if err == nil {
		bs := float64(fi.Size())
		kbs := bs / 1024.0
		mbs := kbs / 1024.0
		if mbs < 1024.0 {
			return fmt.Sprintf("%v M", mbs)
		}
		gbs := mbs / 1024.0
		if gbs < 1024.0 {
			return fmt.Sprintf("%v G", gbs)
		}
		tbs := gbs / 1024.0
		return fmt.Sprintf("%v T", tbs)
	}
	return ""
}

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

//DirIsEmpty 验证目录是否为空
func DirIsEmpty(dir string) bool {
	infos, err := ioutil.ReadDir(dir)
	if len(infos) == 0 || err != nil {
		return true
	}
	return false
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

// Md5file 计算文件md5
func Md5file(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	hash := crypto.MD5.New()
	_, err = io.Copy(hash, r)
	if err != nil {
		return "", err
	}

	out := hex.EncodeToString(hash.Sum(nil))
	return out, nil
}

func Md5FromLocal(localPath string) (string, error) {
	cmd := fmt.Sprintf("md5sum %s | cut -d\" \" -f1", localPath)
	c := exec.Command("sh", "-c", cmd)
	out, err := c.CombinedOutput()
	if err != nil {
		return "", err
	}
	md5 := string(out)
	md5 = strings.ReplaceAll(md5, "\n", "")
	md5 = strings.ReplaceAll(md5, "\r", "")

	return md5, nil
}
