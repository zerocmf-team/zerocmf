package main

import (
	"flag"
	"fmt"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/middleware"

	"zerocmf/service/user/api/internal/config"
	"zerocmf/service/user/api/internal/handler"
	"zerocmf/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	Init.SetSalts(c.Database.AuthCode)

	ctx := svc.NewServiceContext(c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 初始化
	server.Use(middleware.NewSiteMiddleware(ctx.Data).Handle)
	
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
