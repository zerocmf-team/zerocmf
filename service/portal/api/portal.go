package main

import (
	"flag"
	"fmt"
	"zerocmf/common/bootstrap/middleware"

	"zerocmf/service/portal/api/internal/config"
	"zerocmf/service/portal/api/internal/handler"
	"zerocmf/service/portal/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/portal.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 初始化
	server.Use(middleware.NewSiteMiddleware(ctx.Data).Handle)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
