package config

import (
	"zerocmf/common/bootstrap/db"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	App struct{
		Domain string `json:",optional"`
	}
	Database database.Database
}
