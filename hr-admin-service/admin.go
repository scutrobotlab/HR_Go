package main

import (
	"HR_Go/hr-admin-service/internal/cronjob"
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"

	"HR_Go/hr-admin-service/internal/config"
	"HR_Go/hr-admin-service/internal/server"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		hr_admin_service.RegisterHrServiceServer(grpcServer, server.NewHrServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	myCron := cron.New()
	defer myCron.Stop()
	cronjob.InitCronJob(ctx, myCron)
	myCron.Start()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
