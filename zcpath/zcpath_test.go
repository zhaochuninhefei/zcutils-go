package zcpath

import (
	"fmt"
	"testing"
)

func TestCreateDir(t *testing.T) {
	suc, err := CreateDir("testdata/testdir")
	if !suc {
		t.Fatal(err)
	} else {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("首次创建目录成功")
		}
	}
}
