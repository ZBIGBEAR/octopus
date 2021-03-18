package util

import (
	"os"
)

// 判断是不是文件夹
func IsDir(path string) (bool, error) {
	s, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return s.IsDir(), nil
}

// 判断是不是文件
func IsFile(path string) (bool, error) {
	b, err := IsDir(path)
	if err != nil {
		return false, err
	}
	return !b, nil
}


