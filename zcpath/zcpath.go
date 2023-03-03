package zcpath

import (
	"fmt"
	"os"
)

// CreateDir 创建目录
//  @param path 目录路径
//  @return bool 创建成功与否
//  @return error
func CreateDir(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, fmt.Errorf("该目录已经存在: %s", path)
	} else {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

type FileFilterCondition struct {
	FileNamePrefix string // 文件名前缀
	FileNameSuffix string // 文件名后缀
	FileNameRegex  string // 文件名正则表达式
	ContainsHidden bool   // 是否查找隐藏文件
	ContainsDir    bool   // 是否查找目录
}

func FilterFileByCondition(dir string) []string {
	return nil
}
