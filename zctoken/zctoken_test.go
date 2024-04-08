package zctoken

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"gitee.com/zhaochuninhefei/gmgo/sm2"
	"gitee.com/zhaochuninhefei/gmgo/x509"
	"gitee.com/zhaochuninhefei/zcutils-go/zcrandom"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestPrepareKeys(t *testing.T) {
	// 生成sm2公私钥
	sm2Priv, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	sm2Pub := &sm2Priv.PublicKey

	// 生成私钥pem文件
	_, err = x509.WritePrivateKeytoPemFile("testdata/sm2_pri_key.pem", sm2Priv, nil)
	if err != nil {
		t.Fatal(err)
	}
	// 生成公钥pem文件
	_, err = x509.WritePublicKeytoPemFile("testdata/sm2_pub_key.pem", sm2Pub)
	if err != nil {
		t.Fatal(err)
	}
	// 从pem文件读取私钥
	privKey, err := x509.ReadPrivateKeyFromPemFile("testdata/sm2_pri_key.pem", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("读取到sm2私钥 : %v\n", privKey)
	// 从pem文件读取公钥
	pubKey, err := x509.ReadPublicKeyFromPemFile("testdata/sm2_pub_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("读取到sm2公钥 : %v\n", pubKey)
	fmt.Println("测试sm2私钥与公钥文件读写成功")

	// 生成ecdsa公私钥
	ecPri, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	ecPub := &ecPri.PublicKey
	// 生成私钥pem文件
	_, err = x509.WritePrivateKeytoPemFile("testdata/ec_pri_key.pem", ecPri, nil)
	if err != nil {
		t.Fatal(err)
	}
	// 生成公钥pem文件
	_, err = x509.WritePublicKeytoPemFile("testdata/ec_pub_key.pem", ecPub)
	if err != nil {
		t.Fatal(err)
	}

	// 生成ed25519公私钥
	edPub, edPri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	// 生成私钥pem文件
	_, err = x509.WritePrivateKeytoPemFile("testdata/ed_pri_key.pem", edPri, nil)
	if err != nil {
		t.Fatal(err)
	}
	// 生成公钥pem文件
	_, err = x509.WritePublicKeytoPemFile("testdata/ed_pub_key.pem", edPub)
	if err != nil {
		t.Fatal(err)
	}

}

func TestBuildTokenWithGM(t *testing.T) {
	// 从pem文件读取私钥
	privKey, err := x509.ReadPrivateKeyFromPemFile("testdata/sm2_pri_key.pem", nil)
	if err != nil {
		t.Fatal(err)
	}
	// 从pem文件读取公钥
	pubKey, err := x509.ReadPublicKeyFromPemFile("testdata/sm2_pub_key.pem")
	if err != nil {
		t.Fatal(err)
	}

	payloads := CreateStdPayloads("zhaochun", "test", "anyone", "No001", 5)
	// 注意这里exp传入的是 time.Time{} ,即零值，不重置payloads里的exp
	token, err := BuildTokenWithGM(payloads, time.Time{}, privKey.(*sm2.PrivateKey))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", token)

	time.Sleep(time.Second * 3)

	payloadsAfterCheck, err := CheckTokenWithGM(token, pubKey.(*sm2.PublicKey))
	if err != nil {
		t.Fatal(err)
	}
	jsonPayloads, _ := json.Marshal(payloadsAfterCheck)
	fmt.Printf("jsonPayloads: %s\n", jsonPayloads)
}

func TestBuildTokenWithGMTimeout(t *testing.T) {
	// 从pem文件读取私钥
	privKey, err := x509.ReadPrivateKeyFromPemFile("testdata/sm2_pri_key.pem", nil)
	if err != nil {
		t.Fatal(err)
	}
	// 从pem文件读取公钥
	pubKey, err := x509.ReadPublicKeyFromPemFile("testdata/sm2_pub_key.pem")
	if err != nil {
		t.Fatal(err)
	}

	payloads := CreateStdPayloads("zhaochun", "test", "anyone", "No002", 5)
	// 注意这里重置了过期时间
	token, err := BuildTokenWithGM(payloads, time.Now().Add(time.Second*1), privKey.(*sm2.PrivateKey))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", token)

	time.Sleep(time.Second * 3)

	_, err = CheckTokenWithGM(token, pubKey.(*sm2.PublicKey))
	if err != nil {
		if strings.HasPrefix(err.Error(), "[-1]token过期") {
			fmt.Printf("token超时正确返回错误: %s\n", err)
		} else {
			t.Fatalf("token超时未能返回正确的错误: %s\n", err)
		}
	} else {
		t.Fatal("token超时处理不正确")
	}
}

func TestBuildTokenWithSm2Sm3(t *testing.T) {
	// 从pem文件读取私钥pem
	privKeyPem, err := os.ReadFile("testdata/sm2_pri_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	// 从pem文件读取公钥pem
	pubKeyPem, err := os.ReadFile("testdata/sm2_pub_key.pem")
	if err != nil {
		t.Fatal(err)
	}

	//token, err := PrepareStdTokenStruct("zhaochun", "test", "anyone", "No003", 5, ALG_SM2_SM3)
	token, err := PrepareSplTokenStruct("anyone", 5, ALG_SM2_SM3)
	if err != nil {
		t.Fatal(err)
	}
	tokenStr, err := BuildTokenWithECC(token, time.Time{}, privKeyPem)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", tokenStr)

	jsonToken, err := json.Marshal(token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token struct : %s\n", string(jsonToken))

	token1, err := CheckTokenWithECC(tokenStr, pubKeyPem)
	if err != nil {
		t.Fatal(err)
	}
	jsonToken1, err := json.Marshal(token1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token1 struct : %s\n", string(jsonToken1))

	if reflect.DeepEqual(token, token1) {
		fmt.Println("OK")
	} else {
		t.Fatal("NG")
	}
}

func TestBuildTokenWithSm2Sm3Timeout(t *testing.T) {
	// 从pem文件读取私钥pem
	privKeyPem, err := os.ReadFile("testdata/sm2_pri_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	// 从pem文件读取公钥pem
	pubKeyPem, err := os.ReadFile("testdata/sm2_pub_key.pem")
	if err != nil {
		t.Fatal(err)
	}

	token, err := PrepareStdTokenStruct("zhaochun", "test", "anyone", "No003", 5, ALG_SM2_SM3)
	if err != nil {
		t.Fatal(err)
	}
	tokenStr, err := BuildTokenWithECC(token, time.Now().Add(time.Second*1), privKeyPem)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", tokenStr)

	jsonToken, err := json.Marshal(token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token struct : %s\n", string(jsonToken))

	time.Sleep(time.Second * 3)

	_, err = CheckTokenWithECC(tokenStr, pubKeyPem)
	if err != nil {
		if strings.HasPrefix(err.Error(), "[-1]token过期") {
			fmt.Printf("token超时正确返回错误: %s\n", err)
		} else {
			t.Fatalf("token超时未能返回正确的错误: %s\n", err)
		}
	} else {
		t.Fatal("token超时处理不正确")
	}
}

func TestBuildTokenWithEcdsa(t *testing.T) {
	// 从pem文件读取私钥pem
	privKeyPem, err := os.ReadFile("testdata/ec_pri_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	// 从pem文件读取公钥pem
	pubKeyPem, err := os.ReadFile("testdata/ec_pub_key.pem")
	if err != nil {
		t.Fatal(err)
	}

	token, err := PrepareStdTokenStruct("zhaochun", "test", "anyone", "No003", 5, ALG_ECDSA_SHA256)
	if err != nil {
		t.Fatal(err)
	}
	tokenStr, err := BuildTokenWithECC(token, time.Time{}, privKeyPem)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", tokenStr)

	jsonToken, err := json.Marshal(token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token struct : %s\n", string(jsonToken))

	token1, err := CheckTokenWithECC(tokenStr, pubKeyPem)
	if err != nil {
		t.Fatal(err)
	}
	jsonToken1, err := json.Marshal(token1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token1 struct : %s\n", string(jsonToken1))

	if reflect.DeepEqual(token, token1) {
		fmt.Println("OK")
	} else {
		t.Fatal("NG")
	}
}

func TestBuildTokenWithEd25519(t *testing.T) {
	// 从pem文件读取私钥pem
	privKeyPem, err := os.ReadFile("testdata/ed_pri_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	// 从pem文件读取公钥pem
	pubKeyPem, err := os.ReadFile("testdata/ed_pub_key.pem")
	if err != nil {
		t.Fatal(err)
	}

	token, err := PrepareStdTokenStruct("zhaochun", "test", "anyone", "No003", 5, ALG_ED25519_SHA256)
	if err != nil {
		t.Fatal(err)
	}
	tokenStr, err := BuildTokenWithECC(token, time.Time{}, privKeyPem)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", tokenStr)

	jsonToken, err := json.Marshal(token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token struct : %s\n", string(jsonToken))

	token1, err := CheckTokenWithECC(tokenStr, pubKeyPem)
	if err != nil {
		t.Fatal(err)
	}
	jsonToken1, err := json.Marshal(token1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token1 struct : %s\n", string(jsonToken1))

	if reflect.DeepEqual(token, token1) {
		fmt.Println("OK")
	} else {
		t.Fatal("NG")
	}
}

func TestBuildTokenWithHMACSM3(t *testing.T) {
	keyBytes, _ := zcrandom.GenerateRandomBytes(64)

	token, err := PrepareSplTokenStruct("anyone", 5, ALG_HMAC_SM3)
	if err != nil {
		t.Fatal(err)
	}
	tokenStr, err := BuildTokenWithHMAC(token, time.Time{}, keyBytes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", tokenStr)

	jsonToken, err := json.Marshal(token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token struct : %s\n", string(jsonToken))

	token1, err := CheckTokenWithHMAC(tokenStr, keyBytes)
	if err != nil {
		t.Fatal(err)
	}
	jsonToken1, err := json.Marshal(token1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token1 struct : %s\n", string(jsonToken1))

	if reflect.DeepEqual(token, token1) {
		fmt.Println("OK")
	} else {
		t.Fatal("NG")
	}
}

func TestBuildTokenWithHMACSHA256(t *testing.T) {
	//keyBytes, _ := zcrandom.GenerateRandomBytes(64)

	token, err := PrepareSplTokenStruct("anyone", 5, ALG_HMAC_SHA256)
	if err != nil {
		t.Fatal(err)
	}
	tokenStr, err := BuildTokenWithHMAC(token, time.Time{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", tokenStr)

	jsonToken, err := json.Marshal(token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token struct : %s\n", string(jsonToken))

	token1, err := CheckTokenWithHMAC(tokenStr, nil)
	if err != nil {
		t.Fatal(err)
	}
	jsonToken1, err := json.Marshal(token1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token1 struct : %s\n", string(jsonToken1))

	if reflect.DeepEqual(token, token1) {
		fmt.Println("OK")
	} else {
		t.Fatal("NG")
	}
}

func BenchmarkBuildTokenWithSM2SM3(b *testing.B) {
	// 从pem文件读取私钥pem
	privKeyPem, err := os.ReadFile("testdata/sm2_pri_key.pem")
	if err != nil {
		b.Fatal(err)
	}
	token, err := PrepareSplTokenStruct("anyone", 5, ALG_SM2_SM3)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err = BuildTokenWithECC(token, time.Time{}, privKeyPem)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBuildTokenWithECDSA(b *testing.B) {
	// 从pem文件读取私钥pem
	privKeyPem, err := os.ReadFile("testdata/ec_pri_key.pem")
	if err != nil {
		b.Fatal(err)
	}
	token, err := PrepareSplTokenStruct("anyone", 5, ALG_ECDSA_SHA256)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err = BuildTokenWithECC(token, time.Time{}, privKeyPem)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBuildTokenWithED25519(b *testing.B) {
	// 从pem文件读取私钥pem
	privKeyPem, err := os.ReadFile("testdata/ed_pri_key.pem")
	if err != nil {
		b.Fatal(err)
	}
	token, err := PrepareSplTokenStruct("anyone", 5, ALG_ED25519_SHA256)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err = BuildTokenWithECC(token, time.Time{}, privKeyPem)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBuildTokenWithHMACSM3(b *testing.B) {
	keyBytes, _ := zcrandom.GenerateRandomBytes(64)

	token, err := PrepareSplTokenStruct("anyone", 5, ALG_HMAC_SM3)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err = BuildTokenWithHMAC(token, time.Time{}, keyBytes)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBuildTokenWithHMACSHA256(b *testing.B) {
	keyBytes, _ := zcrandom.GenerateRandomBytes(64)

	token, err := PrepareSplTokenStruct("anyone", 5, ALG_HMAC_SHA256)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err = BuildTokenWithHMAC(token, time.Time{}, keyBytes)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCheckTokenWithSM2SM3(b *testing.B) {
	// 从pem文件读取私钥pem
	privKeyPem, err := os.ReadFile("testdata/sm2_pri_key.pem")
	if err != nil {
		b.Fatal(err)
	}
	// 从pem文件读取公钥pem
	pubKeyPem, err := os.ReadFile("testdata/sm2_pub_key.pem")
	if err != nil {
		b.Fatal(err)
	}

	token, err := PrepareSplTokenStruct("anyone", 5, ALG_SM2_SM3)
	if err != nil {
		b.Fatal(err)
	}
	_, err = BuildTokenWithECC(token, time.Time{}, privKeyPem)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := CheckTokenWithECC(token.TokenStr, pubKeyPem)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCheckTokenWithECDSA(b *testing.B) {
	// 从pem文件读取私钥pem
	privKeyPem, err := os.ReadFile("testdata/ec_pri_key.pem")
	if err != nil {
		b.Fatal(err)
	}
	// 从pem文件读取公钥pem
	pubKeyPem, err := os.ReadFile("testdata/ec_pub_key.pem")
	if err != nil {
		b.Fatal(err)
	}

	token, err := PrepareSplTokenStruct("anyone", 5, ALG_ECDSA_SHA256)
	if err != nil {
		b.Fatal(err)
	}
	_, err = BuildTokenWithECC(token, time.Time{}, privKeyPem)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := CheckTokenWithECC(token.TokenStr, pubKeyPem)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCheckTokenWithED25519(b *testing.B) {
	// 从pem文件读取私钥pem
	privKeyPem, err := os.ReadFile("testdata/ed_pri_key.pem")
	if err != nil {
		b.Fatal(err)
	}
	// 从pem文件读取公钥pem
	pubKeyPem, err := os.ReadFile("testdata/ed_pub_key.pem")
	if err != nil {
		b.Fatal(err)
	}

	token, err := PrepareSplTokenStruct("anyone", 5, ALG_ED25519_SHA256)
	if err != nil {
		b.Fatal(err)
	}
	_, err = BuildTokenWithECC(token, time.Time{}, privKeyPem)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := CheckTokenWithECC(token.TokenStr, pubKeyPem)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCheckTokenWithHMACSM3(b *testing.B) {
	keyBytes, _ := zcrandom.GenerateRandomBytes(64)

	token, err := PrepareSplTokenStruct("anyone", 5, ALG_HMAC_SM3)
	if err != nil {
		b.Fatal(err)
	}
	_, err = BuildTokenWithHMAC(token, time.Time{}, keyBytes)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := CheckTokenWithHMAC(token.TokenStr, keyBytes)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCheckTokenWithHMACSHA256(b *testing.B) {
	keyBytes, _ := zcrandom.GenerateRandomBytes(64)

	token, err := PrepareSplTokenStruct("anyone", 5, ALG_HMAC_SHA256)
	if err != nil {
		b.Fatal(err)
	}
	_, err = BuildTokenWithHMAC(token, time.Time{}, keyBytes)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := CheckTokenWithHMAC(token.TokenStr, keyBytes)
		if err != nil {
			b.Fatal(err)
		}
	}
}
