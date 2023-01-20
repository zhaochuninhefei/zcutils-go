package zcrandom

import "crypto/rand"

// GenerateRandomBytes 生成一个指定长度的随机字节数组
//  @param size 长度
//  @return []byte 随机字节数组
//  @return error
func GenerateRandomBytes(size uint) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
