package zcutil

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"gitee.com/zhaochuninhefei/zcutils-go/zcpath"
	"github.com/nxadm/tail"
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

// CallAsyncFuncAndWaitByLog 调用异步函数并根据日志处理函数等待结束
//  - logPath 用于监听异步函数是否执行结束的日志文件,会在调用异步函数funcAsync之前删除
//  - funcAsync 异步函数,注意CallAsyncFuncAndWaitByLog内部会在删除日志文件后直接以同步方式调用该函数.即funcAsync的异步处理是由其内部完成的.
//  - funcHandlerLogLine 日志按行处理函数,该函数以一行日志为入参,根据该函数的返回值确定是否结束等待.返回bool为true时则结束等待,返回error非空时结束等待并返回该error.
//  - timeoutSeconds 等待的超时秒数,超过该时间则立即结束等待并返回超时错误.
// 该函数内部会先删除logPath文件,然后执行funcAsync,然后tail监听logPath文件,将获取到的每一行日志作为入参调用funcHandlerLogLine进行判断,根据结果决定是否继续tail.如果超时则直接返回超时错误.
func CallAsyncFuncAndWaitByLog(logPath string, funcAsync func() error, funcHandlerLogLine func(line string) (bool, error), timeoutSeconds int) error {
	// 检查funcAsync和funcHandlerLogLine是否为nil
	if funcAsync == nil || funcHandlerLogLine == nil {
		return errors.New("funcAsync和funcHandlerLogLine不能为nil")
	}

	// 删除日志文件
	err := zcpath.RemoveFile(logPath)
	if err != nil {
		return err
	}

	// 执行funcAsync
	err = funcAsync()
	if err != nil {
		return err
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
		return err
	}
	for {
		select {
		case <-ctx.Done():
			// 监听日志文件超时
			err = t.Stop()
			if err != nil {
				fmt.Printf(err.Error())
			}
			return fmt.Errorf("tail %s.log timeout", logPath)
		case line := <-t.Lines:
			end, err := funcHandlerLogLine(line.Text)
			if err != nil {
				errStop := t.Stop()
				if errStop != nil {
					fmt.Println(errStop.Error())
				}
				return err
			}
			if end {
				err = t.Stop()
				if err != nil {
					fmt.Println(err.Error())
				}
				return nil
			}
		}
	}
}
