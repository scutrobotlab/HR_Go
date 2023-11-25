package svc

import (
	"HR_Go/common"
	"HR_Go/dal/query"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"HR_Go/hr-service/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type ServiceContext struct {
	Common       common.Config
	Config       config.Config
	Query        *query.Query
	Db           *gorm.DB
	AdminService hr_admin_service.HrServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		Logger: getGormLogger(logger.Warn),
	})
	if err != nil {
		logx.Error("gorm open error", err)
		os.Exit(1)
		return nil
	}
	commonConf, err := common.GetConfig("../common/config.yaml")
	if err != nil {
		logx.Error("common config load error", err)
		os.Exit(1)
		return nil
	}
	adminEtcd := c.Etcd
	adminEtcd.Key = c.AdminServiceKey
	return &ServiceContext{
		Common: commonConf,
		Config: c,
		Query:  query.Use(db),
		Db:     db,
		AdminService: hr_admin_service.NewHrServiceClient(zrpc.MustNewClient(zrpc.RpcClientConf{
			Etcd: adminEtcd,
		}).Conn()),
	}
}

func NewServiceContext4Test(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		Logger: getGormLogger(logger.Info),
	})
	if err != nil {
		return nil
	}
	commonConf, err := common.GetConfig("../../../common/config.yaml")
	if err != nil {
		return nil
	}
	adminEtcd := c.Etcd
	adminEtcd.Key = c.AdminServiceKey
	return &ServiceContext{
		Common: commonConf,
		Config: c,
		Query:  query.Use(db),
		Db:     db,
	}
}

func getGormLogger(logLevel logger.LogLevel) logger.Interface {
	return logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             time.Duration(200) * time.Millisecond,
		Colorful:                  true,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      false,
		LogLevel:                  logLevel,
	})
}
