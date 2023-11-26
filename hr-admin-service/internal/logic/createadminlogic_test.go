package logic

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"testing"
)

func TestCreateAdminLogic_CreateAdmin(t *testing.T) {
	ctx, svcCtx := GetCtxAndSvcCtxForTest()

	logic := NewCreateAdminLogic(ctx, svcCtx)
	resp, err := logic.CreateAdmin(&hr_admin_service.CreateAdminReq{
		Admin: &hr_admin_service.Admin{
			Name: "管理员",
			Groups: []*hr_admin_service.AdminGroup{
				{Name: "机械"},
				{Name: "电控"},
			},
		},
		Password: "12345678",
	})
	if err != nil {
		return
	}

	t.Logf("resp: %v", resp)
}
