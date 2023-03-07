package main

import (
	"fmt"
	"gitee.com/zhaochuninhefei/zcutils-go/zcargs"
	"os"
)

// main zcargs测试用主函数, 在当前目录下执行以下命令，并检查输出是否正确。
//  go run zcargs_main.go -test_one xxx --test2 yYy --test_3 it
//  go run zcargs_main.go -test_one xxx --test2 "yYy" --test_3
//  go run zcargs_main.go -test_one xxx --TEST2 yYy --test_3 IT
//  go run zcargs_main.go -test_one xxx --test2=yyy --test_3
//  go run zcargs_main.go -test_one xxx --test2=yYy --test_3 it --Test2 'YYY'
//  go run zcargs_main.go -test_one xxx --test2 yYy --test_3 --Test2="YYyy"
//  go run zcargs_main.go -test_one xxx --test2 yYy --test_3 IT --Test2
func main() {
	fmt.Println(os.Args)

	test2 := zcargs.TakeArgFromOsArgs("test2", true, "n")
	fmt.Printf("test2: %s\n", test2)

	test3 := zcargs.TakeArgFromOsArgs("TEST_3", false, "it")
	fmt.Printf("test3: %s\n", test3)

	test4 := zcargs.TakeArgFromOsArgs("TEST_4", false, "444")
	fmt.Printf("test4: %s\n", test4)

	fmt.Println(os.Args)
}
