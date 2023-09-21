package svc

import (
	"HR_Go/common"
	"HR_Go/hr-api/internal/config"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
)

type ServiceContext struct {
	Common    common.Config
	Config    config.Config
	HrService hr_service.HrServiceClient
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
		Common:    commonConf,
		Config:    c,
		HrService: hr_service.NewHrServiceClient(client.Conn()),
	}
}
