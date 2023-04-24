package zcpath

import (
	"fmt"
	"os"
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

func TestClearDir(t *testing.T) {
	// 创建测试目录
	_, _ = CreateDir("testdata/testdir")
	// 在测试目录下创建文件
	err := CreateFile("testdata/testdir/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 在测试目录下创建子目录
	_, _ = CreateDir("testdata/testdir/testdir1")
	// 在子目录下创建文件
	err = CreateFile("testdata/testdir/testdir1/testfile2.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 在testdir1下创建子目录
	_, _ = CreateDir("testdata/testdir/testdir1/testdir2")
	// 在testdir2下创建文件
	err = CreateFile("testdata/testdir/testdir1/testdir2/testfile3.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 打印测试目录
	fmt.Println("测试目录结构 打印所有层级:")
	err = PrintDirTree("testdata/testdir", -1, false, true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("测试目录结构 只打印0层:")
	err = PrintDirTree("testdata/testdir", 0, false, true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("测试目录结构 只打印1层:")
	err = PrintDirTree("testdata/testdir", 1, false, true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("测试目录结构 只打印2层:")
	err = PrintDirTree("testdata/testdir", 2, false, true)
	if err != nil {
		t.Fatal(err)
	}
	// 清空测试目录
	err = ClearDir("testdata/testdir")
	if err != nil {
		t.Fatal(err)
	}
	// 打印测试目录
	fmt.Println("测试目录结构 清空后:")
	err = PrintDirTree("testdata/testdir", -1, false, true)
	if err != nil {
		t.Fatal(err)
	}
	// 判断测试目录是否为空
	isEmpty, err := IsDirEmpty("testdata/testdir")
	if err != nil {
		t.Fatal(err)
	}
	if !isEmpty {
		t.Fatal("测试目录不为空")
	} else {
		fmt.Println("测试目录清空成功")
	}
}

func TestFileExists(t *testing.T) {
	// 创建测试文件
	err := CreateFile("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 判断文件是否存在
	exists, err := FileExists("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("文件不存在")
	} else {
		fmt.Println("文件存在")
	}

	// 判断不存在的文件是否存在
	exists, err = FileExists("testdata/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("文件存在")
	} else {
		fmt.Println("文件不存在")
	}

	// path存在但是不是文件
	exists, err = FileExists("testdata")
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("文件存在")
	} else {
		fmt.Println("文件不存在")
	}

	// 删除测试文件
	err = os.Remove("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFilterFileByCondition(t *testing.T) {
	condition := FileFilterCondition{
		FileNamePrefix: "zc",
		FileNameSuffix: ".go",
		FileNameRegex:  "zc.*",
		ContainsHidden: false,
		ContainsDir:    false,
	}
	files := FilterFileByCondition(".", condition)
	for _, file := range files {
		fmt.Println(file)
	}
}

func TestFileCopy(t *testing.T) {
	// 创建测试文件
	err := CreateFile("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 复制文件
	err = FileCopy("testdata/testfile.txt", "testdata/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 判断文件是否存在
	exists, err := FileExists("testdata/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("文件不存在")
	} else {
		fmt.Println("文件存在")
	}
	// 删除测试文件
	err = os.Remove("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 删除测试文件
	err = os.Remove("testdata/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}

}
