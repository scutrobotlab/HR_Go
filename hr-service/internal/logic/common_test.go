package logic

import (
	"HR_Go/hr-service/internal/config"
	"HR_Go/hr-service/internal/svc"
	"context"
)

func GetCtxAndSvcCtxForTest() (context.Context, *svc.ServiceContext) {
	c := config.Config{
		DataSource: "root:123456@(localhost:3306)/hr?charset=utf8mb4&parseTime=True&loc=Local",
	}
	return context.Background(), svc.NewServiceContext4Test(c)
}
