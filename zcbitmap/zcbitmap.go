package zcbitmap

import (
	"fmt"
	"strconv"
)

// BitSet32 32位的位图
type BitSet32 uint32

// SetBit 将位图中第i位(从右开始)设置为 1
//  i 一般取值范围: [1, 32], 超过该范围取对32的求余
func (b BitSet32) SetBit(i uint) {
	mask := uint32(1) << ((i - 1) % 32)
	b |= BitSet32(mask)
}

// ClearBit 将位图中第i位(从右开始)设置为 0
//  i 一般取值范围: [1, 32], 超过该范围取对32的求余
func (b BitSet32) ClearBit(i uint) {
	mask := uint32(1) << ((i - 1) % 32)
	b &= BitSet32(^mask)
}

// CheckBit 检查位图中第 i 位是否为 1
func (b BitSet32) CheckBit(i uint) bool {
	mask := uint32(1) << ((i - 1) % 32)
	return (b & BitSet32(mask)) != 0
}

// ToInt 将位图转换为一个整数
func (b BitSet32) ToInt() int {
	return int(b)
}

// FromInt 将一个整数转换为位图
func FromInt(i int) BitSet32 {
	return BitSet32(i)
}

// ToBinaryStr 将位图转换为一个二进制字符串
func (b BitSet32) ToBinaryStr(paddingZero bool) string {
	if paddingZero {
		return fmt.Sprintf("%032b", b.ToInt())
	}
	return strconv.FormatUint(uint64(b), 2)
}

// FromBinaryStr 将一个二进制字符串转换为位图
func FromBinaryStr(s string) (BitSet32, error) {
	if len(s) > 32 {
		return 0, fmt.Errorf("二进制字符串长度(%d)超过了32", len(s))
	}
	i, err := strconv.ParseUint(s, 2, 0)
	if err != nil {
		return 0, err
	}
	return BitSet32(i), nil
}

// MatchAll 检查当前位图是否完全匹配目标位图的所有非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果仍等于bs则表示满足MatchAll
func (b BitSet32) MatchAll(bs BitSet32) bool {
	newBs := b & bs
	fmt.Printf("newBs: %s\n", newBs.ToBinaryStr(false))
	return newBs == bs
}

// MatchAny 检查当前位图是否有任意一位匹配目标位图对应的非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果大于0则表示满足MatchAny
func (b BitSet32) MatchAny(bs BitSet32) bool {
	newBs := b & bs
	fmt.Printf("newBs: %s\n", newBs.ToBinaryStr(false))
	return newBs > 0
}
