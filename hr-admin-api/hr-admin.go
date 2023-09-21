package main

import (
	"flag"
	"fmt"
	"net/http"

	"HR_Go/hr-admin-api/internal/config"
	"HR_Go/hr-admin-api/internal/handler"
	"HR_Go/hr-admin-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/hr-admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(unauthorizedCallback))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func unauthorizedCallback(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		return
	}
}
