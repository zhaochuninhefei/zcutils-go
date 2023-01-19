package zctime

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()

	timeStr := now.Format(TIME_FORMAT_SIMPLE)
	fmt.Println(timeStr)

	timeObj, err := time.Parse(TIME_FORMAT_SIMPLE, timeStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
}
