package zcbitmap

import (
	"fmt"
	"testing"
)

func TestBitSet32(t *testing.T) {
	var b BitSet32
	b.SetBit(3)
	b.SetBit(7)
	fmt.Printf("%032b\n", b.ToInt())
	fmt.Println(b.ToBinaryStr(true))
	fmt.Println(b.ToBinaryStr(false))

	b.ClearBit(3)
	fmt.Printf("%032b\n", b.ToInt())

	fmt.Println(b.TestBit(3))
	fmt.Println(b.TestBit(7))

	var i int = 255
	b = FromInt(i)
	fmt.Printf("%032b\n", b.ToInt())
}
