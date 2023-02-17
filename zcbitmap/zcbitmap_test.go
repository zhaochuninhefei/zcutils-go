package zcbitmap

import (
	"fmt"
	"testing"
)

func TestBitSet32(t *testing.T) {
	fmt.Println("测试 1000100 :")
	var b BitSet32
	b.SetBit(3)
	b.SetBit(7)
	fmt.Printf("%032b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))

	fmt.Println("将第3位设置为0:")
	b.ClearBit(3)
	fmt.Printf("%032b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))

	fmt.Printf("第3位是否为1: %t\n", b.CheckBit(3))
	fmt.Printf("第7位是否为1: %t\n", b.CheckBit(7))

	fmt.Println("将 255 转为 BitSet32")
	var i = 255
	b = FromInt(i)
	fmt.Printf("%032b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))

	fmt.Println("将 10101100 转为 BitSet32")
	s := "10101100"
	b, _ = FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Printf("%032b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))

}

func TestMatchBitSet32(t *testing.T) {
	status, _ := FromBinaryStr("101")

	fmt.Println("----- 测试 111")
	s := "111"
	b, _ := FromBinaryStr(s)
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
	b, _ = FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 111: %t\n", allMatch)
	if !allMatch {
		t.Fatal("101 MatchAll 111 不应该返回 false")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 111: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 111 不应该返回 false")
	}

	fmt.Println("----- 测试 010")
	s = "010"
	b, _ = FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 111: %t\n", allMatch)
	if allMatch {
		t.Fatal("101 MatchAll 111 不应该返回 true")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 111: %t\n", anyMatch)
	if anyMatch {
		t.Fatal("101 MatchAny 111 不应该返回 true")
	}

	fmt.Println("----- 测试 001")
	s = "001"
	b, _ = FromBinaryStr(s)
	fmt.Printf("对应int值: %d\n", b)
	fmt.Println(b.ToBinaryStr(false))

	allMatch = status.MatchAll(b)
	fmt.Printf("101 MatchAll 111: %t\n", allMatch)
	if !allMatch {
		t.Fatal("101 MatchAll 111 不应该返回 false")
	}
	anyMatch = status.MatchAny(b)
	fmt.Printf("101 MatchAny 111: %t\n", anyMatch)
	if !anyMatch {
		t.Fatal("101 MatchAny 111 不应该返回 false")
	}
}
