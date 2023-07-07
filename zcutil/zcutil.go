package zcutil

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"gitee.com/zhaochuninhefei/zcutils-go/zcpath"
	"github.com/fsnotify/fsnotify"
	"github.com/nxadm/tail"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Int32ToBytes 将int32值转为4个字节的byte数组
func Int32ToBytes(i int32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

// BytesToInt32 将byte数组转为int32
func BytesToInt32(buf []byte) int32 {
	return int32(binary.BigEndian.Uint32(buf))
}

// PrintStack 打印调用栈
func PrintStack(msg string) {
	var pcs [32]uintptr
	n := runtime.Callers(2, pcs[:]) // skip first 3 frames
	frames := runtime.CallersFrames(pcs[:n])
	if msg != "" {
		fmt.Println(msg)
	}
	fmt.Println("当前调用栈如下:")
	for {
		frame, more := frames.Next()
		fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}

// If 判断条件是否为真，如果为真则返回第一个参数，否则返回第二个参数
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

type IfCondition func() bool
type IfTrueVal func() (interface{}, error)
type IfFalseVal func() (interface{}, error)

func IfByFunc(condition IfCondition, trueVal IfTrueVal, falseVal IfFalseVal) (interface{}, error) {
	if condition() {
		return trueVal()
	}
	return falseVal()
}

type IfTrueValNoError func() interface{}
type IfFalseValNoError func() interface{}

func IfByFuncNoErr(condition IfCondition, trueVal IfTrueValNoError, falseVal IfFalseValNoError) interface{} {
	if condition() {
		return trueVal()
	}
	return falseVal()
}

// CallAsyncFuncAndWaitByLog 调用异步函数并根据日志处理函数等待结束
//  - logPath 用于监听异步函数是否执行结束的日志文件,会在调用异步函数funcAsync之前删除
//  - funcAsync 异步函数,注意CallAsyncFuncAndWaitByLog内部会在删除日志文件后直接以同步方式调用该函数.即funcAsync的异步处理是由其内部完成的.
//  - funcHandlerLogLine 日志按行处理函数,该函数以一行日志为入参,根据该函数的返回值确定是否结束等待.返回bool为true时则结束等待,返回error非空时结束等待并返回该error.
//  - timeoutSeconds 等待的超时秒数,超过该时间则立即结束等待并返回超时错误.
// 该函数内部会先删除logPath文件,然后执行funcAsync,然后tail监听logPath文件,将获取到的每一行日志作为入参调用funcHandlerLogLine进行判断,根据结果决定是否继续tail.如果超时则直接返回超时错误.
func CallAsyncFuncAndWaitByLog(logPath string, funcAsync func() error, funcHandlerLogLine func(line string) (bool, error), timeoutSeconds int) error {
	// 检查funcAsync和funcHandlerLogLine是否为nil
	if funcAsync == nil || funcHandlerLogLine == nil {
		return errors.New("[params error]funcAsync和funcHandlerLogLine不能为nil")
	}

	// 删除日志文件
	err := zcpath.RemoveFile(logPath)
	if err != nil {
		return fmt.Errorf("[logfile error]删除日志文件失败: %s", err.Error())
	}

	// 执行funcAsync
	err = funcAsync()
	if err != nil {
		return fmt.Errorf("[funcAsync error]执行异步函数失败: %s", err.Error())
	}

	fmt.Printf("==== tail %s start ====\n", logPath)
	// 配置超时Context, 默认90秒
	// 如果timeoutSeconds<=0, 则使用默认值90秒
	if timeoutSeconds <= 0 {
		timeoutSeconds = 90
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	// 读取日志文件
	t, err := tail.TailFile(logPath, tail.Config{
		Follow: true,
		ReOpen: true,
	})
	if err != nil {
		return fmt.Errorf("[tail error]监听日志文件失败: %s", err.Error())
	}
	for {
		select {
		case <-ctx.Done():
			// 监听日志文件超时
			err = t.Stop()
			if err != nil {
				fmt.Printf(err.Error())
			}
			return fmt.Errorf("[tail timeout]监听 %s.log 超时", logPath)
		case line := <-t.Lines:
			end, err := funcHandlerLogLine(line.Text)
			if err != nil {
				fmt.Printf("==== tail %s stop by error : %s ====\n", logPath, err.Error())
				errStop := t.Stop()
				if errStop != nil {
					fmt.Println(errStop.Error())
				}
				return fmt.Errorf("[funcHandlerLogLine error]处理日志行返回错误: %s", err.Error())
			}
			if end {
				fmt.Printf("==== tail %s finish. ====\n", logPath)
				err = t.Stop()
				if err != nil {
					fmt.Println(err.Error())
				}
				return nil
			}
		}
	}
}

func CallAsyncFuncAndWaitByFlag(flagPath, logPath string, funcAsync func() error, timeoutSeconds int) ([]string, error) {
	// 删除标志文件与日志文件
	if err := zcpath.RemoveFile(flagPath); err != nil {
		return nil, fmt.Errorf("[flagfile error]删除标志文件失败: %s", err.Error())
	}
	if err := zcpath.RemoveFile(logPath); err != nil {
		return nil, fmt.Errorf("[logfile error]删除日志文件失败: %s", err.Error())
	}

	// 创建监听器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("[watcher error]创建监听失败: %s", err.Error())
	}
	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {
			fmt.Printf("[watcher error]关闭监听失败: %s", err.Error())
		}
	}(watcher)
	// 监听flag文件所在目录
	dirPath := filepath.Dir(flagPath)
	if err = watcher.Add(dirPath); err != nil {
		return nil, fmt.Errorf("[watcher error]添加监听目录失败: %s", err.Error())
	}

	// 调用异步函数
	if err = funcAsync(); err != nil {
		return nil, fmt.Errorf("[funcAsync error]调用异步函数失败: %s", err.Error())
	}

	// 配置超时Context, 默认90秒
	// 如果timeoutSeconds<=0, 则使用默认值90秒
	if timeoutSeconds <= 0 {
		timeoutSeconds = 90
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Create == fsnotify.Create {
				fmt.Printf("监听到 %s 被创建\n", event.Name)
				// 检查被创建的文件是否为flag文件
				if event.Name != flagPath {
					continue
				}
				fmt.Printf("标志文件[%s]已创建\n", flagPath)
				// 停止监听
				if err = watcher.Remove(dirPath); err != nil {
					return nil, fmt.Errorf("[watcher error]移除监听目录[%s]失败: %s", dirPath, err.Error())
				}
				// 读取日志文件
				logFile, err := os.Open(logPath)
				if err != nil {
					return nil, fmt.Errorf("[logfile error]打开日志文件[%s]失败: %s", logPath, err.Error())
				}
				//goland:noinspection GoDeferInLoop
				defer func(logFile *os.File) {
					err := logFile.Close()
					if err != nil {
						fmt.Printf("关闭日志文件[%s]失败: %s\n", flagPath, err.Error())
					}
				}(logFile)
				scanner := bufio.NewScanner(logFile)
				var logs []string
				for scanner.Scan() {
					logs = append(logs, scanner.Text())
				}
				if err = scanner.Err(); err != nil {
					return nil, fmt.Errorf("[logfile error]读取日志文件[%s]失败: %s", logPath, err.Error())
				}
				return logs, nil
			}
		case err := <-watcher.Errors:
			errMsg := fmt.Sprintf("[watcher error]监听标志文件[%s]发生错误: %s", flagPath, err.Error())
			fmt.Println(errMsg)
			return nil, errors.New(errMsg)
		case <-ctx.Done():
			errMsg := fmt.Sprintf("[watcher timeout]监听标志文件[%s]超时", flagPath)
			fmt.Println(errMsg)
			// 停止监听
			if err = watcher.Remove(dirPath); err != nil {
				fmt.Printf("[watcher error]移除监听目录[%s]失败: %s\n", dirPath, err.Error())
			}
			return nil, errors.New(errMsg)
		}
	}
}
