package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/redis"
)

type Config struct {
	zrpc.RpcServerConf
	Database  database.Database
	RedisConf redis.Redis
}
