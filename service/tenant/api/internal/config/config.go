package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/redis"
)

type Config struct {
	rest.RestConf
	App struct {
		Domain string `json:",optional"`
	}
	AdminRpc   zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
	PortalRpc  zrpc.RpcClientConf
	LowcodeRpc zrpc.RpcClientConf
	TenantRpc  zrpc.RpcClientConf
	ShopRpc    zrpc.RpcClientConf
	Database   database.Database
	Redis      redis.Redis
	Apisix     apisix.Apisix
}
