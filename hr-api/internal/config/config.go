package config

import (
	"HR_Go/util"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth util.Auth
	Etcd discov.EtcdConf
}
