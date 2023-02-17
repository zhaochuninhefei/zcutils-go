package zcbitmap

import (
	"fmt"
	"testing"
)

// --------------------------------------------------------

func TestBitSet8(t *testing.T) {
	fmt.Println("测试 1000100 :")
	var b BitSet8
	b.SetBit(3)
	b.SetBit(7)
	fmt.Printf("%08b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	bStrAct := b.ToBinaryStr(false)
	fmt.Println(bStrAct)
	if bStrAct != "1000100" {
		t.Fatal("测试 1000100 失败")
	}

	fmt.Println("将第3位设置为0:")
	b.ClearBit(3)
	fmt.Printf("%08b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))

	fmt.Printf("第3位是否为1: %t\n", b.CheckBit(3))
	if b.CheckBit(3) {
		t.Fatal("将第3位设置为0 失败")
	}
	fmt.Printf("第7位是否为1: %t\n", b.CheckBit(7))
	if !b.CheckBit(7) {
		t.Fatal("将第7位设置为1 失败")
	}

	fmt.Println("将 255 转为 BitSet8")
	var i uint8 = 255
	b = ConvBs8FromUInt8(i)
	fmt.Printf("%08b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))
	if "11111111" != b.ToBinaryStr(false) {
		t.Fatal("将 255 转为 BitSet8 失败")
	}

	fmt.Println("将 10101100 转为 BitSet8")
	s := "10101100"
	b = ConvBs8FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Printf("%08b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))
	if "10101100" != b.ToBinaryStr(false) {
		t.Fatal("将 255 转为 BitSet8 失败")
	}
}

func TestMatchBitSet8(t *testing.T) {
	status := ConvBs8FromBinaryStr("101")

	fmt.Println("----- 测试 111")
	s := "111"
	b := ConvBs8FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch := status.MatchAll(b)
	fmt.Printf("101 MatchAll 111: %t\n", allMatch)
	if allMatch {
		t.Fatal("101 MatchAll 111 不应该返回 true")
	}
	anyMatch := status.MatchAny(b)
	fmt.Printf("101 MatchAny 111: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 111 不应该返回 false")
	}

	fmt.Println("----- 测试 100")
	s = "100"
	b = ConvBs8FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 100: %t\n", allMatch)
	if !allMatch {
		t.Fatal("101 MatchAll 100 不应该返回 false")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 100: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 100 不应该返回 false")
	}

	fmt.Println("----- 测试 010")
	s = "010"
	b = ConvBs8FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 010: %t\n", allMatch)
	if allMatch {
		t.Fatal("101 MatchAll 010 不应该返回 true")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 010: %t\n", anyMatch)
	if anyMatch {
		t.Fatal("101 MatchAny 010 不应该返回 true")
	}

	fmt.Println("----- 测试 001")
	s = "001"
	b = ConvBs8FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 001: %t\n", allMatch)
	if !allMatch {
		t.Fatal("101 MatchAll 001 不应该返回 false")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 001: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 001 不应该返回 false")
	}
}

// --------------------------------------------------------

func TestBitSet16(t *testing.T) {
	fmt.Println("测试 1000100 :")
	var b BitSet16
	b.SetBit(3)
	b.SetBit(7)
	fmt.Printf("%016b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	bStrAct := b.ToBinaryStr(false)
	fmt.Println(bStrAct)
	if bStrAct != "1000100" {
		t.Fatal("测试 1000100 失败")
	}

	fmt.Println("将第3位设置为0:")
	b.ClearBit(3)
	fmt.Printf("%016b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))

	fmt.Printf("第3位是否为1: %t\n", b.CheckBit(3))
	if b.CheckBit(3) {
		t.Fatal("将第3位设置为0 失败")
	}
	fmt.Printf("第7位是否为1: %t\n", b.CheckBit(7))
	if !b.CheckBit(7) {
		t.Fatal("将第7位设置为1 失败")
	}

	fmt.Println("将 255 转为 BitSet16")
	var i uint16 = 255
	b = ConvBs16FromUInt16(i)
	fmt.Printf("%016b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))
	if "11111111" != b.ToBinaryStr(false) {
		t.Fatal("将 255 转为 BitSet16 失败")
	}

	fmt.Println("将 10101100 转为 BitSet16")
	s := "10101100"
	b = ConvBs16FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Printf("%016b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))
	if "10101100" != b.ToBinaryStr(false) {
		t.Fatal("将 255 转为 BitSet16 失败")
	}
}

func TestMatchBitSet16(t *testing.T) {
	status := ConvBs16FromBinaryStr("101")

	fmt.Println("----- 测试 111")
	s := "111"
	b := ConvBs16FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch := status.MatchAll(b)
	fmt.Printf("101 MatchAll 111: %t\n", allMatch)
	if allMatch {
		t.Fatal("101 MatchAll 111 不应该返回 true")
	}
	anyMatch := status.MatchAny(b)
	fmt.Printf("101 MatchAny 111: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 111 不应该返回 false")
	}

	fmt.Println("----- 测试 100")
	s = "100"
	b = ConvBs16FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 100: %t\n", allMatch)
	if !allMatch {
		t.Fatal("101 MatchAll 100 不应该返回 false")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 100: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 100 不应该返回 false")
	}

	fmt.Println("----- 测试 010")
	s = "010"
	b = ConvBs16FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 010: %t\n", allMatch)
	if allMatch {
		t.Fatal("101 MatchAll 010 不应该返回 true")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 010: %t\n", anyMatch)
	if anyMatch {
		t.Fatal("101 MatchAny 010 不应该返回 true")
	}

	fmt.Println("----- 测试 001")
	s = "001"
	b = ConvBs16FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 001: %t\n", allMatch)
	if !allMatch {
		t.Fatal("101 MatchAll 001 不应该返回 false")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 001: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 001 不应该返回 false")
	}
}

// --------------------------------------------------------

func TestBitSet32(t *testing.T) {
	fmt.Println("测试 1000100 :")
	var b BitSet32
	b.SetBit(3)
	b.SetBit(7)
	fmt.Printf("%032b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	bStrAct := b.ToBinaryStr(false)
	fmt.Println(bStrAct)
	if bStrAct != "1000100" {
		t.Fatal("测试 1000100 失败")
	}

	fmt.Println("将第3位设置为0:")
	b.ClearBit(3)
	fmt.Printf("%032b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))

	fmt.Printf("第3位是否为1: %t\n", b.CheckBit(3))
	if b.CheckBit(3) {
		t.Fatal("将第3位设置为0 失败")
	}
	fmt.Printf("第7位是否为1: %t\n", b.CheckBit(7))
	if !b.CheckBit(7) {
		t.Fatal("将第7位设置为1 失败")
	}

	fmt.Println("将 255 转为 BitSet32")
	var i uint32 = 255
	b = ConvBs32FromUInt32(i)
	fmt.Printf("%032b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))
	if "11111111" != b.ToBinaryStr(false) {
		t.Fatal("将 255 转为 BitSet32 失败")
	}

	fmt.Println("将 10101100 转为 BitSet32")
	s := "10101100"
	b = ConvBs32FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Printf("%032b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))
	if "10101100" != b.ToBinaryStr(false) {
		t.Fatal("将 255 转为 BitSet32 失败")
	}
}

func TestMatchBitSet32(t *testing.T) {
	status := ConvBs32FromBinaryStr("101")

	fmt.Println("----- 测试 111")
	s := "111"
	b := ConvBs32FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch := status.MatchAll(b)
	fmt.Printf("101 MatchAll 111: %t\n", allMatch)
	if allMatch {
		t.Fatal("101 MatchAll 111 不应该返回 true")
	}
	anyMatch := status.MatchAny(b)
	fmt.Printf("101 MatchAny 111: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 111 不应该返回 false")
	}

	fmt.Println("----- 测试 100")
	s = "100"
	b = ConvBs32FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 100: %t\n", allMatch)
	if !allMatch {
		t.Fatal("101 MatchAll 100 不应该返回 false")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 100: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 100 不应该返回 false")
	}

	fmt.Println("----- 测试 010")
	s = "010"
	b = ConvBs32FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 010: %t\n", allMatch)
	if allMatch {
		t.Fatal("101 MatchAll 010 不应该返回 true")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 010: %t\n", anyMatch)
	if anyMatch {
		t.Fatal("101 MatchAny 010 不应该返回 true")
	}

	fmt.Println("----- 测试 001")
	s = "001"
	b = ConvBs32FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 001: %t\n", allMatch)
	if !allMatch {
		t.Fatal("101 MatchAll 001 不应该返回 false")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 001: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 001 不应该返回 false")
	}
}

// --------------------------------------------------------

func TestBitSet64(t *testing.T) {
	fmt.Println("测试 1000100 :")
	var b BitSet64
	b.SetBit(3)
	b.SetBit(7)
	fmt.Printf("%064b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	bStrAct := b.ToBinaryStr(false)
	fmt.Println(bStrAct)
	if bStrAct != "1000100" {
		t.Fatal("测试 1000100 失败")
	}

	fmt.Println("将第3位设置为0:")
	b.ClearBit(3)
	fmt.Printf("%064b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))

	fmt.Printf("第3位是否为1: %t\n", b.CheckBit(3))
	if b.CheckBit(3) {
		t.Fatal("将第3位设置为0 失败")
	}
	fmt.Printf("第7位是否为1: %t\n", b.CheckBit(7))
	if !b.CheckBit(7) {
		t.Fatal("将第7位设置为1 失败")
	}

	fmt.Println("将 255 转为 BitSet64")
	var i uint64 = 255
	b = ConvBs64FromUInt64(i)
	fmt.Printf("%064b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))
	if "11111111" != b.ToBinaryStr(false) {
		t.Fatal("将 255 转为 BitSet64 失败")
	}

	fmt.Println("将 10101100 转为 BitSet64")
	s := "10101100"
	b = ConvBs64FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Printf("%064b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))
	if "10101100" != b.ToBinaryStr(false) {
		t.Fatal("将 255 转为 BitSet64 失败")
	}
}

func TestMatchBitSet64(t *testing.T) {
	status := ConvBs64FromBinaryStr("101")

	fmt.Println("----- 测试 111")
	s := "111"
	b := ConvBs64FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch := status.MatchAll(b)
	fmt.Printf("101 MatchAll 111: %t\n", allMatch)
	if allMatch {
		t.Fatal("101 MatchAll 111 不应该返回 true")
	}
	anyMatch := status.MatchAny(b)
	fmt.Printf("101 MatchAny 111: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 111 不应该返回 false")
	}

	fmt.Println("----- 测试 100")
	s = "100"
	b = ConvBs64FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 100: %t\n", allMatch)
	if !allMatch {
		t.Fatal("101 MatchAll 100 不应该返回 false")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 100: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 100 不应该返回 false")
	}

	fmt.Println("----- 测试 010")
	s = "010"
	b = ConvBs64FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 010: %t\n", allMatch)
	if allMatch {
		t.Fatal("101 MatchAll 010 不应该返回 true")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 010: %t\n", anyMatch)
	if anyMatch {
		t.Fatal("101 MatchAny 010 不应该返回 true")
	}

	fmt.Println("----- 测试 001")
	s = "001"
	b = ConvBs64FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 001: %t\n", allMatch)
	if !allMatch {
		t.Fatal("101 MatchAll 001 不应该返回 false")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 001: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 001 不应该返回 false")
	}
}

// --------------------------------------------------------------------------
