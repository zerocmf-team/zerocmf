package main

import (
	"flag"
	"fmt"

	"zerocmf/service/lowcode/rpc/internal/config"
	"zerocmf/service/lowcode/rpc/internal/server"
	"zerocmf/service/lowcode/rpc/internal/svc"
	"zerocmf/service/lowcode/rpc/types/lowcode"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/lowcode.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		lowcode.RegisterLowcodeServer(grpcServer, server.NewLowcodeServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
