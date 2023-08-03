package zcpath

import (
	"bufio"
	"fmt"
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"log"
	"os"
	"path"
	"reflect"
	"strconv"
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
	success, err := CreateDir("testdata/subdir0")
	if !success {
		t.Fatal(err)
	}
	err = CreateFile("testdata/subdir0/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 复制文件
	err = FileCopyToDir("testdata/testfile.txt", "testdata/subdir1/subdir11")
	if err != nil {
		t.Fatal(err)
	}
	err = FileWithWildcardCopyToDir("testdata/*/*.txt", "testdata/subdir1/subdir12")
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
	exists, err = FileExists("testdata/subdir1/subdir12/testfile1.txt")
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
	err = os.Remove("testdata/subdir0/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = os.Remove("testdata/subdir1/subdir11/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = os.Remove("testdata/subdir1/subdir12/testfile1.txt")
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

	// 创建测试子目录 testdata/subdir0
	ok, err := CreateDir("testdata/subdir0")
	if !ok && err != nil {
		t.Fatal(err)
	}
	// 创建测试文件 testdata/subdir0/testfile1.txt
	err = CreateFile("testdata/subdir0/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 判断文件是否存在
	exists, err = FileExists("testdata/subdir0/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("文件不存在")
	} else {
		fmt.Println("文件存在")
	}
	// 删除文件
	err = RemoveFileWithWildcard("testdata/subdir0/*.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 判断文件是否存在
	exists, err = FileExists("testdata/subdir0/testfile1.txt")
	if exists {
		t.Fatal("文件存在")
	} else {
		fmt.Println("文件不存在")
	}
}

func TestFileCopyFromDirToDir(t *testing.T) {
	// 创建测试目录
	ok, err := CreateDir("testdata/srcDir")
	if !ok && err != nil {
		t.Fatal(err)
	}
	// 创建测试文件 testfile1
	err = CreateFile("testdata/srcDir/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 创建测试文件 testfile2
	err = CreateFile("testdata/srcDir/testfile2.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 创建测试子目录 subdir1
	ok, err = CreateDir("testdata/srcDir/subdir1")
	if !ok && err != nil {
		t.Fatal(err)
	}
	// 创建测试文件 subdir1/testfile3
	err = CreateFile("testdata/srcDir/subdir1/testfile3.txt")
	if err != nil {
		t.Fatal(err)
	}

	// 打印测试目录
	fmt.Println("源目录结构:")
	err = PrintDirTree("testdata/srcDir", -1, false, true)
	if err != nil {
		t.Fatal(err)
	}

	// 测试FileCopyFromDirToDir
	err = FileCopyFromDirToDir("testdata/srcDir", "testdata/dstDir")
	if err != nil {
		t.Fatal(err)
	}
	// 检查目标目录
	fmt.Println("目标目录结构:")
	err = PrintDirTree("testdata/dstDir", -1, false, true)
	if err != nil {
		t.Fatal(err)
	}

	// 删除测试目录
	err = os.RemoveAll("testdata/srcDir")
	if err != nil {
		t.Fatal(err)
	}
	err = os.RemoveAll("testdata/dstDir")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDirCopy(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				src: "testdata/srcDir",
				dst: "testdata/dstDir",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建测试目录
			ok, err := CreateDir(tt.args.src)
			if !ok && err != nil {
				t.Fatal(err)
			}
			// 创建测试文件 testfile1
			err = CreateFile(path.Join(tt.args.src, "testfile1.txt"))
			if err != nil {
				t.Fatal(err)
			}
			// 创建测试子目录 subdir1
			ok, err = CreateDir(path.Join(tt.args.src, "subdir1"))
			if !ok && err != nil {
				t.Fatal(err)
			}
			// 创建测试文件 subdir1/testfile2
			err = CreateFile(path.Join(tt.args.src, "subdir1", "testfile2.txt"))
			if err != nil {
				t.Fatal(err)
			}
			// 打印测试目录
			fmt.Println("测试目录:")
			err = PrintDirTree(tt.args.src, -1, false, true)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println("目标目录:")
			err = PrintDirTree(tt.args.dst, -1, false, true)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				t.Fatal("目标目录此时不应该存在!")
			}
			if err := DirCopy(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("DirCopy() error = %v, wantErr %v", err, tt.wantErr)
			}
			// 检查目标目录
			fmt.Println("目标目录(测试后):")
			err = PrintDirTree(tt.args.dst, -1, false, true)
			if err != nil {
				t.Fatal(err)
			}
			// 删除测试目录与目标目录
			if err = os.RemoveAll(tt.args.src); err != nil {
				t.Fatal(err)
			}
			if err = os.RemoveAll(tt.args.dst); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestReadFileToLinesAndAll(t *testing.T) {
	// 写入测试数据
	err := writeFile("testdata/testfile.txt", []byte("line1\nline2\nline3\nline4\nline5"))
	if err != nil {
		t.Fatal(err)
	}
	// 读取文件
	lines, all, err := ReadFileToLinesAndAll("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 打印lines
	fmt.Println("按行读取结果:")
	for _, line := range lines {
		fmt.Println(line)
	}
	// 打印all
	fmt.Println("全部读取结果:")
	fmt.Println(all)
	// 检查读取结果
	if len(lines) != 5 {
		t.Fatal("按行读取错误!")
	}
	if all != "line1\nline2\nline3\nline4\nline5" {
		t.Fatal("全部读取错误!")
	}
	// 删除测试文件
	err = os.Remove("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadFileToLinesBySize(t *testing.T) {
	zclog.Level = zclog.LOG_LEVEL_DEBUG
	// 写入测试数据
	err := writeFile("testdata/testfile.txt", []byte("line1\nline2\nline3\nline4\nline5"))
	if err != nil {
		t.Fatal(err)
	}
	// 读取文件
	lines, err := ReadFileToLinesBySize("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 打印lines
	fmt.Println("按行读取结果:")
	for _, line := range lines {
		fmt.Println(line)
	}
	// 检查读取结果
	if len(lines) != 5 {
		t.Fatal("按行读取错误!")
	}
	// 删除测试文件
	err = os.Remove("testdata/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}

	// 通过循环，流式写入大小超过1MB的文件 testdata/testBigFile.txt
	err = writeLargeFile("testdata/testBigFile.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 读取文件
	lines, err = ReadFileToLinesBySize("testdata/testBigFile.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 打印lines
	fmt.Println("按行读取结果(前10行):")
	for i, line := range lines {
		fmt.Println(line)
		if i == 9 {
			break
		}
	}
	fmt.Printf("按行读取行数: %d\n", len(lines))
	// 删除测试文件
	err = os.Remove("testdata/testBigFile.txt")
	if err != nil {
		t.Fatal(err)
	}
}

const megabyte = 1024 * 1024

func writeFile(filePath string, bytesContent []byte) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	_, err = file.Write(bytesContent)
	if err != nil {
		return err
	}
	return nil
}

func writeLargeFile(filePath string) error {
	// 创建文件
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	// 创建带缓存的Writer
	writer := bufio.NewWriterSize(f, 4096)
	// 写入1MB数据
	for i := 0; i < megabyte; i++ {
		_, err = writer.WriteString(strconv.Itoa(i) + "\n")
		if err != nil {
			return err
		}

		// 定期flush以清空缓存
		if i%100 == 0 {
			_ = writer.Flush()
		}
	}
	// 写入剩余缓存数据
	_ = writer.Flush()
	log.Println("Generated file larger than 1MB")
	return nil
}

func TestSplitPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				path: "/a/b/c/d.txt",
			},
			want: []string{"a", "b", "c", "d.txt"},
		},
		{
			name: "test2",
			args: args{
				path: "a/b/c/d.txt",
			},
			want: []string{"a", "b", "c", "d.txt"},
		},
		{
			name: "test3",
			args: args{
				path: "./a/b/c/d.txt",
			},
			want: []string{".", "a", "b", "c", "d.txt"},
		},
		{
			name: "test4",
			args: args{
				path: "../a/b/c/d.txt",
			},
			want: []string{"..", "a", "b", "c", "d.txt"},
		},
		{
			name: "test5",
			args: args{
				path: "C:\\a\\b\\c\\d.txt",
			},
			want: []string{"C:", "a", "b", "c", "d.txt"},
		},
		{
			name: "test6",
			args: args{
				path: "\\a\\b\\c\\d.txt",
			},
			want: []string{"a", "b", "c", "d.txt"},
		},
		{
			name: "test7",
			args: args{
				path: "a\\b\\c\\d.txt",
			},
			want: []string{"a", "b", "c", "d.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				path: "/a/b/c/d.txt",
			},
			want: "a",
		},
		{
			name: "test2",
			args: args{
				path: "a/b/c/d.txt",
			},
			want: "a",
		},
		{
			name: "test3",
			args: args{
				path: "./a/b/c/d.txt",
			},
			want: ".",
		},
		{
			name: "test4",
			args: args{
				path: "../a/b/c/d.txt",
			},
			want: "..",
		},
		{
			name: "test5",
			args: args{
				path: "C:\\a\\b\\c\\d.txt",
			},
			want: "C:",
		},
		{
			name: "test6",
			args: args{
				path: "\\a\\b\\c\\d.txt",
			},
			want: "a",
		},
		{
			name: "test7",
			args: args{
				path: "a\\b\\c\\d.txt",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstDir(tt.args.path); got != tt.want {
				t.Errorf("FirstDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileNotEmpty(t *testing.T) {
	// 创建测试目录
	ok, err := CreateDir("testdata/TestFileNotEmpty")
	if !ok && err != nil {
		t.Fatal(err)
	}
	// 创建测试文件 testfile1.txt size为0
	err = CreateFile(path.Join("testdata/TestFileNotEmpty", "testfile1.txt"))
	if err != nil {
		t.Fatal(err)
	}
	// 创建测试文件 testfile2.txt size大于0
	err = CreateFile(path.Join("testdata/TestFileNotEmpty", "testfile2.txt"))
	if err != nil {
		t.Fatal(err)
	}
	// 向 testfile2.txt 写入数据
	err = writeFile(path.Join("testdata/TestFileNotEmpty", "testfile2.txt"), []byte("test"))
	if err != nil {
		t.Fatal(err)
	}
	// 创建测试子目录 subdir
	ok, err = CreateDir(path.Join("testdata/TestFileNotEmpty", "subdir"))
	if !ok && err != nil {
		t.Fatal(err)
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				path: "testdata/TestFileNotEmpty/testfile1.txt",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				path: "testdata/TestFileNotEmpty/testfile2.txt",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				path: "testdata/TestFileNotEmpty/subdir",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileNotEmpty(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileNotEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileNotEmpty() got = %v, want %v", got, tt.want)
			}
		})
	}
	// 删除测试文件
	err = os.Remove("testdata/TestFileNotEmpty/testfile1.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 删除测试文件
	err = os.Remove("testdata/TestFileNotEmpty/testfile2.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 删除测试子目录 subdir
	err = os.Remove("testdata/TestFileNotEmpty/subdir")
	if err != nil {
		t.Fatal(err)
	}
	// 删除测试目录
	err = os.RemoveAll("testdata/TestFileNotEmpty")
	if err != nil {
		t.Fatal(err)
	}
}
