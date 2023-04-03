package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zerocmf/common/bootstrap/database"
)

type Config struct {
	rest.RestConf
	App struct {
		Domain string `json:",optional"`
	}
	UserRpc  zrpc.RpcClientConf
	AdminRpc zrpc.RpcClientConf
	Database database.Database
}
