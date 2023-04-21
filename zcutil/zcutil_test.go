package zcutil

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"reflect"
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
