package zcstr

import "strings"

// TrimAndUpper 去除字符串首尾空格并转为大写
func TrimAndUpper(str string) string {
	return strings.ToUpper(strings.TrimSpace(str))
}

// TrimAndLower 去除字符串首尾空格并转为小写
func TrimAndLower(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}
