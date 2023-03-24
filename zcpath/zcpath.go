package zcpath

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

// FilterFileByCondition 根据条件过滤文件
func FilterFileByCondition(dir string, condition FileFilterCondition) []string {
	// 使用filepath.Walk遍历dir中的文件，并根据condition进行过滤
	var files []string
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && !condition.ContainsDir {
			return nil
		}
		if !condition.ContainsHidden && info.Name()[0] == '.' {
			return nil
		}
		if condition.FileNamePrefix != "" && !strings.HasPrefix(info.Name(), condition.FileNamePrefix) {
			return nil
		}
		if condition.FileNameSuffix != "" && !strings.HasSuffix(info.Name(), condition.FileNameSuffix) {
			return nil
		}
		if condition.FileNameRegex != "" {
			matched, err := regexp.MatchString(condition.FileNameRegex, info.Name())
			if err != nil {
				return err
			}
			if !matched {
				return nil
			}
		}
		files = append(files, path)
		return nil
	})
	return files
}
