package main

import (
	"flag"
	"fmt"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/service/admin/rpc/internal/config"
	"zerocmf/service/admin/rpc/internal/server"
	"zerocmf/service/admin/rpc/internal/svc"
	"zerocmf/service/admin/rpc/types/admin"

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

	Init.SetSalts(c.Database.AuthCode)

	ctx := svc.NewServiceContext(c)
	svr := server.NewAdminServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		admin.RegisterAdminServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
