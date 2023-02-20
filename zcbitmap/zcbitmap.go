package zcbitmap

import (
	"fmt"
	"strconv"
)

// --------------------------------------------------------------------------

// BitSet8 8位的位图
type BitSet8 uint8

// SetBit 将位图中第i位(从右开始)设置为 1
//  i 一般取值范围: [1, 8], 超过该范围取对8的求余
func (b *BitSet8) SetBit(i uint) {
	mask := uint8(1) << ((i - 1) % 8)
	*b |= BitSet8(mask)
}

// ClearBit 将位图中第i位(从右开始)设置为 0
//  i 一般取值范围: [1, 8], 超过该范围取对8的求余
func (b *BitSet8) ClearBit(i uint) {
	mask := uint8(1) << ((i - 1) % 8)
	*b &= BitSet8(^mask)
}

// CheckBit 检查位图中第 i 位是否为 1
//  i 一般取值范围: [1, 8], 超过该范围取对8的求余
func (b *BitSet8) CheckBit(i uint) bool {
	mask := uint8(1) << ((i - 1) % 8)
	return (*b & BitSet8(mask)) != 0
}

// ToInt 将位图转换为一个整数
func (b *BitSet8) ToInt() uint8 {
	return uint8(*b)
}

// ConvBs8FromUInt8 将一个整数转换为位图
func ConvBs8FromUInt8(i uint8) BitSet8 {
	return BitSet8(i)
}

// ToBinaryStr 将位图转换为一个二进制字符串
func (b *BitSet8) ToBinaryStr(paddingZero bool) string {
	if paddingZero {
		return fmt.Sprintf("%08b", b.ToInt())
	}
	return strconv.FormatUint(uint64(*b), 2)
}

// ConvBs8FromBinaryStr 将一个二进制字符串转换为位图
func ConvBs8FromBinaryStr(s string) BitSet8 {
	if len(s) > 8 {
		s = s[len(s)-8:]
	}
	i, err := strconv.ParseUint(s, 2, 0)
	if err != nil {
		return 0
	}
	return BitSet8(i)
}

// MatchAll 检查当前位图是否完全匹配目标位图的所有非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果仍等于bs则表示满足MatchAll
func (b *BitSet8) MatchAll(bs BitSet8) bool {
	newBs := *b & bs
	return newBs == bs
}

// MatchAny 检查当前位图是否有任意一位匹配目标位图对应的非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果大于0则表示满足MatchAny
func (b *BitSet8) MatchAny(bs BitSet8) bool {
	newBs := *b & bs
	return newBs > 0
}

// --------------------------------------------------------------------------

// BitSet16 16位的位图
type BitSet16 uint16

// SetBit 将位图中第i位(从右开始)设置为 1
//  i 一般取值范围: [1, 16], 超过该范围取对16的求余
func (b *BitSet16) SetBit(i uint) {
	mask := uint16(1) << ((i - 1) % 16)
	*b |= BitSet16(mask)
}

// ClearBit 将位图中第i位(从右开始)设置为 0
//  i 一般取值范围: [1, 16], 超过该范围取对8的求余
func (b *BitSet16) ClearBit(i uint) {
	mask := uint16(1) << ((i - 1) % 16)
	*b &= BitSet16(^mask)
}

// CheckBit 检查位图中第 i 位是否为 1
//  i 一般取值范围: [1, 16], 超过该范围取对8的求余
func (b *BitSet16) CheckBit(i uint) bool {
	mask := uint16(1) << ((i - 1) % 16)
	return (*b & BitSet16(mask)) != 0
}

// ToInt 将位图转换为一个整数
func (b *BitSet16) ToInt() uint16 {
	return uint16(*b)
}

// ConvBs16FromUInt16 将一个整数转换为位图
func ConvBs16FromUInt16(i uint16) BitSet16 {
	return BitSet16(i)
}

// ToBinaryStr 将位图转换为一个二进制字符串
func (b *BitSet16) ToBinaryStr(paddingZero bool) string {
	if paddingZero {
		return fmt.Sprintf("%016b", b.ToInt())
	}
	return strconv.FormatUint(uint64(*b), 2)
}

// ConvBs16FromBinaryStr 将一个二进制字符串转换为位图
func ConvBs16FromBinaryStr(s string) BitSet16 {
	if len(s) > 16 {
		s = s[len(s)-16:]
	}
	i, err := strconv.ParseUint(s, 2, 0)
	if err != nil {
		return 0
	}
	return BitSet16(i)
}

// MatchAll 检查当前位图是否完全匹配目标位图的所有非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果仍等于bs则表示满足MatchAll
func (b *BitSet16) MatchAll(bs BitSet16) bool {
	newBs := *b & bs
	return newBs == bs
}

// MatchAny 检查当前位图是否有任意一位匹配目标位图对应的非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果大于0则表示满足MatchAny
func (b *BitSet16) MatchAny(bs BitSet16) bool {
	newBs := *b & bs
	return newBs > 0
}

// --------------------------------------------------------------------------

// BitSet32 32位的位图
type BitSet32 uint32

