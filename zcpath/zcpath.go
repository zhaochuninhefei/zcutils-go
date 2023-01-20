package zcpath

import (
	"fmt"
	"os"
)

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
