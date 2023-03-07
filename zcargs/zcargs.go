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
//  @param defaultVal : 参数默认值, 参数未设置时返回该值。
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
func TakeArgFromOsArgs(argName string, remove bool, defaultVal string) (val string) {
	// 上同步锁
	LockOsArgs.Lock()
	defer LockOsArgs.Unlock()
	// 检查命令行有没有参数, os.Args第一个参数是当前执行程序全路径
	if len(os.Args) < 2 {
		// 命令行没有传入任何参数，直接返回默认值
		val = defaultVal
		return
	}
	// 检查参数名是否非空
	if argName == "" || strings.TrimSpace(argName) == "" {
		panic("can not pass empty as argName")
	}
	// 获取小写参数名
	argNameLow := zcstr.TrimAndLower(argName)
	// 遍历当前os.Args, 注意从第二个参数开始遍历
	for i := 1; i < len(os.Args); i++ {
		// 当前参数
		argCur := os.Args[i]
		// 不是"-"开头的话，直接跳过
		if !strings.HasPrefix(argCur, "-") {
			continue
		}

		// 判断当前参数是否是os.Args中最后一个参数
		lastArg := i == len(os.Args)-1
		argNext := ""
		if !lastArg {
			// 当前参数不是最后一个时，获取下一个参数
			argNext = os.Args[i+1]
		}

		// 去除当前参数开头的"-"或"--"
		var argTmp string
		if strings.HasPrefix(argCur, "--") {
			argTmp = argCur[2:]
		} else {
			argTmp = argCur[1:]
		}
		// 获取小写的当前参数(已去处"-"或"--")
		argTmpLow := strings.ToLower(argTmp)

		// 判断当前参数类型
		var argType int
		// 当前参数完全匹配
		if argTmpLow == argNameLow {
			if lastArg {
				// 如果当前参数是最后一个参数，那么当前参数就是标识类型
				argType = ArgType1
			} else {
				if strings.HasPrefix(argNext, "-") {
					// 如果下一个参数也是"-"开头，则当前参数是标识类型
					argType = ArgType1
				} else {
					// 如果下一个参数不是"-"开头，则当前参数是键值对类型，下一个参数就是当前参数的值
					argType = ArgType2
				}
			}
		} else if strings.HasPrefix(argTmpLow, argNameLow+"=") {
			// 如果匹配"<参数名>="，则当前参数是带等号的键值对类型，等号后面的内容就是当前参数的值
			argType = ArgType3
		} else {
			// 未能匹配，继续遍历
			continue
		}
		// 设置返回值，以及当前参数的起始与结束索引
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
	// val是零值时，返回传入的默认值
	if val == "" {
		val = defaultVal
	}
	return
}
