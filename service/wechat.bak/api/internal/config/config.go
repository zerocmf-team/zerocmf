package config

import (
	"github.com/zeromicro/go-zero/rest"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/redis"
)

type Config struct {
	rest.RestConf
	App struct {
		Domain string `json:",optional"`
	}
	Database database.Database
	Redis    redis.Redis
}
