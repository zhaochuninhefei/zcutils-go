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

func TestFileCopyToDir(t *testing.T) {
	// 创建测试文件
	err := CreateFile("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 复制文件
	err = FileCopyToDir("testdata/testfile.txt", "testdata/subdir1/subdir11")
	if err != nil {
		t.Fatal(err)
	}
	// 判断文件是否存在
	exists, err := FileExists("testdata/subdir1/subdir11/testfile.txt")
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
	err = os.Remove("testdata/subdir1/subdir11/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestChmodDir(t *testing.T) {
	// 创建测试目录 chmod_dir/subdir1
	ok, err := CreateDir("testdata/chmod_dir/subdir1")
	if !ok && err != nil {
		t.Fatal(err)
	}

	// 创建测试文件 chmod_dir/testfile1.txt
	err = CreateFile("testdata/chmod_dir/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 创建测试文件 chmod_dir/testfile2.txt
	err = CreateFile("testdata/chmod_dir/testfile2.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 创建测试文件 chmod_dir/subdir1/testfile3.txt
	err = CreateFile("testdata/chmod_dir/subdir1/testfile3.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 创建测试目录 chmod_dir/subdir2
	ok, err = CreateDir("testdata/chmod_dir/subdir2")
	if !ok && err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("创建目录失败")
	}
	// PrintDirTreeWithMode
	err = PrintDirTreeWithMode("testdata/chmod_dir", -1, false, true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("测试数据准备完毕")

	// 修改目录权限
	err = ChmodDir("testdata/chmod_dir", os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	// PrintDirTreeWithMode
	err = PrintDirTreeWithMode("testdata/chmod_dir", -1, false, true)
	if err != nil {
		t.Fatal(err)
	}
	// 查看目录权限
	fileInfo, err := os.Stat("testdata/chmod_dir")
	if err != nil {
		t.Fatal(err)
	}
	// 判断目录权限是否是0777
	if (fileInfo.Mode() & os.ModePerm) != 0777 {
		t.Fatal("目录权限不是0777")
	} else {
		fmt.Println("目录权限是0777")
	}

	// 修改目录权限
	err = ChmodDir("testdata/chmod_dir", 0755)
	if err != nil {
		t.Fatal(err)
	}
	// PrintDirTreeWithMode
	err = PrintDirTreeWithMode("testdata/chmod_dir", -1, false, true)
	if err != nil {
		t.Fatal(err)
	}
	// 查看目录权限
	fileInfo, err = os.Stat("testdata/chmod_dir")
	if err != nil {
		t.Fatal(err)
	}
	// 判断目录权限是否是0755
	if (fileInfo.Mode() & os.ModePerm) != 0755 {
		t.Fatal("目录权限不是0755")
	} else {
		fmt.Println("目录权限是0755")
	}

	// 删除测试目录 chmod_dir
	err = os.RemoveAll("testdata/chmod_dir")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveFile(t *testing.T) {
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
	// 删除测试文件
	err = RemoveFile("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 判断文件是否存在
	exists, err = FileExists("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("文件存在")
	} else {
		fmt.Println("文件不存在")
	}
}
