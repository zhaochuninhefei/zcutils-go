// Package zctoken 凭证处理包
package zctoken

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/zhaochuninhefei/gmgo/sm2"
	"gitee.com/zhaochuninhefei/gmgo/sm3"
	"gitee.com/zhaochuninhefei/gmgo/x509"
	"gitee.com/zhaochuninhefei/zcutils-go/zctime"
	"hash"
	"strings"
	"time"
)

// Alg 凭证算法类型，目前支持:"SM2-SM3","ECDSA-SHA256","ED25519-SHA256","HMAC-SM3","HMAC-SHA256"。
//  - "SM2-SM3","ECDSA-SHA256","ED25519-SHA256"，使用椭圆曲线签名算法生成token，算法前半部是签名算法，后半部是散列算法(用于签名前计算凭证内容摘要)。
//  - "HMAC-SM3","HMAC-SHA256"，表示为凭证生成HMAC而不是签名，算法后半部是HMAC对应的散列算法。
type Alg string

// zctoken支持的凭证算法列表、默认算法以及默认凭证类型(目前只有JWT)
//goland:noinspection GoSnakeCaseUsage
const (
	ALG_SM2_SM3        Alg = "SM2-SM3"
	ALG_ECDSA_SHA256   Alg = "ECDSA-SHA256"
	ALG_ED25519_SHA256 Alg = "ED25519-SHA256"
	ALG_HMAC_SM3       Alg = "HMAC-SM3"
	ALG_HMAC_SHA256    Alg = "HMAC-SHA256"

	// ALG_DEFAULT 默认凭证算法
	ALG_DEFAULT = ALG_SM2_SM3
	// TYP_DEFAULT 默认凭证类型
	TYP_DEFAULT = "JWT"

	// HMAC_KEY_DEFAULT_HEX HMAC默认密钥，长度64的字节数组转为hex字符串，使用`zcrandom.GenerateRandomBytes`生成。
	HMAC_KEY_DEFAULT_HEX = "1fa90680817dc824c48140190aa9d3fb1c0643f5efce445a36bab21afd4b2c5328f41adf668e46b86a27087499f2c8cdbe91a3c717dca18430e04c942a0e74aa"
)

// hmacKeyDefault HMAC默认密钥，长度64的字节数组转为hex字符串，使用`zcrandom.GenerateRandomBytes`生成
var hmacKeyDefault []byte

func init() {
	var err error
	hmacKeyDefault, err = hex.DecodeString(HMAC_KEY_DEFAULT_HEX)
	if err != nil {
		panic(err)
	}
}

// Token 凭证结构体
type Token struct {
	// Header 凭证头
	Header TokenHeader `json:"header"`
	// Payloads 凭证有效负载
	Payloads map[string]string `json:"payloads"`
	// TokenStr 凭证字符串
	TokenStr string `json:"token_str"`
}

// TokenHeader 凭证头部
type TokenHeader struct {
	// Alg 凭证算法
	Alg Alg `json:"alg"`
	// Typ 凭证类型
	Typ string `json:"typ"`
}

// CreateTokenHeader 创建凭证头部
//  @param alg 凭证算法
//  @param typ 凭证类型
//  @return *TokenHeader
func CreateTokenHeader(alg Alg, typ string) *TokenHeader {
	return &TokenHeader{
		Alg: alg,
		Typ: typ,
	}
}

// CreateTokenHeaderDefault 使用默认配置创建凭证头部
//  @return *TokenHeader
func CreateTokenHeaderDefault() *TokenHeader {
	return CreateTokenHeader(ALG_DEFAULT, TYP_DEFAULT)
}

// CreateStdPayloads 创建标准凭证有效负载
//  其中，过期时间使用 `当前时间 + expSeconds过期时间秒数` ，生效时间与签发时间均采用当前时间
//
//  @param iss 签发者
//  @param sub 主题
//  @param aud 受众
//  @param jti 编号
//  @param expSeconds 过期时间秒数
//  @return map[string]string 凭证有效负载
func CreateStdPayloads(iss string, sub string, aud string, jti string, expSeconds uint64) map[string]string {
	now := time.Now()
	strNow := now.Format(zctime.TIME_FORMAT_SIMPLE)
	payloads := make(map[string]string)
	// 签发者
	payloads["iss"] = iss
	// 主题
	payloads["sub"] = sub
	// 受众
	payloads["aud"] = aud
	// 编号
	payloads["jti"] = jti
	// 过期时间
	payloads["exp"] = now.Add(time.Second * time.Duration(expSeconds)).Format(zctime.TIME_FORMAT_SIMPLE)
	// 生效时间
	payloads["nbf"] = strNow
	// 签发时间
	payloads["iat"] = strNow
	return payloads
}

