package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/redis"
)

type Config struct {
	rest.RestConf
	AdminRpc  zrpc.RpcClientConf
	UserRpc   zrpc.RpcClientConf
	PortalRpc zrpc.RpcClientConf
	Database  database.Database
	Redis     redis.Redis
	Apisix    struct {
		ApiKey string `json:"apiKey"`
		Host   string `json:"host"`
	}
}
