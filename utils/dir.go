package utils

import "os"

// DirExists 判断目录是否存在
func DirExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	} else if fileInfo.IsDir() {
		return true, nil
	}
	return false, err
}
