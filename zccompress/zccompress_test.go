package zccompress

import (
	"fmt"
	"gitee.com/zhaochuninhefei/zcutils-go/zcpath"
	"testing"
)

func TestCompressDirToTargz(t *testing.T) {
	// 清空测试目录
	err := zcpath.ClearDir("testdata")
	if err != nil {
		t.Fatal(err)
	}

	// 创建一个目录
	ok, err := zcpath.CreateDir("testdata/testDir")
	if !ok && err != nil {
		t.Fatal(err)
	}
	// 创建一个文件
	err = zcpath.CreateFile("testdata/testDir/testFile.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 创建一个子目录 subdir1
	ok, err = zcpath.CreateDir("testdata/testDir/subdir1")
	if !ok && err != nil {
		t.Fatal(err)
	}
	// 创建一个文件 subdir1/testFile1.txt
	err = zcpath.CreateFile("testdata/testDir/subdir1/testFile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 创建一个子目录 subdir2
	ok, err = zcpath.CreateDir("testdata/testDir/subdir2")
	if !ok && err != nil {
		t.Fatal(err)
	}
	// 打印测试目录结构
	fmt.Println("测试目录结构如下:")
	err = zcpath.PrintDirTree("testdata/testDir", -1, false, true)
	if err != nil {
		t.Fatal(err)
	}

	// 测试CompressDirToTargz
	err = CompressDirToTargz("testdata/testDir", "testdata/XXX.tar.gz")
	if err != nil {
		t.Fatal(err)
	}
	// 检查压缩文件是否存在
	if isFile, err := zcpath.FileExists("testdata/XXX.tar.gz"); !isFile || err != nil {
		t.Fatal("压缩文件不存在")
	}

	// 测试重复CompressDirToTargz
	err = CompressDirToTargz("testdata/testDir", "testdata/XXX.tar.gz")
	if err != nil {
		t.Fatal(err)
	}
	// 检查压缩文件是否存在
	if isFile, err := zcpath.FileExists("testdata/XXX.tar.gz"); !isFile || err != nil {
		t.Fatal("压缩文件不存在")
	}

	// 测试UnCompressTargzToDir
	err = UnCompressTargzToDir("testdata/XXX.tar.gz", "testdata/YYY")
	if err != nil {
		t.Fatal(err)
	}
	// 打印解压缩后的目录结构
	fmt.Println("解压缩后的目录结构如下:")
	err = zcpath.PrintDirTree("testdata/YYY", -1, false, true)
	if err != nil {
		t.Fatal(err)
	}

	// 测试重复解压缩
	err = UnCompressTargzToDir("testdata/XXX.tar.gz", "testdata/YYY")
	if err != nil {
		t.Fatal(err)
	}
	// 打印解压缩后的目录结构
	fmt.Println("重复解压缩后的目录结构如下:")
	err = zcpath.PrintDirTree("testdata/YYY", -1, false, true)
	if err != nil {
		t.Fatal(err)
	}
}
