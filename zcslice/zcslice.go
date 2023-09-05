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
//  注意没有编辑sliceAll，而是返回了一个新的切片。
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

// TrimAndRmSpace 对sliceBefore的每个string做trimspace，然后去除空的元素
//  注意没有编辑sliceBefore，而是返回了一个新的切片。
func TrimAndRmSpace(sliceBefore []string) []string {
	var sliceAfter []string
	for _, item := range sliceBefore {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			sliceAfter = append(sliceAfter, trimmed)
		}
	}
	return sliceAfter
}

// ReverseBytes 反转字节切片
func ReverseBytes(input []byte) []byte {
	length := len(input)
	reversed := make([]byte, length)
	for i := 0; i < length; i++ {
		reversed[i] = input[length-1-i]
	}
	return reversed
}
