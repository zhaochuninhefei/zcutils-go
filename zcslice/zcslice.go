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