// CreateSplPayloads 创建简单版凭证有效负载
//
//  @param aud 受众
//  @param expSeconds
//  @return map[string]string
func CreateSplPayloads(aud string, expSeconds uint64) map[string]string {
	now := time.Now()
	payloads := make(map[string]string)
	// 受众
	payloads["aud"] = aud
	// 过期时间
	payloads["exp"] = now.Add(time.Second * time.Duration(expSeconds)).Format(zctime.TIME_FORMAT_SIMPLE)
	return payloads
}

// PrepareStdTokenStruct 准备标准凭证结构体
//
//  @param iss 签发者
//  @param sub 主题
//  @param aud 受众
//  @param jti 编号
//  @param expSeconds 过期时间秒数
//  @param alg 凭证算法
//  @return *Token 凭证结构体(指针)
//  @return error
func PrepareStdTokenStruct(
	iss string,
	sub string,
	aud string,
	jti string,
	expSeconds uint64,
	alg Alg) (*Token, error) {
	// 创建token结构体
	token := &Token{
		Header: TokenHeader{
			Alg: alg,
			Typ: TYP_DEFAULT,
		},
		Payloads: CreateStdPayloads(iss, sub, aud, jti, expSeconds),
	}
	return token, nil
}

// PrepareSplTokenStruct 准备标准凭证结构体
//
//  @param aud 受众
//  @param expSeconds 过期时间秒数
//  @param alg 凭证算法
//  @return *Token 凭证结构体(指针)
//  @return error
func PrepareSplTokenStruct(aud string, expSeconds uint64, alg Alg) (*Token, error) {
	// 创建token结构体
	token := &Token{
		Header: TokenHeader{
			Alg: alg,
			Typ: TYP_DEFAULT,
		},
		Payloads: CreateSplPayloads(aud, expSeconds),
	}
	return token, nil
}

