package zctime

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()

	timeStr := now.Format(TIME_FORMAT_SIMPLE.String())
	fmt.Println(timeStr)

	timeStr = now.Format(TIME_FORMAT_SIMPLE_MILLI.String())
	fmt.Println(timeStr)

	timeObj, err := time.Parse(TIME_FORMAT_SIMPLE.String(), timeStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
}
