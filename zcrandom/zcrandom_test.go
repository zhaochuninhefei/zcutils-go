package zcrandom

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGenerateRandomBytes(t *testing.T) {
	for i := 0; i < 10; i++ {
		rbytes, _ := GenerateRandomBytes(64)
		fmt.Println(hex.EncodeToString(rbytes))
	}
}
