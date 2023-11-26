package common

import (
	"crypto/sha256"
	"fmt"
)

// EncryptPassword 加密密码
func EncryptPassword(id int64, password string) string {
	src := fmt.Sprintf("%s-%d", password, id)
	bytes := sha256.Sum256([]byte(src))
	return fmt.Sprintf("%x", bytes)
}

// VerifyPassword 验证密码
func VerifyPassword(id int64, password string, encryptedPassword string) bool {
	return EncryptPassword(id, password) == encryptedPassword
}
