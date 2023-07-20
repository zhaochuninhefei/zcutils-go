package zcslice

import "strings"

// Contains 判断一个字符串是否在一个字符串切片中
func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

// Subtract 从sliceAll中移除sliceToSub中的元素
func Subtract(sliceAll, sliceToSub []string) []string {
	m := make(map[string]bool)
	for _, item := range sliceToSub {
		m[item] = true
	}
	var rest []string
	for _, item := range sliceAll {
		if _, ok := m[item]; !ok {
			rest = append(rest, item)
		}
	}
	return rest
}

func TrimAndRmSpace(slice1 []string) []string {
	var slice2 []string
	for _, item := range slice1 {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			slice2 = append(slice2, trimmed)
		}
	}
	return slice2
}
