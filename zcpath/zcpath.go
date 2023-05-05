package zcpath

import (
	"fmt"
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
	if _, err := os.Stat(dir); err != nil {
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