// BuildTokenWithECC 使用椭圆曲线签名算法构建凭证
//  @param token 凭证结构体
//  @param exp 凭证过期时间，如果不打算重置token.Payloads中的过期时间，则这里传入time零值(`time.Time{}`)即可。
//  @param priKeyPem 私钥pem
//  @return error
func BuildTokenWithECC(token *Token, exp time.Time, priKeyPem []byte) error {
	if token == nil {
		return errors.New("[-9]token不可传nil")
	}
	// 创建默认token头部
	tokenHeader := &token.Header
	if tokenHeader == nil {
		return errors.New("[-9]token头不可为nil")
	}
	// 将token头转为json
	jsonTokenHeader, err := json.Marshal(&tokenHeader)
	if err != nil {
		return fmt.Errorf("[-9]token头json序列化失败: %s", err)
	}
	// 对token头做base64编码
	headerBase64 := base64.RawURLEncoding.EncodeToString(jsonTokenHeader)

	payloads := token.Payloads
	if payloads == nil {
		return errors.New("[-9]token有效负载不可为nil")
	}
	// 重置凭证过期时间
	if !exp.IsZero() {
		payloads["exp"] = exp.Format(zctime.TIME_FORMAT_SIMPLE)
	}
	// 将token的有效负载转为json
	jsonPayloads, err := json.Marshal(payloads)
	if err != nil {
		return fmt.Errorf("[-9]token有效负载json序列化失败: %s", err)
	}
	// 对token的有效负载做base64编码
	payloadsBase64 := base64.RawURLEncoding.EncodeToString(jsonPayloads)
	// 拼接token内容
	content := strings.Join([]string{headerBase64, payloadsBase64}, ".")

	priKey, err := x509.ReadPrivateKeyFromPem(priKeyPem, nil)
	if err != nil {
		return fmt.Errorf("[-9]私钥pem读取失败: %s", err)
	}
	// 根据Alg选择算法对content做哈希并签名
	var digest, sign []byte
	switch tokenHeader.Alg {
	case ALG_SM2_SM3:
		digest = sm3.Sm3Sum([]byte(content))
		switch priKey := priKey.(type) {
		case *sm2.PrivateKey:
			// 对摘要做sm2签名
			sign, err = priKey.Sign(rand.Reader, digest, nil)
			if err != nil {
				return fmt.Errorf("[-9]token签名失败: %s", err)
			}
		default:
			return fmt.Errorf("[-9]凭证算法(%s)与私钥pem不匹配", tokenHeader.Alg)
		}

	case ALG_ECDSA_SHA256:
		sum := sha256.Sum256([]byte(content))
		digest = sum[:]
		switch priKey := priKey.(type) {
		case *ecdsa.PrivateKey:
			// 对摘要做sm2签名
			sign, err = priKey.Sign(rand.Reader, digest, nil)
			if err != nil {
				return fmt.Errorf("[-9]token签名失败: %s", err)
			}
		default:
			return fmt.Errorf("[-9]凭证算法(%s)与私钥pem不匹配", tokenHeader.Alg)
		}

	case ALG_ED25519_SHA256:
		sum := sha256.Sum256([]byte(content))
		digest = sum[:]
		switch priKey := priKey.(type) {
		case ed25519.PrivateKey:
			// 对摘要做sm2签名
			sign, err = priKey.Sign(rand.Reader, digest, crypto.Hash(0))
			if err != nil {
				return fmt.Errorf("[-9]token签名失败: %s", err)
			}
		default:
			return fmt.Errorf("[-9]凭证算法(%s)与私钥pem不匹配", tokenHeader.Alg)
		}
	default:
		return fmt.Errorf("[-9]BuildTokenWithECC不支持的凭证算法: %s", tokenHeader.Alg)
	}

	// 将签名转为hex字符串
	signStr := hex.EncodeToString(sign)
	// 拼接凭证
	tokenStr := strings.Join([]string{content, signStr}, ".")
	token.TokenStr = tokenStr

	return nil
}

