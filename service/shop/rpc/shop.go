package main

import (
	"flag"
	"fmt"

	"zerocmf/service/shop/rpc/internal/config"
	categoryserviceServer "zerocmf/service/shop/rpc/internal/server/categoryservice"
	productattrserviceServer "zerocmf/service/shop/rpc/internal/server/productattrservice"
	productserviceServer "zerocmf/service/shop/rpc/internal/server/productservice"
	shopserviceServer "zerocmf/service/shop/rpc/internal/server/shopservice"
	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/shop.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		shop.RegisterShopServiceServer(grpcServer, shopserviceServer.NewShopServiceServer(ctx))
		shop.RegisterCategoryServiceServer(grpcServer, categoryserviceServer.NewCategoryServiceServer(ctx))
		shop.RegisterProductServiceServer(grpcServer, productserviceServer.NewProductServiceServer(ctx))
		shop.RegisterProductAttrServiceServer(grpcServer, productattrserviceServer.NewProductAttrServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
