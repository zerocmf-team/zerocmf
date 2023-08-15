package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/redis"

	"github.com/zerocmf/wechatEasySdk"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	TenantRpc zrpc.RpcClientConf
	Database  database.Database
	Apisix    apisix.Apisix
	Redis     redis.Redis
	Wechat    wechatEasySdk.Config
}
