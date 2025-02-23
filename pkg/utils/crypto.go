package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// HashPassword 密码加盐哈希
func HashPassword(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

// GenerateSalt 生成随机盐值
func GenerateSalt() string {
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err)
	}
	return hex.EncodeToString(salt)
}
