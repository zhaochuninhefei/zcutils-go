package zctime

import (
	"fmt"
	"strings"
	"time"
)

// Format zctime时间格式
type Format string

//goland:noinspection GoSnakeCaseUsage,GoUnusedConst
const (
	TIME_FORMAT_SIMPLE       Format = "20060102150405"          // 简单格式: yyyyMMddHHmmss
	TIME_FORMAT_SLASH        Format = "2006/01/02 15:04:05"     // 斜杠格式: yyyy/MM/dd HH:mm:ss
	TIME_FORMAT_DASH         Format = "2006-01-02 15:04:05"     // 横杠格式: yyyy-MM-dd HH:mm:ss
	TIME_FORMAT_SIMPLE_MILLI Format = "20060102150405.000"      // 带毫秒的简单格式: yyyyMMddHHmmss.SSS
	TIME_FORMAT_SLASH_MILLI  Format = "2006/01/02 15:04:05.000" // 带毫秒的斜杠格式: yyyy/MM/dd HH:mm:ss.SSS
	TIME_FORMAT_DASH_MILLI   Format = "2006-01-02 15:04:05.000" // 带毫秒的横杠格式: yyyy-MM-dd HH:mm:ss.SSS
)

// String 转为string
func (ztf Format) String() string {
	return string(ztf)
}

// FormatTimeToStr 将time转为Format指定格式的字符串
func (ztf Format) FormatTimeToStr(time time.Time) string {
	return time.Format(ztf.String())
}

// FormatTimeToSimpleMilliNoPoint 将time转为带毫秒的简单格式字符串，并去除点
func (ztf Format) FormatTimeToSimpleMilliNoPoint(time time.Time) (string, error) {
	if ztf != TIME_FORMAT_SIMPLE_MILLI {
		return "", fmt.Errorf("当前ZcTimeFormat不支持FormatTimeToSimpleMilliNoPoint方法: [%s]", ztf.String())
	}
	timeStr := time.Format(ztf.String())
	return strings.ReplaceAll(timeStr, ".", ""), nil
}

// ParseStrToTime 将Format指定格式的字符串转为time
func (ztf Format) ParseStrToTime(timeStr string) (time.Time, error) {
	return time.Parse(ztf.String(), timeStr)
}

// ParseSimpleMilliNoPointToTime 将带毫秒的简单格式字符串(已去除点)转为time
func (ztf Format) ParseSimpleMilliNoPointToTime(timeStr string) (time.Time, error) {
	if ztf != TIME_FORMAT_SIMPLE_MILLI {
		return time.Time{}, fmt.Errorf("当前ZcTimeFormat不支持ParseSimpleMilliNoPointToTime方法: [%s]", ztf.String())
	}
	if len(timeStr) != 17 {
		return time.Time{}, fmt.Errorf("试图转换的时间字符串长度不是17: %s", timeStr)
	}
	timeNewStr := timeStr[:14] + "." + timeStr[14:]
	return time.Parse(ztf.String(), timeNewStr)
}

// GetNowYMDHMS 获取当前时间戳字符串,格式为: yyyyMMddHHmmss
func GetNowYMDHMS() string {
	return TIME_FORMAT_SIMPLE.FormatTimeToStr(time.Now())
}
