package zcslice

// Contains 判断一个字符串是否在一个字符串切片中
func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

// Diff 计算两个字符串切片的差集
func Diff(slice1, slice2 []string) []string {
	m := make(map[string]bool)

	for _, item := range slice2 {
		m[item] = true
	}

	var diff []string
	for _, item := range slice1 {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}

	return diff
}
