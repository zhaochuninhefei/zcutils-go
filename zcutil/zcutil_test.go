package zcutil

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
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
