package zctoken

import (
	"crypto/rand"
	"fmt"
	"gitee.com/zhaochuninhefei/gmgo/sm2"
	"gitee.com/zhaochuninhefei/gmgo/x509"
	"testing"
)

func TestCreateSM2Key(t *testing.T) {
	// 生成sm2密钥对
	priv, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	pub := &priv.PublicKey

	// 生成私钥pem文件
	_, err = x509.WritePrivateKeytoPemFile("testdata/pri_key.pem", priv, nil)
	if err != nil {
		t.Fatal(err)
	}
	// 生成公钥pem文件
	_, err = x509.WritePublicKeytoPemFile("testdata/pub_key.pem", pub)
	if err != nil {
		t.Fatal(err)
	}
	// 从pem文件读取私钥
	privKey, err := x509.ReadPrivateKeyFromPemFile("testdata/pri_key.pem", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("读取到sm2私钥 : %v\n", privKey)
	// 从pem文件读取公钥
	pubKey, err := x509.ReadPublicKeyFromPemFile("testdata/pub_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("读取到sm2公钥 : %v\n", pubKey)
	fmt.Println("测试sm2私钥与公钥文件读写成功")
}
