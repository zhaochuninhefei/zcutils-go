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

func TestNormalFormatAndParse(t *testing.T) {
	now := time.Now()
	fmt.Printf("now : %s\n", now)

	testNormalFormatAndParse(t, now, TIME_FORMAT_SIMPLE)
	testNormalFormatAndParse(t, now, TIME_FORMAT_DASH)
	testNormalFormatAndParse(t, now, TIME_FORMAT_SLASH)

	testNormalFormatAndParse(t, now, TIME_FORMAT_SIMPLE_MILLI)
	testNormalFormatAndParse(t, now, TIME_FORMAT_DASH_MILLI)
	testNormalFormatAndParse(t, now, TIME_FORMAT_SLASH_MILLI)

	fmt.Println()
}

func testNormalFormatAndParse(t *testing.T, now time.Time, timeFormat Format) {
	fmt.Printf("\ntimeFormat: %s\n", timeFormat.String())

	timeStr := timeFormat.FormatTimeToStr(now)
	fmt.Printf("timeStr: %s\n", timeStr)
	timeNew, err := timeFormat.ParseStrToTime(timeStr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("timeNew : %s\n", timeNew)
}

func TestSimpleMilliNoPoint(t *testing.T) {
	now := time.Now()
	fmt.Printf("now : %s\n", now)

	simpleMilliNoPoint, err := TIME_FORMAT_SIMPLE_MILLI.FormatTimeToSimpleMilliNoPoint(now)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("simpleMilliNoPoint: %s\n", simpleMilliNoPoint)

	timeNew, err := TIME_FORMAT_SIMPLE_MILLI.ParseSimpleMilliNoPointToTime(simpleMilliNoPoint)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("timeNew: %s\n", timeNew)

	timeNew2, err := TIME_FORMAT_SIMPLE_MILLI.ParseSimpleMilliNoPointToTime("99991231235959100")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("timeNew2: %s\n", timeNew2)

	_, err = TIME_FORMAT_SIMPLE_MILLI.ParseSimpleMilliNoPointToTime("9999123123595910")
	if err == nil {
		t.Fatal("应该返回错误")
	} else {
		fmt.Println(err.Error())
	}

	_, err = TIME_FORMAT_SLASH_MILLI.FormatTimeToSimpleMilliNoPoint(now)
	if err == nil {
		t.Fatal("应该返回错误")
	} else {
		fmt.Println(err.Error())
	}

	_, err = TIME_FORMAT_SLASH_MILLI.ParseSimpleMilliNoPointToTime(simpleMilliNoPoint)
	if err == nil {
		t.Fatal("应该返回错误")
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println()
}

func TestGetNowYMDHMS(t *testing.T) {
	timeNow := GetNowYMDHMS()
	fmt.Println(timeNow)
}
