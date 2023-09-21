package svc

import (
	"HR_Go/common"
	"HR_Go/hr-admin-api/internal/config"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
)

type ServiceContext struct {
	Common       common.Config
	Config       config.Config
	AdminService hr_admin_service.HrServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: c.Etcd,
	})
	commonConf, err := common.GetConfig("../common/config.yaml")
	if err != nil {
		logx.Error("common config load error", err)
		os.Exit(1)
		return nil
	}
	return &ServiceContext{
		Common:       commonConf,
		Config:       c,
		AdminService: hr_admin_service.NewHrServiceClient(client.Conn()),
	}
}
