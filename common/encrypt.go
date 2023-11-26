package common

import (
	"crypto/sha256"
	"fmt"
)

// EncryptPassword 加密密码
func EncryptPassword(salt string, id int64, password string) string {
	src := fmt.Sprintf("%s-%s-%d", salt, password, id)
	bytes := sha256.Sum256([]byte(src))
	return fmt.Sprintf("%x", bytes)
}

// VerifyPassword 验证密码
func VerifyPassword(salt string, id int64, password string, encryptedPassword string) bool {
	return EncryptPassword(salt, id, password) == encryptedPassword
}
