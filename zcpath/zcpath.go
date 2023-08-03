package zcpath

import (
	"bufio"
	"fmt"
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// CreateDir 创建目录,若目录已经存在，会返回错误信息但bool返回true。
//  @param path 目录路径
//  @return bool 创建成功与否
//  @return error
func CreateDir(path string) (bool, error) {
	if dirInfo, err := os.Stat(path); err == nil && dirInfo.IsDir() {
		return true, fmt.Errorf("该目录已经存在: %s", path)
	} else {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// ClearDir 清空目录
//  该函数会删除目录下的所有文件和子目录, 但是不会删除目录本身
func ClearDir(dir string) error {
	// 先判断dir是否存在, 如果不存在, 则创建
	if dirInfo, err := os.Stat(dir); err != nil || !dirInfo.IsDir() {
		err1 := os.MkdirAll(dir, 0755)
		if err1 != nil {
			return err1
		}
	}
	// Open the directory
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer func(d *os.File) {
		err := d.Close()
		if err != nil {
			panic(err)
		}
	}(d)

	// Read the directory entries
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	// Loop over the entries
	for _, name := range names {
		// Join the directory path and the entry name
		fullPath := filepath.Join(dir, name)

		// Remove the entry using RemoveAll, which works for both files and directories
		err = os.RemoveAll(fullPath)
		if err != nil {
			return err
		}
	}

	return nil
}

// IsDirEmpty 判断目录是否为空
func IsDirEmpty(dirPath string) (bool, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return false, err
	}

	if len(files) == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// CreateFile 创建文件
func CreateFile(filePath string) error {
	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	// 关闭文件
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

// RemoveFile 删除文件, 如果文件不存在，不报错
func RemoveFile(filePath string) error {
	err := os.Remove(filePath)
	// 如果文件不存在，不报错
	if err != nil && os.IsNotExist(err) {
		return nil
	} else {
		return err
	}
}

// RemoveFileWithWildcard 根据带通配符的文件路径删除文件
//  - filePath: 要删除的文件path, 如"/a/b/*.txt"
func RemoveFileWithWildcard(filePath string) error {
	// 使用通配符查找匹配的文件
	matches, err := filepath.Glob(filePath)
	if err != nil {
		return fmt.Errorf("无法匹配文件:[%s], 发生错误: %s", filePath, err)
	}
	// 遍历匹配的文件
	for _, match := range matches {
		err = RemoveFile(match)
		if err != nil {
			return err
		}
	}
	return nil
}

// PrintDirTree 打印目录树
//  @param root 根目录
//  @param level 打印层级, 0只打印根目录自身，1表示打印一层(根目录下的文件与目录), 2表示打印两层, 依次类推。-1表示打印所有层级。
//  @param onlyDir 只打印目录
//  @param showHidden 是否显示隐藏文件
func PrintDirTree(root string, level int, onlyDir bool, showHidden bool) error {
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !showHidden && info.Name()[0] == '.' {
			if info.IsDir() {
				return filepath.SkipDir
			} else {
				return nil
			}
		}
		if info.IsDir() {
			if path == root || level == -1 || strings.Count(path[len(root):], string(os.PathSeparator)) <= level {
				fmt.Println(path)
			} else {
				return filepath.SkipDir
			}
		} else {
			if !onlyDir && (level == -1 || strings.Count(path[len(root):], string(os.PathSeparator)) <= level) {
				fmt.Println(path)
			}
		}
		return nil
	}
	err := filepath.Walk(root, walkFunc)
	if err != nil {
		return err
	}
	return nil
}

func PrintDirTreeWithMode(root string, level int, onlyDir bool, showHidden bool) error {
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !showHidden && info.Name()[0] == '.' {
			if info.IsDir() {
				return filepath.SkipDir
			} else {
				return nil
			}
		}
		if info.IsDir() {
			if path == root || level == -1 || strings.Count(path[len(root):], string(os.PathSeparator)) <= level {
				// 获取文件权限
				mode := info.Mode()
				// 打印文件权限
				fmt.Printf("%s %s\n", path, mode.String())
			} else {
				return filepath.SkipDir
			}
		} else {
			if !onlyDir && (level == -1 || strings.Count(path[len(root):], string(os.PathSeparator)) <= level) {
				// 获取文件权限
				mode := info.Mode()
				// 打印文件权限
				fmt.Printf("%s %s\n", path, mode.String())
			}
		}
		return nil
	}
	err := filepath.Walk(root, walkFunc)
	if err != nil {
		return err
	}
	return nil
}

// FileExists 判断文件是否存在，且是不是文件
//  @param path 文件路径
func FileExists(path string) (bool, error) {
	// 判断文件是否存在
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	// 判断是否是目录
	if fileInfo.IsDir() {
		return false, nil
	}
	return true, nil
}

// FileNotEmpty 判断文件是否非空
//  @param path 文件路径
//  文件不存在、或不是文件而是目录、或文件size为0时都返回 false;
//  文件存在、且是文件、且size>0时返回 true;
//  获取文件信息失败时返回 false和error。
func FileNotEmpty(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	// 判断是否是目录
	if fileInfo.IsDir() {
		return false, nil
	}
	// 判断文件size
	if fileInfo.Size() > 0 {
		return true, nil
	}
	return false, nil
}

// FileCopy 拷贝文件
//  @param src 源文件路径
//  @param dst 目标文件路径
func FileCopy(src string, dst string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			panic(err)
		}
	}(srcFile)
	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			panic(err)
		}
	}(dstFile)
	// 拷贝文件
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

