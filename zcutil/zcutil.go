package zcutil

import (
	"encoding/binary"
	"fmt"
	"runtime"
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
