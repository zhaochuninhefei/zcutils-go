package zcsync

import (
	"fmt"
	"testing"
	"time"
)

// 测试CallSync
func TestCallSync(t *testing.T) {
	// 创建一个Caller
	caller := NewCaller()
	// 执行CallSync
	err := caller.CallSync(func() error {
		fmt.Println("目标函数执行开始")
		time.Sleep(3 * time.Second)
		fmt.Println("目标函数执行结束")
		return nil
	})
	// 判断结果
	if err != nil {
		t.Error(err)
	}
}

// 测试CallSync并发
func TestCallSyncConcurrent(t *testing.T) {
	// 创建一个Caller
	caller := NewCaller()
	// 并发执行CallSync
	for i := 0; i < 10; i++ {
		go func(index int) {
			err := caller.CallSync(func() error {
				fmt.Printf("目标函数 %d 执行开始\n", index)
				time.Sleep(1 * time.Second)
				fmt.Printf("目标函数 %d 执行结束\n", index)
				return nil
			})
			if err != nil {
				t.Error(err)
			}
		}(i)
	}
	// 等待15秒
	time.Sleep(15 * time.Second)
	fmt.Println("测试CallSync并发结束")
}