// FileCopyToDir 拷贝文件到目录
//  @param src 源文件路径
//  @param dstDir 目标目录
func FileCopyToDir(src string, dstDir string) error {
	// 获取源文件名
	srcFileName := filepath.Base(src)
	// 拼接目标文件路径
	dstFilePath := filepath.Join(dstDir, srcFileName)
	// 创建目标目录
	ok, err := CreateDir(dstDir)
	if !ok && err != nil {
		return err
	}
	// 拷贝文件
	err = FileCopy(src, dstFilePath)
	if err != nil {
		return err
	}
	return nil
}

// FileCopyFromDirToDir 拷贝源目录下所有文件到目标目录，包括子目录及子目录下的文件
//  @param srcDir 源目录
//  @param dstDir 目标目录
func FileCopyFromDirToDir(srcDir string, dstDir string) error {
	// 检查srcDir是否存在,是否是目录
	srcInfo, err := os.Stat(srcDir)
	if err != nil {
		return err
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("srcDir(%s)不是目录", srcDir)
	}
	// 创建目标目录
	ok, err := CreateDir(dstDir)
	if !ok && err != nil {
		return err
	}
	// 获取源目录下所有文件和目录
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}
	// 遍历所有文件和目录
	for _, file := range files {
		srcPath := filepath.Join(srcDir, file.Name())
		destPath := filepath.Join(dstDir, file.Name())

		if file.IsDir() {
			// 如果是目录，递归拷贝目录
			err = FileCopyFromDirToDir(srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			// 如果是文件，拷贝文件
			data, err := ioutil.ReadFile(srcPath)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(destPath, data, file.Mode())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// FileWithWildcardCopyToDir 将含有通配符的文件路径指向的文件拷贝到目标目录下
//  - sourcePath: 含有通配符(*)的文件路径，如: /a/b/*.txt
//  - dstDir: 目标目录，如: /c/d
//
// 效果类似: cp /a/b/*.txt /c/d
func FileWithWildcardCopyToDir(sourcePath string, dstDir string) error {
	// 使用通配符查找匹配的文件
	matches, err := filepath.Glob(sourcePath)
	if err != nil {
		return fmt.Errorf("无法匹配文件:[%s], 发生错误: %s", sourcePath, err)
	}
	// 遍历匹配的文件
	for _, match := range matches {
		err = FileCopyToDir(match, dstDir)
		if err != nil {
			return err
		}
	}
	return nil
}

// DirCopy 拷贝目录
//  例如: DirCopy("/x/y/z", "/a/b/c") 效果: 将子目录z拷贝到目标目录/a/b/c下面，即"/a/b/c/z"
func DirCopy(src string, dst string) error {
	// 获取源目录名
	srcDirName := filepath.Base(src)
	// 拼接目标目录路径
	dstDirPath := filepath.Join(dst, srcDirName)
	// 拷贝源目录下所有文件到目标目录，包括子目录及子目录下的文件
	return FileCopyFromDirToDir(src, dstDirPath)
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

// ChmodDir 修改目录权限
func ChmodDir(dir string, mode os.FileMode) error {
	// 修改目录权限
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		err = os.Chmod(path, mode)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// ReadFileToLinesAndAll 从filePath读取文件内容，返回每一行的内容和所有内容
//  使用ioutil.ReadFile一次性读取文件全部字节，建议只在读取小文件(<1MB)时使用该函数。
func ReadFileToLinesAndAll(filePath string) ([]string, string, error) {
	// 读取文件内容
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, "", err
	}
	// 将文件内容转换为字符串
	fileContent := string(fileBytes)
	// 将文件内容按行存储到字符串切片中
	lines := strings.Split(fileContent, "\n")
	return lines, fileContent, nil
}

// 大文件阈值
const maxFileSize = 1 << 20 // 1MB

// ReadFileToLinesBySize 根据文件大小读取行
//  当文件大小不足1MB时，整个读入内存;
//  当文件大小达到或超过1MB时，按行读取，注意每行token数量不能超过65536.
func ReadFileToLinesBySize(filename string) ([]string, error) {
	// 获取文件信息
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	// 根据大小选择读取方式
	var lines []string
	if fi.Size() < maxFileSize {
		zclog.Debugf("按小文件读取, size: %d", fi.Size())
		// 小文件,直接读取到内存
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		lines = strings.Split(string(content), "\n")
	} else {
		zclog.Debugf("按大文件读取, size: %d", fi.Size())
		// 大文件,流式读取
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}

	return lines, nil
}

// SplitPath 切分路径
//  @param path 路径
//  例如:
//  SplitPath("/a/b/c/d.txt") = []string{"a", "b", "c", "d.txt"}
//  SplitPath("a/b/c/d.txt") = []string{"a", "b", "c", "d.txt"}
//  SplitPath("./a/b/c/d.txt") = []string{".", "a", "b", "c", "d.txt"}
//  SplitPath("../a/b/c/d.txt") = []string{"..", "a", "b", "c", "d.txt"}
//  SplitPath("C:\a\b\c\d.txt") = []string{"C:", "a", "b", "c", "d.txt"}
//  SplitPath("\a\b\c\d.txt") = []string{"a", "b", "c", "d.txt"}
//  SplitPath("a\b\c\d.txt") = []string{"a", "b", "c", "d.txt"}
func SplitPath(path string) []string {
	// 将路径统一转换为/,方便后续处理
	path = strings.ReplaceAll(path, "\\", "/")
	// 如果 path以"/"开头，去除开头的"/"
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	// 切分路径
	parts := strings.Split(path, "/")
	return parts
}

// FirstDir 获取path的第一个目录
//  @param path 路径
//  例如:
//  FirstDir("/a/b/c/d.txt") = "a"
//  FirstDir("a/b/c/d.txt") = "a"
//  FirstDir("./a/b/c/d.txt") = "."
//  FirstDir("../a/b/c/d.txt") = ".."
//  FirstDir("C:\a\b\c\d.txt") = "C:"
//  FirstDir("\a\b\c\d.txt") = "a"
//  FirstDir("a\b\c\d.txt") = "a"
func FirstDir(path string) string {
	parts := SplitPath(path)
	return parts[0]
}
