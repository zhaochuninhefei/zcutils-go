package zctime

import (
	"fmt"
	"strings"
	"time"
)

type ZcTimeFormat string

//goland:noinspection GoSnakeCaseUsage,GoUnusedConst
const (
	TIME_FORMAT_SIMPLE       ZcTimeFormat = "20060102150405"          // 简单格式: yyyyMMddHHmmss
	TIME_FORMAT_SLASH        ZcTimeFormat = "2006/01/02 15:04:05"     // 斜杠格式: yyyy/MM/dd HH:mm:ss
	TIME_FORMAT_DASH         ZcTimeFormat = "2006-01-02 15:04:05"     // 横杠格式: yyyy-MM-dd HH:mm:ss
	TIME_FORMAT_SIMPLE_MILLI ZcTimeFormat = "20060102150405.000"      // 带毫秒的简单格式: yyyyMMddHHmmss.SSS
	TIME_FORMAT_SLASH_MILLI  ZcTimeFormat = "2006/01/02 15:04:05.000" // 斜杠格式: yyyy/MM/dd HH:mm:ss.SSS
	TIME_FORMAT_DASH_MILLI   ZcTimeFormat = "2006-01-02 15:04:05.000" // 横杠格式: yyyy-MM-dd HH:mm:ss.SSS
)

func (ztf ZcTimeFormat) String() string {
	return string(ztf)
}

func (ztf ZcTimeFormat) FormatTimeToStr(time time.Time) string {
	return time.Format(ztf.String())
}

func (ztf ZcTimeFormat) FormatTimeToSimpleMilliNoPoint(time time.Time) (string, error) {
	if ztf != TIME_FORMAT_SIMPLE_MILLI {
		return "", fmt.Errorf("当前ZcTimeFormat无法实现ParseSimpleMilliNoPointToTime方法: %s", ztf.String())
	}
	timeStr := time.Format(ztf.String())
	return strings.ReplaceAll(timeStr, ".", ""), nil
}

func (ztf ZcTimeFormat) ParseStrToTime(timeStr string) (time.Time, error) {
	return time.Parse(ztf.String(), timeStr)
}

func (ztf ZcTimeFormat) ParseSimpleMilliNoPointToTime(timeStr string) (time.Time, error) {
	if ztf != TIME_FORMAT_SIMPLE_MILLI {
		return time.Time{}, fmt.Errorf("当前ZcTimeFormat无法实现ParseSimpleMilliNoPointToTime方法: %s", ztf.String())
	}
	if len(timeStr) != 17 {
		return time.Time{}, fmt.Errorf("试图转换的时间字符串长度不是17: %s", timeStr)
	}
	timeNewStr := timeStr[:14] + "." + timeStr[15:]
	return time.Parse(ztf.String(), timeNewStr)
}