// CheckTokenWithECC 使用椭圆曲线签名算法校验凭证
//  @param tokenStr 凭证字符串
//  @param pubKeyPem 验签公钥pem
//  @return *Token 凭证结构体(指针)
//  @return error
func CheckTokenWithECC(tokenStr string, pubKeyPem []byte) (*Token, error) {
	token := &Token{
		TokenStr: tokenStr,
	}

	tmpArr := strings.Split(tokenStr, ".")
	if len(tmpArr) != 3 {
		return nil, errors.New("[-5]token格式错误")
	}
	headerBase64 := tmpArr[0]
	payloadsBase64 := tmpArr[1]
	signStr := tmpArr[2]

	// 检查token头
	jsonTokenHeader, err := base64.RawURLEncoding.DecodeString(headerBase64)
	if err != nil {
		return nil, fmt.Errorf("[-5]token头base64解码失败: %s", err)
	}
	var tokenHeader TokenHeader
	err = json.Unmarshal(jsonTokenHeader, &tokenHeader)
	if err != nil {
		return nil, fmt.Errorf("[-5]token头json反序列化失败: %s", err)
	}
	token.Header = tokenHeader

	// 签名解码
	sign, err := hex.DecodeString(signStr)
	if err != nil {
		return nil, fmt.Errorf("[-5]token签名hex解码失败: %s", err)
	}
	// 读取公钥pem
	pubKey, err := x509.ReadPublicKeyFromPem(pubKeyPem)
	if err != nil {
		return nil, fmt.Errorf("[-9]公钥pem读取失败: %s", err)
	}
	// 签名内容
	content := strings.Join([]string{headerBase64, payloadsBase64}, ".")
	var digest []byte
	switch tokenHeader.Alg {
	case ALG_SM2_SM3:
		// 签名摘要
		digest = sm3.Sm3Sum([]byte(content))
		// 检查公钥类型并验签
		switch pubKey := pubKey.(type) {
		case *sm2.PublicKey:
			if !pubKey.Verify(digest, sign) {
				return nil, fmt.Errorf("[-5]token验签失败: %s", err)
			}
		default:
			return nil, fmt.Errorf("[-9]凭证算法(%s)与公钥pem不匹配", tokenHeader.Alg)
		}

	case ALG_ECDSA_SHA256:
		// 签名摘要
		sum := sha256.Sum256([]byte(content))
		digest = sum[:]
		// 检查公钥类型并验签
		switch pubKey := pubKey.(type) {
		case *ecdsa.PublicKey:
			if !ecdsa.VerifyASN1(pubKey, digest, sign) {
				return nil, fmt.Errorf("[-5]token验签失败: %s", err)
			}
		default:
			return nil, fmt.Errorf("[-9]凭证算法(%s)与私钥pem不匹配", tokenHeader.Alg)
		}

	case ALG_ED25519_SHA256:
		// 签名摘要
		sum := sha256.Sum256([]byte(content))
		digest = sum[:]
		// 检查公钥类型并验签
		switch pubKey := pubKey.(type) {
		case ed25519.PublicKey:
			if !ed25519.Verify(pubKey, digest, sign) {
				return nil, fmt.Errorf("[-5]token验签失败: %s", err)
			}
		default:
			return nil, fmt.Errorf("[-9]凭证算法(%s)与私钥pem不匹配", tokenHeader.Alg)
		}
	default:
		return nil, fmt.Errorf("[-9]不支持的凭证算法: %s", tokenHeader.Alg)
	}

	// 解析有效负载
	jsonPayloads, err := base64.RawURLEncoding.DecodeString(payloadsBase64)
	if err != nil {
		return nil, fmt.Errorf("[-5]token有效负载base64解码失败: %s", err)
	}
	var payloads map[string]string
	err = json.Unmarshal(jsonPayloads, &payloads)
	if err != nil {
		return nil, fmt.Errorf("[-5]token有效负载json反序列化失败: %s", err)
	}
	token.Payloads = payloads

	// 凭证过期检查
	expVal := payloads["exp"]
	if expVal != "" {
		now := time.Now()
		exp, err := time.ParseInLocation(zctime.TIME_FORMAT_SIMPLE, expVal, time.Local)
		if err != nil {
			return nil, fmt.Errorf("[-5]token过期时间反序列化失败: %s", err)
		}
		if now.After(exp) {
			return nil, fmt.Errorf("[-1]token过期,过期时间: %s, 检查时间: %s", expVal, now.Format(zctime.TIME_FORMAT_SIMPLE))
		}
	}

	return token, nil
}

// BuildTokenWithGM 使用SM2-SM3算法构建凭证
//  @param payloads 凭证有效负载
//  @param exp 凭证过期时间，如果不打算重置payloads中的过期时间，则这里传入time零值(`time.Time{}`)即可。
//  @param priKey 签名私钥(sm2)
//  @return string 凭证字符串
//  @return error
func BuildTokenWithGM(payloads map[string]string, exp time.Time, priKey *sm2.PrivateKey) (string, error) {
	if payloads == nil {
		return "", errors.New("[-1]凭证有效负载不可为nil")
	}
	if priKey == nil {
		return "", errors.New("[-1]签名私钥(sm2)不可为nil")
	}

	// 创建默认token头部
	tokenHeader := CreateTokenHeaderDefault()
	// 将token头转为json
	jsonTokenHeader, err := json.Marshal(&tokenHeader)
	if err != nil {
		return "", fmt.Errorf("[-9]token头json序列化失败: %s", err)
	}
	// 对token头做base64编码
	headerBase64 := base64.RawURLEncoding.EncodeToString(jsonTokenHeader)

	// 重置凭证过期时间
	if !exp.IsZero() {
		payloads["exp"] = exp.Format(zctime.TIME_FORMAT_SIMPLE)
	}
	// 将token的有效负载转为json
	jsonPayloads, err := json.Marshal(payloads)
	if err != nil {
		return "", fmt.Errorf("[-9]token有效负载json序列化失败: %s", err)
	}
	// 对token的有效负载做base64编码
	payloadsBase64 := base64.RawURLEncoding.EncodeToString(jsonPayloads)
	// 拼接token内容
	content := strings.Join([]string{headerBase64, payloadsBase64}, ".")
	// 对token内容做sm3摘要计算
	digest := sm3.Sm3Sum([]byte(content))
	// 对摘要做sm2签名
	sign, err := priKey.Sign(rand.Reader, digest, nil)
	if err != nil {
		return "", fmt.Errorf("[-9]token签名失败: %s", err)
	}
	// 将签名转为hex字符串
	signStr := hex.EncodeToString(sign)
	// 拼接凭证
	token := strings.Join([]string{content, signStr}, ".")
	return token, nil
}

