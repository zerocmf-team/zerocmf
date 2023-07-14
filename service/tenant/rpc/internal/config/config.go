package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zerocmf/common/bootstrap/database"
)

type Config struct {
	zrpc.RpcServerConf
	AdminRpc zrpc.RpcClientConf
	Database database.Database
}