// SetBit 将位图中第i位(从右开始)设置为 1
//  i 一般取值范围: [1, 32], 超过该范围取对32的求余
func (b *BitSet32) SetBit(i uint) {
	mask := uint32(1) << ((i - 1) % 32)
	*b |= BitSet32(mask)
}

// ClearBit 将位图中第i位(从右开始)设置为 0
//  i 一般取值范围: [1, 32], 超过该范围取对32的求余
func (b *BitSet32) ClearBit(i uint) {
	mask := uint32(1) << ((i - 1) % 32)
	*b &= BitSet32(^mask)
}

// CheckBit 检查位图中第 i 位是否为 1
func (b *BitSet32) CheckBit(i uint) bool {
	mask := uint32(1) << ((i - 1) % 32)
	return (*b & BitSet32(mask)) != 0
}

// ToInt 将位图转换为一个整数
func (b *BitSet32) ToInt() int {
	return int(*b)
}

// ConvBs32FromUInt32 将一个整数转换为位图
func ConvBs32FromUInt32(i uint32) BitSet32 {
	return BitSet32(i)
}

// ToBinaryStr 将位图转换为一个二进制字符串
func (b *BitSet32) ToBinaryStr(paddingZero bool) string {
	if paddingZero {
		return fmt.Sprintf("%032b", b.ToInt())
	}
	return strconv.FormatUint(uint64(*b), 2)
}

// ConvBs32FromBinaryStr 将一个二进制字符串转换为位图
func ConvBs32FromBinaryStr(s string) BitSet32 {
	if len(s) > 32 {
		s = s[len(s)-32:]
	}
	i, err := strconv.ParseUint(s, 2, 0)
	if err != nil {
		return 0
	}
	return BitSet32(i)
}

// MatchAll 检查当前位图是否完全匹配目标位图的所有非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果仍等于bs则表示满足MatchAll
func (b *BitSet32) MatchAll(bs BitSet32) bool {
	newBs := *b & bs
	fmt.Printf("newBs: %s\n", newBs.ToBinaryStr(false))
	return newBs == bs
}

// MatchAny 检查当前位图是否有任意一位匹配目标位图对应的非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果大于0则表示满足MatchAny
func (b *BitSet32) MatchAny(bs BitSet32) bool {
	newBs := *b & bs
	fmt.Printf("newBs: %s\n", newBs.ToBinaryStr(false))
	return newBs > 0
}

// --------------------------------------------------------------------------

// BitSet64 64位的位图
type BitSet64 uint64

// SetBit 将位图中第i位(从右开始)设置为 1
//  i 一般取值范围: [1, 64], 超过该范围取对64的求余
func (b *BitSet64) SetBit(i uint) {
	mask := uint64(1) << ((i - 1) % 64)
	*b |= BitSet64(mask)
}

// ClearBit 将位图中第i位(从右开始)设置为 0
//  i 一般取值范围: [1, 64], 超过该范围取对64的求余
func (b *BitSet64) ClearBit(i uint) {
	mask := uint64(1) << ((i - 1) % 64)
	*b &= BitSet64(^mask)
}

// CheckBit 检查位图中第 i 位是否为 1
//  i 一般取值范围: [1, 64], 超过该范围取对64的求余
func (b *BitSet64) CheckBit(i uint) bool {
	mask := uint64(1) << ((i - 1) % 64)
	return (*b & BitSet64(mask)) != 0
}

// ToInt 将位图转换为一个整数
func (b *BitSet64) ToInt() int {
	return int(*b)
}

// ConvBs64FromUInt64 将一个整数转换为位图
func ConvBs64FromUInt64(i uint64) BitSet64 {
	return BitSet64(i)
}

// ToBinaryStr 将位图转换为一个二进制字符串
func (b *BitSet64) ToBinaryStr(paddingZero bool) string {
	if paddingZero {
		return fmt.Sprintf("%064b", b.ToInt())
	}
	return strconv.FormatUint(uint64(*b), 2)
}

// ConvBs64FromBinaryStr 将一个二进制字符串转换为位图
func ConvBs64FromBinaryStr(s string) BitSet64 {
	if len(s) > 64 {
		s = s[len(s)-64:]
	}
	i, err := strconv.ParseUint(s, 2, 0)
	if err != nil {
		return 0
	}
	return BitSet64(i)
}

// MatchAll 检查当前位图是否完全匹配目标位图的所有非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果仍等于bs则表示满足MatchAll
func (b *BitSet64) MatchAll(bs BitSet64) bool {
	newBs := *b & bs
	fmt.Printf("newBs: %s\n", newBs.ToBinaryStr(false))
	return newBs == bs
}

// MatchAny 检查当前位图是否有任意一位匹配目标位图对应的非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果大于0则表示满足MatchAny
func (b *BitSet64) MatchAny(bs BitSet64) bool {
	newBs := *b & bs
	fmt.Printf("newBs: %s\n", newBs.ToBinaryStr(false))
	return newBs > 0
}

// --------------------------------------------------------------------------