// CheckTokenWithGM 使用SM2-SM3算法校验凭证
//  @param token 凭证字符串
//  @param pubKey 验签公钥(sm2)
//  @return map[string]string 凭证有效负载
//  @return error
func CheckTokenWithGM(token string, pubKey *sm2.PublicKey) (map[string]string, error) {
	tmpArr := strings.Split(token, ".")
	if len(tmpArr) != 3 {
		return nil, errors.New("[-5]token格式错误")
	}
	headerBase64 := tmpArr[0]
	payloadsBase64 := tmpArr[1]
	signStr := tmpArr[2]

	// 检查token头
	jsonTokenHeader, err := base64.RawURLEncoding.DecodeString(headerBase64)
	if err != nil {
		return nil, fmt.Errorf("[-5]token头base64解码失败: %s", err)
	}
	var tokenHeader TokenHeader
	err = json.Unmarshal(jsonTokenHeader, &tokenHeader)
	if err != nil {
		return nil, fmt.Errorf("[-5]token头json反序列化失败: %s", err)
	}

	// 检查签名
	content := strings.Join([]string{headerBase64, payloadsBase64}, ".")
	digest := sm3.Sm3Sum([]byte(content))
	sign, err := hex.DecodeString(signStr)
	if err != nil {
		return nil, fmt.Errorf("[-5]token签名hex解码失败: %s", err)
	}
	if !pubKey.Verify(digest, sign) {
		return nil, fmt.Errorf("[-5]token验签失败: %s", err)
	}

	// 解析有效负载
	jsonPayloads, err := base64.RawURLEncoding.DecodeString(payloadsBase64)
	if err != nil {
		return nil, fmt.Errorf("[-5]token有效负载base64解码失败: %s", err)
	}
	var payloads map[string]string
	err = json.Unmarshal(jsonPayloads, &payloads)
	if err != nil {
		return nil, fmt.Errorf("[-5]token有效负载json反序列化失败: %s", err)
	}
	// 凭证过期检查
	expVal := payloads["exp"]
	if expVal != "" {
		now := time.Now()
		exp, err := time.ParseInLocation(zctime.TIME_FORMAT_SIMPLE, expVal, time.Local)
		if err != nil {
			return nil, fmt.Errorf("[-5]token过期时间反序列化失败: %s", err)
		}
		if now.After(exp) {
			return nil, fmt.Errorf("[-1]token过期,过期时间: %s, 检查时间: %s", expVal, now.Format(zctime.TIME_FORMAT_SIMPLE))
		}
	}

	return payloads, nil
}

// BuildTokenWithHMAC 使用HMAC算法构建凭证
//  @param token 凭证结构体
//  @param exp 凭证过期时间，如果不打算重置payloads中的过期时间，则这里传入time零值(`time.Time{}`)即可。
//  @param keyBytes HMAC密钥
//  @return error
func BuildTokenWithHMAC(token *Token, exp time.Time, keyBytes []byte) error {
	if token == nil {
		return errors.New("[-9]token不可传nil")
	}
	if keyBytes == nil {
		keyBytes = hmacKeyDefault
	}

	// 创建默认token头部
	tokenHeader := &token.Header
	if tokenHeader == nil {
		return errors.New("[-9]token头不可为nil")
	}
	// 将token头转为json
	jsonTokenHeader, err := json.Marshal(&tokenHeader)
	if err != nil {
		return fmt.Errorf("[-9]token头json序列化失败: %s", err)
	}
	// 对token头做base64编码
	headerBase64 := base64.RawURLEncoding.EncodeToString(jsonTokenHeader)

	payloads := token.Payloads
	if payloads == nil {
		return errors.New("[-9]token有效负载不可为nil")
	}
	// 重置凭证过期时间
	if !exp.IsZero() {
		payloads["exp"] = exp.Format(zctime.TIME_FORMAT_SIMPLE)
	}
	// 将token的有效负载转为json
	jsonPayloads, err := json.Marshal(payloads)
	if err != nil {
		return fmt.Errorf("[-9]token有效负载json序列化失败: %s", err)
	}
	// 对token的有效负载做base64编码
	payloadsBase64 := base64.RawURLEncoding.EncodeToString(jsonPayloads)
	// 拼接校验内容
	content := strings.Join([]string{headerBase64, payloadsBase64}, ".")

	var hasher hash.Hash
	switch tokenHeader.Alg {
	case ALG_HMAC_SM3:
		hasher = hmac.New(sm3.New, keyBytes)
	case ALG_HMAC_SHA256:
		hasher = hmac.New(sha256.New, keyBytes)
	default:
		return fmt.Errorf("[-9]BuildTokenWithHMAC不支持的凭证算法: %s", tokenHeader.Alg)
	}
	hasher.Write([]byte(content))
	sum := hasher.Sum(nil)
	// 将校验和转为hex字符串
	sumStr := hex.EncodeToString(sum)
	// 拼接凭证
	tokenStr := strings.Join([]string{content, sumStr}, ".")
	token.TokenStr = tokenStr

	return nil
}

