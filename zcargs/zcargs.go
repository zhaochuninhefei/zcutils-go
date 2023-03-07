package zcargs

import (
	"gitee.com/zhaochuninhefei/zcutils-go/zcstr"
	"os"
	"strconv"
	"strings"
	"sync"
)

//goland:noinspection GoUnusedConst
const (
	argTypeUnknown = iota // 参数格式未知
	ArgType1              // 参数格式 `--test`或`-test`
	ArgType2              // 参数格式 `--test xxx`或`-test xxx`
	ArgType3              // 参数格式 `--test=xxx`或`-test=xxx`
)

var LockOsArgs sync.Mutex

// TakeArgFromOsArgs 从os.Args中获取指定参数
//  @param argName : 参数名, 忽略大小写, 不可传零值。
//  @param remove : 是否从os.Args中移除该参数。
//  @return val 指定参数值。
//
// 参数类型支持以下三种:
//  - 参数格式 `--test`或`-test` : 标识类型参数，如果设置了，就返回"y"，否则返回""。
//  - 参数格式 `--test xxx`或`-test xxx` : 键值对类型参数，如果设置了，就返回对应的设值。
//  - 参数格式 `--test=xxx`或`-test=xxx` : 带等号的键值对类型参数，如果设置了，就返回"="设值。
//
// 注意:
//  - 如果参数重复设置，后面的重复参数会覆盖前面的相同参数。
//  - 如果参数重复设置且需要移除，则会移除全部的该参数。
//  - 该函数会使用`zcargs.LockOsArgs`上同步锁, 建议程序中其他读写`os.Args`的goroutine也使用该锁。
func TakeArgFromOsArgs(argName string, remove bool) (val string) {
	LockOsArgs.Lock()
	defer LockOsArgs.Unlock()

	if len(os.Args) < 2 {
		return
	}
	if argName == "" || strings.TrimSpace(argName) == "" {
		panic("can not pass empty as argName")
	}
	argNameLow := zcstr.TrimAndLower(argName)

	for i := 1; i < len(os.Args); i++ {
		argCur := os.Args[i]
		if !strings.HasPrefix(argCur, "-") {
			continue
		}
		lastArg := i == len(os.Args)-1
		argNext := ""
		if !lastArg {
			argNext = os.Args[i+1]
		}

		if strings.HasPrefix(argCur, "--") || strings.HasPrefix(argCur, "-") {
			var argTmp string
			if strings.HasPrefix(argCur, "--") {
				argTmp = argCur[2:]
			} else {
				argTmp = argCur[1:]
			}

			argTmpLow := strings.ToLower(argTmp)

			var argType int
			if argTmpLow == argNameLow {
				if lastArg {
					argType = ArgType1
				} else {
					if strings.HasPrefix(argNext, "-") {
						argType = ArgType1
					} else {
						argType = ArgType2
					}
				}
			} else if strings.HasPrefix(argTmpLow, argNameLow+"=") {
				argType = ArgType3
			} else {
				continue
			}
			var startIndex, stopIndex int
			switch argType {
			case ArgType1:
				// 该参数为标识类参数，返回"y"即可
				val = "y"
				// 该参数的起始与结束索引
				startIndex = i
				stopIndex = i + 1
			case ArgType2:
				// 该参数为键值对类型，返回下一个arg即可
				val = argNext
				// 该参数的起始与结束索引
				startIndex = i
				stopIndex = i + 2
			case ArgType3:
				// test=xxx
				val = argTmp[len(argNameLow)+1:]
				// 该参数的起始与结束索引
				startIndex = i
				stopIndex = i + 1
			default:
				panic("代码逻辑有误,参数类型识别失败:" + strconv.Itoa(argType))
			}
			if remove {
				// 从os.Args中去除该参数
				os.Args = append(os.Args[:startIndex], os.Args[stopIndex:]...)
				i--
			}
		}
	}
	return
}
