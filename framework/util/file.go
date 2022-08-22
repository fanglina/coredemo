package util

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool  {
	_, err := os.Stat(path) // os.Stat 获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//路径是否是隐藏路径
func IsHiddenDirectory(path string) bool  {
	return len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".")
}

//输出所有子目录，目录名
func SubDir(folder string) ([]string, error)  {
	subs, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	ret := []string{}
	for _, sub := range subs {
		if sub.IsDir() {
			ret = append(ret, sub.Name())
		}
	}
	return ret, nil
}

// DownloadFile 下载文件
func DownloadFile(filepath string, url string) error  {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}