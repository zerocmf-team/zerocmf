package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	TenantRpc zrpc.RpcClientConf
	ShopRpc   zrpc.RpcClientConf
}
