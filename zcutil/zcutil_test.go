package zcutil

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestInt32ToBytes(t *testing.T) {
	b := Int32ToBytes(41)
	fmt.Println(hex.EncodeToString(b))
}

func TestBytesToInt32(t *testing.T) {
	var num int32 = 41
	b := Int32ToBytes(num)
	fmt.Println(hex.EncodeToString(b))

	numNew := BytesToInt32(b)
	fmt.Println(numNew)

	if num != numNew {
		t.Fatal("BytesToInt32转换后不相等")
	}
}

func TestTempDir(t *testing.T) {

	tmpKeyStore, err := ioutil.TempDir("testdata", "msp-keystore")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tmpKeyStore)
}

func TestPrintStack(t *testing.T) {
	testPrintStack03()
}

func testPrintStack03() {
	testPrintStack02()
}

func testPrintStack02() {
	testPrintStack01()
}

func testPrintStack01() {
	PrintStack("")
}

func TestIf(t *testing.T) {
	// 定义一个map
	m := make(map[string]string)
	// 往map中添加键值对, 有 name 和 age, 没有 level
	m["name"] = "张三"
	m["age"] = "20"
	type args struct {
		condition bool
		trueVal   interface{}
		falseVal  interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "test1",
			args: args{
				condition: m["name"] == "张三",
				trueVal:   m["name"],
				falseVal:  "无名",
			},
			want: "张三",
		},
		{
			name: "test2",
			args: args{
				condition: m["age"] != "",
				trueVal:   m["age"],
				falseVal:  "18",
			},
			want: "20",
		},
		{
			name: "test3",
			args: args{
				condition: m["level"] != "",
				trueVal:   m["level"],
				falseVal:  "0",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.condition, tt.args.trueVal, tt.args.falseVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfByFunc(t *testing.T) {
	val, err := IfByFunc(func() bool {
		return true
	}, func() (interface{}, error) {
		return 1, nil
	}, func() (interface{}, error) {
		return 0, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(val)

	fmt.Println(IfByFuncNoErr(func() bool {
		return false
	}, func() interface{} {
		return 1
	}, func() interface{} {
		return 0
	}))
}

func TestCallAsyncFuncAndWaitByLog(t *testing.T) {
	// 日志文件路径
	logFilePath := "testdata/test.log"

	// 测试正常结束
	err := CallAsyncFuncAndWaitByLog(logFilePath, func() error {
		fmt.Println("aysnc func start...")
		// goruntine 执行日志写入
		go func() {
			for i := 0; i < 10; i++ {
				// 向logFilePath写入日志
				_ = writeLog(logFilePath, fmt.Sprintf("test log %d", i))
			}
			_ = writeLog(logFilePath, "success")
		}()
		return nil
	}, func(line string) (bool, error) {
		fmt.Printf("读取到日志line: %s\n", line)
		if line == "success" {
			return true, nil
		}
		if strings.HasPrefix(line, "[err]") {
			return true, errors.New(line)
		}
		return false, nil
	}, 20)
	if err != nil {
		t.Fatal(err)
	}

	// 测试错误结束
	err = CallAsyncFuncAndWaitByLog(logFilePath, func() error {
		fmt.Println("aysnc func start...")
		// goruntine 执行日志写入
		go func() {
			for i := 0; i < 10; i++ {
				// 向logFilePath写入日志
				_ = writeLog(logFilePath, fmt.Sprintf("test log %d", i))
			}
			_ = writeLog(logFilePath, "[err]发生错误")
		}()
		return nil
	}, func(line string) (bool, error) {
		fmt.Printf("读取到日志line: %s\n", line)
		if line == "success" {
			return true, nil
		}
		if strings.HasPrefix(line, "[err]") {
			return true, errors.New(line)
		}
		return false, nil
	}, 20)
	if err == nil {
		t.Fatal("未能返回错误")
	} else {
		if err.Error() == "[funcHandlerLogLine error]处理日志行返回错误: [err]发生错误" {
			fmt.Printf("成功获取错误消息: %s", err.Error())
		} else {
			t.Fatal("错误消息与预期不符: " + err.Error())
		}
	}

	// 测试超时
	err = CallAsyncFuncAndWaitByLog(logFilePath, func() error {
		fmt.Println("aysnc func start...")
		// goruntine 执行日志写入
		go func() {
			for i := 0; i < 10; i++ {
				// 向logFilePath写入日志
				_ = writeLog(logFilePath, fmt.Sprintf("test log %d", i))
			}
		}()
		return nil
	}, func(line string) (bool, error) {
		fmt.Printf("读取到日志line: %s\n", line)
		if line == "success" {
			return true, nil
		}
		if strings.HasPrefix(line, "[err]") {
			return true, errors.New(line)
		}
		return false, nil
	}, 3)
	if err == nil {
		t.Fatal("未能返回错误")
	} else {
		if strings.HasPrefix(err.Error(), "[tail timeout]") {
			fmt.Printf("成功获取超时错误消息: %s\n", err.Error())
		} else {
			t.Fatal("错误消息与预期不符: " + err.Error())
		}
	}
}

func TestCallAsyncFuncAndWaitByFlag(t *testing.T) {
	// 标志文件与日志文件路径
	flagFilePath := "testdata/test.over"
	logFilePath := "testdata/test.log"

	// 测试正常结束
	lines, err := CallAsyncFuncAndWaitByFlag(flagFilePath, logFilePath, func() error {
		fmt.Println("aysnc func start...")
		// goruntine 执行日志写入
		go func() {
			for i := 0; i < 10; i++ {
				// 向logFilePath写入日志
				_ = writeLog(logFilePath, fmt.Sprintf("test log %d", i))
			}
			_ = writeLog(flagFilePath, "over")
		}()
		return nil
	}, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 10 {
		t.Fatal("未能读取到所有日志")
	}

	// 测试超时
	lines, err = CallAsyncFuncAndWaitByFlag(flagFilePath, logFilePath, func() error {
		fmt.Println("aysnc func start...")
		// goruntine 执行日志写入
		go func() {
			for i := 0; i < 10; i++ {
				// 向logFilePath写入日志
				_ = writeLog(logFilePath, fmt.Sprintf("test log %d", i))
			}
		}()
		return nil
	}, 3)
	if err == nil {
		t.Fatal("未能返回错误")
	} else {
		if strings.HasPrefix(err.Error(), "[watcher timeout]") {
			fmt.Printf("成功获取超时错误消息: %s\n", err.Error())
		} else {
			t.Fatal("错误消息与预期不符: " + err.Error())
		}
	}
}

// writeLog 向日志文件logPath追加写入line并换行,如果日志文件不存在就创建新文件
func writeLog(logPath string, line string) error {
	// 打开日志文件
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	// 写入日志行并换行
	if _, err = file.WriteString(line + "\n"); err != nil {
		return err
	}
	return nil

}
