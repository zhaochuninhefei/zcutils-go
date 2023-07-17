package zcwaiter

import (
	"errors"
	"fmt"
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"time"
)

// Waiter 等待器
type Waiter struct {
	waitMaxTimes  int
	waitMSPerTime time.Duration
}

// NewWaiter 创建等待器，采用默认的waitMaxTimes与waitMSPerTime
//  waitMaxTimes: 最大等待次数，默认300
//  waitMSPerTime: 每次等待时间(毫秒)，默认1000
func NewWaiter() *Waiter {
	return &Waiter{
		waitMaxTimes:  300,
		waitMSPerTime: 1000,
	}
}

// NewCustomWaiter 创建等待器，采用指定的waitMaxTimes与waitMSPerTime
//  waitMaxTimes: 最大等待次数，必须大于0
//  waitMSPerTime: 每次等待时间(毫秒)，必须大于0
func NewCustomWaiter(waitMaxTimes int, waitMSPerTime time.Duration) (*Waiter, error) {
	if waitMaxTimes <= 0 {
		return nil, errors.New("waitMaxTimes must be greater than 0")
	}
	if waitMSPerTime <= 0 {
		return nil, errors.New("waitMSPerTime must be greater than 0")
	}
	return &Waiter{
		waitMaxTimes:  waitMaxTimes,
		waitMSPerTime: waitMSPerTime,
	}, nil
}

func (w *Waiter) WaitUntil(supplier func() bool) error {
	zclog.Printf("waitMaxTimes : %d; waitMSPerTime : %d", w.waitMaxTimes, w.waitMSPerTime)
	ticker := time.NewTicker(w.waitMSPerTime * time.Millisecond)
	defer ticker.Stop()
	for i := 0; i < w.waitMaxTimes; i++ {
		select {
		case <-ticker.C:
			if supplier() {
				zclog.Println("处理结束。")
				return nil
			} else {
				zclog.Println("处理尚未结束，请等待。。。")
			}
		}
	}
	return fmt.Errorf("处理超时，在等待了 %d 次(每次等待 %d 毫秒)之后指定条件仍然未满足", w.waitMaxTimes, w.waitMSPerTime)
}
