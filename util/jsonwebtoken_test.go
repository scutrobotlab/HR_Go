package util

import (
	"HR_Go/common"
	"fmt"
	"testing"
)

// 不常用
func TestGenerateJwtToken(t *testing.T) {
	auth := getAuth()
	payload := make(map[string]any)
	payload["username"] = "debug"
	payload["permission"] = "admin"

	token, err := GetJwtToken(auth, payload)
	if err != nil {
		t.Failed()
		return
	}
	fmt.Println(token)
}

// 每天执行一次，并将Token复制到Apifox
func TestGetLoginJwtToken(t *testing.T) {
	auth := getAuth()

	type TokenConfig struct {
		UserId         int64
		UserPermission int32
		TokenTitle     string
	}
	tokenConfigs := []TokenConfig{
		{1, common.SuperAdmin, "SuperAdmin"},
		{2, common.Admin, "Admin"},
		{10, common.Applicant, "Applicant"},
	}

	for _, config := range tokenConfigs {
		userInfo := common.UserInfo{config.UserId, config.UserPermission}
		token, err := GetLoginJwtToken(auth, userInfo)
		if err != nil {
			t.Failed()
			return
		}
		fmt.Printf("%s:\n%s\n", config.TokenTitle, token)
	}
}

func getAuth() Auth {
	accessSecret := "12345678.debug"
	accessExpire := int64(12 * 60 * 60)
	return Auth{
		AccessSecret: accessSecret,
		AccessExpire: accessExpire,
	}
}
