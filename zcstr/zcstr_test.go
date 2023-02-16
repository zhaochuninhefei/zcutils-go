package zcstr

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimAndUpper(t *testing.T) {
	str1 := " asdf_天行健_fdsa "
	str2 := "	asdf_天行健_fdsa  "
	str3 := "\nasdf_天行健_fdsa\f"
	exp := "ASDF_天行健_FDSA"

	fmt.Printf("str1:[%s]\n", str1)
	rel1 := TrimAndUpper(str1)
	fmt.Printf("rel1:[%s]\n", rel1)
	assert.Equal(t, exp, rel1)

	fmt.Printf("str2:[%s]\n", str2)
	rel2 := TrimAndUpper(str2)
	fmt.Printf("rel2:[%s]\n", rel2)
	assert.Equal(t, exp, rel2)

	fmt.Printf("str3:[%s]\n", str3)
	rel3 := TrimAndUpper(str3)
	fmt.Printf("rel3:[%s]\n", rel3)
	assert.Equal(t, exp, rel3)

}

func TestTrimAndLower(t *testing.T) {
	str1 := " asdf_天行健_FDSA "
	str2 := "	ASDF_天行健_fdsa  "
	str3 := "\nASDF_天行健_FDSA\f"
	exp := "asdf_天行健_fdsa"

	fmt.Printf("str1:[%s]\n", str1)
	rel1 := TrimAndLower(str1)
	fmt.Printf("rel1:[%s]\n", rel1)
	assert.Equal(t, exp, rel1)

	fmt.Printf("str2:[%s]\n", str2)
	rel2 := TrimAndLower(str2)
	fmt.Printf("rel2:[%s]\n", rel2)
	assert.Equal(t, exp, rel2)

	fmt.Printf("str3:[%s]\n", str3)
	rel3 := TrimAndLower(str3)
	fmt.Printf("rel3:[%s]\n", rel3)
	assert.Equal(t, exp, rel3)
}
