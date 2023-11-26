package svc

import (
	"HR_Go/common"
	"HR_Go/dal/query"
	"HR_Go/hr-admin-service/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type ServiceContext struct {
	Common common.Config
	Config config.Config
	Query  *query.Query
	Db     *gorm.DB
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
	err = common.AutoMigrate(db)
	if err != nil {
		logx.Error("gorm auto migrate error", err)
		return nil
	}
	commonConf, err := common.GetConfig("../common/config.yaml")
	if err != nil {
		logx.Error("common config load error", err)
		os.Exit(1)
		return nil
	}
	return &ServiceContext{
		Common: commonConf,
		Config: c,
		Query:  query.Use(db),
		Db:     db,
	}
}

func NewServiceContext4Test(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		Logger: getGormLogger(logger.Info),
	})
	if err != nil {
		return nil
	}
	err = common.AutoMigrate(db)
	if err != nil {
		logx.Error("gorm auto migrate error", err)
		return nil
	}
	commonConf, err := common.GetConfig("../../../common/config.yaml")
	if err != nil {
		return nil
	}
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
