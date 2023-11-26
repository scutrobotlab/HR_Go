package config

import (
	"HR_Go/common"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth common.Auth
	Etcd discov.EtcdConf
}
