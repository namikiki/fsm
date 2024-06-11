package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
)

// RandBytesSlice  返回随机一定长度的byte数组
func RandBytesSlice(lens int) []byte {
	Bytes := make([]byte, lens)

	for i := 0; i < lens; i++ {
		num := 48 + rand.Intn(74)
		Bytes[i] = byte(num)
	}
	return Bytes
}

func PasswordHash(password string, salt []byte) string {
	//计算 salt 和密码组合的SHA1摘要
	h := sha1.New()
	h.Write([]byte(password))
	h.Write(salt)

	return hex.EncodeToString(h.Sum(nil))
}
