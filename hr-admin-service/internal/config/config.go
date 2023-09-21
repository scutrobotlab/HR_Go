package config

import (
	"HR_Go/common/sms"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource   string
	SmsBaoConfig sms.SmsBaoConfig
}
