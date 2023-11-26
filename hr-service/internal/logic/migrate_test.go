package logic

import (
	"HR_Go/common"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	_, svcCtx := GetCtxAndSvcCtxForTest()
	db := svcCtx.Db

	// 实际上创建 ServiceContext 时已经自动迁移了
	err := common.AutoMigrate(db)
	if err != nil {
		t.Error("gorm auto migrate error", err)
		return
	}
}
