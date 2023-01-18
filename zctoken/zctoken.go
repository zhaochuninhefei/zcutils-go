// Package zctoken 凭证处理包
package zctoken

import (
	"encoding/base64"
	"encoding/json"
	"gitee.com/zhaochuninhefei/gmgo/sm3"
)

//goland:noinspection GoSnakeCaseUsage
const (
	// ALG_DEFAULT 默认凭证算法
	ALG_DEFAULT = "SM2-with-SM3"
	// TYP_DEFAULT 默认凭证类型
	TYP_DEFAULT = "JWT"
)

// TokenHeader 凭证头部
type TokenHeader struct {
	// Alg 凭证算法
	Alg string `json:"alg"`
	// Typ 凭证类型
	Typ string `json:"typ"`
}

// CreateTokenHeader 创建凭证头部
//  @param alg 凭证算法
//  @param typ 凭证类型
//  @return *TokenHeader
func CreateTokenHeader(alg string, typ string) *TokenHeader {
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
func BuildToken(payloads map[string]string, priKeyBase64 string) (string, error) {
	tokenHeader := CreateTokenHeaderDefault()
	jsonTokenHeader, err := json.Marshal(&tokenHeader)
	if err != nil {
		return "", err
	}
	headerBase64 := base64.URLEncoding.EncodeToString(jsonTokenHeader)

	jsonPayloads, err := json.Marshal(payloads)
	if err != nil {
		return "", err
	}
	payloadsBase64 := base64.URLEncoding.EncodeToString(jsonPayloads)

	content := headerBase64 + "." + payloadsBase64

	sm3.Sm3Sum([]byte(content))

	_, err = base64.URLEncoding.DecodeString(priKeyBase64)
	if err != nil {
		return "", err
	}

	//sign, err := sm2.Sm2Sign(privKey, digest, rand.Reader)

	return "", nil
}
