#!/bin/bash
# 执行各个包的测试用例

set -e

# protoreflect
echo "protoreflect 测试用例"
cd protobuf/protoreflect
go test
cd ../../

# 等待控制台输入任意字符继续
echo
read -rp "protoreflect 测试用例 结束，按下任意按键继续..." -n 1
echo

# protobuffer
echo "protobuffer 测试用例"
cd protobuf/protobuffer
go test
cd ../../

# 等待控制台输入任意字符继续
echo
read -rp "protobuffer 测试用例 结束，按下任意按键继续..." -n 1
echo

# zcargs
echo "zcargs 测试用例"
cd zcargs/test
go run zcargs_main.go
cd ../../

# 等待控制台输入任意字符继续
echo
read -rp "zcargs 测试用例 结束，按下任意按键继续..." -n 1
echo

# zcbitmap
echo "zcbitmap 测试用例"
cd zcbitmap
go test
cd ../

# 等待控制台输入任意字符继续
echo
read -rp "zcbitmap 测试用例 结束，按下任意按键继续..." -n 1
echo

# zccompress
echo "zccompress 测试用例"
cd zccompress
go test
cd ../

# 等待控制台输入任意字符继续
echo
read -rp "zccompress 测试用例 结束，按下任意按键继续..." -n 1
echo

# zcpath
echo "zcpath 测试用例"
cd zcpath
go test
cd ../

# 等待控制台输入任意字符继续
echo
read -rp "zcpath 测试用例 结束，按下任意按键继续..." -n 1
echo

# zcrandom
echo "zcrandom 测试用例"
cd zcrandom
go test
cd ../

# 等待控制台输入任意字符继续
echo
read -rp "zcrandom 测试用例 结束，按下任意按键继续..." -n 1
echo

# zcslice
echo "zcslice 测试用例"
cd zcslice
go test
cd ../

# 等待控制台输入任意字符继续
echo
read -rp "zcslice 测试用例 结束，按下任意按键继续..." -n 1
echo

# zcssh
echo "zcssh 测试用例"
cd zcssh
go test
cd ../

# 等待控制台输入任意字符继续
echo
read -rp "zcssh 测试用例 结束，按下任意按键继续..." -n 1
echo

# zcstr
echo "zcstr 测试用例"
cd zcstr
go test
cd ../

# 等待控制台输入任意字符继续
echo
read -rp "zcstr 测试用例 结束，按下任意按键继续..." -n 1
echo

# zcsync
echo "zcsync 测试用例"
cd zcsync
go test
cd ../

# 等待控制台输入任意字符继续
echo
read -rp "zcsync 测试用例 结束，按下任意按键继续..." -n 1
echo

# zctime
echo "zctime 测试用例"
cd zctime
go test
cd ../

# 等待控制台输入任意字符继续
echo
read -rp "zctime 测试用例 结束，按下任意按键继续..." -n 1
echo

# zctoken
echo "zctoken 测试用例"
cd zctoken
go test
cd ../

# 等待控制台输入任意字符继续
echo
read -rp "zctoken 测试用例 结束，按下任意按键继续..." -n 1
echo

# zcutil
echo "zcutil 测试用例"
cd zcutil
go test
cd ../

# 等待控制台输入任意字符继续
echo
read -rp "zcutil 测试用例 结束，按下任意按键继续..." -n 1
echo


# zcwaiter
echo "zcwaiter 测试用例"
cd zcwaiter
go test
cd ../
echo "zcwaiter 测试用例 结束"

echo "全部测试用例 结束"
echo