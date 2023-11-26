package common

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Auth struct {
	AccessSecret string
	AccessExpire int64
}

// GetLoginJwtToken 获取登录Token
func GetLoginJwtToken(auth Auth, userInfo UserInfo) (string, error) {
	payload := make(map[string]any)
	payload["id"] = userInfo.Id
	payload["per"] = userInfo.Permission
	return GetJwtToken(auth, payload)
}

// GetJwtToken
// 获取 Jwt Token
func GetJwtToken(auth Auth, payload map[string]any) (string, error) {
	token, err := generateJwtToken(auth.AccessSecret, time.Now().Unix(), auth.AccessExpire, payload)
	return token, err
}

// generateJwtToken
// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func generateJwtToken(secretKey string, iat, seconds int64, payload map[string]any) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payload {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
