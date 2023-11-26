package logic

import (
	"HR_Go/common"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	_, svcCtx := GetCtxAndSvcCtxForTest()
	db := svcCtx.Db

	err := common.AutoMigrate(db)
	if err != nil {
		t.Error("gorm auto migrate error", err)
		return
	}
}
