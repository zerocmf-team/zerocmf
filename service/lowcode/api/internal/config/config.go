package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/redis"
)

type Config struct {
	rest.RestConf
	TenantRpc zrpc.RpcClientConf
	MongoDB   database.Mongo
	Redis     redis.Redis
	App       struct {
		Domain string `json:",optional"`
	}
}
