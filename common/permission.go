package common

import (
	"context"
	"encoding/json"
)

const (
	Unknown    = iota // 没有登录
	Applicant         // 面试者
	Admin             // 管理员（面试官）
	SuperAdmin        // 超级管理员（虎王id=1)
)

type (
	UserInfo struct {
		Id         int64 `json:"id"`
		Permission int32 `json:"per"`
	}
)

func GetUserInfo(ctx context.Context) *UserInfo {
	id := ctx.Value("id")
	var idValue int64
	if id == nil {
		idValue = Unknown
	} else {
		idValue, _ = id.(json.Number).Int64()
	}

	permission := ctx.Value("per")
	var permissionValue int64
	if permission == nil {
		permissionValue = Unknown
	} else {
		permissionValue, _ = permission.(json.Number).Int64()
	}

	return &UserInfo{Id: idValue, Permission: int32(permissionValue)}
}
