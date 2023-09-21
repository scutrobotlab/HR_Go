package logic

import (
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"testing"
)

func TestAdminLoginLogic_AdminLogin(t *testing.T) {
	ctx, svcCtx := GetCtxAndSvcCtxForTest()
	logic := NewAdminLoginLogic(ctx, svcCtx)

	resp, err := logic.AdminLogin(&hr_admin_service.AdminLoginReq{
		Code: "TYloe9WmXFmtKK129Tg2ZYuBUp8sK95g",
	})
	if err != nil {
		return
	}

	t.Logf("resp: %+v", resp)
}
