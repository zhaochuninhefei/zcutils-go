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

func TestTempDir(t *testing.T) {
	tmpKeyStore, err := ioutil.TempDir("testdata", "msp-keystore")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tmpKeyStore)
}