// CheckTokenWithHMAC 使用HMAC算法校验凭证
//  @param tokenStr 凭证字符串
//  @param keyBytes HMAC密钥
//  @return *Token 凭证结构体(指针)
//  @return error
func CheckTokenWithHMAC(tokenStr string, keyBytes []byte) (*Token, error) {
	if keyBytes == nil {
		keyBytes = hmacKeyDefault
	}

	token := &Token{
		TokenStr: tokenStr,
	}

	tmpArr := strings.Split(tokenStr, ".")
	if len(tmpArr) != 3 {
		return nil, errors.New("[-5]token格式错误")
	}
	headerBase64 := tmpArr[0]
	payloadsBase64 := tmpArr[1]
	sumStr := tmpArr[2]

	// 检查token头
	jsonTokenHeader, err := base64.RawURLEncoding.DecodeString(headerBase64)
	if err != nil {
		return nil, fmt.Errorf("[-5]token头base64解码失败: %s", err)
	}
	var tokenHeader TokenHeader
	err = json.Unmarshal(jsonTokenHeader, &tokenHeader)
	if err != nil {
		return nil, fmt.Errorf("[-5]token头json反序列化失败: %s", err)
	}
	token.Header = tokenHeader

	// 校验和解码
	sum, err := hex.DecodeString(sumStr)
	if err != nil {
		return nil, fmt.Errorf("[-5]token签名hex解码失败: %s", err)
	}
	// 校验内容
	content := strings.Join([]string{headerBase64, payloadsBase64}, ".")
	var hasher hash.Hash
	switch tokenHeader.Alg {
	case ALG_HMAC_SM3:
		hasher = hmac.New(sm3.New, keyBytes)
	case ALG_HMAC_SHA256:
		hasher = hmac.New(sha256.New, keyBytes)
	default:
		return nil, fmt.Errorf("[-9]CheckTokenWithHMAC不支持的凭证算法: %s", tokenHeader.Alg)
	}
	hasher.Write([]byte(content))
	if !hmac.Equal(sum, hasher.Sum(nil)) {
		return nil, fmt.Errorf("[-5]token验签失败: %s", err)
	}

	// 解析有效负载
	jsonPayloads, err := base64.RawURLEncoding.DecodeString(payloadsBase64)
	if err != nil {
		return nil, fmt.Errorf("[-5]token有效负载base64解码失败: %s", err)
	}
	var payloads map[string]string
	err = json.Unmarshal(jsonPayloads, &payloads)
	if err != nil {
		return nil, fmt.Errorf("[-5]token有效负载json反序列化失败: %s", err)
	}
	token.Payloads = payloads

	// 凭证过期检查
	expVal := payloads["exp"]
	if expVal != "" {
		now := time.Now()
		exp, err := time.ParseInLocation(zctime.TIME_FORMAT_SIMPLE, expVal, time.Local)
		if err != nil {
			return nil, fmt.Errorf("[-5]token过期时间反序列化失败: %s", err)
		}
		if now.After(exp) {
			return nil, fmt.Errorf("[-1]token过期,过期时间: %s, 检查时间: %s", expVal, now.Format(zctime.TIME_FORMAT_SIMPLE))
		}
	}

	return token, nil
}
