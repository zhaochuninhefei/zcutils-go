package zcstr

import "strings"

func TrimAndUpper(str string) string {
	return strings.ToUpper(strings.TrimSpace(str))
}

func TrimAndLower(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}
