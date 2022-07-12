package config

import (
	"zerocmf/common/bootstrap/database"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	App struct{
		Domain string `json:",optional"`
	}
	UserRpc zrpc.RpcClientConf
	AdminRpc zrpc.RpcClientConf
	Database database.Database
}
