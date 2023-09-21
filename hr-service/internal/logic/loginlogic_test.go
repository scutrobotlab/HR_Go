package logic

import (
	hr_service "HR_Go/hr-service/pb/hr-service"
	"testing"
)

func TestLoginLogic_Login(t *testing.T) {
	ctx, svcCtx := GetCtxAndSvcCtxForTest()
	logic := NewLoginLogic(ctx, svcCtx)

	_, err := logic.Login(&hr_service.LoginReq{
		Token: "",
	})
	if err != nil {
		return
	}
}
